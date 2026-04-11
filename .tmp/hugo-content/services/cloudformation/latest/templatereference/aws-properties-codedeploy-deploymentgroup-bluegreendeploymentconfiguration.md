This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeDeploy::DeploymentGroup BlueGreenDeploymentConfiguration

Information about blue/green deployment options for a deployment group.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DeploymentReadyOption" : DeploymentReadyOption,
  "GreenFleetProvisioningOption" : GreenFleetProvisioningOption,
  "TerminateBlueInstancesOnDeploymentSuccess" : BlueInstanceTerminationOption
}

```

### YAML

```yaml

  DeploymentReadyOption:
    DeploymentReadyOption
  GreenFleetProvisioningOption:
    GreenFleetProvisioningOption
  TerminateBlueInstancesOnDeploymentSuccess:
    BlueInstanceTerminationOption

```

## Properties

`DeploymentReadyOption`

Information about the action to take when newly provisioned instances are ready to
receive traffic in a blue/green deployment.

_Required_: No

_Type_: [DeploymentReadyOption](aws-properties-codedeploy-deploymentgroup-deploymentreadyoption.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GreenFleetProvisioningOption`

Information about how instances are provisioned for a replacement environment in a
blue/green deployment.

_Required_: No

_Type_: [GreenFleetProvisioningOption](aws-properties-codedeploy-deploymentgroup-greenfleetprovisioningoption.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TerminateBlueInstancesOnDeploymentSuccess`

Information about whether to terminate instances in the original fleet during a
blue/green deployment.

_Required_: No

_Type_: [BlueInstanceTerminationOption](aws-properties-codedeploy-deploymentgroup-blueinstanceterminationoption.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AutoRollbackConfiguration

BlueInstanceTerminationOption

All content copied from https://docs.aws.amazon.com/.
