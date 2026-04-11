This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Logs::Transformer ParseWAF

Use this processor to parse AWS WAF vended logs, extract fields, and and
convert them into a JSON format. This processor always processes the entire log event
message. For more information about this processor including examples, see [parseWAF](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation-processors.md#CloudWatch-Logs-Transformation-parsePostGres).

For more information about AWS WAF log format, see [Log\
examples for web ACL traffic](../../../waf/latest/developerguide/logging-examples.md).

###### Important

If you use this processor, it must be the first processor in your
transformer.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Source" : String
}

```

### YAML

```yaml

  Source: String

```

## Properties

`Source`

Omit this parameter and the whole log message will be processed by this processor. No
other value than `@message` is allowed for `source`.

_Required_: No

_Type_: String

_Pattern_: `^.*[a-zA-Z0-9]+.*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ParseVPC

Processor

All content copied from https://docs.aws.amazon.com/.
