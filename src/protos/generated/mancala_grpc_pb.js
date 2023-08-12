// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var mancala_pb = require('./mancala_pb.js');

function serialize_mancala_HandshakeRequest(arg) {
  if (!(arg instanceof mancala_pb.HandshakeRequest)) {
    throw new Error('Expected argument of type mancala.HandshakeRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_mancala_HandshakeRequest(buffer_arg) {
  return mancala_pb.HandshakeRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_mancala_HandshakeResponse(arg) {
  if (!(arg instanceof mancala_pb.HandshakeResponse)) {
    throw new Error('Expected argument of type mancala.HandshakeResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_mancala_HandshakeResponse(buffer_arg) {
  return mancala_pb.HandshakeResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_mancala_MoveRequest(arg) {
  if (!(arg instanceof mancala_pb.MoveRequest)) {
    throw new Error('Expected argument of type mancala.MoveRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_mancala_MoveRequest(buffer_arg) {
  return mancala_pb.MoveRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_mancala_MoveResponse(arg) {
  if (!(arg instanceof mancala_pb.MoveResponse)) {
    throw new Error('Expected argument of type mancala.MoveResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_mancala_MoveResponse(buffer_arg) {
  return mancala_pb.MoveResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_mancala_UpdateRequest(arg) {
  if (!(arg instanceof mancala_pb.UpdateRequest)) {
    throw new Error('Expected argument of type mancala.UpdateRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_mancala_UpdateRequest(buffer_arg) {
  return mancala_pb.UpdateRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_mancala_UpdateResponse(arg) {
  if (!(arg instanceof mancala_pb.UpdateResponse)) {
    throw new Error('Expected argument of type mancala.UpdateResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_mancala_UpdateResponse(buffer_arg) {
  return mancala_pb.UpdateResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var MancalaServiceService = exports.MancalaServiceService = {
  gameHandshake: {
    path: '/mancala.MancalaService/GameHandshake',
    requestStream: false,
    responseStream: false,
    requestType: mancala_pb.HandshakeRequest,
    responseType: mancala_pb.HandshakeResponse,
    requestSerialize: serialize_mancala_HandshakeRequest,
    requestDeserialize: deserialize_mancala_HandshakeRequest,
    responseSerialize: serialize_mancala_HandshakeResponse,
    responseDeserialize: deserialize_mancala_HandshakeResponse,
  },
  makeMove: {
    path: '/mancala.MancalaService/MakeMove',
    requestStream: false,
    responseStream: false,
    requestType: mancala_pb.MoveRequest,
    responseType: mancala_pb.MoveResponse,
    requestSerialize: serialize_mancala_MoveRequest,
    requestDeserialize: deserialize_mancala_MoveRequest,
    responseSerialize: serialize_mancala_MoveResponse,
    responseDeserialize: deserialize_mancala_MoveResponse,
  },
  requestUpdate: {
    path: '/mancala.MancalaService/RequestUpdate',
    requestStream: false,
    responseStream: false,
    requestType: mancala_pb.UpdateRequest,
    responseType: mancala_pb.UpdateResponse,
    requestSerialize: serialize_mancala_UpdateRequest,
    requestDeserialize: deserialize_mancala_UpdateRequest,
    responseSerialize: serialize_mancala_UpdateResponse,
    responseDeserialize: deserialize_mancala_UpdateResponse,
  },
};

exports.MancalaServiceClient = grpc.makeGenericClientConstructor(MancalaServiceService);
