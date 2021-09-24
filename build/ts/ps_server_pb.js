// source: ps_server.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var common_ps_common_status_pb = require('./common/ps_common_status_pb.js');
goog.object.extend(proto, common_ps_common_status_pb);
var common_ps_common_auth_pb = require('./common/ps_common_auth_pb.js');
goog.object.extend(proto, common_ps_common_auth_pb);
var opts_ps_opts_server_pb = require('./opts/ps_opts_server_pb.js');
goog.object.extend(proto, opts_ps_opts_server_pb);
goog.exportSymbol('proto.protos.GetServerOptionsRequest', null, global);
goog.exportSymbol('proto.protos.GetServerOptionsResponse', null, global);
goog.exportSymbol('proto.protos.SetServerOptionsRequest', null, global);
goog.exportSymbol('proto.protos.SetServerOptionsResponse', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.protos.GetServerOptionsRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, 500, null, null);
};
goog.inherits(proto.protos.GetServerOptionsRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.protos.GetServerOptionsRequest.displayName = 'proto.protos.GetServerOptionsRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.protos.GetServerOptionsResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.protos.GetServerOptionsResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.protos.GetServerOptionsResponse.displayName = 'proto.protos.GetServerOptionsResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.protos.SetServerOptionsRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, 500, null, null);
};
goog.inherits(proto.protos.SetServerOptionsRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.protos.SetServerOptionsRequest.displayName = 'proto.protos.SetServerOptionsRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.protos.SetServerOptionsResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, 500, null, null);
};
goog.inherits(proto.protos.SetServerOptionsResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.protos.SetServerOptionsResponse.displayName = 'proto.protos.SetServerOptionsResponse';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.protos.GetServerOptionsRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.protos.GetServerOptionsRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.protos.GetServerOptionsRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.GetServerOptionsRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    auth: (f = msg.getAuth()) && common_ps_common_auth_pb.Auth.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.protos.GetServerOptionsRequest}
 */
proto.protos.GetServerOptionsRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.protos.GetServerOptionsRequest;
  return proto.protos.GetServerOptionsRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.protos.GetServerOptionsRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.protos.GetServerOptionsRequest}
 */
proto.protos.GetServerOptionsRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 9999:
      var value = new common_ps_common_auth_pb.Auth;
      reader.readMessage(value,common_ps_common_auth_pb.Auth.deserializeBinaryFromReader);
      msg.setAuth(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.protos.GetServerOptionsRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.protos.GetServerOptionsRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.protos.GetServerOptionsRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.GetServerOptionsRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getAuth();
  if (f != null) {
    writer.writeMessage(
      9999,
      f,
      common_ps_common_auth_pb.Auth.serializeBinaryToWriter
    );
  }
};


/**
 * optional common.Auth auth = 9999;
 * @return {?proto.protos.common.Auth}
 */
proto.protos.GetServerOptionsRequest.prototype.getAuth = function() {
  return /** @type{?proto.protos.common.Auth} */ (
    jspb.Message.getWrapperField(this, common_ps_common_auth_pb.Auth, 9999));
};


/**
 * @param {?proto.protos.common.Auth|undefined} value
 * @return {!proto.protos.GetServerOptionsRequest} returns this
*/
proto.protos.GetServerOptionsRequest.prototype.setAuth = function(value) {
  return jspb.Message.setWrapperField(this, 9999, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.protos.GetServerOptionsRequest} returns this
 */
proto.protos.GetServerOptionsRequest.prototype.clearAuth = function() {
  return this.setAuth(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.protos.GetServerOptionsRequest.prototype.hasAuth = function() {
  return jspb.Message.getField(this, 9999) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.protos.GetServerOptionsResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.protos.GetServerOptionsResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.protos.GetServerOptionsResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.GetServerOptionsResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    serverOptions: (f = msg.getServerOptions()) && opts_ps_opts_server_pb.ServerOptions.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.protos.GetServerOptionsResponse}
 */
proto.protos.GetServerOptionsResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.protos.GetServerOptionsResponse;
  return proto.protos.GetServerOptionsResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.protos.GetServerOptionsResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.protos.GetServerOptionsResponse}
 */
proto.protos.GetServerOptionsResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new opts_ps_opts_server_pb.ServerOptions;
      reader.readMessage(value,opts_ps_opts_server_pb.ServerOptions.deserializeBinaryFromReader);
      msg.setServerOptions(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.protos.GetServerOptionsResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.protos.GetServerOptionsResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.protos.GetServerOptionsResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.GetServerOptionsResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getServerOptions();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      opts_ps_opts_server_pb.ServerOptions.serializeBinaryToWriter
    );
  }
};


/**
 * optional opts.ServerOptions server_options = 1;
 * @return {?proto.protos.opts.ServerOptions}
 */
proto.protos.GetServerOptionsResponse.prototype.getServerOptions = function() {
  return /** @type{?proto.protos.opts.ServerOptions} */ (
    jspb.Message.getWrapperField(this, opts_ps_opts_server_pb.ServerOptions, 1));
};


/**
 * @param {?proto.protos.opts.ServerOptions|undefined} value
 * @return {!proto.protos.GetServerOptionsResponse} returns this
*/
proto.protos.GetServerOptionsResponse.prototype.setServerOptions = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.protos.GetServerOptionsResponse} returns this
 */
