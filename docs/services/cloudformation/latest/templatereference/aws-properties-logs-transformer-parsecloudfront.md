This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Logs::Transformer ParseCloudfront

This processor parses CloudFront vended logs, extract fields, and convert them into
JSON format. Encoded field values are decoded. Values that are integers and doubles are
treated as such. For more information about this processor including examples, see
[parseCloudfront](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation-processors.md#CloudWatch-Logs-Transformation-parseCloudfront)

For more information about CloudFront log format, see [Configure and use\
standard logs (access logs)](../../../amazoncloudfront/latest/developerguide/accesslogs.md).

If you use this processor, it must be the first processor in your transformer.

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

MoveKeys

ParseJSON

All content copied from https://docs.aws.amazon.com/.
