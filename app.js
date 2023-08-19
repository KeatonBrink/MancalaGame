// Proto requires
var parseArgs = require('minimist');
var messages = require('./src/protos/generated/mancala_pb');
var services = require('./src/protos/generated/mancala_grpc_pb');
var grpc = require('@grpc/grpc-js');

// File serving requirments
const http = require('http');
const fs = require('fs');

// Back facing server components
var client
var request
function initBackendCommunication() {
  var argv = parseArgs(process.argv.slice(2), {
    string: 'target'
  });
  var target;
  if (argv.target) {
    target = argv.target;
  } else {
    target = 'localhost:50051';
  }
  client = new services.MancalaServiceClient(target, grpc.credentials.createInsecure());
  // var request = new messages.HandshakeRequest();
  // var user;
  // if (argv._.length > 0) {
  //   user = argv._[0]; 
  // } else {
  //   user = 'world';
  // }
  // request.setUsername(user);
  // client.gameHandshake(request, function(err, response) {
  //   console.log('Greeting:', response.getMessage());
  // });
}
initBackendCommunication();

//In Progress
function findGame(UserName) {
  request = new messages.HandshakeRequest();
  request.setUsername(UserName)
  client.gameHandshake(request, function(err, response) {
    if (err) {
      // TODO: Should probably display error to user
      console.log('An Error: ')
      console.log(err)
    }
    console.log("Got a response, I think: " + JSON.parse.stringify(response))
    return response
  })
  console.log("findGame end")
}


//Front facing server
const server = http.createServer((req, res) => {
  // Static file server
  // Get index.html
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
  // Get client.js
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
  // Get style.css
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

    // Ajax incoming request
    // Send a move request to server
    // TODO:  Finish implementation
  } else if (req.method === 'POST' && req.url === '/MakeMove') {
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
      // Find game to join
  } else if (req.method === 'POST' && req.url === '/FindGame') {
    let body = '';
    req.on('data', (chunk) => {
      body += chunk.toString();
    });
    req.on('end', async () => {
        const data = JSON.parse(body);
        const userName = data.userName;
        if (userName.length > 0) {
          request = new messages.HandshakeRequest();
          request.setUsername(userName)
          var res
          await client.gameHandshake(request, function(err, response) {
            if (err) {
              // TODO: Should probably display error to user
              console.log('An Error: ')
              console.log(err)
            }
            console.log("Got a response, I think: " + JSON.parse.stringify(response))
            res = response
          })
          console.log("tree " + res)

          // TODO: I should probably expand the possible errors
          if (res.getErrorcode != 0) {
            message = 'Error Please Try Again: ' + res.getErrormessage
          } else {
            message = res
          }
        } else {
            message = 'Invalid value: ' + userName + ' Length: ' + userName.length;
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