proto.protos.GetServerOptionsResponse.prototype.clearServerOptions = function() {
  return this.setServerOptions(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.protos.GetServerOptionsResponse.prototype.hasServerOptions = function() {
  return jspb.Message.getField(this, 1) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.protos.SetServerOptionsRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.protos.SetServerOptionsRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.protos.SetServerOptionsRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.SetServerOptionsRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    auth: (f = msg.getAuth()) && common_ps_common_auth_pb.Auth.toObject(includeInstance, f),
    vcserviceToken: jspb.Message.getFieldWithDefault(msg, 1, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.protos.SetServerOptionsRequest}
 */
proto.protos.SetServerOptionsRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.protos.SetServerOptionsRequest;
  return proto.protos.SetServerOptionsRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.protos.SetServerOptionsRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.protos.SetServerOptionsRequest}
 */
proto.protos.SetServerOptionsRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 9999:
      var value = new common_ps_common_auth_pb.Auth;
      reader.readMessage(value,common_ps_common_auth_pb.Auth.deserializeBinaryFromReader);
      msg.setAuth(value);
      break;
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setVcserviceToken(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.protos.SetServerOptionsRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.protos.SetServerOptionsRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.protos.SetServerOptionsRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.SetServerOptionsRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getAuth();
  if (f != null) {
    writer.writeMessage(
      9999,
      f,
      common_ps_common_auth_pb.Auth.serializeBinaryToWriter
    );
  }
  f = message.getVcserviceToken();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional common.Auth auth = 9999;
 * @return {?proto.protos.common.Auth}
 */
proto.protos.SetServerOptionsRequest.prototype.getAuth = function() {
  return /** @type{?proto.protos.common.Auth} */ (
    jspb.Message.getWrapperField(this, common_ps_common_auth_pb.Auth, 9999));
};


/**
 * @param {?proto.protos.common.Auth|undefined} value
 * @return {!proto.protos.SetServerOptionsRequest} returns this
*/
proto.protos.SetServerOptionsRequest.prototype.setAuth = function(value) {
  return jspb.Message.setWrapperField(this, 9999, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.protos.SetServerOptionsRequest} returns this
 */
proto.protos.SetServerOptionsRequest.prototype.clearAuth = function() {
  return this.setAuth(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.protos.SetServerOptionsRequest.prototype.hasAuth = function() {
  return jspb.Message.getField(this, 9999) != null;
};


/**
 * optional string vcservice_token = 1;
 * @return {string}
 */
proto.protos.SetServerOptionsRequest.prototype.getVcserviceToken = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.SetServerOptionsRequest} returns this
 */
proto.protos.SetServerOptionsRequest.prototype.setVcserviceToken = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.protos.SetServerOptionsResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.protos.SetServerOptionsResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.protos.SetServerOptionsResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.SetServerOptionsResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    serverOptions: (f = msg.getServerOptions()) && opts_ps_opts_server_pb.ServerOptions.toObject(includeInstance, f),
    status: (f = msg.getStatus()) && common_ps_common_status_pb.Status.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.protos.SetServerOptionsResponse}
 */
proto.protos.SetServerOptionsResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.protos.SetServerOptionsResponse;
  return proto.protos.SetServerOptionsResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.protos.SetServerOptionsResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.protos.SetServerOptionsResponse}
 */
proto.protos.SetServerOptionsResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new opts_ps_opts_server_pb.ServerOptions;
      reader.readMessage(value,opts_ps_opts_server_pb.ServerOptions.deserializeBinaryFromReader);
      msg.setServerOptions(value);
      break;
    case 1000:
      var value = new common_ps_common_status_pb.Status;
      reader.readMessage(value,common_ps_common_status_pb.Status.deserializeBinaryFromReader);
      msg.setStatus(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.protos.SetServerOptionsResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.protos.SetServerOptionsResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.protos.SetServerOptionsResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.SetServerOptionsResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getServerOptions();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      opts_ps_opts_server_pb.ServerOptions.serializeBinaryToWriter
    );
  }
  f = message.getStatus();
  if (f != null) {
    writer.writeMessage(
      1000,
      f,
      common_ps_common_status_pb.Status.serializeBinaryToWriter
    );
  }
};


/**
 * optional opts.ServerOptions server_options = 1;
 * @return {?proto.protos.opts.ServerOptions}
 */
proto.protos.SetServerOptionsResponse.prototype.getServerOptions = function() {
  return /** @type{?proto.protos.opts.ServerOptions} */ (
    jspb.Message.getWrapperField(this, opts_ps_opts_server_pb.ServerOptions, 1));
};


/**
 * @param {?proto.protos.opts.ServerOptions|undefined} value
 * @return {!proto.protos.SetServerOptionsResponse} returns this
*/
proto.protos.SetServerOptionsResponse.prototype.setServerOptions = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.protos.SetServerOptionsResponse} returns this
 */
proto.protos.SetServerOptionsResponse.prototype.clearServerOptions = function() {
  return this.setServerOptions(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.protos.SetServerOptionsResponse.prototype.hasServerOptions = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional common.Status status = 1000;
 * @return {?proto.protos.common.Status}
 */
proto.protos.SetServerOptionsResponse.prototype.getStatus = function() {
  return /** @type{?proto.protos.common.Status} */ (
    jspb.Message.getWrapperField(this, common_ps_common_status_pb.Status, 1000));
};


/**
 * @param {?proto.protos.common.Status|undefined} value
 * @return {!proto.protos.SetServerOptionsResponse} returns this
*/
proto.protos.SetServerOptionsResponse.prototype.setStatus = function(value) {
  return jspb.Message.setWrapperField(this, 1000, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.protos.SetServerOptionsResponse} returns this
 */
proto.protos.SetServerOptionsResponse.prototype.clearStatus = function() {
  return this.setStatus(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.protos.SetServerOptionsResponse.prototype.hasStatus = function() {
  return jspb.Message.getField(this, 1000) != null;
};


goog.object.extend(exports, proto.protos);