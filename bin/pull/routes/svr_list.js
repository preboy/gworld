var express = require('express');
var router = express.Router();

const config = require('../../config.json')


// ----------------------------------------------------------------------------

let server_list = {}

// ----------------------------------------------------------------------------

function load_server_list() {
    for (let svr in config.games) {
        let tab = config.games[svr]
        server_list[svr] = {
            ip:     tab.host,
            svr:    svr,
            port:   tab.port,
            name:   svr + "'s name",
        }
    }
}

function init() {
    load_server_list()
}

// ----------------------------------------------------------------------------

router.get("/reload", function(req, res) {
    server_list = {}
    load_server_list()
})


router.get("/", function(req, res) {
    res.json(server_list);
});


// ----------------------------------------------------------------------------

init()
module.exports = router;