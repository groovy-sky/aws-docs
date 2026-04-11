This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream HttpEndpointCommonAttribute

Describes the metadata that's delivered to the specified HTTP endpoint destination.
Kinesis Firehose supports any custom HTTP endpoint or HTTP endpoints owned by supported
third-party service providers, including Datadog, MongoDB, and New Relic.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AttributeName" : String,
  "AttributeValue" : String
}

```

### YAML

```yaml

  AttributeName: String
  AttributeValue: String

```

## Properties

`AttributeName`

The name of the HTTP endpoint common attribute.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AttributeValue`

The value of the HTTP endpoint common attribute.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HiveJsonSerDe

HttpEndpointConfiguration

All content copied from https://docs.aws.amazon.com/.
