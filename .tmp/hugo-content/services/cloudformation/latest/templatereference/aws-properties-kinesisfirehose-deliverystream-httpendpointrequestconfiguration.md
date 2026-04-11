This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream HttpEndpointRequestConfiguration

The configuration of the HTTP endpoint request. Kinesis Firehose supports any custom
HTTP endpoint or HTTP endpoints owned by supported third-party service providers, including
Datadog, MongoDB, and New Relic.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CommonAttributes" : [ HttpEndpointCommonAttribute, ... ],
  "ContentEncoding" : String
}

```

### YAML

```yaml

  CommonAttributes:
    - HttpEndpointCommonAttribute
  ContentEncoding: String

```

## Properties

`CommonAttributes`

Describes the metadata sent to the HTTP endpoint destination.

_Required_: No

_Type_: Array of [HttpEndpointCommonAttribute](aws-properties-kinesisfirehose-deliverystream-httpendpointcommonattribute.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContentEncoding`

Kinesis Data Firehose uses the content encoding to compress the body of a request
before sending the request to the destination. For more information, see Content-Encoding
in MDN Web Docs, the official Mozilla documentation.

_Required_: No

_Type_: String

_Allowed values_: `NONE | GZIP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HttpEndpointDestinationConfiguration

IcebergDestinationConfiguration

All content copied from https://docs.aws.amazon.com/.
