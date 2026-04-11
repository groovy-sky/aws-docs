This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot CloudWatchLogGroupLogDestination

The Amazon CloudWatch Logs log group where the text and metadata logs are
delivered. The log group must exist before you enable logging.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CloudWatchLogGroupArn" : String,
  "LogPrefix" : String
}

```

### YAML

```yaml

  CloudWatchLogGroupArn: String
  LogPrefix: String

```

## Properties

`CloudWatchLogGroupArn`

The Amazon Resource Name (ARN) of the log group where text and
metadata logs are delivered.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogPrefix`

The prefix of the log stream name within the log group that you
specified

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Button

CodeHookSpecification

All content copied from https://docs.aws.amazon.com/.
