import base64
import time
import os
import json
import pip

def main(args):
    body = args["__ow_body"]
    if args["__ow_headers"]["content-type"] == "application/json":
        body = base64.b64decode(body).decode("utf-8")
    body = json.loads(body)
    return  { "body": body }

