const express = require('express');
const router = express.Router();
const fs = require('fs');


module.exports = function(req, res, next) {
    let q = req.body;

    // save to file
    // TODO

    res.json({
        code:   0, 
        msg:    'report ok'
    });
};
