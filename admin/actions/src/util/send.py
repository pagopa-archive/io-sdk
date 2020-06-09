import base64
import time
import os
import json
import pip
import zlib
import redis

def main(args):
    dest = args.get("fiscal_code")
    subj = args.get("subject")
    mesg = args.get("markdown") 
    if dest and subj and mesg:
        id = str(zlib.crc32(dest.encode("utf-8")))
        red =  redis.Redis(host=os.environ.get("__OW_REDIS", "127.0.0.1"))
        data = {"subject": subj, "markdown": mesg, "fiscal_code": dest}
        data = json.dumps(data).encode("utf-8")
        red.set("sent:%s" % dest, data)
        return {"body": {"id": id} }

    return { "body": { "detail": "validation errors"}}
