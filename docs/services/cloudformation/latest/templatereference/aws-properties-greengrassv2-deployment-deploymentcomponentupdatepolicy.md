This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GreengrassV2::Deployment DeploymentComponentUpdatePolicy

Contains information about a deployment's policy that defines when components are safe to
update.

Each component on a device can report whether or not it's ready to update. After a
component and its dependencies are ready, they can apply the update in the deployment. You can
configure whether or not the deployment notifies components of an update and waits for a
response. You specify the amount of time each component has to respond to the update
notification.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Action" : String,
  "TimeoutInSeconds" : Integer
}

```

### YAML

```yaml

  Action: String
  TimeoutInSeconds: Integer

```

## Properties

`Action`

Whether or not to notify components and wait for components to become safe to update.
Choose from the following options:

- `NOTIFY_COMPONENTS` – The deployment notifies each component before it
stops and updates that component. Components can use the [SubscribeToComponentUpdates](../../../greengrass/v2/developerguide/interprocess-communication.md#ipc-operation-subscribetocomponentupdates) IPC operation to receive these notifications. Then,
components can respond with the [DeferComponentUpdate](../../../greengrass/v2/developerguide/interprocess-communication.md#ipc-operation-defercomponentupdate) IPC operation. For more information, see the [Create\
deployments](../../../greengrass/v2/developerguide/create-deployments.md) in the _AWS IoT Greengrass V2 Developer_
_Guide_.

- `SKIP_NOTIFY_COMPONENTS` – The deployment doesn't notify components or wait
for them to be safe to update.

Default: `NOTIFY_COMPONENTS`

_Required_: No

_Type_: String

_Allowed values_: `NOTIFY_COMPONENTS | SKIP_NOTIFY_COMPONENTS`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TimeoutInSeconds`

The amount of time in seconds that each component on a device has to report that it's safe
to update. If the component waits for longer than this timeout, then the deployment proceeds
on the device.

Default: `60`

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `2147483647`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ComponentRunWith

DeploymentConfigurationValidationPolicy

All content copied from https://docs.aws.amazon.com/.
