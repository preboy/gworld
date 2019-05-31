var express = require('express');
var router = express.Router();

const config = require('../../config.json')
const dbmgr = require('../../modules/dbmgr');


// ----------------------------------------------------------------------------

let svr_list = {}

// ----------------------------------------------------------------------------

function load_server_list() {
    for (let svr in config.games) {
        let tab = config.games[svr]
        svr_list[svr] = {
            ip:     tab.host,
            svr:    svr,
            port:   tab.port,
            name:   svr,
            stat:   0,      // 0:维护  1:正常  2:火热  3:新服
        }
    }
}


function load_server_stat() {

    let db = dbmgr.get('c').db();

    db.collection('server_info').find({}).toArray((err, docs) => {

        if (err) {
            return;
        }

        for (let i = 0; i < docs.length; i++) {

            let doc = docs[i];
            let svr = doc._id;

            if (svr_list[svr]) {
                svr_list[svr].name = doc.name;
                svr_list[svr].stat = doc.stat;
            }
        }
    });
}


function init() {
    setTimeout(()=>{
        load_server_list();
        load_server_stat();
    }, 2000)
}

// ----------------------------------------------------------------------------

router.get("/reload", function(req, res) {

    svr_list = {};

    load_server_list();
    load_server_stat();

    res.end("OK");
})


router.get("/", function(req, res) {
    res.json(svr_list);
});


// ----------------------------------------------------------------------------

init()

module.exports = router;

