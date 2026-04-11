This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::InferenceComponent InferenceComponentRollingUpdatePolicy

Specifies a rolling deployment strategy for updating a SageMaker AI inference
component.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaximumBatchSize" : InferenceComponentCapacitySize,
  "MaximumExecutionTimeoutInSeconds" : Integer,
  "RollbackMaximumBatchSize" : InferenceComponentCapacitySize,
  "WaitIntervalInSeconds" : Integer
}

```

### YAML

```yaml

  MaximumBatchSize:
    InferenceComponentCapacitySize
  MaximumExecutionTimeoutInSeconds: Integer
  RollbackMaximumBatchSize:
    InferenceComponentCapacitySize
  WaitIntervalInSeconds: Integer

```

## Properties

`MaximumBatchSize`

The batch size for each rolling step in the deployment process. For each step, SageMaker AI provisions capacity on the new endpoint fleet, routes traffic to that fleet,
and terminates capacity on the old endpoint fleet. The value must be between 5% to 50% of
the copy count of the inference component.

_Required_: No

_Type_: [InferenceComponentCapacitySize](aws-properties-sagemaker-inferencecomponent-inferencecomponentcapacitysize.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaximumExecutionTimeoutInSeconds`

The time limit for the total deployment. Exceeding this limit causes a timeout.

_Required_: No

_Type_: Integer

_Minimum_: `600`

_Maximum_: `28800`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RollbackMaximumBatchSize`

The batch size for a rollback to the old endpoint fleet. If this field is absent, the
value is set to the default, which is 100% of the total capacity. When the default is used,
SageMaker AI provisions the entire capacity of the old fleet at once during
rollback.

_Required_: No

_Type_: [InferenceComponentCapacitySize](aws-properties-sagemaker-inferencecomponent-inferencecomponentcapacitysize.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WaitIntervalInSeconds`

The length of the baking period, during which SageMaker AI monitors alarms for each
batch on the new fleet.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `3600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InferenceComponentDeploymentConfig

InferenceComponentRuntimeConfig

All content copied from https://docs.aws.amazon.com/.
