This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelCard InferenceSpecification

Defines how to perform inference generation after a training job is run.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Containers" : [ Container, ... ]
}

```

### YAML

```yaml

  Containers:
    - Container

```

## Properties

`Containers`

The Amazon ECR registry path of the Docker image that contains the inference code.

_Required_: Yes

_Type_: Array of [Container](aws-properties-sagemaker-modelcard-container.md)

_Minimum_: `1`

_Maximum_: `15`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InferenceEnvironment

IntendedUses

All content copied from https://docs.aws.amazon.com/.
