function main(args) {
    if("__ow_body" in args) {
        let buf = Buffer.from(args["__ow_body"], "base64")
        args["__ow_body"] = buf.toString()
    }
    return {
        "body": args
    }
}
