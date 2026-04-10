This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::MonitoringSchedule MonitoringInput

The inputs for a monitoring job.

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

_Type_: [BatchTransformInput](aws-properties-sagemaker-monitoringschedule-batchtransforminput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EndpointInput`

The endpoint for a monitoring job.

_Required_: No

_Type_: [EndpointInput](aws-properties-sagemaker-monitoringschedule-endpointinput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MonitoringExecutionSummary

MonitoringJobDefinition

All content copied from https://docs.aws.amazon.com/.
