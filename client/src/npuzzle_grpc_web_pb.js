/**
 * @fileoverview gRPC-Web generated client stub for npuzzle
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.npuzzle = require('./npuzzle_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.npuzzle.NpuzzleClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.npuzzle.NpuzzlePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.npuzzle.Handshake,
 *   !proto.npuzzle.Handshake>}
 */
const methodDescriptor_Npuzzle_Greets = new grpc.web.MethodDescriptor(
  '/npuzzle.Npuzzle/Greets',
  grpc.web.MethodType.UNARY,
  proto.npuzzle.Handshake,
  proto.npuzzle.Handshake,
  /** @param {!proto.npuzzle.Handshake} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.npuzzle.Handshake.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.npuzzle.Handshake,
 *   !proto.npuzzle.Handshake>}
 */
const methodInfo_Npuzzle_Greets = new grpc.web.AbstractClientBase.MethodInfo(
  proto.npuzzle.Handshake,
  /** @param {!proto.npuzzle.Handshake} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.npuzzle.Handshake.deserializeBinary
);


/**
 * @param {!proto.npuzzle.Handshake} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.npuzzle.Handshake)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.npuzzle.Handshake>|undefined}
 *     The XHR Node Readable Stream
 */
proto.npuzzle.NpuzzleClient.prototype.greets =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/npuzzle.Npuzzle/Greets',
      request,
      metadata || {},
      methodDescriptor_Npuzzle_Greets,
      callback);
};


/**
 * @param {!proto.npuzzle.Handshake} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.npuzzle.Handshake>}
 *     A native promise that resolves to the response
 */
proto.npuzzle.NpuzzlePromiseClient.prototype.greets =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/npuzzle.Npuzzle/Greets',
      request,
      metadata || {},
      methodDescriptor_Npuzzle_Greets);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.npuzzle.Problem,
 *   !proto.npuzzle.Empty>}
 */
const methodDescriptor_Npuzzle_Solve = new grpc.web.MethodDescriptor(
  '/npuzzle.Npuzzle/Solve',
  grpc.web.MethodType.UNARY,
  proto.npuzzle.Problem,
  proto.npuzzle.Empty,
  /** @param {!proto.npuzzle.Problem} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.npuzzle.Empty.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.npuzzle.Problem,
 *   !proto.npuzzle.Empty>}
 */
const methodInfo_Npuzzle_Solve = new grpc.web.AbstractClientBase.MethodInfo(
  proto.npuzzle.Empty,
  /** @param {!proto.npuzzle.Problem} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.npuzzle.Empty.deserializeBinary
);


/**
 * @param {!proto.npuzzle.Problem} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.npuzzle.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.npuzzle.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.npuzzle.NpuzzleClient.prototype.solve =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/npuzzle.Npuzzle/Solve',
      request,
      metadata || {},
      methodDescriptor_Npuzzle_Solve,
      callback);
};


/**
 * @param {!proto.npuzzle.Problem} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.npuzzle.Empty>}
 *     A native promise that resolves to the response
 */
proto.npuzzle.NpuzzlePromiseClient.prototype.solve =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/npuzzle.Npuzzle/Solve',
      request,
      metadata || {},
      methodDescriptor_Npuzzle_Solve);
};


module.exports = proto.npuzzle;

