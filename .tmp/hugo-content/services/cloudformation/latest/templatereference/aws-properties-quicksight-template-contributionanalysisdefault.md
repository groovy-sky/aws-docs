This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template ContributionAnalysisDefault

The contribution analysis visual display for a line, pie, or bar chart.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContributorDimensions" : [ ColumnIdentifier, ... ],
  "MeasureFieldId" : String
}

```

### YAML

```yaml

  ContributorDimensions:
    - ColumnIdentifier
  MeasureFieldId: String

```

## Properties

`ContributorDimensions`

The dimensions columns that are used in the contribution analysis,
usually a list of `ColumnIdentifiers`.

_Required_: Yes

_Type_: Array of [ColumnIdentifier](aws-properties-quicksight-template-columnidentifier.md)

_Minimum_: `1`

_Maximum_: `4`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MeasureFieldId`

The measure field that is used in the contribution analysis.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ContextMenuOption

CurrencyDisplayFormatConfiguration

All content copied from https://docs.aws.amazon.com/.
