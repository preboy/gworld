const express = require('express');
const router = express.Router();
const fs = require('fs');


// 客户端上传lua堆栈错误，服务端保存文件
router.post("/", function(req, res) {
    let q = req.body;

    if (q.key && q.msg) {
        let key = `./runtime/stack/${q.key}_${+new Date()}`
        fs.writeFile(key, q.msg, (err)=>{});
    }

    res.end("OK");
});


module.exports = router;
