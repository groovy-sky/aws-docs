This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataAutomationProject ImageStandardExtraction

Settings for generating data from images.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BoundingBox" : ImageBoundingBox,
  "Category" : ImageExtractionCategory
}

```

### YAML

```yaml

  BoundingBox:
    ImageBoundingBox
  Category:
    ImageExtractionCategory

```

## Properties

`BoundingBox`

Settings for generating bounding boxes.

_Required_: Yes

_Type_: [ImageBoundingBox](aws-properties-bedrock-dataautomationproject-imageboundingbox.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Category`

Settings for generating categorical data.

_Required_: Yes

_Type_: [ImageExtractionCategory](aws-properties-bedrock-dataautomationproject-imageextractioncategory.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ImageOverrideConfiguration

ImageStandardGenerativeField

All content copied from https://docs.aws.amazon.com/.
