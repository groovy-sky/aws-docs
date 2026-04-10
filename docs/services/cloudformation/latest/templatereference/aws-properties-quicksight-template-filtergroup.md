This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template FilterGroup

A grouping of individual filters. Filter groups are applied to the same group of visuals.

For more information, see [Adding filter conditions (group filters) with AND and OR operators](../../../quicksight/latest/user/add-a-compound-filter.md) in the _Amazon Quick Suite User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CrossDataset" : String,
  "FilterGroupId" : String,
  "Filters" : [ Filter, ... ],
  "ScopeConfiguration" : FilterScopeConfiguration,
  "Status" : String
}

```

### YAML

```yaml

  CrossDataset: String
  FilterGroupId: String
  Filters:
    - Filter
  ScopeConfiguration:
    FilterScopeConfiguration
  Status: String

```

## Properties

`CrossDataset`

The filter new feature which can apply filter group to all data sets. Choose one of the following options:

- `ALL_DATASETS`

- `SINGLE_DATASET`

_Required_: Yes

_Type_: String

_Allowed values_: `ALL_DATASETS | SINGLE_DATASET`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterGroupId`

The value that uniquely identifies a `FilterGroup` within a dashboard, template, or analysis.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Filters`

The list of filters that are present in a `FilterGroup`.

_Required_: Yes

_Type_: Array of [Filter](aws-properties-quicksight-template-filter.md)

_Minimum_: `0`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScopeConfiguration`

The configuration that specifies what scope to apply to a `FilterGroup`.

This is a union type structure. For this structure to be valid, only one of the attributes can be defined.

_Required_: Yes

_Type_: [FilterScopeConfiguration](aws-properties-quicksight-template-filterscopeconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

The status of the `FilterGroup`.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FilterDropDownControl

FilterListConfiguration

All content copied from https://docs.aws.amazon.com/.
