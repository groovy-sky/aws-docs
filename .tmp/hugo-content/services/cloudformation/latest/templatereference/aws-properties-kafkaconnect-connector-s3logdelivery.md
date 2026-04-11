This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KafkaConnect::Connector S3LogDelivery

Details about delivering logs to Amazon S3.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Bucket" : String,
  "Enabled" : Boolean,
  "Prefix" : String
}

```

### YAML

```yaml

  Bucket: String
  Enabled: Boolean
  Prefix: String

```

## Properties

`Bucket`

The name of the S3 bucket that is the destination for log delivery.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Enabled`

Specifies whether connector logs get sent to the specified Amazon S3 destination.

_Required_: Yes

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Prefix`

The S3 prefix that is the destination for log delivery.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ProvisionedCapacity

ScaleInPolicy

All content copied from https://docs.aws.amazon.com/.
