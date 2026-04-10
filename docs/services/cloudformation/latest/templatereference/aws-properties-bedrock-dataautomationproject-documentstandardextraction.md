This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataAutomationProject DocumentStandardExtraction

Settings for generating data from documents.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BoundingBox" : DocumentBoundingBox,
  "Granularity" : DocumentExtractionGranularity
}

```

### YAML

```yaml

  BoundingBox:
    DocumentBoundingBox
  Granularity:
    DocumentExtractionGranularity

```

## Properties

`BoundingBox`

Whether to generate bounding boxes.

_Required_: Yes

_Type_: [DocumentBoundingBox](aws-properties-bedrock-dataautomationproject-documentboundingbox.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Granularity`

Which granularities to generate data for.

_Required_: Yes

_Type_: [DocumentExtractionGranularity](aws-properties-bedrock-dataautomationproject-documentextractiongranularity.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DocumentOverrideConfiguration

DocumentStandardGenerativeField

All content copied from https://docs.aws.amazon.com/.
