let fs = require("fs")

let f = fs.openSync("./FILES", "w")

fs.readdirSync(".").forEach(v=>{
    if (v.endsWith(".proto")) {
        fs.writeSync(f, v + " ")
    }
})

fs.closeSync(f)
