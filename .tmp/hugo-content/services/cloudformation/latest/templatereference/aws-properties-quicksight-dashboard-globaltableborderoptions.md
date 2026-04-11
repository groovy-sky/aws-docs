This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard GlobalTableBorderOptions

Determines the border options for a table visual.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SideSpecificBorder" : TableSideBorderOptions,
  "UniformBorder" : TableBorderOptions
}

```

### YAML

```yaml

  SideSpecificBorder:
    TableSideBorderOptions
  UniformBorder:
    TableBorderOptions

```

## Properties

`SideSpecificBorder`

Determines the options for side specific border.

_Required_: No

_Type_: [TableSideBorderOptions](aws-properties-quicksight-dashboard-tablesideborderoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UniformBorder`

Determines the options for uniform border.

_Required_: No

_Type_: [TableBorderOptions](aws-properties-quicksight-dashboard-tableborderoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GeospatialWindowOptions

GradientColor

All content copied from https://docs.aws.amazon.com/.
