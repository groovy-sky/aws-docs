This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Endpoint DeploymentConfig

The deployment configuration for an endpoint, which contains the desired deployment
strategy and rollback configurations.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AutoRollbackConfiguration" : AutoRollbackConfig,
  "BlueGreenUpdatePolicy" : BlueGreenUpdatePolicy,
  "RollingUpdatePolicy" : RollingUpdatePolicy
}

```

### YAML

```yaml

  AutoRollbackConfiguration:
    AutoRollbackConfig
  BlueGreenUpdatePolicy:
    BlueGreenUpdatePolicy
  RollingUpdatePolicy:
    RollingUpdatePolicy

```

## Properties

`AutoRollbackConfiguration`

Automatic rollback configuration for handling endpoint deployment failures and
recovery.

_Required_: No

_Type_: [AutoRollbackConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-endpoint-autorollbackconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BlueGreenUpdatePolicy`

Update policy for a blue/green deployment. If this update policy is specified, SageMaker
creates a new fleet during the deployment while maintaining the old fleet. SageMaker flips
traffic to the new fleet according to the specified traffic routing configuration. Only
one update policy should be used in the deployment configuration. If no update policy is
specified, SageMaker uses a blue/green deployment strategy with all at once traffic shifting
by default.

_Required_: No

_Type_: [BlueGreenUpdatePolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-endpoint-bluegreenupdatepolicy.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RollingUpdatePolicy`

Specifies a rolling deployment strategy for updating a SageMaker endpoint.

_Required_: No

_Type_: [RollingUpdatePolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-endpoint-rollingupdatepolicy.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CapacitySize

RollingUpdatePolicy
