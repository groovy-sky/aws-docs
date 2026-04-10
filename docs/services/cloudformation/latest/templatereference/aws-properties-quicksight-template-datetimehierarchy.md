This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template DateTimeHierarchy

The option that determines the hierarchy of any `DateTime` fields.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DrillDownFilters" : [ DrillDownFilter, ... ],
  "HierarchyId" : String
}

```

### YAML

```yaml

  DrillDownFilters:
    - DrillDownFilter
  HierarchyId: String

```

## Properties

`DrillDownFilters`

The option that determines the drill down filters for the
`DateTime` hierarchy.

_Required_: No

_Type_: Array of [DrillDownFilter](aws-properties-quicksight-template-drilldownfilter.md)

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HierarchyId`

The hierarchy ID of the `DateTime` hierarchy.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DateTimeFormatConfiguration

DateTimeParameterDeclaration

All content copied from https://docs.aws.amazon.com/.
