// @generated by protoc-gen-es v1.3.0 with parameter "target=ts"
// @generated from file mgmt/v1alpha1/transformer.proto (package mgmt.v1alpha1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message mgmt.v1alpha1.Transformer
 */
export class Transformer extends Message<Transformer> {
  /**
   * @generated from field: string title = 1;
   */
  title = "";

  /**
   * @generated from field: string value = 2;
   */
  value = "";

  constructor(data?: PartialMessage<Transformer>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "mgmt.v1alpha1.Transformer";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "title", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "value", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Transformer {
    return new Transformer().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Transformer {
    return new Transformer().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Transformer {
    return new Transformer().fromJsonString(jsonString, options);
  }

  static equals(a: Transformer | PlainMessage<Transformer> | undefined, b: Transformer | PlainMessage<Transformer> | undefined): boolean {
    return proto3.util.equals(Transformer, a, b);
  }
}

/**
 * @generated from message mgmt.v1alpha1.GetTransformersRequest
 */
export class GetTransformersRequest extends Message<GetTransformersRequest> {
  constructor(data?: PartialMessage<GetTransformersRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "mgmt.v1alpha1.GetTransformersRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetTransformersRequest {
    return new GetTransformersRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetTransformersRequest {
    return new GetTransformersRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetTransformersRequest {
    return new GetTransformersRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetTransformersRequest | PlainMessage<GetTransformersRequest> | undefined, b: GetTransformersRequest | PlainMessage<GetTransformersRequest> | undefined): boolean {
    return proto3.util.equals(GetTransformersRequest, a, b);
  }
}

/**
 * @generated from message mgmt.v1alpha1.GetTransformersResponse
 */
export class GetTransformersResponse extends Message<GetTransformersResponse> {
  /**
   * @generated from field: repeated mgmt.v1alpha1.Transformer transformers = 1;
   */
  transformers: Transformer[] = [];

  constructor(data?: PartialMessage<GetTransformersResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "mgmt.v1alpha1.GetTransformersResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "transformers", kind: "message", T: Transformer, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetTransformersResponse {
    return new GetTransformersResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetTransformersResponse {
    return new GetTransformersResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetTransformersResponse {
    return new GetTransformersResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetTransformersResponse | PlainMessage<GetTransformersResponse> | undefined, b: GetTransformersResponse | PlainMessage<GetTransformersResponse> | undefined): boolean {
    return proto3.util.equals(GetTransformersResponse, a, b);
  }
}
