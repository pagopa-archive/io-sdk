function main(args) {
    if(args.url) {
        return new Promise(function(resolve) {
            let url = new URL(args.url);
            if(args.username)
                url.username = args.username;
            if(args.password)
                url.password = args.password
            require('https').get(url, function handleHttp(resp)  {
                let data = '';
                resp.on('data', (chunk) => { data += chunk; });
                resp.on('end', () => { resolve({"body": data}) });
              }
            ).on("error", (err) => {resolve({"error": err.message})})
        })
    }
    return { "body": { "form": [
        {
            "type": "message",
            "name": "note",
            "description": "Use url 'https://raw.githubusercontent.com/pagopa/io-sdk/master/docs/sample.json' for sample data"
        },
        {
            "name": "url",
            "description": "URL",
            "type": "string",
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
            "required": true
        }]
    } }
}