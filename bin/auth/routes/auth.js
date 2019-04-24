var express = require('express');
var router = express.Router();

const handlers = require('./auth_handlers.js')

const gtab = require('../../modules/gtab');


router.get('/', function(req, res) {
    let q = req.query;

    let sdk = gtab.sdks[q.sdk];
    let ret =
    {
        msg:  "",
        code: -1,
    }

    if (!q.sdk || !sdk) {
        ret.msg = `Not Found SDK: ${q.sdk}`;
        ret.code = 1;
        res.json(ret);
        console.log(ret.msg);
        return;
    }

    let h = handlers[q.sdk];
    if (!h) {
        ret.msg = `${q.sdk} Not Found HANDLER`;
        ret.code = 2;
        res.json(ret);
        console.log(ret.msg);
        return;
    }

    h(q, sdk, res, ret);
});


module.exports = router;
