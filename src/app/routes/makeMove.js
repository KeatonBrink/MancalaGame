var messages = require("../../protos/generated/mancala_pb");
var services = require("../../protos/generated/mancala_grpc_pb");

async function makeMoveRoute(req, res, client) {
    try {
        const data = req.body;
        //Prepare the request to the gRPC server
        response = makeMove(data.userHash, data.move);
        response = makeMovegRPCToJSON(response);
        res.status(200).json({ message: "Move processed successfully" });
    } catch (error) {
        res.status(500).json({ error: "An error occurred" });
    }
}

async function makeMove(UserHash, Move) {
    return new Promise((resolve, reject) => {
        request = new messages.MoveRequest();
        request.setUserhash(UserHash);
        request.setMove(Move);

        client.makeMove(request, (err, response) => {
            if (err) {
                console.log("An Error: ");
                console.log(err);
                reject(err); // Reject the promise with the error
            } else {
                resolve(response); // Resolve the promise with the response
            }
        });
    });
}

function makeMovegRPCToJSON(gRPCObject) {
    if (typeof gRPCObject === "object") {
        return {
            errorCode: gRPCObject.getErrorcode(),
            errorMessage: gRPCObject.getErrormessage(),
            message: gRPCObject.getMessage(),
            board: gRPCObject.getBoard(),
        };
    } else {
        // The gRPC encountered an error and could not be returned correctly
        return {
            errorCode: 2,
            errorMessage: "",
            message: "",
            board: "",
        };
    }
}

module.exports = makeMoveRoute;
