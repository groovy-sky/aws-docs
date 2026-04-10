This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeDeploy::DeploymentGroup AutoRollbackConfiguration

The `AutoRollbackConfiguration` property type configures automatic rollback for
an AWS CodeDeploy deployment group when a deployment is not completed successfully.
For more information, see [Automatic Rollbacks](../../../codedeploy/latest/userguide/deployments-rollback-and-redeploy.md#deployments-rollback-and-redeploy-automatic-rollbacks) in the _AWS CodeDeploy User_
_Guide_.

`AutoRollbackConfiguration` is a property of the [DeploymentGroup](../userguide/aws-resource-codedeploy-deploymentgroup.md) resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enabled" : Boolean,
  "Events" : [ String, ... ]
}

```

### YAML

```yaml

  Enabled: Boolean
  Events:
    - String

```

## Properties

`Enabled`

Indicates whether a defined automatic rollback configuration is currently
enabled.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Events`

The event type or types that trigger a rollback. Valid values are
`DEPLOYMENT_FAILURE`, `DEPLOYMENT_STOP_ON_ALARM`, or
`DEPLOYMENT_STOP_ON_REQUEST`.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AlarmConfiguration

BlueGreenDeploymentConfiguration

All content copied from https://docs.aws.amazon.com/.
