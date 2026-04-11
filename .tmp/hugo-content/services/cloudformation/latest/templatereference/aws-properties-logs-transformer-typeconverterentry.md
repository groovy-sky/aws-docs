This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Logs::Transformer TypeConverterEntry

This object defines one value type that will be converted using the [typeConverter](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation-processors.md#CloudWatch-Logs-Transformation-typeConverter) processor.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Type" : String
}

```

### YAML

```yaml

  Key: String
  Type: String

```

## Properties

`Key`

The key with the value that is to be converted to a different type.

_Required_: Yes

_Type_: String

_Pattern_: `^.*[a-zA-Z0-9]+.*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type to convert the field value to. Valid values are `integer`,
`double`, `string` and `boolean`.

_Required_: Yes

_Type_: String

_Allowed values_: `boolean | integer | double | string`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TypeConverter

UpperCaseString

All content copied from https://docs.aws.amazon.com/.
