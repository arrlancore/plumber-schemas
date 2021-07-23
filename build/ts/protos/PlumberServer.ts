// Original file: protos/base.proto

import type * as grpc from '@grpc/grpc-js'
import type { MethodDefinition } from '@grpc/proto-loader'
import type { CreateConnectionRequest as _protos_CreateConnectionRequest, CreateConnectionRequest__Output as _protos_CreateConnectionRequest__Output } from '../protos/CreateConnectionRequest';
import type { CreateConnectionResponse as _protos_CreateConnectionResponse, CreateConnectionResponse__Output as _protos_CreateConnectionResponse__Output } from '../protos/CreateConnectionResponse';
import type { CreateRelayRequest as _protos_CreateRelayRequest, CreateRelayRequest__Output as _protos_CreateRelayRequest__Output } from '../protos/CreateRelayRequest';
import type { CreateRelayResponse as _protos_CreateRelayResponse, CreateRelayResponse__Output as _protos_CreateRelayResponse__Output } from '../protos/CreateRelayResponse';
import type { DeleteConnectionRequest as _protos_DeleteConnectionRequest, DeleteConnectionRequest__Output as _protos_DeleteConnectionRequest__Output } from '../protos/DeleteConnectionRequest';
import type { DeleteConnectionResponse as _protos_DeleteConnectionResponse, DeleteConnectionResponse__Output as _protos_DeleteConnectionResponse__Output } from '../protos/DeleteConnectionResponse';
import type { DeleteRelayRequest as _protos_DeleteRelayRequest, DeleteRelayRequest__Output as _protos_DeleteRelayRequest__Output } from '../protos/DeleteRelayRequest';
import type { DeleteRelayResponse as _protos_DeleteRelayResponse, DeleteRelayResponse__Output as _protos_DeleteRelayResponse__Output } from '../protos/DeleteRelayResponse';
import type { GetAllConnectionsRequest as _protos_GetAllConnectionsRequest, GetAllConnectionsRequest__Output as _protos_GetAllConnectionsRequest__Output } from '../protos/GetAllConnectionsRequest';
import type { GetAllConnectionsResponse as _protos_GetAllConnectionsResponse, GetAllConnectionsResponse__Output as _protos_GetAllConnectionsResponse__Output } from '../protos/GetAllConnectionsResponse';
import type { GetConnectionRequest as _protos_GetConnectionRequest, GetConnectionRequest__Output as _protos_GetConnectionRequest__Output } from '../protos/GetConnectionRequest';
import type { GetConnectionResponse as _protos_GetConnectionResponse, GetConnectionResponse__Output as _protos_GetConnectionResponse__Output } from '../protos/GetConnectionResponse';
import type { ResumeRelayRequest as _protos_ResumeRelayRequest, ResumeRelayRequest__Output as _protos_ResumeRelayRequest__Output } from '../protos/ResumeRelayRequest';
import type { ResumeRelayResponse as _protos_ResumeRelayResponse, ResumeRelayResponse__Output as _protos_ResumeRelayResponse__Output } from '../protos/ResumeRelayResponse';
import type { StartReadRequest as _protos_StartReadRequest, StartReadRequest__Output as _protos_StartReadRequest__Output } from '../protos/StartReadRequest';
import type { StartReadResponse as _protos_StartReadResponse, StartReadResponse__Output as _protos_StartReadResponse__Output } from '../protos/StartReadResponse';
import type { StopReadRequest as _protos_StopReadRequest, StopReadRequest__Output as _protos_StopReadRequest__Output } from '../protos/StopReadRequest';
import type { StopReadResponse as _protos_StopReadResponse, StopReadResponse__Output as _protos_StopReadResponse__Output } from '../protos/StopReadResponse';
import type { StopRelayRequest as _protos_StopRelayRequest, StopRelayRequest__Output as _protos_StopRelayRequest__Output } from '../protos/StopRelayRequest';
import type { StopRelayResponse as _protos_StopRelayResponse, StopRelayResponse__Output as _protos_StopRelayResponse__Output } from '../protos/StopRelayResponse';
import type { TestConnectionRequest as _protos_TestConnectionRequest, TestConnectionRequest__Output as _protos_TestConnectionRequest__Output } from '../protos/TestConnectionRequest';
import type { TestConnectionResponse as _protos_TestConnectionResponse, TestConnectionResponse__Output as _protos_TestConnectionResponse__Output } from '../protos/TestConnectionResponse';
import type { UpdateConnectionRequest as _protos_UpdateConnectionRequest, UpdateConnectionRequest__Output as _protos_UpdateConnectionRequest__Output } from '../protos/UpdateConnectionRequest';
import type { UpdateConnectionResponse as _protos_UpdateConnectionResponse, UpdateConnectionResponse__Output as _protos_UpdateConnectionResponse__Output } from '../protos/UpdateConnectionResponse';
import type { UpdateRelayRequest as _protos_UpdateRelayRequest, UpdateRelayRequest__Output as _protos_UpdateRelayRequest__Output } from '../protos/UpdateRelayRequest';
import type { UpdateRelayResponse as _protos_UpdateRelayResponse, UpdateRelayResponse__Output as _protos_UpdateRelayResponse__Output } from '../protos/UpdateRelayResponse';
import type { WriteRequest as _protos_WriteRequest, WriteRequest__Output as _protos_WriteRequest__Output } from '../protos/WriteRequest';
import type { WriteResponse as _protos_WriteResponse, WriteResponse__Output as _protos_WriteResponse__Output } from '../protos/WriteResponse';

