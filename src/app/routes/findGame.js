var messages = require("../../protos/generated/mancala_pb");

// Make the makeMove function async
async function findGameRoute(req, res, client) {
    try {
        const userName = req.body.userName;
        let errorMessage = "";
        let response = "";
        if (userName.length < 1) {
            errorMessage =
                "Invalid value: " + userName + " Length: " + userName.length;
        } else {
            try {
                response = await findGame(userName, client);
                response = gameHandshakegRPCToJSON(response);
            } catch (error) {
                // Handle unexpected errors
                console.error(error);
            }
        }

        // Send the response with the appropriate message
        console.log("Responding to find game request: " + userName);
        res.status(200).json({ data: response, error: errorMessage });
    } catch (error) {
        console.log(error);
        res.status(500).json({
            data: gameHandshakegRPCToJSON(""),
            error: "An error occurred",
        });
    }
}

// Make the findGame function async
async function findGame(UserName, client) {
    return new Promise((resolve, reject) => {
        request = new messages.HandshakeRequest();
        request.setUsername(UserName);

        client.gameHandshake(request, (err, response) => {
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

function gameHandshakegRPCToJSON(gRPCObject) {
    if (typeof gRPCObject === "object") {
        return {
            errorCode: gRPCObject.getErrorcode(),
            errorMessage: gRPCObject.getErrormessage(),
            message: gRPCObject.getMessage(),
            userHash: gRPCObject.getUserhash(),
            serverWebSocketAddress: gRPCObject.getServerwebsocketaddress(),
        };
    } else {
        // The gRPC encountered an error and could not be returned correctly
        return {
            errorCode: 2,
            errorMessage: "",
            message: "",
            userHash: "",
            serverWebSocketAddress: "",
        };
    }
}

// Export the findGameRoute function
module.exports = findGameRoute;
