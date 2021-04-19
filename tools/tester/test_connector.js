const net = require('net');
const ss = require('./session')

const connections = 9;

let sessions = [];
let sid = 1;
let _creator_tid;

_creator_tid = setInterval(() => {

    let c = net.createConnection(9002, "115.159.6.66", () => {
        sessions[c.sid] = new ss.Session(c);
    });

    c.sid = sid;

    c.on('error', (err) => {
        console.log("connection failed: sid = ", c.sid, err);
    });

    if (sid >= connections) {
        clearInterval(_creator_tid);
        _creator_tid = null;
        console.log("Total", sid, "connection created !");
    } else {
        sid++;
    }

}, 50);