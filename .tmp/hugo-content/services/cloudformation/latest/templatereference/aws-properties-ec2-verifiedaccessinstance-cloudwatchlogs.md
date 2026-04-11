This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::VerifiedAccessInstance CloudWatchLogs

Options for CloudWatch Logs as a logging destination.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enabled" : Boolean,
  "LogGroup" : String
}

```

### YAML

```yaml

  Enabled: Boolean
  LogGroup: String

```

## Properties

`Enabled`

Indicates whether logging is enabled.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogGroup`

The ID of the CloudWatch Logs log group.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::EC2::VerifiedAccessInstance

KinesisDataFirehose

All content copied from https://docs.aws.amazon.com/.
