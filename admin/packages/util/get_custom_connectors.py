from nimbella import redis
import json


def main(args):

    red = redis()

    connectors = red.get('connectors')

    connectors = connectors.decode("utf-8").split(",") if connectors else []

    return {"body": {"details": {
        "connectors": connectors
        }}}
