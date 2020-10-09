#%%
import requests
import json
import sys

#%%
def send(url, key, msg):
    hdr = {"Ocp-Apim-Subscription-Key": key}
    print("curl -X POST -d '%s' -H 'Ocp-Apim-Subscription-Key: %s' %s" %( json.dumps(msg), key, url))
    r = requests.post(url, json=msg, headers=hdr)
    if r.status_code == 201:
        return {"body": json.loads(r.text) }
    return { "body": {"error": r.text}}

#%%
def main(args):
    """
    >>> args = {}
    >>> print(main(args))
    {'body': {'error': "missing argument 'io-messages'"}}
    >>> args['io-messages'] = ""
    >>> print(main(args))
    {'body': {'error': "missing argument 'io-apikey'"}}
    >>> args['io-apikey'] = "483b7b1f3a974b45b5c44a43538c9255"
    >>> args['subject'] = "Welcome new user !"
    >>> args['markdown'] = "This is a markdown header to show how easily markdown can be converted to **HTML** Remember: this has to be a long text."
    >>> print(main(args))
    {'body': {'error': "missing argument 'fiscal_code'"}}
    >>> args['fiscal_code'] = "DGRMLL66R65H769R"
    >>> print(len(main(args)['body']['id']))
    26
    >>> args["amount"] = "a"
    >>> print(main(args))
    {'body': {'error': 'conversion error'}}
    >>> args["amount"] = 1
    >>> print(len(main(args)['body']['id']))
    26
    >>> args["due_date"] = "2021-01-01T00:00:00.000Z"
    >>> args["notice_number"]: "000000000000000001"
    >>> print(len(main(args)['body']['id']))
    26
    >>> #args['io-apikey'] = "c64b38f22e8344a18d63d7c524b171cc"
    >>> #args['fiscal_code'] = "SCCNDR68T05L483L"
    >>> #print(main(args))

    """
    try:
        url = args['io-messages']
        key = args['io-apikey']
        msg = {
            "content": {
                "subject": args['subject'],
                "markdown": args['markdown'],
            },
            "fiscal_code": args['fiscal_code']
        }
        if "amount" in args and args["amount"] != "":
            pd = {
                "amount": int(args["amount"]),
                "notice_number": "000000000000000001",
            }
            if "notice_number" in args:
                pd["notice_number"] = ("000000000000000000" + args["notice_number"])[-18:]
            if "due_date" in args and args["due_date"] !="" :
                msg["content"]["due_date"] = args["due_date"]
                if "invalid_after_due_date" in args and args["invalid_after_due_date"] !="":
                    pd["invalid_after_due_date"]= bool(args['invalid_after_due_date'])
            msg["content"]["payment_data"] = pd
        return send(url, key, msg)
    except KeyError as e:
        return { "body": { "error": "missing argument %s" % str(e)}}
    except ValueError as e:
        return { "body": { "error": "conversion error"}}
    except Exception as e:
        return { "body": { "error": str(e) } }

if __name__ == "__main__":
    import doctest
    doctest.testmod()
