var express = require('express');
var router = express.Router();


router.get("/", function(req, res) {
    res.json({
        title:      'title',
        content:    'thie is notice content'
    });
});


module.exports = router;
