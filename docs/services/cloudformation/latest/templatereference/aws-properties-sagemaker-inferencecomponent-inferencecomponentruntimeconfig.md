This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::InferenceComponent InferenceComponentRuntimeConfig

Runtime settings for a model that is deployed with an inference component.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CopyCount" : Integer,
  "CurrentCopyCount" : Integer,
  "DesiredCopyCount" : Integer
}

```

### YAML

```yaml

  CopyCount: Integer
  CurrentCopyCount: Integer
  DesiredCopyCount: Integer

```

## Properties

`CopyCount`

The number of runtime copies of the model container to deploy with the inference
component. Each copy can serve inference requests.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CurrentCopyCount`

The current number of copies of the model deployed for the inference component.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DesiredCopyCount`

The desired number of copies of the model to deploy for the inference component.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

InferenceComponentRollingUpdatePolicy

InferenceComponentSpecification
