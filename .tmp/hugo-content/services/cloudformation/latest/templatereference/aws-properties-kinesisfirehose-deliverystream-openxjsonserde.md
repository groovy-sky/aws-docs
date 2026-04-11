This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream OpenXJsonSerDe

The OpenX SerDe. Used by Firehose for deserializing data, which means
converting it from the JSON format in preparation for serializing it to the Parquet or ORC
format. This is one of two deserializers you can choose, depending on which one offers the
functionality you need. The other option is the native Hive / HCatalog JsonSerDe.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CaseInsensitive" : Boolean,
  "ColumnToJsonKeyMappings" : {Key: Value, ...},
  "ConvertDotsInJsonKeysToUnderscores" : Boolean
}

```

### YAML

```yaml

  CaseInsensitive: Boolean
  ColumnToJsonKeyMappings:
    Key: Value
  ConvertDotsInJsonKeysToUnderscores: Boolean

```

## Properties

`CaseInsensitive`

When set to `true`, which is the default, Firehose converts
JSON keys to lowercase before deserializing them.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColumnToJsonKeyMappings`

Maps column names to JSON keys that aren't identical to the column names. This is
useful when the JSON contains keys that are Hive keywords. For example,
`timestamp` is a Hive keyword. If you have a JSON key named
`timestamp`, set this parameter to `{"ts": "timestamp"}` to map
this key to a column named `ts`.

_Required_: No

_Type_: Object of String

_Pattern_: `[a-zA-Z0-9]+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConvertDotsInJsonKeysToUnderscores`

When set to `true`, specifies that the names of the keys include dots and
that you want Firehose to replace them with underscores. This is useful
because Apache Hive does not allow dots in column names. For example, if the JSON contains
a key whose name is "a.b", you can define the column name to be "a\_b" when using this
option.

The default is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MSKSourceConfiguration

OrcSerDe

All content copied from https://docs.aws.amazon.com/.
