This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Logs::Transformer Grok

This processor uses pattern matching to parse and structure unstructured data. This
processor can also extract fields from log messages.

For more information about this processor including examples, see [grok](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation-processors.md#CloudWatch-Logs-Transformation-Grok) in the _CloudWatch Logs User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Match" : String,
  "Source" : String
}

```

### YAML

```yaml

  Match: String
  Source: String

```

## Properties

`Match`

The grok pattern to match against the log event. For a list of supported grok
patterns, see [Supported grok patterns](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation-processors.md#Grok-Patterns).

_Required_: Yes

_Type_: String

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Source`

The path to the field in the log event that you want to parse. If you omit this value,
the whole log message is parsed.

_Required_: No

_Type_: String

_Pattern_: `^.*[a-zA-Z0-9]+.*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteKeys

ListToMap

All content copied from https://docs.aws.amazon.com/.
