This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template ExplicitHierarchy

The option that determines the hierarchy of the fields that are built within a visual's field wells. These fields can't be duplicated to other visuals.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Columns" : [ ColumnIdentifier, ... ],
  "DrillDownFilters" : [ DrillDownFilter, ... ],
  "HierarchyId" : String
}

```

### YAML

```yaml

  Columns:
    - ColumnIdentifier
  DrillDownFilters:
    - DrillDownFilter
  HierarchyId: String

```

## Properties

`Columns`

The list of columns that define the explicit hierarchy.

_Required_: Yes

_Type_: Array of [ColumnIdentifier](aws-properties-quicksight-template-columnidentifier.md)

_Minimum_: `2`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DrillDownFilters`

The option that determines the drill down filters for the explicit hierarchy.

_Required_: No

_Type_: Array of [DrillDownFilter](aws-properties-quicksight-template-drilldownfilter.md)

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HierarchyId`

The hierarchy ID of the explicit hierarchy.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExcludePeriodConfiguration

FieldBasedTooltip

All content copied from https://docs.aws.amazon.com/.
