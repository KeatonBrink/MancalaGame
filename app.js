// Proto requires
var parseArgs = require('minimist');
var messages = require('./src/protos/generated/mancala_pb');
var services = require('./src/protos/generated/mancala_grpc_pb');
var grpc = require('@grpc/grpc-js');

// File serving requirments
const express = require('express')
const bodyParser = require('body-parser')

const app = express()

// Middleware to parse JSON body
app.use(bodyParser.json())

// Serve static files from the 'src/static' directory
app.use(express.static('src/static'))

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
}
initBackendCommunication();


// Make the findGame function async
async function findGame(UserName) {
  return new Promise((resolve, reject) => {
    request = new messages.HandshakeRequest();
    request.setUsername(UserName);

    client.gameHandshake(request, (err, response) => {
      if (err) {
        console.log('An Error: ');
        console.log(err);
        reject(err); // Reject the promise with the error
      } else {
        resolve(response); // Resolve the promise with the response
      }
    });
  });
}
// Refactored route for finding a game
app.post('/FindGame', async (req, res) => {
  try {
    const userName = req.body.userName;
    let responseMessage = ''; // Initialize an empty response message

    if (userName.length < 1) {
      responseMessage = 'Invalid value: ' + userName + ' Length: ' + userName.length;
    } else {
      try {
        const response = await findGame(userName);

        if (response.getErrorcode() === 1) {
          const errorMessage = response.getErrormessage();
          responseMessage = `Error: ${errorMessage}`;
        } else {
          // Handle successful response
          responseMessage = response.getMessage();
        }
      } catch (error) {
        // Handle unexpected errors
        responseMessage = 'An error occurred';
        console.error(error);
      }
    }

    // Send the response with the appropriate message
    res.status(200).json({ message: responseMessage });
  } catch (error) {
    res.status(500).json({ error: 'An error occurred' });
  }
});

// Refactored route for making a move
app.post('/MakeMove', (req, res) => {
  try {
    const data = req.body;
    // Your move processing logic...

    res.status(200).json({ message: 'Move processed successfully' });
  } catch (error) {
    res.status(500).json({ error: 'An error occurred' });
  }
});

// Start the server
const port = process.env.PORT || 3000;
app.listen(port, () => {
  console.log(`Server listening on port ${port}`);
});






// //Front facing server
// const server = http.createServer((req, res) => {
//   // Static file server
//   // Get index.html
//   if (req.method === 'GET' && req.url === '/') {
//     fs.readFile('src/static/index.html', (err, data) => {
//         if (err) {
//             res.writeHead(500, { 'Content-Type': 'text/plain' });
//             res.end('Internal Server Error');
//         } else {
//             res.writeHead(200, { 'Content-Type': 'text/html' });
//             res.end(data);
//         }
//     });
//   // Get client.js
//   } else if (req.method === 'GET' && req.url === '/client.js') {
//     fs.readFile('src/static/client.js', (err, data) => {
//       if (err) {
//           res.writeHead(500, { 'Content-Type': 'text/plain' });
//           res.end('Internal Server Error');
//       } else {
//           res.writeHead(200, { 'Content-Type': 'text/javascript' });
//           res.end(data);
//       }
//     });
//   // Get style.css
//   } else if (req.method === 'GET' && req.url === '/style.css') {
//     fs.readFile('src/static/style.css', (err, data) => {
//       if (err) {
//           res.writeHead(500, { 'Content-Type': 'text/plain' });
//           res.end('Internal Server Error');
//       } else {
//           res.writeHead(200, { 'Content-Type': 'text/css' });
//           res.end(data);
//       }
//     }); 

//     // Ajax incoming request
//     // Send a move request to server
//     // TODO:  Finish implementation
//   } else if (req.method === 'POST' && req.url === '/MakeMove') {
//       let body = '';
//       req.on('data', (chunk) => {
//           body += chunk.toString();
//       });
//       req.on('end', () => {
//           const data = JSON.parse(body);
//           const value = data.pitIndex;
//           let message = '';
//           if (value >= 0 && value < 12) {
//               message = 'Hello World ' + value;
//           } else {
//               message = 'Invalid value';
//           }
//           res.writeHead(200, { 'Content-Type': 'application/json' });
//           res.end(JSON.stringify({ message }));
//       });
//       // Find game to join
//   } else if (req.method === 'POST' && req.url === '/FindGame') {
//     let body = '';
//     req.on('data', (chunk) => {
//       body += chunk.toString();
//     });
//     req.on('end', async () => {
//         const data = JSON.parse(body);
//         const userName = data.userName;
//         if (userName.length > 0) {
//           request = new messages.HandshakeRequest();
//           request.setUsername(userName)
//           var res
//           await client.gameHandshake(request, function(err, response) {
//             if (err) {
//               // TODO: Should probably display error to user
//               console.log('An Error: ')
//               console.log(err)
//             }
//             console.log("Got a response, I think: " + JSON.parse.stringify(response))
//             res = response
//           })
//           console.log("tree " + res)

//           // TODO: I should probably expand the possible errors
//           if (res.getErrorcode != 0) {
//             message = 'Error Please Try Again: ' + res.getErrormessage
//           } else {
//             message = res
//           }
//         } else {
//             message = 'Invalid value: ' + userName + ' Length: ' + userName.length;
//         }
//         res.writeHead(200, { 'Content-Type': 'application/json' });
//         res.end(JSON.stringify({ message }));
//     });
//     } else {
//       res.writeHead(404, { 'Content-Type': 'text/plain' });
//       res.end('Not Found');
//     }
// });