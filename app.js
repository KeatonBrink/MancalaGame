let board = [
    [0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0]
]

const basicBoard = [
    [0, 4, 4, 4, 4, 4, 4],
    [4, 4, 4, 4, 4, 4, 0]
]

let currentPlayer = 1


const pits = document.querySelectorAll('.mancala-pit')

pits.forEach((pit, index) => {
    pit.addEventListener('click', () => {
        // Send a gRPC request to your Go back-end with the clicked pit index
        grpcClient.makeMove({ pitIndex: index }, (error, response) => {
            if (error) {
                console.error('Error making move:', error);
            } else {
                console.log('Move successful:', response.message);
                // Update the UI based on the response from the back-end
            }
        });
    });
});