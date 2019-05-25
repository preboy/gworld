
let sdks = {};

require('../sdk.json').forEach((v)=>{
    sdks[v.name] = v;
});


module.exports = {
    sdks,
}

