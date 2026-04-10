This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataBrew::Job CsvOutputOptions

Represents a set of options that define how DataBrew will write a
comma-separated value (CSV) file.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Delimiter" : String
}

```

### YAML

```yaml

  Delimiter: String

```

## Properties

`Delimiter`

A single character that specifies the delimiter used to create CSV job output.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ColumnStatisticsConfiguration

DatabaseOutput

All content copied from https://docs.aws.amazon.com/.
