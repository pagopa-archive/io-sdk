function main(args) {
    // check if login data
    if(args.login) {
        if(args.login == "demo" && args.password =="demo") {
            res = { "data": sampleData }
        } else {
            res = { "error": "bad login or password"}
        }
    } else {
        res = {"form": sampleForm}
    }
    return {
        "body": JSON.stringify(res)
    }
}

const sampleForm = [
    {
        "type": "message",
        "name": "note",
        "description": "Sample user is 'demo' with password 'demo'",
    },
    {
        "name": "login",
        "description": "Login",
        "type": "string",
        "required": true
    },
    {
        "name": "password",
        "description": "Password",
        "type": "password",
        "required": true
    },
    {
        "type": "message",
        "name": "note",
        "description": "Replace this example deploying a custom importer.",
    },
]

const sampleData = [
    {
    "amount": 0,
    "due_date": 44197,
    "fiscal_code": "ISPXNB32R82Y766F",
    "invalid_after_due_date": false,
    "markdown": "# Welcome, Giovanni Rossi\\\\nYour fiscal code is ISPXNB32R82Y766F\\\\nI hope you will enjoy IO.",
    "notice_number": 1,
    "subject": "Welcome to IO, Giovanni",
  },
  {
    "amount": 0,
    "due_date": 44197,
    "fiscal_code": "ISPXNB32R82Y766D",
    "invalid_after_due_date": false,
    "markdown": "# Welcome, Luca Rossi\\\\nYour fiscal code is ISPXNB32R82Y766D\\\\n.I hope you will enjoy IO.",
    "notice_number": 1,
    "subject": "Welcome to IO, Luca",
  }
]