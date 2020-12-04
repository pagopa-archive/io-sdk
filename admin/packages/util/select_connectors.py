from requests import get
from base64 import b64encode


class RequestError(Exception):

    def __init__(self, message="Error during request"):
        super().__init__(message)


class GithubService():

    GITHUB_REPO_HOST = 'https://sf3ris.github.io'
    GITHUB_REPO_NAME = '/io-gateway-connectors/'
    GITHUB_REPO_INDEX = 'index.json'

    def __init__(self):
        pass

    def get_connectors(self):

        url = self.GITHUB_REPO_HOST \
            + self.GITHUB_REPO_NAME \
            + self.GITHUB_REPO_INDEX

        response = get(url)

        if response.status_code != 200:
            raise RequestError

        connectors = response.json()
        return connectors

    def get_connector_file(self, connector_name):
        url = self.GITHUB_REPO_HOST + self.GITHUB_REPO_NAME + connector_name
        response = get(url)

        if response.status_code != 200:
            raise RequestError("Error downloading binary connector file")

        return response.text


def main(args):

    connectors = args.get('connectors')

    if not connectors:
        return {"body": {"detail": "connectors list is empty"}}

    github_service = GithubService()
    connectors_response = github_service.get_connectors()
    valid_connectors = [connector for connector in connectors_response
                        if connector['name'] in connectors
                        ]

    if not valid_connectors:
        return {"body": {"detail": "no valid connectors names provided"}}

    for connector in valid_connectors:

        binary_connector = github_service.get_connector_file(
            connector['file_path']
        )
        return {"body": {"detail": "asd"}}
