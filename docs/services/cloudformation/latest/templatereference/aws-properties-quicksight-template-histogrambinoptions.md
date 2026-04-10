This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template HistogramBinOptions

The options that determine the presentation of histogram bins.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BinCount" : BinCountOptions,
  "BinWidth" : BinWidthOptions,
  "SelectedBinType" : String,
  "StartValue" : Number
}

```

### YAML

```yaml

  BinCount:
    BinCountOptions
  BinWidth:
    BinWidthOptions
  SelectedBinType: String
  StartValue: Number

```

## Properties

`BinCount`

The options that determine the bin count of a histogram.

_Required_: No

_Type_: [BinCountOptions](aws-properties-quicksight-template-bincountoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BinWidth`

The options that determine the bin width of a histogram.

_Required_: No

_Type_: [BinWidthOptions](aws-properties-quicksight-template-binwidthoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelectedBinType`

The options that determine the selected bin type.

_Required_: No

_Type_: String

_Allowed values_: `BIN_COUNT | BIN_WIDTH`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartValue`

The options that determine the bin start value.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HistogramAggregatedFieldWells

HistogramConfiguration

All content copied from https://docs.aws.amazon.com/.
