"use strict"

let fs = require("fs")
let path = require("path")
let child_process = require("child_process")

const options   = { encoding: 'utf8', flag: 'a+' }
const reg_const = /\s*(\w+)\s*=\s*(\d+)\s*\/\/\s*(.*)$/
const reg_proto = /\s*message\s+(\w+)\s*{\s*\/\/\s*opcode\:\s*(\d+)/

// ----------------------------------------------------------------------------

let proto_files   = []
let proto_matches = {}

// ----------------------------------------------------------------------------

function find_proto_files() {
    fs.readdirSync(".").forEach(fn => {
        if (fn.endsWith(".proto")) {
            proto_files.push(fn)
            proto_matches[fn] = []
        }
    })
    proto_files.sort()
    return proto_files
}


function gen_handler_go() {
    proto_files.forEach(fn => {
        if (fn == "0.type.proto" || fn == "1_session.proto") {
            return
        }

        let name = `../../../server/player/handler_${path.basename(fn, ".proto")}.go`

        if (!fs.existsSync(name)) {
            let file = fs.openSync(name, "w")
            let head =
`package player

import (
    "core/tcp"
    "github.com/gogo/protobuf/proto"
    "public/protocol"
    "public/protocol/msg"
)

`
            fs.writeFileSync(file, head, options)
            fs.closeSync(file)
        }

        proto_matches[fn].forEach(match => {
            if (!match[1].endsWith("Request")) {
                return
            }

            let handler = `handler_${match[1]}`
            if (!fs.readFileSync(name, "utf8").includes(handler)) {
                let data =
`func ${handler}(plr *Player, packet *tcp.Packet) {
    req := &msg.${match[1]}{}
    res := &msg.${match.resp}{}
    proto.Unmarshal(packet.Data, req)

    // TODO

    plr.SendPacket(protocol.MSG_SC_${match.resp}, res)
}

`
                fs.writeFileSync(name, data, { flag: "a+" })
            }
        })
    })
}


function gen_handler_lua() {
    proto_files.forEach(fn => {
        if (fn == "0.type.proto") {
            return
        }

        let name = `../../../../../2dgame/simulator/win32/src/message/msg_${path.basename(fn, ".proto")}.lua`

        if (!fs.existsSync(name)) {
            let head =
`require "message.opcode"

local Event      = require "core.event"
local EventMgr   = require "core.event_mgr"

local md         = MessageDispatcher
local Opcode     = Opcode

`
            fs.writeFileSync(name, head, options)
        }

        proto_matches[fn].forEach(match => {
            if (!match[1].endsWith("Response") && !match[1].endsWith("Update")) {
                return
            }

            let handler = `md[Opcode.${match.opcode}]`
            if (!fs.readFileSync(name, "utf8").includes(handler)) {
                let data =
`
local function ${match.opcode}(tab)
    print("msg:${match.opcode}")
    -- TODO
end
${handler} = ${match.opcode}
`
                fs.writeFileSync(name, data, { flag: "a+" })
            }
        })
    })
}


// 处理一个文件
function extract_proto_file(file) {
    fs.readFileSync(file, "utf8").split("\n").forEach(line => {
        var matched = line.match(reg_proto)
        if (!matched) {
            return
        }

        if (matched[1].endsWith("Request")) {
            matched.opcode = `MSG_CS_${matched[1]}`
            matched.resp = matched[1].replace(/Request$/, "Response")
        } else if (matched[1].endsWith("Response") || matched[1].endsWith("Update")) {
            matched.opcode = `MSG_SC_${matched[1]}`
        } else {
            return
        }
        proto_matches[file].push(matched)
    })
}


function gen_register_go(fn) {
    let file = fs.openSync(fn, "w")
    let head =
`// Do NOT edit this file manually

package player

import (
    "public/protocol"
)

func init() {
`
    fs.writeFileSync(file, head, options)

    proto_files.forEach(f => {
        if (f == "1_session.proto") {
            return
        }
        proto_matches[f].forEach(match => {
            if (!match[1].endsWith("Request")) {
                return
            }
            let line = `\tregister_handler(protocol.${match.opcode}, handler_${match[1]})\n`
            fs.writeFileSync(file, line, options)
        })
    })

    fs.writeFileSync(file, "}\n", options)
    fs.closeSync(file)

    child_process.execFileSync("go", ["fmt", fn])
}


function gen_register_lua(fn) {
    let file = fs.openSync(fn, "w")
    let head =
`-- Do NOT edit this file manually

require "message.opcode"

`
    fs.writeFileSync(file, head, options)

    proto_files.forEach(name => {
        if (name == "0.type.proto") {
            return
        }
        let line = `require "message.msg_${path.basename(name, ".proto")}"\n`
        fs.writeFileSync(file, line, options)
    })

    fs.closeSync(file)
}


function gen_proto_file_go(fn) {
    let file = fs.openSync(fn, "w")
    let head =
`// Do NOT edit this file manually

package protocol

const (
`
    fs.writeFileSync(file, head, options)

    proto_files.forEach(f => {
        proto_matches[f].forEach(match => {
            let line = `\t${match.opcode} uint16 = ${match[2]}\n`
            fs.writeFileSync(file, line, options)
        })
    })

    fs.writeFileSync(file, ")\n", options)
    fs.closeSync(file)

    child_process.execFileSync("go", ["fmt", fn])
}


function gen_proto_file_lua(fn) {
    let span = 32
    let file = fs.openSync(fn, "w")
    let head =
`-- Do NOT edit this file manually

cc.exports.Opcode  = {}
cc.exports.MsgName = {}

local Opcode  = Opcode
local MsgName = MsgName

`
    fs.writeFileSync(file, head, options)

    proto_files.forEach(f => {
        proto_matches[f].forEach(match => {
            let d = span - match[1].length
            let s = ' '.repeat(d)
            let line = `Opcode.${match.opcode}${s}= ${match[2]};\t\tMsgName[Opcode.${match.opcode}]${s}= "msg.${match[1]}";\n`
            fs.writeFileSync(file, line, options)
        })
    })

    fs.closeSync(file)
}


function export_const_file(go_file, lua_file) {
    let span = 40
    let file = fs.openSync(lua_file, "w")
    let head =
`-- Do NOT edit this file manually

cc.exports.Error =
{
`
    fs.writeFileSync(file, head, options)

    fs.readFileSync(go_file, "utf8").split("\n").forEach(line => {
        var matched = line.match(reg_const)
        if (matched) {
            let d = span - matched[1].length
            let s = ' '.repeat(d)

            let content = `\t${matched[1]}${s}= ${matched[2]},\t\t-- ${matched[3]}\n`
            fs.writeFileSync(file, content, options)
        }
    })

    fs.writeFileSync(file, "}", options)

    fs.closeSync(file)
}


function export_proto_files() {
    find_proto_files().forEach(file => {
        extract_proto_file(file)
    })

    gen_proto_file_go("../opcode.go")
    gen_proto_file_lua("../../../../../2dgame/simulator/win32/src/message/opcode.lua")

    gen_register_go("../../../server/player/handler_0_init.go")
    gen_register_lua("../../../../../2dgame/simulator/win32/src/message/init.lua")

    gen_handler_go()
    gen_handler_lua()
}


// ----------------------------------------------------------------------------

export_const_file("../../ec/error_code.go", "../../../../../2dgame/simulator/win32/src/error_code.lua")

export_proto_files()