This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet OutputColumnNameOverride

Specifies a mapping to override the name of an output column from a transform operation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OutputColumnName" : String,
  "SourceColumnName" : String
}

```

### YAML

```yaml

  OutputColumnName: String
  SourceColumnName: String

```

## Properties

`OutputColumnName`

The new name to assign to the column in the output.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceColumnName`

The original name of the column from the source transform operation.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OutputColumn

OverrideDatasetParameterOperation

All content copied from https://docs.aws.amazon.com/.
