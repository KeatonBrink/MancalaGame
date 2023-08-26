// Proto requires
var parseArgs = require("minimist");
var messages = require("./src/protos/generated/mancala_pb");
var services = require("./src/protos/generated/mancala_grpc_pb");
var grpc = require("@grpc/grpc-js");

// File serving requirments
const express = require("express");
const bodyParser = require("body-parser");

// Bring in post routes
const findGameRoute = require("./src/app/routes/findGame");
const makeMoveRoute = require("./src/app/routes/makeMove");

// Create the express app
const app = express();

// Middleware to parse JSON body
app.use(bodyParser.json());

// Serve static files from the 'src/static' directory
app.use(express.static("src/static"));

// Back facing server components
var client;
var request;
function initBackendCommunication() {
    var argv = parseArgs(process.argv.slice(2), {
        string: "target",
    });
    var target;
    if (argv.target) {
        target = argv.target;
    } else {
        target = "localhost:50051";
    }
    client = new services.MancalaServiceClient(
        target,
        grpc.credentials.createInsecure()
    );
}
initBackendCommunication();

// Refactored route for finding a game
app.post("/FindGame", async (req, res) => {
    await findGameRoute(req, res, client);
});

// Refactored route for making a move
app.post("/MakeMove", async (req, res) => {
    await makeMoveRoute(req, res, client);
});

// Start the server
const port = process.env.PORT || 3000;
app.listen(port, () => {
    console.log(`Server listening on port ${port}`);
});
