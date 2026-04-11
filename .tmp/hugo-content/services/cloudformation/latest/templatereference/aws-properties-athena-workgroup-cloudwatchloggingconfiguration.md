This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Athena::WorkGroup CloudWatchLoggingConfiguration

Configuration settings for delivering logs to Amazon CloudWatch log groups.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enabled" : Boolean,
  "LogGroup" : String,
  "LogStreamNamePrefix" : String,
  "LogTypes" : {Key: Value, ...}
}

```

### YAML

```yaml

  Enabled: Boolean
  LogGroup: String
  LogStreamNamePrefix: String
  LogTypes:
    Key: Value

```

## Properties

`Enabled`

Enables CloudWatch logging.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogGroup`

The name of the log group in Amazon CloudWatch Logs where you want to publish
your logs.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9._/-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogStreamNamePrefix`

Prefix for the CloudWatch log stream name.

_Required_: No

_Type_: String

_Pattern_: `^[^:*]*$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogTypes`

The types of logs that you want to publish to CloudWatch.

_Required_: No

_Type_: Object of Array

_Pattern_: `^.+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Classification

CustomerContentEncryptionConfiguration

All content copied from https://docs.aws.amazon.com/.
