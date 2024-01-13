// Require the port number from app.js
const PORT = require("../../../app").port;
const e = require("express");
const WebSocket = require("ws");

// Update Move.js handles websockets that update clients when a successful move is made

// Grpc response
// message GameStatusResponse {
//     int32 errorCode = 1;
//     string errorMessage = 2;
//     string message = 3;
//     string gameID = 4;
//     string board = 5;
//     string player1Hash = 6;
//     string player2Hash = 7;
//     string playerTurn = 8;
// }

activeGames = {};

function GetGameData(gameData) {
    return {
        gameID: gameData.gameID,
        board: gameData.board,
        player1Hash: gameData.player1Hash,
        player2Hash: gameData.player2Hash,
        playerTurn: gameData.playerTurn,
    };
}

// Recieve websocket connection from client
const wss = new WebSocket.Server({ port: PORT });
// Server-side code
const activeConnections = new Map(); // Map to store WebSocket connections

wss.on("connection", (ws, req) => {
    // Extract the username from the query parameters
    const userhash = new URL(
        req.url,
        `http://${req.headers.host}`
    ).searchParams.get("userhash");

    // Generate a unique identifier based on the username (you can combine it with other factors if needed)
    const connectionId = userhash;

    // Store the WebSocket connection with its identifier
    activeConnections.set(connectionId, ws);
});

function addActiveGame(gameData) {
    if (gameData.errorCode != 0) {
        console.log(
            "Error " +
                gameData.errorCode +
                " adding active game: " +
                gameData.errorMessage
        );
        return;
    }
    relaventGameData = GetGameData(gameData);
    activeGames[gameData.gameID] = relaventGameData;
    console.log("Added active game: " + gameData.gameID);
}

function removeActiveGame(gameID) {
    delete activeGames[gameID];
    console.log("Removed active game: " + gameID);
    // TODO: Send a message to the clients that the game has ended and close websockets
}

// Need to finish this function
function updatePlayerTurn(gameState, playerHash) {
    if (activeGames[gameID].player1Hash == playerHash) {
        activeGames[gameID].playerTurn = activeGames[gameID].player2Hash;
    } else {
        activeGames[gameID].playerTurn = activeGames[gameID].player1Hash;
    }
}

// Export functions
module.exports = {
    addActiveGame: addActiveGame,
    removeActiveGame: removeActiveGame,
    updatePlayerTurn: updatePlayerTurn,
};
