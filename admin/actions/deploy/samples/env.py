import os
def main(args):
    res = {}
    env = os.environ
    for k in env:
        if k.startswith("__OW"):
            res[k] = env[k]
    return res
