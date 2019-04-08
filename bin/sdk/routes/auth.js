var express = require('express');
var router = express.Router();

const TokenGenerator = require('tokgen');
const generator = new TokenGenerator();

const dbmgr = require('../modules/dbmgr');

let tokens = {} // uid -> [token, time]

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

function tran_name(name) {
    // todo
    return name
}

router.post('/register', function(req, res) {
    let q = req.body;

    let name;
    let passwd;
    let pass = false;

    do {
        if (!q.name || !q.passwd) {
            break
        }

        name = q.name.toLowerCase();
        passwd = q.passwd;

        if(!verify_ansi(name)) {
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
        _id:    name,
        passwd: passwd,
    };

    db.collection('account').insertOne(doc, (err, r) => {
        let msg = err == null ? "ok" : "error";
        res.send(msg);
    })
});

router.post('/login', function(req, res) {
    let q = req.body;

    let name;
    let passwd;

    let pass = false;

    do {
        if (!q.name || !q.passwd) {
            break
        }

        name = q.name.toLowerCase();
        passwd = q.passwd;

        if(!verify_ansi(name)) {
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

    var cond = {
        _id:    name,
        passwd: passwd,
    };

    db.collection('account').findOne(cond, {}, (err, r) => {
        let ret = {
            msg: "error",
        };

        if (r != null) {
            let key = tran_name(r._id);
            let val = generator.generate();
            tokens[key] = [val, (new Date()).valueOf()];

            ret.msg = "ok";
            ret.uid = key;
            ret.token = val;
        }

        res.json(ret);
    });
});

router.post('/verify', function(req, res) {
    let q = req.body;

    let pass = false;

    do {
        if (!q.uid || !q.token) {
            break;
        }

        let t = tokens[q.uid];
        if (!t || t[0] != q.token) {
            break;
        }

        let now = (new Date()).valueOf()
        if (now -t[1] >= 60*1000) {
            break;
        }

        pass = true;

        tokens[q.uid] = null;
    }
    while(false);

    res.json({
        ret : pass,
    });
});


module.exports = router
