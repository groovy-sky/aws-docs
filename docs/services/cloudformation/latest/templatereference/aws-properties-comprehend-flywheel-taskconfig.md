This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Comprehend::Flywheel TaskConfig

Configuration about the model associated with a flywheel.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DocumentClassificationConfig" : DocumentClassificationConfig,
  "EntityRecognitionConfig" : EntityRecognitionConfig,
  "LanguageCode" : String
}

```

### YAML

```yaml

  DocumentClassificationConfig:
    DocumentClassificationConfig
  EntityRecognitionConfig:
    EntityRecognitionConfig
  LanguageCode: String

```

## Properties

`DocumentClassificationConfig`

Configuration required for a document classification model.

_Required_: No

_Type_: [DocumentClassificationConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-comprehend-flywheel-documentclassificationconfig.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EntityRecognitionConfig`

Configuration required for an entity recognition model.

_Required_: No

_Type_: [EntityRecognitionConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-comprehend-flywheel-entityrecognitionconfig.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LanguageCode`

Language code for the language that the model supports.

_Required_: Yes

_Type_: String

_Allowed values_: `en | es | fr | it | de | pt`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

VpcConfig
