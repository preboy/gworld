var express = require('express');
var router = express.Router();

const config = require('../../config.json')

router.get("/", function(req, res) {
    res.json({
        game1: {
            ip: "127.0.0.1",
            port: 9001,
            name: "飞龙在天",
            status: 0,
        },
        game2: {
            ip: "127.0.0.1",
            port: 9002,
            name: "亢龙有悔",
            status: 0,
        }
    });
});


module.exports = router;