This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataAutomationProject VideoStandardExtraction

Settings for generating data from video.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BoundingBox" : VideoBoundingBox,
  "Category" : VideoExtractionCategory
}

```

### YAML

```yaml

  BoundingBox:
    VideoBoundingBox
  Category:
    VideoExtractionCategory

```

## Properties

`BoundingBox`

Settings for generating bounding boxes.

_Required_: Yes

_Type_: [VideoBoundingBox](aws-properties-bedrock-dataautomationproject-videoboundingbox.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Category`

Settings for generating categorical data.

_Required_: Yes

_Type_: [VideoExtractionCategory](aws-properties-bedrock-dataautomationproject-videoextractioncategory.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VideoOverrideConfiguration

VideoStandardGenerativeField

All content copied from https://docs.aws.amazon.com/.
