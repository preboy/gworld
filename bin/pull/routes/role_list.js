var express = require('express');
var router = express.Router();

const config = require('../../config.json')

router.get("/", function(req, res) {
    let ret = {};

    if (!q.pseudo) {
        res.json(ret);
        return
    }

    const cond = {
        acct: q.pseudo,
    }

    db.collection('role_list').find(cond).toArray((err, docs) => {
        if (err) {
            res.json(ret);
        }

        res.json(docs);
    });
});


module.exports = router;