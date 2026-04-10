This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream HttpEndpointConfiguration

Describes the configuration of the HTTP endpoint to which Kinesis Firehose delivers
data. Kinesis Firehose supports any custom HTTP endpoint or HTTP endpoints owned by
supported third-party service providers, including Datadog, MongoDB, and New Relic.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccessKey" : String,
  "Name" : String,
  "Url" : String
}

```

### YAML

```yaml

  AccessKey: String
  Name: String
  Url: String

```

## Properties

`AccessKey`

The access key required for Kinesis Firehose to authenticate with the HTTP endpoint
selected as the destination.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the HTTP endpoint selected as the destination.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Url`

The URL of the HTTP endpoint selected as the destination.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HttpEndpointCommonAttribute

HttpEndpointDestinationConfiguration

All content copied from https://docs.aws.amazon.com/.
