var express = require('express');
var router = express.Router();


router.get('/', function(req, res) {
    let q = req.query;

    res.json({
        code:   0,
        msg:    'auth ok'
    });

});


module.exports = router;