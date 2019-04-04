var express = require('express');
var router = express.Router();


module.exports = function(req, res, next) {
    res.json({
        title:      'title', 
        content:    'thie is notice content'
    });
};
