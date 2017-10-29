
exports.dispatch = (session, packet) => {
	let opcode = packet.readUInt16LE();
	let message = packet.slice(2);
	console.log("新的消息，opcode=", opcode, "length=", message.byteLength );
	switch (opcode){
	default:
		console.log("未知的协议ID:", opcode);
	}
	// session.send(packet);
}





const net = require('net');
const common = require('./common.js');

exports.make_packet = () => {

	var len = common.getRandomInt(1, 100);
	var pid = common.getRandomInt(1, 4000);
	var buff = Buffer.allocUnsafe(2+2+len);
	
    buff.writeUInt16LE(len, 0);
	buff.writeUInt16LE(pid, 2);

	for(var i = 0; i < len; i++){
		var ch = common.getRandomInt(65, 125);
		buff.writeUInt8(ch, i+4);
	}

	return [buff, pid, len];
}
