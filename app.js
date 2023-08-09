// var PROTO_PATH = 'src/protos/route_guide.proto';
// var grpc = require('@grpc/grpc-js');
// var protoLoader = require('@grpc/proto-loader');
// // Suggested options for similarity to existing grpc.load behavior
// var packageDefinition = protoLoader.loadSync(
//     PROTO_PATH,
//     {keepCase: true,
//      longs: String,
//      enums: String,
//      defaults: true,
//      oneofs: true
//     });
// var protoDescriptor = grpc.loadPackageDefinition(packageDefinition);
// // The protoDescriptor object has the full package hierarchy
// var routeguide = protoDescriptor.routeguide;

// let board = [
//     [0, 0, 0, 0, 0, 0, 0],
//     [0, 0, 0, 0, 0, 0, 0]
// ]

// const basicBoard = [
//     [0, 4, 4, 4, 4, 4, 4],
//     [4, 4, 4, 4, 4, 4, 0]
// ]

// let currentPlayer = 1

// console.log("hello world")
// const pits = document.querySelectorAll('.mancala-pit')

// pits.forEach((pit, index) => {
//     pit.addEventListener('click', () => {
//         console.log('Pit: ', index)
//         // Send a gRPC request to your Go back-end with the clicked pit index
//         // grpcClient.makeMove({ pitIndex: index }, (error, response) => {
//         //     if (error) {
//         //         console.error('Error making move:', error);
//         //     } else {
//         //         console.log('Move successful:', response.message);
//         //         // Update the UI based on the response from the back-end
//         //     }
//         // });
//     });
// });
// import * as mancala from './src/protos/mancala_pb.js';
// import * as MancalaServiceClient from './src/protos/mancala_grpc_web_pb.js';

// // Create a gRPC-Web client
// const client = new MancalaServiceClient.MancalaServiceClient('http://localhost:50051');

// // Prepare the request
// var request = new mancala.HandshakeRequest();
// request.setUserName('world');

// // Call the GameHandshake method using gRPC-Web
// client.gameHandshake(request, {}, (err, response) => {
//   if (err) {
//     console.error('Error:', err);
//     return;
//   }
//   console.log('Greeting:', response.getMessage());
// });

import { MancalaServiceClient } from './src/protos/mancala_grpc_web_pb.js';

const client = new MancalaServiceClient('http://localhost:8080');

  const request = new HelloRequest();
  request.setName('World');

  client.sayHello(request, {}, (err, response) => {
    if (err) {
      console.error('gRPC-Web Error:', err);
    } else {
      console.log('gRPC-Web Response:', response.getMessage());
    }
  });