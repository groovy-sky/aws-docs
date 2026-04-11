This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template BodySectionRepeatDimensionConfiguration

Describes the dataset column and constraints for the dynamic values used to repeat the contents of a section. The dataset column is either **Category** or **Numeric** column configuration

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DynamicCategoryDimensionConfiguration" : BodySectionDynamicCategoryDimensionConfiguration,
  "DynamicNumericDimensionConfiguration" : BodySectionDynamicNumericDimensionConfiguration
}

```

### YAML

```yaml

  DynamicCategoryDimensionConfiguration:
    BodySectionDynamicCategoryDimensionConfiguration
  DynamicNumericDimensionConfiguration:
    BodySectionDynamicNumericDimensionConfiguration

```

## Properties

`DynamicCategoryDimensionConfiguration`

Describes the **Category** dataset column and constraints around the dynamic values that will be used in repeating the section contents.

_Required_: No

_Type_: [BodySectionDynamicCategoryDimensionConfiguration](aws-properties-quicksight-template-bodysectiondynamiccategorydimensionconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DynamicNumericDimensionConfiguration`

Describes the **Numeric** dataset column and constraints around the dynamic values used to repeat the contents of a section.

_Required_: No

_Type_: [BodySectionDynamicNumericDimensionConfiguration](aws-properties-quicksight-template-bodysectiondynamicnumericdimensionconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BodySectionRepeatConfiguration

BodySectionRepeatPageBreakConfiguration

All content copied from https://docs.aws.amazon.com/.
