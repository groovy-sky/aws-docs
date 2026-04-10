This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot TextLogDestination

Defines the Amazon CloudWatch Logs destination log group for
conversation text logs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CloudWatch" : CloudWatchLogGroupLogDestination
}

```

### YAML

```yaml

  CloudWatch:
    CloudWatchLogGroupLogDestination

```

## Properties

`CloudWatch`

Defines the Amazon CloudWatch Logs log group where text and metadata logs are
delivered.

_Required_: Yes

_Type_: [CloudWatchLogGroupLogDestination](aws-properties-lex-bot-cloudwatchloggrouplogdestination.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TextInputSpecification

TextLogSetting

All content copied from https://docs.aws.amazon.com/.
