This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MSK::Cluster CloudWatchLogs

Details of the CloudWatch Logs destination for broker logs.

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

Specifies whether broker logs get sent to the specified CloudWatch Logs destination.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogGroup`

The CloudWatch log group that is the destination for broker logs.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ClientAuthentication

ConfigurationInfo

All content copied from https://docs.aws.amazon.com/.
