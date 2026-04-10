This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard ComboChartFieldWells

The field wells of the visual.

This is a union type structure. For this structure to be valid, only one of the attributes can be defined.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ComboChartAggregatedFieldWells" : ComboChartAggregatedFieldWells
}

```

### YAML

```yaml

  ComboChartAggregatedFieldWells:
    ComboChartAggregatedFieldWells

```

## Properties

`ComboChartAggregatedFieldWells`

The aggregated field wells of a combo chart. Combo charts only have aggregated field wells. Columns in a combo chart are aggregated by category.

_Required_: No

_Type_: [ComboChartAggregatedFieldWells](aws-properties-quicksight-dashboard-combochartaggregatedfieldwells.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ComboChartConfiguration

ComboChartSortConfiguration

All content copied from https://docs.aws.amazon.com/.
