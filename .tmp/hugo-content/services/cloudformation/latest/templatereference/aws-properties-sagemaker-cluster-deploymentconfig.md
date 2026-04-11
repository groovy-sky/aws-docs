This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Cluster DeploymentConfig

The deployment configuration for an endpoint, which contains the desired deployment
strategy and rollback configurations.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AutoRollbackConfiguration" : [ AlarmDetails, ... ],
  "RollingUpdatePolicy" : RollingUpdatePolicy,
  "WaitIntervalInSeconds" : Integer
}

```

### YAML

```yaml

  AutoRollbackConfiguration:
    - AlarmDetails
  RollingUpdatePolicy:
    RollingUpdatePolicy
  WaitIntervalInSeconds: Integer

```

## Properties

`AutoRollbackConfiguration`

Automatic rollback configuration for handling endpoint deployment failures and
recovery.

_Required_: No

_Type_: Array of [AlarmDetails](aws-properties-sagemaker-cluster-alarmdetails.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RollingUpdatePolicy`

Specifies a rolling deployment strategy for updating a SageMaker endpoint.

_Required_: No

_Type_: [RollingUpdatePolicy](aws-properties-sagemaker-cluster-rollingupdatepolicy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WaitIntervalInSeconds`

The wait interval in seconds between deployment batches.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `3600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ClusterSlurmConfig

EnvironmentConfig

All content copied from https://docs.aws.amazon.com/.
