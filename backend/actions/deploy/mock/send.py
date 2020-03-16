import base64
import time
import os
import redis

def main(args):
    red =  redis.Redis(host=os.environ.get("__OW_REDIS", "127.0.0.1"))
    body = args
    code = args.fiscal_code
    red.set(code, args["__ow_body"])
    id = str(int(time.time()*1000))
    return { "id": id}

