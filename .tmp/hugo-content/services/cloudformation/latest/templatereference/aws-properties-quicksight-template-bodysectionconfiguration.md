This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template BodySectionConfiguration

The configuration of a body section.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Content" : BodySectionContent,
  "PageBreakConfiguration" : SectionPageBreakConfiguration,
  "RepeatConfiguration" : BodySectionRepeatConfiguration,
  "SectionId" : String,
  "Style" : SectionStyle
}

```

### YAML

```yaml

  Content:
    BodySectionContent
  PageBreakConfiguration:
    SectionPageBreakConfiguration
  RepeatConfiguration:
    BodySectionRepeatConfiguration
  SectionId: String
  Style:
    SectionStyle

```

## Properties

`Content`

The configuration of content in a body section.

_Required_: Yes

_Type_: [BodySectionContent](aws-properties-quicksight-template-bodysectioncontent.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PageBreakConfiguration`

The configuration of a page break for a section.

_Required_: No

_Type_: [SectionPageBreakConfiguration](aws-properties-quicksight-template-sectionpagebreakconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RepeatConfiguration`

Describes the configurations that are required to declare a section as repeating.

_Required_: No

_Type_: [BodySectionRepeatConfiguration](aws-properties-quicksight-template-bodysectionrepeatconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SectionId`

The unique identifier of a body section.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Style`

The style options of a body section.

_Required_: No

_Type_: [SectionStyle](aws-properties-quicksight-template-sectionstyle.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BinWidthOptions

BodySectionContent

All content copied from https://docs.aws.amazon.com/.
