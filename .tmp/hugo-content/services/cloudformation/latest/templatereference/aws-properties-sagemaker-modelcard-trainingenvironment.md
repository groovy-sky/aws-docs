This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelCard TrainingEnvironment

SageMaker AI training image.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContainerImage" : [ String, ... ]
}

```

### YAML

```yaml

  ContainerImage:
    - String

```

## Properties

`ContainerImage`

SageMaker AI inference image URI.

_Required_: No

_Type_: Array of String

_Maximum_: `1024 | 15`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TrainingDetails

TrainingHyperParameter

All content copied from https://docs.aws.amazon.com/.
