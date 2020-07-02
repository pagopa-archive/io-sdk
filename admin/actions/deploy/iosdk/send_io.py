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
        return send(url, hdr, msg)
    except KeyError as e:
        return { "body": { "error": "missing argument %s" % str(e)}}
    except:
        return { "body": { "error": sys.exc_info()[0] } }

if __name__ == "__main__":
    import doctest
    doctest.testmod()
