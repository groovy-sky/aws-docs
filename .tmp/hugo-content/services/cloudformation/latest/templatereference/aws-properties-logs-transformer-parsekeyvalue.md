This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Logs::Transformer ParseKeyValue

This processor parses a specified field in the original log event into key-value
pairs.

For more information about this processor including examples, see [parseKeyValue](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-parseKeyValue) in the _CloudWatch Logs User_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Destination" : String,
  "FieldDelimiter" : String,
  "KeyPrefix" : String,
  "KeyValueDelimiter" : String,
  "NonMatchValue" : String,
  "OverwriteIfExists" : Boolean,
  "Source" : String
}

```

### YAML

```yaml

  Destination: String
  FieldDelimiter: String
  KeyPrefix: String
  KeyValueDelimiter: String
  NonMatchValue: String
  OverwriteIfExists: Boolean
  Source: String

```

## Properties

`Destination`

The destination field to put the extracted key-value pairs into

_Required_: No

_Type_: String

_Pattern_: `^.*[a-zA-Z0-9]+.*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldDelimiter`

The field delimiter string that is used between key-value pairs in the original log
events. If you omit this, the ampersand `&` character is used.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KeyPrefix`

If you want to add a prefix to all transformed keys, specify it here.

_Required_: No

_Type_: String

_Pattern_: `^.*[a-zA-Z0-9]+.*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KeyValueDelimiter`

The delimiter string to use between the key and value in each pair in the transformed
log event.

If you omit this, the equal `=` character is used.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NonMatchValue`

A value to insert into the value field in the result, when a key-value pair is not
successfully split.

_Required_: No

_Type_: String

_Pattern_: `^.*[a-zA-Z0-9]+.*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OverwriteIfExists`

Specifies whether to overwrite the value if the destination key already exists. If you
omit this, the default is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Source`

Path to the field in the log event that will be parsed. Use dot notation to access
child fields. For example, `store.book`

_Required_: No

_Type_: String

_Pattern_: `^.*[a-zA-Z0-9]+.*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ParseJSON

ParsePostgres

All content copied from https://docs.aws.amazon.com/.
