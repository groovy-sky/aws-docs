This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template WordCloudChartConfiguration

The configuration of a word cloud visual.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CategoryLabelOptions" : ChartAxisLabelOptions,
  "FieldWells" : WordCloudFieldWells,
  "Interactions" : VisualInteractionOptions,
  "SortConfiguration" : WordCloudSortConfiguration,
  "WordCloudOptions" : WordCloudOptions
}

```

### YAML

```yaml

  CategoryLabelOptions:
    ChartAxisLabelOptions
  FieldWells:
    WordCloudFieldWells
  Interactions:
    VisualInteractionOptions
  SortConfiguration:
    WordCloudSortConfiguration
  WordCloudOptions:
    WordCloudOptions

```

## Properties

`CategoryLabelOptions`

The label options (label text, label visibility, and sort icon visibility) for the word cloud category.

_Required_: No

_Type_: [ChartAxisLabelOptions](aws-properties-quicksight-template-chartaxislabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldWells`

The field wells of the visual.

_Required_: No

_Type_: [WordCloudFieldWells](aws-properties-quicksight-template-wordcloudfieldwells.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Interactions`

The general visual interactions setup for a visual.

_Required_: No

_Type_: [VisualInteractionOptions](aws-properties-quicksight-template-visualinteractionoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SortConfiguration`

The sort configuration of a word cloud visual.

_Required_: No

_Type_: [WordCloudSortConfiguration](aws-properties-quicksight-template-wordcloudsortconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WordCloudOptions`

The options for a word cloud visual.

_Required_: No

_Type_: [WordCloudOptions](aws-properties-quicksight-template-wordcloudoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WordCloudAggregatedFieldWells

WordCloudFieldWells

All content copied from https://docs.aws.amazon.com/.
