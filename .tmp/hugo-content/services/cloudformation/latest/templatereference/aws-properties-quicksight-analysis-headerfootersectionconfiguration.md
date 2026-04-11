This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis HeaderFooterSectionConfiguration

The configuration of a header or footer section.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Layout" : SectionLayoutConfiguration,
  "SectionId" : String,
  "Style" : SectionStyle
}

```

### YAML

```yaml

  Layout:
    SectionLayoutConfiguration
  SectionId: String
  Style:
    SectionStyle

```

## Properties

`Layout`

The layout configuration of the header or footer section.

_Required_: Yes

_Type_: [SectionLayoutConfiguration](aws-properties-quicksight-analysis-sectionlayoutconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SectionId`

The unique identifier of the header or footer section.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Style`

The style options of a header or footer section.

_Required_: No

_Type_: [SectionStyle](aws-properties-quicksight-analysis-sectionstyle.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GrowthRateComputation

HeatMapAggregatedFieldWells

All content copied from https://docs.aws.amazon.com/.
