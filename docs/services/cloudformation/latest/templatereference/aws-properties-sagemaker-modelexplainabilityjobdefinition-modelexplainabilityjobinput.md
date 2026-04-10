This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelExplainabilityJobDefinition ModelExplainabilityJobInput

Inputs for the model explainability job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BatchTransformInput" : BatchTransformInput,
  "EndpointInput" : EndpointInput
}

```

### YAML

```yaml

  BatchTransformInput:
    BatchTransformInput
  EndpointInput:
    EndpointInput

```

## Properties

`BatchTransformInput`

Input object for the batch transform job.

_Required_: No

_Type_: [BatchTransformInput](aws-properties-sagemaker-modelexplainabilityjobdefinition-batchtransforminput.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EndpointInput`

Input object for the endpoint

_Required_: No

_Type_: [EndpointInput](aws-properties-sagemaker-modelexplainabilityjobdefinition-endpointinput.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ModelExplainabilityBaselineConfig

MonitoringOutput

All content copied from https://docs.aws.amazon.com/.
