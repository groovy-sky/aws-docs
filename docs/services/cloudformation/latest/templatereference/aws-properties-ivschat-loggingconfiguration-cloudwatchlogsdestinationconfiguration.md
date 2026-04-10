This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IVSChat::LoggingConfiguration CloudWatchLogsDestinationConfiguration

The CloudWatchLogsDestinationConfiguration property type specifies a CloudWatch Logs location where chat logs will be stored.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LogGroupName" : String
}

```

### YAML

```yaml

  LogGroupName: String

```

## Properties

`LogGroupName`

Name of the Amazon Cloudwatch Logs destination where chat activity will be logged.

_Required_: Yes

_Type_: String

_Pattern_: `^[\.\-_/#A-Za-z0-9]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::IVSChat::LoggingConfiguration

DestinationConfiguration

All content copied from https://docs.aws.amazon.com/.
