
const net = require("net")

const dispatcher = require("./dispatcher.js");

class Session {

	constructor(c, sid){
		this.socket		 = c;
		this.sid		 = sid;
		this.recv_buffer = null;
        this.sending     = false;
		
        c._session		 = this;

		c.on('end', ()=>{
			console.log("client disconnected", this.game_status);
            this.ending = true;
		});

		c.on('data', (data)=>{
            // console.log("new data:", data.byteLength, data[0], data[1], data[2], data[3]);

			if (this.recv_buffer){
				data = Buffer.concat([this.recv_buffer, data]);
				this.recv_buffer = null;
			}
			if (data.byteLength < 4){
				this.recv_buffer = data;
				return;
			}
			// read header
			let packet_size = data.readUInt16LE(0);
            let opcode = data.readUInt16LE(2);

			if (data.byteLength-4 < packet_size){
				this.recv_buffer = data;
				return;
			}
			let packet = Buffer.allocUnsafe(packet_size);
			let bytes = data.copy(packet, 0, 4, packet_size+4);
			if (bytes != packet_size){
				console.log("拷贝数据失败", bytes, packet_size);
				c.end("拷贝数据失败");
				return;
			}
            
            let left = data.byteLength - 4 - packet_size
            if ( left == 0){
				this.recv_buffer = null;
			}
            else{
                let left_data = Buffer.allocUnsafe(left);
                data.copy(left_data, 0, 4+packet_size, data.byteLength);
                this.recv_buffer = left_data;
            }

			// dispatcher.dispatch(this, packet);
            if (packet.compare(this._packet) == 0){
                // console.log("recv check: pass");
                setTimeout(()=>{
                    this.SendPacket();
                    clearTimeout(this._tid);
                    delete this._tid;
                }, 3000);
            }
            else{
                console.log("数据出错了");
                console.log(packet);
                console.log(this._packet);
            }

		});

		c.on('drain', ()=>{
			this.sending = false;
		});

		c.on('error', (err)=>{
			console.log("socket errored, sid = ", c._session.sid, err);
			c.end();
		});

        // 发始发送数据
        this.SendPacket();
	}


	SendPacket(){
        if (this.ending){
            console.log("Already closed ~!!");
            return;
        }

        if(this.sending){
            console.log("sorry, sending in pending !!!");
            return;
        }
        
        let [packet, pid, len] = dispatcher.make_packet();
		let ret = this.socket.write(packet);
		if (!ret){
            this.sending = true;
			console.log("data queued in user memory!");
		}

        this._packet = Buffer.allocUnsafe(len);
        packet.copy(this._packet, 0, 4, len+4);
        console.log("send: pid, len = ", pid, len);
        console.log("send:", packet);
	}

    Close(){
        if(this._tid){
            clearTimeout(this._tid);
            delete this._tid;
        }
        this.socket.end();
    }

}


exports.Session = Session;
