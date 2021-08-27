// package: protos.backends
// file: backends/aws-sns.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class AWSSNSConn extends jspb.Message { 
    getAwsRegion(): string;
    setAwsRegion(value: string): AWSSNSConn;
    getAwsAccessKeyId(): string;
    setAwsAccessKeyId(value: string): AWSSNSConn;
    getAwsSecretAccessKey(): string;
    setAwsSecretAccessKey(value: string): AWSSNSConn;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AWSSNSConn.AsObject;
    static toObject(includeInstance: boolean, msg: AWSSNSConn): AWSSNSConn.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AWSSNSConn, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AWSSNSConn;
    static deserializeBinaryFromReader(message: AWSSNSConn, reader: jspb.BinaryReader): AWSSNSConn;
}

export namespace AWSSNSConn {
    export type AsObject = {
        awsRegion: string,
        awsAccessKeyId: string,
        awsSecretAccessKey: string,
    }
}

export class AWSSNSReadArgs extends jspb.Message { 
    getTopicArn(): string;
    setTopicArn(value: string): AWSSNSReadArgs;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AWSSNSReadArgs.AsObject;
    static toObject(includeInstance: boolean, msg: AWSSNSReadArgs): AWSSNSReadArgs.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AWSSNSReadArgs, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AWSSNSReadArgs;
    static deserializeBinaryFromReader(message: AWSSNSReadArgs, reader: jspb.BinaryReader): AWSSNSReadArgs;
}

export namespace AWSSNSReadArgs {
    export type AsObject = {
        topicArn: string,
    }
}

export class AWSSNSWriteArgs extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AWSSNSWriteArgs.AsObject;
    static toObject(includeInstance: boolean, msg: AWSSNSWriteArgs): AWSSNSWriteArgs.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AWSSNSWriteArgs, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AWSSNSWriteArgs;
    static deserializeBinaryFromReader(message: AWSSNSWriteArgs, reader: jspb.BinaryReader): AWSSNSWriteArgs;
}

export namespace AWSSNSWriteArgs {
    export type AsObject = {
    }
}
