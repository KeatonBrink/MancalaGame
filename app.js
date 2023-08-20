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
    let errorMessage = ''
    let response = ''
    if (userName.length < 1) {
      errorMessage = 'Invalid value: ' + userName + ' Length: ' + userName.length;
    } else {
      try {
        response = await findGame(userName);
        console.log("Find game response here")
        response = gameHandshakegRPCToJSON(response)
        console.log(JSON.stringify(response))

      } catch (error) {
        // Handle unexpected errors
        console.error(error);
      }
    }

    // Send the response with the appropriate message
    res.status(200).json({ data: response, error: errorMessage });
  } catch (error) {
    console.log(error)
    res.status(500).json({ data: gameHandshakegRPCToJSON(''), error: 'An error occurred' });
  }
});

function gameHandshakegRPCToJSON(gRPCObject) {
  if (typeof gRPCObject === "object"){
    return {
      'errorCode' : gRPCObject.getErrorcode(),
      'errorMessage' : gRPCObject.getErrormessage(),
      'message': gRPCObject.getMessage(),
      'serverWebSocketAddress' : gRPCObject.getServerwebsocketaddress()
    }
  } else {
    // The gRPC encountered an error and could not be returned correctly 
    return {
      'errorCode' : 2,
      'errorMessage' : '',
      'message': '',
      'serverWebSocketAddress' : ''
    }
  }
}

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