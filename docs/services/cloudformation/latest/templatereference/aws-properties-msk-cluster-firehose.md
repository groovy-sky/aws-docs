This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MSK::Cluster Firehose

Firehose details for BrokerLogs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DeliveryStream" : String,
  "Enabled" : Boolean
}

```

### YAML

```yaml

  DeliveryStream: String
  Enabled: Boolean

```

## Properties

`DeliveryStream`

The Kinesis Data Firehose delivery stream that is the destination for broker logs.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

Specifies whether broker logs get send to the specified Kinesis Data Firehose delivery stream.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EncryptionInTransit

Iam

All content copied from https://docs.aws.amazon.com/.
