import json, os

try: import redis
except:
    import pip
    pip.main(["install", "redis"])
    import redis

def main(args):
    red =  redis.Redis(host=os.environ.get("__OW_REDIS", "127.0.0.1"))
    res = {}
    if "set" in args:
        arg = args["set"]
        [k,v] = arg.split("=", 1)
        v = json.dumps(v).encode("utf-8")
        #print{}}
        res = {}
        if red.set(k,v):
            res["set:%s"%arg] = True
        else:
            res["error"] = "cannot set %s to %s" % (k,v)
    if "get" in args:
        k = args["get"]
        v = red.get(k)
        if v:
            v = json.loads(v.decode("utf-8"))
            try: res[k] = json.loads(v)
            except: res[k] = v
        else:
            res["error"] = "cannot find %s" % k
    if "del" in args:
        k = args["del"]
        if red.delete(k):
            res["del:%s"%k]=True
        else:
            res["error"] = "cannot delete %s" % k
    if "scan" in args:
        pattern = args["scan"]
        (_, ls) = red.scan(match=pattern)
        ls = [ i.decode("utf-8") for i in ls]
        res["scan"] = ls
    if "__ow_method" in args:
        return { "body": json.dumps(res) }
    return res

if __name__=="__main__":
    import sys
    args = {}
    for arg in sys.argv[1:]:
        [k,v] = arg.split(":", 1)
        args[k] = v
    print(json.dumps(main(args), indent=2))