export interface PlumberServerClient extends grpc.Client {
  CreateConnection(argument: _protos_CreateConnectionRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_CreateConnectionResponse__Output) => void): grpc.ClientUnaryCall;
  CreateConnection(argument: _protos_CreateConnectionRequest, metadata: grpc.Metadata, callback: (error?: grpc.ServiceError, result?: _protos_CreateConnectionResponse__Output) => void): grpc.ClientUnaryCall;
  CreateConnection(argument: _protos_CreateConnectionRequest, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_CreateConnectionResponse__Output) => void): grpc.ClientUnaryCall;
  CreateConnection(argument: _protos_CreateConnectionRequest, callback: (error?: grpc.ServiceError, result?: _protos_CreateConnectionResponse__Output) => void): grpc.ClientUnaryCall;
  createConnection(argument: _protos_CreateConnectionRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_CreateConnectionResponse__Output) => void): grpc.ClientUnaryCall;
  createConnection(argument: _protos_CreateConnectionRequest, metadata: grpc.Metadata, callback: (error?: grpc.ServiceError, result?: _protos_CreateConnectionResponse__Output) => void): grpc.ClientUnaryCall;
  createConnection(argument: _protos_CreateConnectionRequest, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_CreateConnectionResponse__Output) => void): grpc.ClientUnaryCall;
  createConnection(argument: _protos_CreateConnectionRequest, callback: (error?: grpc.ServiceError, result?: _protos_CreateConnectionResponse__Output) => void): grpc.ClientUnaryCall;
  
  CreateRelay(argument: _protos_CreateRelayRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_CreateRelayResponse__Output) => void): grpc.ClientUnaryCall;
  CreateRelay(argument: _protos_CreateRelayRequest, metadata: grpc.Metadata, callback: (error?: grpc.ServiceError, result?: _protos_CreateRelayResponse__Output) => void): grpc.ClientUnaryCall;
  CreateRelay(argument: _protos_CreateRelayRequest, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_CreateRelayResponse__Output) => void): grpc.ClientUnaryCall;
  CreateRelay(argument: _protos_CreateRelayRequest, callback: (error?: grpc.ServiceError, result?: _protos_CreateRelayResponse__Output) => void): grpc.ClientUnaryCall;
  createRelay(argument: _protos_CreateRelayRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_CreateRelayResponse__Output) => void): grpc.ClientUnaryCall;
  createRelay(argument: _protos_CreateRelayRequest, metadata: grpc.Metadata, callback: (error?: grpc.ServiceError, result?: _protos_CreateRelayResponse__Output) => void): grpc.ClientUnaryCall;
  createRelay(argument: _protos_CreateRelayRequest, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_CreateRelayResponse__Output) => void): grpc.ClientUnaryCall;
  createRelay(argument: _protos_CreateRelayRequest, callback: (error?: grpc.ServiceError, result?: _protos_CreateRelayResponse__Output) => void): grpc.ClientUnaryCall;
  
  DeleteConnection(argument: _protos_DeleteConnectionRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_DeleteConnectionResponse__Output) => void): grpc.ClientUnaryCall;
  DeleteConnection(argument: _protos_DeleteConnectionRequest, metadata: grpc.Metadata, callback: (error?: grpc.ServiceError, result?: _protos_DeleteConnectionResponse__Output) => void): grpc.ClientUnaryCall;
  DeleteConnection(argument: _protos_DeleteConnectionRequest, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_DeleteConnectionResponse__Output) => void): grpc.ClientUnaryCall;
  DeleteConnection(argument: _protos_DeleteConnectionRequest, callback: (error?: grpc.ServiceError, result?: _protos_DeleteConnectionResponse__Output) => void): grpc.ClientUnaryCall;
  deleteConnection(argument: _protos_DeleteConnectionRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_DeleteConnectionResponse__Output) => void): grpc.ClientUnaryCall;
  deleteConnection(argument: _protos_DeleteConnectionRequest, metadata: grpc.Metadata, callback: (error?: grpc.ServiceError, result?: _protos_DeleteConnectionResponse__Output) => void): grpc.ClientUnaryCall;
  deleteConnection(argument: _protos_DeleteConnectionRequest, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_DeleteConnectionResponse__Output) => void): grpc.ClientUnaryCall;
  deleteConnection(argument: _protos_DeleteConnectionRequest, callback: (error?: grpc.ServiceError, result?: _protos_DeleteConnectionResponse__Output) => void): grpc.ClientUnaryCall;
  
  DeleteRelay(argument: _protos_DeleteRelayRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_DeleteRelayResponse__Output) => void): grpc.ClientUnaryCall;
  DeleteRelay(argument: _protos_DeleteRelayRequest, metadata: grpc.Metadata, callback: (error?: grpc.ServiceError, result?: _protos_DeleteRelayResponse__Output) => void): grpc.ClientUnaryCall;
  DeleteRelay(argument: _protos_DeleteRelayRequest, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_DeleteRelayResponse__Output) => void): grpc.ClientUnaryCall;
  DeleteRelay(argument: _protos_DeleteRelayRequest, callback: (error?: grpc.ServiceError, result?: _protos_DeleteRelayResponse__Output) => void): grpc.ClientUnaryCall;
  deleteRelay(argument: _protos_DeleteRelayRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_DeleteRelayResponse__Output) => void): grpc.ClientUnaryCall;
  deleteRelay(argument: _protos_DeleteRelayRequest, metadata: grpc.Metadata, callback: (error?: grpc.ServiceError, result?: _protos_DeleteRelayResponse__Output) => void): grpc.ClientUnaryCall;
  deleteRelay(argument: _protos_DeleteRelayRequest, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_DeleteRelayResponse__Output) => void): grpc.ClientUnaryCall;
  deleteRelay(argument: _protos_DeleteRelayRequest, callback: (error?: grpc.ServiceError, result?: _protos_DeleteRelayResponse__Output) => void): grpc.ClientUnaryCall;
  
  GetAllConnections(argument: _protos_GetAllConnectionsRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_GetAllConnectionsResponse__Output) => void): grpc.ClientUnaryCall;
  GetAllConnections(argument: _protos_GetAllConnectionsRequest, metadata: grpc.Metadata, callback: (error?: grpc.ServiceError, result?: _protos_GetAllConnectionsResponse__Output) => void): grpc.ClientUnaryCall;
  GetAllConnections(argument: _protos_GetAllConnectionsRequest, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_GetAllConnectionsResponse__Output) => void): grpc.ClientUnaryCall;
  GetAllConnections(argument: _protos_GetAllConnectionsRequest, callback: (error?: grpc.ServiceError, result?: _protos_GetAllConnectionsResponse__Output) => void): grpc.ClientUnaryCall;
  getAllConnections(argument: _protos_GetAllConnectionsRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_GetAllConnectionsResponse__Output) => void): grpc.ClientUnaryCall;
  getAllConnections(argument: _protos_GetAllConnectionsRequest, metadata: grpc.Metadata, callback: (error?: grpc.ServiceError, result?: _protos_GetAllConnectionsResponse__Output) => void): grpc.ClientUnaryCall;
  getAllConnections(argument: _protos_GetAllConnectionsRequest, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_GetAllConnectionsResponse__Output) => void): grpc.ClientUnaryCall;
  getAllConnections(argument: _protos_GetAllConnectionsRequest, callback: (error?: grpc.ServiceError, result?: _protos_GetAllConnectionsResponse__Output) => void): grpc.ClientUnaryCall;
  
  GetConnection(argument: _protos_GetConnectionRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_GetConnectionResponse__Output) => void): grpc.ClientUnaryCall;
  GetConnection(argument: _protos_GetConnectionRequest, metadata: grpc.Metadata, callback: (error?: grpc.ServiceError, result?: _protos_GetConnectionResponse__Output) => void): grpc.ClientUnaryCall;
  GetConnection(argument: _protos_GetConnectionRequest, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_GetConnectionResponse__Output) => void): grpc.ClientUnaryCall;
  GetConnection(argument: _protos_GetConnectionRequest, callback: (error?: grpc.ServiceError, result?: _protos_GetConnectionResponse__Output) => void): grpc.ClientUnaryCall;
  getConnection(argument: _protos_GetConnectionRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_GetConnectionResponse__Output) => void): grpc.ClientUnaryCall;
  getConnection(argument: _protos_GetConnectionRequest, metadata: grpc.Metadata, callback: (error?: grpc.ServiceError, result?: _protos_GetConnectionResponse__Output) => void): grpc.ClientUnaryCall;
  getConnection(argument: _protos_GetConnectionRequest, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_GetConnectionResponse__Output) => void): grpc.ClientUnaryCall;
  getConnection(argument: _protos_GetConnectionRequest, callback: (error?: grpc.ServiceError, result?: _protos_GetConnectionResponse__Output) => void): grpc.ClientUnaryCall;
  
  ResumeRelay(argument: _protos_ResumeRelayRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_ResumeRelayResponse__Output) => void): grpc.ClientUnaryCall;
  ResumeRelay(argument: _protos_ResumeRelayRequest, metadata: grpc.Metadata, callback: (error?: grpc.ServiceError, result?: _protos_ResumeRelayResponse__Output) => void): grpc.ClientUnaryCall;
  ResumeRelay(argument: _protos_ResumeRelayRequest, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_ResumeRelayResponse__Output) => void): grpc.ClientUnaryCall;
  ResumeRelay(argument: _protos_ResumeRelayRequest, callback: (error?: grpc.ServiceError, result?: _protos_ResumeRelayResponse__Output) => void): grpc.ClientUnaryCall;
  resumeRelay(argument: _protos_ResumeRelayRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_ResumeRelayResponse__Output) => void): grpc.ClientUnaryCall;
  resumeRelay(argument: _protos_ResumeRelayRequest, metadata: grpc.Metadata, callback: (error?: grpc.ServiceError, result?: _protos_ResumeRelayResponse__Output) => void): grpc.ClientUnaryCall;
  resumeRelay(argument: _protos_ResumeRelayRequest, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_ResumeRelayResponse__Output) => void): grpc.ClientUnaryCall;
  resumeRelay(argument: _protos_ResumeRelayRequest, callback: (error?: grpc.ServiceError, result?: _protos_ResumeRelayResponse__Output) => void): grpc.ClientUnaryCall;
  
  StartRead(argument: _protos_StartReadRequest, metadata: grpc.Metadata, options?: grpc.CallOptions): grpc.ClientReadableStream<_protos_StartReadResponse__Output>;
  StartRead(argument: _protos_StartReadRequest, options?: grpc.CallOptions): grpc.ClientReadableStream<_protos_StartReadResponse__Output>;
  startRead(argument: _protos_StartReadRequest, metadata: grpc.Metadata, options?: grpc.CallOptions): grpc.ClientReadableStream<_protos_StartReadResponse__Output>;
  startRead(argument: _protos_StartReadRequest, options?: grpc.CallOptions): grpc.ClientReadableStream<_protos_StartReadResponse__Output>;
  
  StopRead(argument: _protos_StopReadRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_StopReadResponse__Output) => void): grpc.ClientUnaryCall;
  StopRead(argument: _protos_StopReadRequest, metadata: grpc.Metadata, callback: (error?: grpc.ServiceError, result?: _protos_StopReadResponse__Output) => void): grpc.ClientUnaryCall;
  StopRead(argument: _protos_StopReadRequest, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_StopReadResponse__Output) => void): grpc.ClientUnaryCall;
  StopRead(argument: _protos_StopReadRequest, callback: (error?: grpc.ServiceError, result?: _protos_StopReadResponse__Output) => void): grpc.ClientUnaryCall;
  stopRead(argument: _protos_StopReadRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_StopReadResponse__Output) => void): grpc.ClientUnaryCall;
  stopRead(argument: _protos_StopReadRequest, metadata: grpc.Metadata, callback: (error?: grpc.ServiceError, result?: _protos_StopReadResponse__Output) => void): grpc.ClientUnaryCall;
  stopRead(argument: _protos_StopReadRequest, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_StopReadResponse__Output) => void): grpc.ClientUnaryCall;
  stopRead(argument: _protos_StopReadRequest, callback: (error?: grpc.ServiceError, result?: _protos_StopReadResponse__Output) => void): grpc.ClientUnaryCall;
  
  StopRelay(argument: _protos_StopRelayRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_StopRelayResponse__Output) => void): grpc.ClientUnaryCall;
  StopRelay(argument: _protos_StopRelayRequest, metadata: grpc.Metadata, callback: (error?: grpc.ServiceError, result?: _protos_StopRelayResponse__Output) => void): grpc.ClientUnaryCall;
  StopRelay(argument: _protos_StopRelayRequest, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_StopRelayResponse__Output) => void): grpc.ClientUnaryCall;
  StopRelay(argument: _protos_StopRelayRequest, callback: (error?: grpc.ServiceError, result?: _protos_StopRelayResponse__Output) => void): grpc.ClientUnaryCall;
  stopRelay(argument: _protos_StopRelayRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_StopRelayResponse__Output) => void): grpc.ClientUnaryCall;
  stopRelay(argument: _protos_StopRelayRequest, metadata: grpc.Metadata, callback: (error?: grpc.ServiceError, result?: _protos_StopRelayResponse__Output) => void): grpc.ClientUnaryCall;
  stopRelay(argument: _protos_StopRelayRequest, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_StopRelayResponse__Output) => void): grpc.ClientUnaryCall;
  stopRelay(argument: _protos_StopRelayRequest, callback: (error?: grpc.ServiceError, result?: _protos_StopRelayResponse__Output) => void): grpc.ClientUnaryCall;
  
  TestConnection(argument: _protos_TestConnectionRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_TestConnectionResponse__Output) => void): grpc.ClientUnaryCall;
  TestConnection(argument: _protos_TestConnectionRequest, metadata: grpc.Metadata, callback: (error?: grpc.ServiceError, result?: _protos_TestConnectionResponse__Output) => void): grpc.ClientUnaryCall;
  TestConnection(argument: _protos_TestConnectionRequest, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_TestConnectionResponse__Output) => void): grpc.ClientUnaryCall;
  TestConnection(argument: _protos_TestConnectionRequest, callback: (error?: grpc.ServiceError, result?: _protos_TestConnectionResponse__Output) => void): grpc.ClientUnaryCall;
  testConnection(argument: _protos_TestConnectionRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_TestConnectionResponse__Output) => void): grpc.ClientUnaryCall;
  testConnection(argument: _protos_TestConnectionRequest, metadata: grpc.Metadata, callback: (error?: grpc.ServiceError, result?: _protos_TestConnectionResponse__Output) => void): grpc.ClientUnaryCall;
  testConnection(argument: _protos_TestConnectionRequest, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_TestConnectionResponse__Output) => void): grpc.ClientUnaryCall;
  testConnection(argument: _protos_TestConnectionRequest, callback: (error?: grpc.ServiceError, result?: _protos_TestConnectionResponse__Output) => void): grpc.ClientUnaryCall;
  
  UpdateConnection(argument: _protos_UpdateConnectionRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_UpdateConnectionResponse__Output) => void): grpc.ClientUnaryCall;
  UpdateConnection(argument: _protos_UpdateConnectionRequest, metadata: grpc.Metadata, callback: (error?: grpc.ServiceError, result?: _protos_UpdateConnectionResponse__Output) => void): grpc.ClientUnaryCall;
  UpdateConnection(argument: _protos_UpdateConnectionRequest, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_UpdateConnectionResponse__Output) => void): grpc.ClientUnaryCall;
  UpdateConnection(argument: _protos_UpdateConnectionRequest, callback: (error?: grpc.ServiceError, result?: _protos_UpdateConnectionResponse__Output) => void): grpc.ClientUnaryCall;
  updateConnection(argument: _protos_UpdateConnectionRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_UpdateConnectionResponse__Output) => void): grpc.ClientUnaryCall;
  updateConnection(argument: _protos_UpdateConnectionRequest, metadata: grpc.Metadata, callback: (error?: grpc.ServiceError, result?: _protos_UpdateConnectionResponse__Output) => void): grpc.ClientUnaryCall;
  updateConnection(argument: _protos_UpdateConnectionRequest, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_UpdateConnectionResponse__Output) => void): grpc.ClientUnaryCall;
  updateConnection(argument: _protos_UpdateConnectionRequest, callback: (error?: grpc.ServiceError, result?: _protos_UpdateConnectionResponse__Output) => void): grpc.ClientUnaryCall;
  
  UpdateRelay(argument: _protos_UpdateRelayRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_UpdateRelayResponse__Output) => void): grpc.ClientUnaryCall;
  UpdateRelay(argument: _protos_UpdateRelayRequest, metadata: grpc.Metadata, callback: (error?: grpc.ServiceError, result?: _protos_UpdateRelayResponse__Output) => void): grpc.ClientUnaryCall;
  UpdateRelay(argument: _protos_UpdateRelayRequest, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_UpdateRelayResponse__Output) => void): grpc.ClientUnaryCall;
  UpdateRelay(argument: _protos_UpdateRelayRequest, callback: (error?: grpc.ServiceError, result?: _protos_UpdateRelayResponse__Output) => void): grpc.ClientUnaryCall;
  updateRelay(argument: _protos_UpdateRelayRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_UpdateRelayResponse__Output) => void): grpc.ClientUnaryCall;
  updateRelay(argument: _protos_UpdateRelayRequest, metadata: grpc.Metadata, callback: (error?: grpc.ServiceError, result?: _protos_UpdateRelayResponse__Output) => void): grpc.ClientUnaryCall;
  updateRelay(argument: _protos_UpdateRelayRequest, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_UpdateRelayResponse__Output) => void): grpc.ClientUnaryCall;
  updateRelay(argument: _protos_UpdateRelayRequest, callback: (error?: grpc.ServiceError, result?: _protos_UpdateRelayResponse__Output) => void): grpc.ClientUnaryCall;
  
  Write(argument: _protos_WriteRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_WriteResponse__Output) => void): grpc.ClientUnaryCall;
  Write(argument: _protos_WriteRequest, metadata: grpc.Metadata, callback: (error?: grpc.ServiceError, result?: _protos_WriteResponse__Output) => void): grpc.ClientUnaryCall;
  Write(argument: _protos_WriteRequest, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_WriteResponse__Output) => void): grpc.ClientUnaryCall;
  Write(argument: _protos_WriteRequest, callback: (error?: grpc.ServiceError, result?: _protos_WriteResponse__Output) => void): grpc.ClientUnaryCall;
  write(argument: _protos_WriteRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_WriteResponse__Output) => void): grpc.ClientUnaryCall;
  write(argument: _protos_WriteRequest, metadata: grpc.Metadata, callback: (error?: grpc.ServiceError, result?: _protos_WriteResponse__Output) => void): grpc.ClientUnaryCall;
  write(argument: _protos_WriteRequest, options: grpc.CallOptions, callback: (error?: grpc.ServiceError, result?: _protos_WriteResponse__Output) => void): grpc.ClientUnaryCall;
  write(argument: _protos_WriteRequest, callback: (error?: grpc.ServiceError, result?: _protos_WriteResponse__Output) => void): grpc.ClientUnaryCall;
  
}

