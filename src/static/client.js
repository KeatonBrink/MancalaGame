let board = [
    [0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0],
];

const basicBoard = [
    [0, 4, 4, 4, 4, 4, 4],
    [4, 4, 4, 4, 4, 4, 0],
];

let currentState = "Awaiting Name";

let playerID = 1;

let userHash = "";

const statusText = document.querySelector("#status-text");
const pits = document.querySelectorAll(".mancala-pit");

pits.forEach((pit, index) => {
    pit.addEventListener("click", () => {
        if (currentState != "Your Turn") {
            return;
        } else if (board[Math.floor(index / 7)][index % 7] == 0) {
            return;
        } else if (index < 6) {
            return;
        }
        console.log("Pit: ", index - 6);
        const xhr = new XMLHttpRequest();
        //
        xhr.open("POST", "/MakeMove", true);
        xhr.setRequestHeader("Content-Type", "application/json");
        xhr.onreadystatechange = function () {
            if (xhr.readyState === 4 && xhr.status === 200) {
                const response = JSON.parse(xhr.responseText);
                console.log(JSON.stringify(response));
                statusText.textContent = response.message;
                if (response.data.message == "Game Over") {
                    statusText.textContent = "Game Over";
                    currentState = "Game Over";
                    board = stringBoardTo2DArray(response.data.board);
                    updateHTMLBoard();
                } else if (response.data.message == "Opponent's Turn") {
                    statusText.textContent = "Opponent Turn";
                    currentState = "Opponent Turn";
                    board = stringBoardTo2DArray(response.data.board);
                    updateHTMLBoard();
                }
            }
        };
        console.log("client.js /Makemove User Hash: " + userHash);
        const payload = {
            pitIndex: index - 6,
            userHash: userHash,
        };
        xhr.send(JSON.stringify(payload));
    });
});

function stringBoardTo2DArray(stringBoard) {
    for (let i = 0; i < 2; i++) {
        stringBoard = stringBoard.replace("[", "");
        stringBoard = stringBoard.replace("]", "");
    }
    let tempBoard = stringBoard.split(";");
    tempBoard = tempBoard.reverse();
    tempBoard.forEach((row, index) => {
        tempBoard[index] = row.split(" ");
    });
    tempBoard[0] = tempBoard[0].reverse();
    return tempBoard;
}

const userNameSubmission = document.querySelector("#user-name-submission");
const userNameLocation = document.getElementById("user-name");

userNameSubmission.addEventListener("click", () => {
    if (currentState != "Awaiting Name") {
        return;
    }
    var userName = userNameLocation.value;
    console.log("Username: " + userName);

    if (userName.length == 0) {
        statusText.textContent =
            "Username must be at least one character long!";
        return;
    }
    statusText.textContent = "Attempting to join game";
    const xhr = new XMLHttpRequest();
    xhr.open("POST", "/FindGame", true);
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            const response = JSON.parse(xhr.responseText);
            console.log(JSON.stringify(response));
            statusText.textContent = response.message;
            if (response.data.message == "Waiting for Second Player") {
                statusText.textContent = "Waiting for second player";
                currentState = "Waiting for Second Player";
                userHash = response.data.userHash;
                webSocketToServer(response.data.serverWebSocketAddress);
            } else if (response.data.message == "Game Starting Soon") {
                statusText.textContent = "Opponent found, their turn";
                currentState = "Opponent Turn";
                board = basicBoard;
                updateHTMLBoard();
                playerID = 2;
                userHash = response.data.userHash;
            } else if (response.data.message == "Name already in use") {
                statusText.textContent = "Name already in use, try again";
                currentState = "Awaiting Name";
            }
        } else {
            statusText.textContent = "Error contacting server, try again";
        }
    };
    xhr.send(JSON.stringify({ userName: userName }));
});

function webSocketToServer(webSocketAddress) {
    let url = "ws://localhost:" + webSocketAddress + "/ws";
    let socket = new WebSocket(url);

    socket.onopen = () => {
        console.log("Connected to WebSocket");
        socket.send(userHash);
    };

    socket.onmessage = (event) => {
        console.log("Received: " + event.data);
        if (event.data == "Opponent Found") {
            statusText.textContent = "Opponent Found, Your Turn";
            currentState = "Your Turn";
            board = basicBoard;
            updateHTMLBoard();
            playerID = 1;
        } else {
            statusText.textContent = "Sever Error";
            currentState = "Awaiting Name";
        }
    };

    socket.onclose = (event) => {
        if (event.wasClean) {
            console.log(
                `Closed cleanly, code=${event.code}, reason=${event.reason}`
            );
        } else {
            console.log("Connection died");
        }
    };
}

function turnUpdateWebSocket(webSocketAddress) {
    let url = "ws://localhost:" + webSocketAddress + "/ws?userhash=";
    let socket = new WebSocket(url + userHash);

    socket.onopen = () => {
        console.log("Connected to Turn WebSocket");
    };

    socket.onmessage = (event) => {
        console.log("Received: " + event.data);
        if (currentState == "Your Turn") {
            board = stringBoardTo2DArray(event.data);
            updateHTMLBoard();
        } else if (event.data == "Your Turn") {
            statusText.textContent = "Your Turn";
            currentState = "Your Turn";
        } else if (event.data == "Game Over") {
            statusText.textContent = "Game Over";
            currentState = "Game Over";
        } else if (event.data == "Opponent's Turn") {
            statusText.textContent = "Opponent's Turn";
            currentState = "Opponent's Turn";
        } else {
            statusText.textContent = "Sever Error";
            currentState = "Awaiting Name";
        }
    };

    socket.onclose = (event) => {
        if (event.wasClean) {
            console.log(
                `Closed cleanly, code=${event.code}, reason=${event.reason}`
            );
        } else {
            console.log("Connection died");
        }
    };

    socket.onerror = (error) => {
        console.log("Error: " + error.message);
    };
}

const boardPits = document.querySelectorAll(".board-pit");
function updateHTMLBoard() {
    boardPits.forEach((pit, index) => {
        pit.textContent = board[Math.floor(index / 7)][index % 7];
    });
}
