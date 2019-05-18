var express = require('express');
var router = express.Router();

const config = require('../../config.json');
const dbmgr = require('../../modules/dbmgr');

router.get("/", function(req, res) {
    let q = req.query;

    if (!q.pseudo) {
        res.json({});
        return
    }

    const cond = {
        acct: q.pseudo,
    }

    let db = dbmgr.get('c').db();

    db.collection('player_info').find(cond).toArray((err, docs) => {
        if (err) {
            res.json({});
        }

        res.json(docs);
    });
});


module.exports = router;