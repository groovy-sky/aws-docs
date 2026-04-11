This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard CascadingControlSource

The source controls that are used in a `CascadingControlConfiguration`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ColumnToMatch" : ColumnIdentifier,
  "SourceSheetControlId" : String
}

```

### YAML

```yaml

  ColumnToMatch:
    ColumnIdentifier
  SourceSheetControlId: String

```

## Properties

`ColumnToMatch`

The column identifier that determines which column to look up for the source sheet control.

_Required_: No

_Type_: [ColumnIdentifier](aws-properties-quicksight-dashboard-columnidentifier.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceSheetControlId`

The source sheet control ID of a `CascadingControlSource`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CascadingControlConfiguration

CategoricalDimensionField

All content copied from https://docs.aws.amazon.com/.
