
let board = [
    [0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0]
]

const basicBoard = [
    [0, 4, 4, 4, 4, 4, 4],
    [4, 4, 4, 4, 4, 4, 0]
]

let currentState = "Awaiting Name"

let currentPlayer = 1

const statusText = document.querySelector('#status-text')
const pits = document.querySelectorAll('.mancala-pit')

pits.forEach((pit, index) => {
    pit.addEventListener('click', () => {
        console.log('Pit: ', index)
        const xhr = new XMLHttpRequest();
        xhr.open('POST', '/MakeMove', true)
        xhr.setRequestHeader('Content-Type', 'application/json');
        xhr.onreadystatechange = function() {
            if (xhr.readyState === 4 && xhr.status === 200) {
                const response = JSON.parse(xhr.responseText);
                console.log(response.message)
            }
        };
        xhr.send(JSON.stringify({ pitIndex: index }));
    });
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

const userNameSubmission = document.querySelector('#user-name-submission')
const userNameLocation = document.getElementById('user-name')

userNameSubmission.addEventListener('click', () => {
    var userName = userNameLocation.value
    console.log('Username: ' + userName)

    if (userName.length == 0) {
        statusText.textContent = 'Username must be at least one character long!'
        return
    }
    statusText.textContent = 'Attempting to join game'
    const xhr = new XMLHttpRequest();
    xhr.open('POST', '/FindGame', true)
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.onreadystatechange = function() {
        if (xhr.readyState === 4 && xhr.status === 200) {
            const response = JSON.parse(xhr.responseText);
            console.log(JSON.stringify(response))
            statusText.textContent = response.message
            if (response.data.message == "Waiting for Second Player") {
                webSocketToServer(response.data.serverWebSocketAddress)
            }
        } else {
            statusText.textContent = 'Error contacting server, try again'
        }
    };
    xhr.send(JSON.stringify({ userName: userName }));
})

function webSocketToServer(webSocketAddress) {

	let socket = new WebSocket("ws://localhost" + webSocketAddress + "/ws");

    socket.onopen = () => {
        console.log("Connected to WebSocket")
    };

    socket.onmessage = event => {
        console.log("Received: " + event.data)
    };

    socket.onclose = event => {
        if (event.wasClean) {
            console.log(`Closed cleanly, code=${event.code}, reason=${event.reason}`)
        } else {
            console.log("Connection died")
        }
    };
}