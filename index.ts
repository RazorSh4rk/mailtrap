const { SMTPServer } = require("smtp-server");

const options = {
    authOptional: true,
    allowInsecureAuth: true,
    disableReverseLookup: true,
    logger: true,

    onData: function(stream, session, callback){
        stream.pipe(process.stdout);
        stream.on('end', callback);
    },
    onRcptTo: function(address, session, callback){
        console.log("RcptTo: "+address)
        return callback();
    } ,
    onMailFrom: function(address, session, callback){
        console.log("MailFrom: "+address)
        return callback();
    },
    onAuth: function(auth, session, callback){
        console.log("Auth: "+auth.method+' by '+auth.username)
        callback(null, {})
    }
}

const server = new SMTPServer({});

server.listen(25);