This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::Distribution GrpcConfig

Amazon CloudFront supports gRPC, an open-source remote procedure call (RPC) framework built on
HTTP/2. gRPC offers bi-directional streaming and binary protocol that buffers payloads,
making it suitable for applications that require low latency communications.

To enable your distribution to handle gRPC requests, you must include HTTP/2 as one of the supported `HTTP` versions and allow `HTTP` methods, including `POST`.

For more information, see [Using gRPC with CloudFront distributions](../../../amazoncloudfront/latest/developerguide/distribution-using-grpc.md) in the
_Amazon CloudFront Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enabled" : Boolean
}

```

### YAML

```yaml

  Enabled: Boolean

```

## Properties

`Enabled`

Enables your CloudFront distribution to receive gRPC requests and to proxy them directly to your
origins.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GeoRestriction

LambdaFunctionAssociation

All content copied from https://docs.aws.amazon.com/.
