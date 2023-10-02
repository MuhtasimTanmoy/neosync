// @generated by protoc-gen-es v1.3.1 with parameter "target=ts"
// @generated from file mgmt/v1alpha1/transformer.proto (package mgmt.v1alpha1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

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
   * @generated from field: repeated mgmt.v1alpha1.Transformers transformers = 1;
   */
  transformers: Transformers[] = [];

  constructor(data?: PartialMessage<GetTransformersResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "mgmt.v1alpha1.GetTransformersResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "transformers", kind: "message", T: Transformers, repeated: true },
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

/**
 * @generated from message mgmt.v1alpha1.Transformers
 */
export class Transformers extends Message<Transformers> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string name = 2;
   */
  name = "";

  /**
   * @generated from field: string description = 3;
   */
  description = "";

  /**
   * @generated from field: mgmt.v1alpha1.TransformerConfig config = 4;
   */
  config?: TransformerConfig;

  constructor(data?: PartialMessage<Transformers>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "mgmt.v1alpha1.Transformers";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "config", kind: "message", T: TransformerConfig },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Transformers {
    return new Transformers().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Transformers {
    return new Transformers().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Transformers {
    return new Transformers().fromJsonString(jsonString, options);
  }

  static equals(a: Transformers | PlainMessage<Transformers> | undefined, b: Transformers | PlainMessage<Transformers> | undefined): boolean {
    return proto3.util.equals(Transformers, a, b);
  }
}

/**
 * @generated from message mgmt.v1alpha1.TransformerConfig
 */
export class TransformerConfig extends Message<TransformerConfig> {
  /**
   * @generated from oneof mgmt.v1alpha1.TransformerConfig.config
   */
  config: {
    /**
     * @generated from field: mgmt.v1alpha1.EmailConfig email_config = 1;
     */
    value: EmailConfig;
    case: "emailConfig";
  } | { case: undefined; value?: undefined } = { case: undefined };

  constructor(data?: PartialMessage<TransformerConfig>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "mgmt.v1alpha1.TransformerConfig";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "email_config", kind: "message", T: EmailConfig, oneof: "config" },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TransformerConfig {
    return new TransformerConfig().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TransformerConfig {
    return new TransformerConfig().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TransformerConfig {
    return new TransformerConfig().fromJsonString(jsonString, options);
  }

  static equals(a: TransformerConfig | PlainMessage<TransformerConfig> | undefined, b: TransformerConfig | PlainMessage<TransformerConfig> | undefined): boolean {
    return proto3.util.equals(TransformerConfig, a, b);
  }
}

/**
 * @generated from message mgmt.v1alpha1.EmailConfig
 */
export class EmailConfig extends Message<EmailConfig> {
  /**
   * @generated from field: bool preserve_domain = 1;
   */
  preserveDomain = false;

  /**
   * @generated from field: bool preserve_length = 2;
   */
  preserveLength = false;

  constructor(data?: PartialMessage<EmailConfig>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "mgmt.v1alpha1.EmailConfig";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "preserve_domain", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 2, name: "preserve_length", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): EmailConfig {
    return new EmailConfig().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): EmailConfig {
    return new EmailConfig().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): EmailConfig {
    return new EmailConfig().fromJsonString(jsonString, options);
  }

  static equals(a: EmailConfig | PlainMessage<EmailConfig> | undefined, b: EmailConfig | PlainMessage<EmailConfig> | undefined): boolean {
    return proto3.util.equals(EmailConfig, a, b);
  }
}

