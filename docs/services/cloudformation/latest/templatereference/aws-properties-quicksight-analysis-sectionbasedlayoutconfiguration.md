This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis SectionBasedLayoutConfiguration

The configuration for a
section-based layout.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BodySections" : [ BodySectionConfiguration, ... ],
  "CanvasSizeOptions" : SectionBasedLayoutCanvasSizeOptions,
  "FooterSections" : [ HeaderFooterSectionConfiguration, ... ],
  "HeaderSections" : [ HeaderFooterSectionConfiguration, ... ]
}

```

### YAML

```yaml

  BodySections:
    - BodySectionConfiguration
  CanvasSizeOptions:
    SectionBasedLayoutCanvasSizeOptions
  FooterSections:
    - HeaderFooterSectionConfiguration
  HeaderSections:
    - HeaderFooterSectionConfiguration

```

## Properties

`BodySections`

A list of body section configurations.

_Required_: Yes

_Type_: Array of [BodySectionConfiguration](aws-properties-quicksight-analysis-bodysectionconfiguration.md)

_Minimum_: `0`

_Maximum_: `28`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CanvasSizeOptions`

The options for the canvas of a section-based layout.

_Required_: Yes

_Type_: [SectionBasedLayoutCanvasSizeOptions](aws-properties-quicksight-analysis-sectionbasedlayoutcanvassizeoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FooterSections`

A list of footer section configurations.

_Required_: Yes

_Type_: Array of [HeaderFooterSectionConfiguration](aws-properties-quicksight-analysis-headerfootersectionconfiguration.md)

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HeaderSections`

A list of header section configurations.

_Required_: Yes

_Type_: Array of [HeaderFooterSectionConfiguration](aws-properties-quicksight-analysis-headerfootersectionconfiguration.md)

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SectionBasedLayoutCanvasSizeOptions

SectionBasedLayoutPaperCanvasSizeOptions

All content copied from https://docs.aws.amazon.com/.
