var express = require('express');
var router = express.Router();

const TokenGenerator = require('tokgen');
const generator = new TokenGenerator();

const dbmgr = require('../../modules/dbmgr');

let tokens = {} // pseudo -> [token, time]

function verify_ansi(str) {
    if(str.length == 0 || str.length > 16) {
        return false;
    }

    for (var i = 0; i < str.length; i++) {
        let char = str.charCodeAt(i)

        if (char >= 48 && char <= 57) {
            continue;
        }

        if (char >= 65 && char <= 90) {
            continue;
        }

        if (char >= 97 && char <= 122) {
            continue;
        }

        return false;
    }

    return true
}

function tran_acct(acct) {
    return `dx_${acct}`
}

router.post('/register', function(req, res) {
    let q = req.body;

    let acct;
    let passwd;

    let pass = false;

    do {
        if (!q.acct || !q.passwd) {
            break
        }

        acct = q.acct.toLowerCase();
        passwd = q.passwd;

        if(!verify_ansi(acct)) {
            break
        }

        if (!verify_ansi(passwd)) {
            break
        }

        pass = true;
    }
    while(false);

    if (!pass) {
        res.send("error");
        return;
    }

    let db = dbmgr.get('sdk').db();

    var doc = {
        _id:    acct,
        passwd: passwd,
    };

    db.collection('account').insertOne(doc, (err, r) => {
        let msg = err == null ? "ok" : "error";
        res.send(msg);
    })
});

router.post('/login', function(req, res) {
    let q = req.body;

    let acct;
    let passwd;

    let pass = false;

    do {
        if (!q.acct || !q.passwd) {
            break
        }

        acct = q.acct.toLowerCase();
        passwd = q.passwd;

        if(!verify_ansi(acct)) {
            break
        }

        if (!verify_ansi(passwd)) {
            break
        }

        pass = true;
    }
    while(false);

    const ret = {
        ret: false,
        msg: "error",
    }

    if (!pass) {
        res.json(ret);
        return;
    }

    let db = dbmgr.get('sdk').db();

    var cond = {
        _id:    acct,
        passwd: passwd,
    };

    db.collection('account').findOne(cond, {}, (err, r) => {
        if (r != null) {
            let key = tran_acct(r._id);
            let val = generator.generate();
            tokens[key] = [val, (new Date()).valueOf()];

            ret.ret = true;
            ret.msg = "ok";
            ret.token = val;
            ret.pseudo = key;
        }

        res.json(ret);
    });
});

router.post('/verify', function(req, res) {
    let q = req.body;

    let pass = false;

    do {
        if (!q.pseudo || !q.token) {
            break;
        }

        let t = tokens[q.pseudo];
        if (!t || t[0] != q.token) {
            break;
        }

        let now = (new Date()).valueOf()
        if (now -t[1] >= 60*1000) {
            break;
        }

        pass = true;

        tokens[q.pseudo] = null;
    }
    while(false);

    res.json({
        ret : pass,
    });
});


module.exports = router
