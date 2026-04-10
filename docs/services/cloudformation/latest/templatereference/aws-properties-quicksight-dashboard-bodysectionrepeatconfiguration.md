This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard BodySectionRepeatConfiguration

Describes the configurations that are required to declare a section as repeating.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DimensionConfigurations" : [ BodySectionRepeatDimensionConfiguration, ... ],
  "NonRepeatingVisuals" : [ String, ... ],
  "PageBreakConfiguration" : BodySectionRepeatPageBreakConfiguration
}

```

### YAML

```yaml

  DimensionConfigurations:
    - BodySectionRepeatDimensionConfiguration
  NonRepeatingVisuals:
    - String
  PageBreakConfiguration:
    BodySectionRepeatPageBreakConfiguration

```

## Properties

`DimensionConfigurations`

List of `BodySectionRepeatDimensionConfiguration` values that describe the dataset column and constraints for the column used to repeat the contents of a section.

_Required_: No

_Type_: Array of [BodySectionRepeatDimensionConfiguration](aws-properties-quicksight-dashboard-bodysectionrepeatdimensionconfiguration.md)

_Minimum_: `0`

_Maximum_: `3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NonRepeatingVisuals`

List of visuals to exclude from repetition in repeating sections. The visuals will render identically, and ignore the repeating configurations in all repeating instances.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 0`

_Maximum_: `512 | 20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PageBreakConfiguration`

Page break configuration to apply for each repeating instance.

_Required_: No

_Type_: [BodySectionRepeatPageBreakConfiguration](aws-properties-quicksight-dashboard-bodysectionrepeatpagebreakconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BodySectionDynamicNumericDimensionConfiguration

BodySectionRepeatDimensionConfiguration

All content copied from https://docs.aws.amazon.com/.
