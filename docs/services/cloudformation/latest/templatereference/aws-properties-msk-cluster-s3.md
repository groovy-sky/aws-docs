This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MSK::Cluster S3

The details of the Amazon S3 destination for broker logs.

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

The name of the S3 bucket that is the destination for broker logs.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

Specifies whether broker logs get sent to the specified Amazon S3 destination.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Prefix`

The S3 prefix that is the destination for broker logs.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Rebalancing

Sasl

All content copied from https://docs.aws.amazon.com/.
