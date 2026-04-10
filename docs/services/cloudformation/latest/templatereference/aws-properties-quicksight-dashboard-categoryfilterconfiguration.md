This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard CategoryFilterConfiguration

The configuration for a `CategoryFilter`.

This is a union type structure. For this structure to be valid, only one of the attributes can be defined.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomFilterConfiguration" : CustomFilterConfiguration,
  "CustomFilterListConfiguration" : CustomFilterListConfiguration,
  "FilterListConfiguration" : FilterListConfiguration
}

```

### YAML

```yaml

  CustomFilterConfiguration:
    CustomFilterConfiguration
  CustomFilterListConfiguration:
    CustomFilterListConfiguration
  FilterListConfiguration:
    FilterListConfiguration

```

## Properties

`CustomFilterConfiguration`

A custom filter that filters based on a single value. This filter can be partially matched.

_Required_: No

_Type_: [CustomFilterConfiguration](aws-properties-quicksight-dashboard-customfilterconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomFilterListConfiguration`

A list of custom filter values. In the Quick Sight console, this filter type is called a custom filter list.

_Required_: No

_Type_: [CustomFilterListConfiguration](aws-properties-quicksight-dashboard-customfilterlistconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterListConfiguration`

A list of filter configurations. In the Quick Sight console, this filter type is called a filter list.

_Required_: No

_Type_: [FilterListConfiguration](aws-properties-quicksight-dashboard-filterlistconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CategoryFilter

CategoryInnerFilter

All content copied from https://docs.aws.amazon.com/.
