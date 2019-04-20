

let sdks = {};
require('../config/sdk.json').forEach((v)=>{
    sdks[v.name] = v;
});


module.exports = {
    sdks,
}

