This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream DirectPutSourceConfiguration

The structure that configures parameters such as `ThroughputHintInMBs` for a stream configured with
Direct PUT as a source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ThroughputHintInMBs" : Integer
}

```

### YAML

```yaml

  ThroughputHintInMBs: Integer

```

## Properties

`ThroughputHintInMBs`

The value that you configure for this parameter is for information purpose only and
does not affect Firehose delivery throughput limit. You can use the [Firehose Limits form](https://support.console.aws.amazon.com/support/home) to request a throughput limit increase.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DestinationTableConfiguration

DocumentIdOptions

All content copied from https://docs.aws.amazon.com/.
