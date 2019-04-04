var express = require('express');
var router = express.Router();


module.exports = function(req, res, next) {
    res.json({
        game1: {
            ip: "127.0.0.1",
            port: 9999,
            name: "亢龙有悔",
            status: 0,
        }
    });
};
