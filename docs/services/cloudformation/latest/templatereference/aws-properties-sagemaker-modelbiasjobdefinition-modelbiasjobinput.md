This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelBiasJobDefinition ModelBiasJobInput

Inputs for the model bias job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BatchTransformInput" : BatchTransformInput,
  "EndpointInput" : EndpointInput,
  "GroundTruthS3Input" : MonitoringGroundTruthS3Input
}

```

### YAML

```yaml

  BatchTransformInput:
    BatchTransformInput
  EndpointInput:
    EndpointInput
  GroundTruthS3Input:
    MonitoringGroundTruthS3Input

```

## Properties

`BatchTransformInput`

Input object for the batch transform job.

_Required_: No

_Type_: [BatchTransformInput](aws-properties-sagemaker-modelbiasjobdefinition-batchtransforminput.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EndpointInput`

Input object for the endpoint

_Required_: No

_Type_: [EndpointInput](aws-properties-sagemaker-modelbiasjobdefinition-endpointinput.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`GroundTruthS3Input`

Location of ground truth labels to use in model bias job.

_Required_: Yes

_Type_: [MonitoringGroundTruthS3Input](aws-properties-sagemaker-modelbiasjobdefinition-monitoringgroundtruths3input.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ModelBiasBaselineConfig

MonitoringGroundTruthS3Input

All content copied from https://docs.aws.amazon.com/.
