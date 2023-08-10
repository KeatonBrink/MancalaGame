// Proto requires
var parseArgs = require('minimist');
var messages = require('./src/protos/generated/mancala_pb');
var services = require('./src/protos/generated/mancala_grpc_pb');
var grpc = require('@grpc/grpc-js');

// File serving requirments
const http = require('http');
const fs = require('fs');

//Front facing server
const server = http.createServer((req, res) => {
    if (req.method === 'GET' && req.url === '/') {
      fs.readFile('src/static/index.html', (err, data) => {
          if (err) {
              res.writeHead(500, { 'Content-Type': 'text/plain' });
              res.end('Internal Server Error');
          } else {
              res.writeHead(200, { 'Content-Type': 'text/html' });
              res.end(data);
          }
      });
    } else if (req.method === 'GET' && req.url === '/client.js') {
      fs.readFile('src/static/client.js', (err, data) => {
        if (err) {
            res.writeHead(500, { 'Content-Type': 'text/plain' });
            res.end('Internal Server Error');
        } else {
            res.writeHead(200, { 'Content-Type': 'text/javascript' });
            res.end(data);
        }
      });
    } else if (req.method === 'GET' && req.url === '/style.css') {
      fs.readFile('src/static/style.css', (err, data) => {
        if (err) {
            res.writeHead(500, { 'Content-Type': 'text/plain' });
            res.end('Internal Server Error');
        } else {
            res.writeHead(200, { 'Content-Type': 'text/css' });
            res.end(data);
        }
      }); 
    } else if (req.method === 'POST' && req.url === '/hello') {
        let body = '';
        req.on('data', (chunk) => {
            body += chunk.toString();
        });
        req.on('end', () => {
            const data = JSON.parse(body);
            const value = data.pitIndex;
            let message = '';
            if (value >= 0 && value < 12) {
                message = 'Hello World ' + value;
            } else {
                message = 'Invalid value';
            }
            res.writeHead(200, { 'Content-Type': 'application/json' });
            res.end(JSON.stringify({ message }));
        });
    } else {
      res.writeHead(404, { 'Content-Type': 'text/plain' });
      res.end('Not Found');
    }
});

const port = 3000;
server.listen(port, () => {
    console.log(`Server listening on port ${port}`);
});

// Back facing server components
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
  var client = new services.MancalaServiceClient(target, grpc.credentials.createInsecure());
  var request = new messages.HandshakeRequest();
  var user;
  if (argv._.length > 0) {
    user = argv._[0]; 
  } else {
    user = 'world';
  }
  request.setUsername(user);
  client.gameHandshake(request, function(err, response) {
    console.log('Greeting:', response.getMessage());
  });
}

main();
