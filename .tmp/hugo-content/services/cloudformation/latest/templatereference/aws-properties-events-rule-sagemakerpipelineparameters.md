This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Events::Rule SageMakerPipelineParameters

These are custom parameters to use when the target is a SageMaker AI Model Building
Pipeline that starts based on EventBridge events.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PipelineParameterList" : [ SageMakerPipelineParameter, ... ]
}

```

### YAML

```yaml

  PipelineParameterList:
    - SageMakerPipelineParameter

```

## Properties

`PipelineParameterList`

List of Parameter names and values for SageMaker AI Model Building Pipeline
execution.

_Required_: No

_Type_: Array of [SageMakerPipelineParameter](aws-properties-events-rule-sagemakerpipelineparameter.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SageMakerPipelineParameter

SqsParameters

All content copied from https://docs.aws.amazon.com/.
