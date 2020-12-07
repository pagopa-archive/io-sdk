from requests import get, put
from base64 import b64encode
from os import environ
from json import loads
from nimbella import redis


def main(args):

    connectors = loads(args['__ow_body'])['connectors']

    if not connectors:
        return {"body": {"detail": "connectors list is empty"}}

    github_service = GithubService()
    connectors_response = github_service.get_connectors()
    valid_connectors = [connector for connector in connectors_response
                        if connector['name'] in connectors
                        ]

    if not valid_connectors:
        return {"body": {"detail": "no valid connectors names provided"}}

    ow_service = OWService()
    red = redis()

    responses = []
    created_connectors = []
    for connector in valid_connectors:

        binary_connector = github_service.get_connector_file(
            connector['file_path']
        )

        create_response = ow_service.create_action(
            connector['name'],
            binary_connector
        )

        responses.append(create_response)

        if 'error' not in create_response:
            created_connectors.append(connector['name'])

    red.set('connectors', ",".join(created_connectors))
    return {"body": {"detail": responses}}


class RequestError(Exception):

    def __init__(self, message="Error during request"):
        super().__init__(message)


class GithubService():
    """
    Class responsible to communicate with the Github Api

    Raises:
        RequestError: Error during request against Github Api Server
        RequestError: Error during request against Github Api Server
    """
    GITHUB_REPO_HOST = 'https://sf3ris.github.io'
    GITHUB_REPO_NAME = '/io-gateway-connectors/'
    GITHUB_REPO_INDEX = 'index.json'

    def __init__(self):
        pass

    def get_connectors(self):
        """
        Retrieves a list of the current available connectors

        Raises:
            RequestError: Error during request against Github Api Server

        Returns:
            list: List of available connectors
        """
        url = self.GITHUB_REPO_HOST \
            + self.GITHUB_REPO_NAME \
            + self.GITHUB_REPO_INDEX

        response = get(url)

        if response.status_code != 200:
            raise RequestError

        connectors = response.json()
        return connectors

    def get_connector_file(self, connector_name):
        """
        Retrive the connector binary data

        Args:
            connector_name (string): Connector's name

        Raises:
            RequestError: Error during request against Github Api Server

        Returns:
            bytes: Connector's binary data
        """
        url = self.GITHUB_REPO_HOST + self.GITHUB_REPO_NAME + connector_name
        response = get(url)

        if response.status_code != 200:
            raise RequestError("Error downloading binary connector file")

        return response.content


class OWService():
    """
    Class responsible to communicate with the OpenWhisk REST API
    """
    def __init__(self):
        self.__host = environ["__OW_API_HOST"]
        self.__api_key = environ["__OW_API_KEY"].split(":")
        self.__namespace = "guest"

    def create_action(self, action_name, binary_data):
        """
        Creates a new connection

        Args:
            action_name (string): Action's name to be create
            binary_data (bytes): Connector's binary data
        """
        action_name = "util/{}".format(action_name)
        url = "{}/api/v1/namespaces/{}/actions/{}?overwrite=true&web=true".format(
            self.__host,
            self.__namespace,
            action_name
        )

        data = {
            "namespace": self.__namespace,
            "name": "123",
            "exec": {
                "kind": self.__get_connector_kind(action_name),
                "binary": "true",
                "code": b64encode(binary_data).decode("utf-8", "ignore"),
                "main": "Main"
            }
        }
        headers = {'Content-type': 'application/json'}

        try:

            response = put(
                url,
                json=data,
                auth=(self.__api_key[0], self.__api_key[1]),
                headers=headers
            )
            if('error' in response.text):
                raise Exception(response.text)

            return "{} connector's action created successfully.".format(
                action_name
            )

        except Exception as e:
            return {
                    "message": "Error creating action for connector {}".format(
                        action_name
                    ),
                    "error": str(e)
            }

    def __get_connector_kind(self, connector_name):

        if('java' in connector_name):
            return 'java:8'
        elif('python' in connector_name):
            return 'python:3'
        elif('node' in connector_name):
            return 'nodejs:10'
        elif('php' in connector_name):
            return 'php:7.4'
        elif('go' in connector_name):
            return 'go:1.11'
        else:
            return 'blackbox'
