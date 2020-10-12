import datetime
import json

def main(args):
  try:
    print(args)
    count = args.get("count", 1)
    fiscal_code = args.get("fiscal_code", "")
    if fiscal_code == "":
        fiscal_code = "AAAAAA00A00A000A"

    rec = {}
    if "amount" in args and args["amount"] != "":
        try: rec["amount"]=int(args["amount"])   
        except Exception as e: print(str(e))
    if "due_date" in args and args["due_date"] != "":
        try: rec["due_date"] = datetime.datetime.fromisoformat(args["due_date"]).isoformat()
        except Exception as e: print(str(e))

    res = []
    for i in range(0, int(count)):
        code = "%s:%d" %  (fiscal_code, i)
        rec1 = rec.copy()
        rec1["fiscal_code"]  = code
        rec1["subject"] =  "Benvenuto  %s" % code
        rec1["markdown"] = "# Benvenuto, %s\nTi diamo il benvenuto. Questo è un messaggio di *benvenuto* per mostrare come generare markdown in **HTML**. Ricordare che deve essere un testo lungo." % code
        res.append(rec1)
    
    return { "body": json.dumps(res, indent=2)}
  except Exception as e:
    return { "body": { "error": str(e)}}


    