var express = require('express');
var router = express.Router();

let handlers_path = './auth_handlers.js';

let handlers = require(handlers_path);

const gtab = require('../../modules/gtab');


router.get('/reload', function(req, res) {

    let path = require.resolve(handlers_path);

    delete require.cache[path];

    handlers = require(handlers_path);

    res.end("OK");
}


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