export interface PlumberServerHandlers extends grpc.UntypedServiceImplementation {
  CreateConnection: grpc.handleUnaryCall<_protos_CreateConnectionRequest__Output, _protos_CreateConnectionResponse>;
  
  CreateRelay: grpc.handleUnaryCall<_protos_CreateRelayRequest__Output, _protos_CreateRelayResponse>;
  
  DeleteConnection: grpc.handleUnaryCall<_protos_DeleteConnectionRequest__Output, _protos_DeleteConnectionResponse>;
  
  DeleteRelay: grpc.handleUnaryCall<_protos_DeleteRelayRequest__Output, _protos_DeleteRelayResponse>;
  
  GetAllConnections: grpc.handleUnaryCall<_protos_GetAllConnectionsRequest__Output, _protos_GetAllConnectionsResponse>;
  
  GetConnection: grpc.handleUnaryCall<_protos_GetConnectionRequest__Output, _protos_GetConnectionResponse>;
  
  ResumeRelay: grpc.handleUnaryCall<_protos_ResumeRelayRequest__Output, _protos_ResumeRelayResponse>;
  
  StartRead: grpc.handleServerStreamingCall<_protos_StartReadRequest__Output, _protos_StartReadResponse>;
  
  StopRead: grpc.handleUnaryCall<_protos_StopReadRequest__Output, _protos_StopReadResponse>;
  
