var createError = require('http-errors');
var express = require('express');
var logger = require('morgan');

var auth = require('./routes/auth');

var app = express();

app.use(logger('dev'));
app.use(express.json());
app.use(express.urlencoded({ extended: false }));

app.use(function(req, res, next){
    console.log([req._startTime.toLocaleString(), req.socket.remoteAddress||'no ip', req.url, req.body, req.query]);
    next();
});

app.use('/api', auth);

// catch 404 and forward to error handler
app.use(function(req, res, next) {
  next(createError(404));
});

// error handler
app.use(function(err, req, res, next) {
  // set locals, only providing error in development
  res.locals.message = err.message;
  res.locals.error = req.app.get('env') === 'development' ? err : {};

  // render the error page
  res.status(err.status || 500);
  res.end('Not Found Page');
});

// ----------------------------------------------------------------------------
// process events

process.on('uncaughtException', (err) => {
    console.log('uncaughtException', err);
});


// ----------------------------------------------------------------------------
module.exports = app;
