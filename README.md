# MancalaGame

Mancala Game attempts to leverage multiple languages and protocols to provide a seemless UI/UX.

The basic structure has a Nodejs server to serve client browser files and forward ajax requests to the backend. The Go backend handles game and move requests to mediate the experience of many clients via the Nodejs server using gRPC. Other software utilized include Expressjs and Websockets.

TODO:
Finish movement logic.

Start gameover logic
