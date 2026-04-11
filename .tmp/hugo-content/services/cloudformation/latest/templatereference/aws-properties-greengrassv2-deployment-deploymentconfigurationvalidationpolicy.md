This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GreengrassV2::Deployment DeploymentConfigurationValidationPolicy

Contains information about how long a component on a core device can validate its
configuration updates before it times out. Components can use the [SubscribeToValidateConfigurationUpdates](../../../greengrass/v2/developerguide/interprocess-communication.md#ipc-operation-subscribetovalidateconfigurationupdates) IPC operation to receive notifications when
a deployment specifies a configuration update. Then, components can respond with the [SendConfigurationValidityReport](../../../greengrass/v2/developerguide/interprocess-communication.md#ipc-operation-sendconfigurationvalidityreport) IPC operation. For more information, see the [Create\
deployments](../../../greengrass/v2/developerguide/create-deployments.md) in the _AWS IoT Greengrass V2 Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TimeoutInSeconds" : Integer
}

```

### YAML

```yaml

  TimeoutInSeconds: Integer

```

## Properties

`TimeoutInSeconds`

The amount of time in seconds that a component can validate its configuration updates. If
the validation time exceeds this timeout, then the deployment proceeds for the device.

Default: `30`

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `2147483647`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeploymentComponentUpdatePolicy

DeploymentIoTJobConfiguration

All content copied from https://docs.aws.amazon.com/.
