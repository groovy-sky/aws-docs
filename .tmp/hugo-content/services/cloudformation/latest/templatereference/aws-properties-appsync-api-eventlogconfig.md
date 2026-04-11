This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppSync::Api EventLogConfig

Describes the CloudWatch Logs configuration for the Event API.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CloudWatchLogsRoleArn" : String,
  "LogLevel" : String
}

```

### YAML

```yaml

  CloudWatchLogsRoleArn: String
  LogLevel: String

```

## Properties

`CloudWatchLogsRoleArn`

The IAM service role that AWS AppSync assumes to publish
CloudWatch Logs in your account.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogLevel`

The type of information to log for the Event API.

_Required_: Yes

_Type_: String

_Allowed values_: `NONE | ERROR | ALL | INFO | DEBUG`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EventConfig

LambdaAuthorizerConfig

All content copied from https://docs.aws.amazon.com/.
