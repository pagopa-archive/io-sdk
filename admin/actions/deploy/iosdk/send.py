#%%
import requests
import json
import sys

#%%
def send(url, hdr, msg):
    r = requests.post(url, json=msg, headers=hdr)
    #print(r)
    if r.status_code == 201:
        return {"body": json.loads(r.text) }
    return { "body": {"error": r.text}}

#%%
def main(args):
    """
    >>> args = {}
    >>> print(main(args))
    {'body': {'error': "missing argument 'io-messages'"}}
    >>> args['io-messages'] = "https://api.io.italia.it/api/v1/messages"
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
    """
    try:
        url = args['io-messages']
        hdr = {"Ocp-Apim-Subscription-Key": args['io-apikey']}
        msg = {
            "content": {
                "subject": args['subject'],
                "markdown": args['markdown'],
            },
            "fiscal_code": args['fiscal_code']
        }
        if "amount" in args:
            pd = {
                "amount": int(args["amount"]),
                "notice_number": "000000000000000001",
            }
            if "notice_number" in args:
                pd["notice_number"] = args["notice_number"]
            if "due_date" in args:
                msg["due_date"] = args["due_date"]
                if "invalid_after_due_date" in args:
                    pd["invalid_after_due_date"]= bool(args['invalid_after_due_date'])
            msg["content"]["payment_data"] = pd
        return send(url, hdr, msg)
    except KeyError as e:
        return { "body": { "error": "missing argument %s" % str(e)}}
    except ValueError as e:
        return { "body": { "error": "conversion error"}}
    except:
        return { "body": { "error": str(sys.exc_info()[0]) } }

if __name__ == "__main__":
    import doctest
    doctest.testmod()