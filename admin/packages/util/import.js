function main(args) {
    if(args.url) {
        let url = args.url

        // use http or https
        let client = require("https")
        if(args.url.startsWith("http:"))
            client = require("http")
        
        // fix calls to API HOST
        if(url.startsWith("http://localhost:3280")) 
            url = process.env.__OW_API_HOST + url.substring(21)

        // build url
        let authURL = new URL(url);
        if(args.username)
            authURL.username = args.username;
        if(args.password)
            authURL.password = args.password

        let body = args["jsonargs"] || "{}"
        return new Promise(function(resolve) {
            let handleHttp = (resp) => {
                resp.setEncoding('utf8');
                let data = '';
                resp.on('data', (chunk) => { data += chunk; });
                resp.on('end', () => { resolve({"body": { "data": JSON.parse(data) }})});
            }
            let handleError = (err) => {resolve({"error": err.message})}
            if(args.useget == "true") {
                console.log("GET "+url)
                client.get(authURL, handleHttp).on("error", handleError)
            } else {
                console.log("POST "+url+" data="+body)
                let req = client.request(authURL, {
                    method: "POST",
                    headers: {
                        'Content-Type': 'application/json',
                        'Content-Length': body.length
                    }
                }, handleHttp)
                req.on("error", handleError)
                req.write(body)
                req.end()
            }
        })
    }
    return { "body": { "form": [
        {
            "type": "message",
            "name": "note",
            "description": "Import from JSON source. Replace the sample endpoint with your service."
        },
        {
            "name": "url",
            "description": "URL",
            "type": "string",
            "value": "http://localhost:3280/api/v1/web/guest/util/sample",
            "required": false
        },
        {
            "name": "username",
            "description": "Username",
            "type": "string",
            "required": false
        },
        {
            "name": "password",
            "description": "Password",
            "type": "password",
            "required": false
        },
        {
            "name": "jsonargs",
            "description": "Args in JSON format to POST:",
            "type": "textarea",
            "value": "{\"count\":1,\n\"fiscal_code\":\"AAAAAA00A00A000A\",\n\"amount\":0,  \"due_date\":\"\", \"notice_number\":\"\"\n}",
            "required": false
        },
        {
            "name": "useget",
            "description": "Use GET (ignoring JSON args):",
            "type": "checkbox",
            "required": false
        }
    ]
    } }
}