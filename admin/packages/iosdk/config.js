function main(args) {
    return {
       "io-apikey": args["io-apikey"],
       "whisk-apihost-local": args["whisk-apihost-local"],
       "whisk-apihost-docker": args["whisk-apihost-docker"],
       "whisk-apikey": args["whisk-apikey"],
       "whisk-namespace": args["whisk-namespace"],
       "env": process.env
    }
}
