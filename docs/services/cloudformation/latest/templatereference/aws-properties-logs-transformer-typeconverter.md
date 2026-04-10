This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Logs::Transformer TypeConverter

Use this processor to convert a value type associated with the specified key to the
specified type. It's a casting processor that changes the types of the specified fields.
Values can be converted into one of the following datatypes: `integer`,
`double`, `string` and `boolean`.

For more information about this processor including examples, see [trimString](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation-processors.md#CloudWatch-Logs-Transformation-trimString) in the _CloudWatch Logs User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Entries" : [ TypeConverterEntry, ... ]
}

```

### YAML

```yaml

  Entries:
    - TypeConverterEntry

```

## Properties

`Entries`

An array of `TypeConverterEntry` objects, where each object contains the
information about one field to change the type of.

_Required_: Yes

_Type_: Array of [TypeConverterEntry](aws-properties-logs-transformer-typeconverterentry.md)

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TrimString

TypeConverterEntry

All content copied from https://docs.aws.amazon.com/.
