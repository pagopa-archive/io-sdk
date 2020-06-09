import base64
import time
import os
import json
import pip
import redis

def main(args):
    red =  redis.Redis(host=os.environ.get("__OW_REDIS", "127.0.0.1"))
    body = args["__ow_body"]
    if args["__ow_headers"]["content-type"] == "application/json":
        body = base64.b64decode(body).decode("utf-8")
    body = json.loads(body)
    # with open("import.json", "r") as f: body = json.loads(f.read())
    data = [ ("message:%s" % m["fiscal_code"], m)  for m in body["data"] ]
    res = {}
    for (k,v) in data:
        v = json.dumps(v).encode("utf-8")
        res[k] = red.set(k,v)
    return { "body": res }
