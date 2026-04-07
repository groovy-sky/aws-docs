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

_Type_: [BatchTransformInput](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-modelbiasjobdefinition-batchtransforminput.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EndpointInput`

Input object for the endpoint

_Required_: No

_Type_: [EndpointInput](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-modelbiasjobdefinition-endpointinput.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`GroundTruthS3Input`

Location of ground truth labels to use in model bias job.

_Required_: Yes

_Type_: [MonitoringGroundTruthS3Input](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-modelbiasjobdefinition-monitoringgroundtruths3input.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ModelBiasBaselineConfig

MonitoringGroundTruthS3Input
