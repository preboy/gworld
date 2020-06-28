const net = require("net")
const dispatcher = require("./dispatcher.js");

let session_cnt = 0;

class Session {

    constructor(c) {
        this.socket = c;
        this.recv_buffer = null;
        this.sending = false;

        c._session = this;

        session_cnt++;
        // console.log("NEW SESSION:", c.sid, session_cnt);

        this._tick = setInterval(()=>{
            if (this.ending) {
                clearInterval(this._tick);
                this._tick = null;
                return;
            }

            let now = +new Date() / 1000;

            if (now - this._send_ts > 5) {
                console.log(c.sid, `${now - this._send_ts} 秒未收到返回数据`);
            }

        }, 1000);

        c.on('end', () => {
            this.ending = true;
            session_cnt--;
            console.log("client disconnected ! SESSION COUNT:", session_cnt);
        });

        c.on('data', (data) => {
            // console.log("new data:", data.byteLength, data);

            if (this.recv_buffer) {
                data = Buffer.concat([this.recv_buffer, data]);
                this.recv_buffer = null;
            }
            if (data.byteLength < 4) {
                this.recv_buffer = data;
                return;
            }

            // read header
            let packet_size = data.readUInt16LE(0);
            if (packet_size > 0x1000) {
                console.log("超大的包", packet_size);
                c.end("超大的包11");
                return;
            }

            if (data.byteLength - 4 < packet_size) {
                this.recv_buffer = data;
                return;
            }

            let packet = Buffer.allocUnsafe(packet_size+4);
            let bytes = data.copy(packet, 0, 0, 4+packet_size);
            if (bytes != packet_size+4) {
                console.log("拷贝数据失败", bytes, packet);
                c.end("拷贝数据失败");
                return;
            }

            this.recv_buffer = null;
            if (data.byteLength > packet_size + 4) {
                this.recv_buffer = Buffer.allocUnsafe(data.byteLength - (packet_size + 4));
                let bytes = data.copy(this.recv_buffer, 0, packet_size + 4, data.byteLength);
                if (bytes != data.byteLength - (packet_size + 4)) {
                    console.log("拷贝数据失败2", bytes, data.byteLength, packet_size);
                    c.end("拷贝数据失败2");
                    return;
                }
            }

            // dispatcher.dispatch(this, packet);

            if (this._packet != null && packet.compare(this._packet) == 0) {
                this._packet = null;

                this._recv_ts = +new Date() / 1000;
                // console.log(c.sid, this._recv_ts - this._send_ts);
                // console.log("收到包", packet.byteLength, packet);

                this._tid = setTimeout(() => {
                    this.SendPacket();
                    this._tid = null;
                }, 3000);
            } else {
                console.log("收到了与发送的不一致的数据包");
                console.log("recv:", packet);
                console.log("send:", this._packet);
            }
        });

        c.on('drain', () => {
            this.sending = false;
        });

        c.removeAllListeners('error');

        c.on('error', (err) => {
            console.log("socket error, sid = ", c.sid, err);
            c.end();
        });

        // 发始发送数据
        this.SendPacket();
    }

    SendPacket() {
        if (this.ending) {
            console.log("Already closed ~!!");
            return;
        }

        if (this.sending) {
            console.log("sorry, sending in pending !!!");
            return;
        }

        let [packet, op, data_len] = dispatcher.make_packet();
        let ret = this.socket.write(packet);
        if (!ret) {
            this.sending = true;
            console.log("data queued in user memory!");
        }

        this._packet = Buffer.allocUnsafe(data_len + 4);
        packet.copy(this._packet, 0, 0, 4 + data_len);

        this._send_ts = +new Date() / 1000;
        // console.log("send: size = ", data_len + 8);
    }

    Close() {
        if (this._tid) {
            clearTimeout(this._tid);
            this._tid = null;
        }
        this.socket.end();
    }

}


exports.Session = Session;