const express = require('express');
const router = express.Router();
const fs = require('fs');


// 客户端上传lua堆栈错误，服务端保存文件
router.post("/", function(req, res) {

    let q = req.body;
    let ret = "failed";

    do {

        if ( !q.uid || !q.data) {
            break;
        }

        let name = q.uid || "name";
        let data = q.data;

        fs.writeFileSync(name, data);

        ret = "successful";
    }
    while(false);

    res.end(ret);
});


module.exports = router;