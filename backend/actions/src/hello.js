function main(args) {
    var name = args.name || "world"
    var msg = "Hello, " + name;
    return { "body":  msg };
}
