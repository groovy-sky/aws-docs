This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::InferenceComponent InferenceComponentDeploymentConfig

The deployment configuration for an endpoint that hosts inference components. The
configuration includes the desired deployment strategy and rollback settings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AutoRollbackConfiguration" : AutoRollbackConfiguration,
  "RollingUpdatePolicy" : InferenceComponentRollingUpdatePolicy
}

```

### YAML

```yaml

  AutoRollbackConfiguration:
    AutoRollbackConfiguration
  RollingUpdatePolicy:
    InferenceComponentRollingUpdatePolicy

```

## Properties

`AutoRollbackConfiguration`

Configuration for automatic rollback during inference component deployment.

_Required_: No

_Type_: [AutoRollbackConfiguration](aws-properties-sagemaker-inferencecomponent-autorollbackconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RollingUpdatePolicy`

Specifies a rolling deployment strategy for updating a SageMaker AI
endpoint.

_Required_: No

_Type_: [InferenceComponentRollingUpdatePolicy](aws-properties-sagemaker-inferencecomponent-inferencecomponentrollingupdatepolicy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InferenceComponentContainerSpecification

InferenceComponentRollingUpdatePolicy

All content copied from https://docs.aws.amazon.com/.