  StopRelay: grpc.handleUnaryCall<_protos_StopRelayRequest__Output, _protos_StopRelayResponse>;
  
  TestConnection: grpc.handleUnaryCall<_protos_TestConnectionRequest__Output, _protos_TestConnectionResponse>;
  
  UpdateConnection: grpc.handleUnaryCall<_protos_UpdateConnectionRequest__Output, _protos_UpdateConnectionResponse>;
  
  UpdateRelay: grpc.handleUnaryCall<_protos_UpdateRelayRequest__Output, _protos_UpdateRelayResponse>;
  
  Write: grpc.handleUnaryCall<_protos_WriteRequest__Output, _protos_WriteResponse>;
  
}

export interface PlumberServerDefinition extends grpc.ServiceDefinition {
  CreateConnection: MethodDefinition<_protos_CreateConnectionRequest, _protos_CreateConnectionResponse, _protos_CreateConnectionRequest__Output, _protos_CreateConnectionResponse__Output>
  CreateRelay: MethodDefinition<_protos_CreateRelayRequest, _protos_CreateRelayResponse, _protos_CreateRelayRequest__Output, _protos_CreateRelayResponse__Output>
  DeleteConnection: MethodDefinition<_protos_DeleteConnectionRequest, _protos_DeleteConnectionResponse, _protos_DeleteConnectionRequest__Output, _protos_DeleteConnectionResponse__Output>
  DeleteRelay: MethodDefinition<_protos_DeleteRelayRequest, _protos_DeleteRelayResponse, _protos_DeleteRelayRequest__Output, _protos_DeleteRelayResponse__Output>
  GetAllConnections: MethodDefinition<_protos_GetAllConnectionsRequest, _protos_GetAllConnectionsResponse, _protos_GetAllConnectionsRequest__Output, _protos_GetAllConnectionsResponse__Output>
  GetConnection: MethodDefinition<_protos_GetConnectionRequest, _protos_GetConnectionResponse, _protos_GetConnectionRequest__Output, _protos_GetConnectionResponse__Output>
  ResumeRelay: MethodDefinition<_protos_ResumeRelayRequest, _protos_ResumeRelayResponse, _protos_ResumeRelayRequest__Output, _protos_ResumeRelayResponse__Output>
  StartRead: MethodDefinition<_protos_StartReadRequest, _protos_StartReadResponse, _protos_StartReadRequest__Output, _protos_StartReadResponse__Output>
  StopRead: MethodDefinition<_protos_StopReadRequest, _protos_StopReadResponse, _protos_StopReadRequest__Output, _protos_StopReadResponse__Output>
  StopRelay: MethodDefinition<_protos_StopRelayRequest, _protos_StopRelayResponse, _protos_StopRelayRequest__Output, _protos_StopRelayResponse__Output>
  TestConnection: MethodDefinition<_protos_TestConnectionRequest, _protos_TestConnectionResponse, _protos_TestConnectionRequest__Output, _protos_TestConnectionResponse__Output>
  UpdateConnection: MethodDefinition<_protos_UpdateConnectionRequest, _protos_UpdateConnectionResponse, _protos_UpdateConnectionRequest__Output, _protos_UpdateConnectionResponse__Output>
  UpdateRelay: MethodDefinition<_protos_UpdateRelayRequest, _protos_UpdateRelayResponse, _protos_UpdateRelayRequest__Output, _protos_UpdateRelayResponse__Output>
  Write: MethodDefinition<_protos_WriteRequest, _protos_WriteResponse, _protos_WriteRequest__Output, _protos_WriteResponse__Output>
}