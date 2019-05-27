
var mongodb = require('mongodb');
var async   = require('async');

// ============================================================================

var pool = {};

// ============================================================================

/*
    pools: [
        [key, cnnstr, size],
        ...
    ]
*/
function init(pools) {
    async.each(
        pools,
        (p, cb) => {
            var key    = p[0];
            var cnnstr = p[1];
            var size   = p[2];

            mongodb.connect(
                cnnstr,
                {
                    poolSize:        size,
                    reconnectTries:  Number.MAX_SAFE_INTEGER,
                    socketTimeoutMS: 0,
                    useNewUrlParser: true,
                },
                (err, db) => {
                    if (!err) {
                        pool[key] = db;
                    }

                    cb(err);
                }
            );
        },
        (err) => {
            if (err) {
                console.error("init db pool failed:", err);
                process.exit(1);
            }

            console.log("db pool init OK");
        }
    );
}

function get(key) {
    return pool[key];
}

// ============================================================================

module.exports = {
    init,
    get,
};
