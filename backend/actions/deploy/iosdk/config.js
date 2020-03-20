function main(args) {
    return {
       "io-apikey": args["io-apikey"],
       "whisk-apihost": args["whisk-apihost"],
       "whisk-apikey": args["whisk-apikey"],
       "whisk-namespace": args["whisk-namespace"],
       "apihost": process.env["__OW_API_HOST"]
    }
}
