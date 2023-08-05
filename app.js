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

let board = [
    [0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0]
]

const basicBoard = [
    [0, 4, 4, 4, 4, 4, 4],
    [4, 4, 4, 4, 4, 4, 0]
]

let currentPlayer = 1

console.log("hello world")
const pits = document.querySelectorAll('.mancala-pit')

pits.forEach((pit, index) => {
    pit.addEventListener('click', () => {
        console.log('Pit: ', index)
        // Send a gRPC request to your Go back-end with the clicked pit index
        // grpcClient.makeMove({ pitIndex: index }, (error, response) => {
        //     if (error) {
        //         console.error('Error making move:', error);
        //     } else {
        //         console.log('Move successful:', response.message);
        //         // Update the UI based on the response from the back-end
        //     }
        // });
    });
});