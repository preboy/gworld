const net = require('net');
const common = require('./common.js');


exports.dispatch = (session, packet) => {
    let opcode = packet.readUInt16LE();
    let message = packet.slice(2);
    console.log("新的消息，opcode=", opcode, "length=", message.byteLength);
    switch (opcode) {
        default: console.log("未知的协议ID:", opcode);
    }
}


exports.make_packet = () => {
    var op  = common.getRandomInt(4000, 4100);
    var data_len = common.getRandomInt(10, 20);

    var buff = Buffer.allocUnsafe(4 + data_len);

    buff.writeUInt16LE(data_len, 0);
    buff.writeUInt16LE(op, 2);

    for (var i = 0; i < data_len; i++) {
        let ch = common.getRandomInt(65, 125);
        buff.writeUInt8(ch, i+4);
    }

    return [buff, op, data_len];
}
