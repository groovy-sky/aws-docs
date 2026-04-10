This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Logs::Transformer Csv

The `CSV` processor parses comma-separated values (CSV) from the log events
into columns.

For more information about this processor including examples, see [csv](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-csv) in the _CloudWatch Logs User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Columns" : [ String, ... ],
  "Delimiter" : String,
  "QuoteCharacter" : String,
  "Source" : String
}

```

### YAML

```yaml

  Columns:
    - String
  Delimiter: String
  QuoteCharacter: String
  Source: String

```

## Properties

`Columns`

An array of names to use for the columns in the transformed log event.

If you omit this, default column names ( `[column_1, column_2 ...]`) are
used.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Delimiter`

The character used to separate each column in the original comma-separated value log
event. If you omit this, the processor looks for the comma `,` character as
the delimiter.

_Required_: No

_Type_: String

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QuoteCharacter`

The character used used as a text qualifier for a single column of data. If you omit
this, the double quotation mark `"` character is used.

_Required_: No

_Type_: String

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Source`

The path to the field in the log event that has the comma separated values to be
parsed. If you omit this value, the whole log message is processed.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CopyValueEntry

DateTimeConverter

All content copied from https://docs.aws.amazon.com/.
