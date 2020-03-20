import base64
import time
import os
import json
import pip

try: import redis
except:    
    pip.main(["install", "redis"])
    import redis

def main(args):
    body = args["__ow_body"]
    if args["__ow_headers"]["content-type"] == "application/json":
        body = base64.b64decode(body).decode("utf-8")
    body = json.loads(body)

    if ( body.get("fiscal_code") and 
         body.get("content") and
         body.get("content").get("subject") and
         body.get("content").get("markdown") and
         body.get("content").get("due_date") ):
            id = str(int(time.time()*1000))
            red =  redis.Redis(host=os.environ.get("__OW_REDIS", "127.0.0.1"))
            data = json.dumps(body).encode("utf-8")
            red.set("message:%s" % body["fiscal_code"], data)
            return {"body": {"id": id} }

    return { "body": { "detail": "validation errors"}}


