const request = require('request')

const handlers = {

    "dx_and" : function(q, sdk, res, ret) {
        const form = {
            token:      q.token,
            pseudo:     q.pseudo,
        }

        request.post(
            {
                url:  "http://115.159.6.66:8100/acct/verify",
                form: form,
                useQuerystring: true,
            },
            function (err, resp, body) {
                if (err) {
                    res.err = err;
                    res.json(ret);
                    console.log(`${q.sdk} request err: ${err}`);
                    return;
                }

                try {
                    body = JSON.parse(body);
                } catch(e) {
                    res.err = e;
                    res.json(ret);
                    console.log(`${q.sdk} JSON.parse err: ${e}, ${body}`);
                    return;
                }

                if (body.ret) {
                    ret.msg = "success";
                    ret.code = 0;
                } else {
                    ret.msg = "failed";
                    ret.code = 10;
                    console.log(`${q.sdk} verify-failed: ${body}`);
                }

                res.json(ret);
            }
        );
    }
}

module.exports = handlers;
