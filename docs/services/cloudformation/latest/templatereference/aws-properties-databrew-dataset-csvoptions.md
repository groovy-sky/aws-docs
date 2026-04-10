This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataBrew::Dataset CsvOptions

Represents a set of options that define how DataBrew will read a
comma-separated value (CSV) file when creating a dataset from that file.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Delimiter" : String,
  "HeaderRow" : Boolean
}

```

### YAML

```yaml

  Delimiter: String
  HeaderRow: Boolean

```

## Properties

`Delimiter`

A single character that specifies the delimiter being used in the CSV file.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HeaderRow`

A variable that specifies whether the first row in the file is parsed as the
header. If this value is false, column names are auto-generated.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::DataBrew::Dataset

DatabaseInputDefinition

All content copied from https://docs.aws.amazon.com/.
