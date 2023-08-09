
var PROTO_PATH = __dirname + '/src/protos/mancala.proto';

var parseArgs = require('minimist');
var grpc = require('@grpc/grpc-js');


function main() {
  var argv = parseArgs(process.argv.slice(2), {
    string: 'target'
  });
  var target;
  if (argv.target) {
    target = argv.target;
  } else {
    target = 'localhost:50051';
  }
  var client = new mancala_proto.MancalaService(target, grpc.credentials.createInsecure());
  var user;
  if (argv._.length > 0) {
    user = argv._[0];
  } else {
    user = 'world';
  }
  client.GameHandshake({userName: user}, function(err, response) {
    console.log(err)
    console.log('Greeting:', response.message);
  });
}

main();
