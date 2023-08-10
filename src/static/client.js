
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
        const xhr = new XMLHttpRequest();
        xhr.open('POST', '/hello', true)
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
