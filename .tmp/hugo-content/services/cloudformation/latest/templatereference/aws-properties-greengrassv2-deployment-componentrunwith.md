This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GreengrassV2::Deployment ComponentRunWith

Contains information system user and group that the AWS IoT Greengrass Core software uses
to run component processes on the core device. For more information, see [Configure the user and group that run components](../../../greengrass/v2/developerguide/configure-greengrass-core-v2.md#configure-component-user) in the _AWS IoT Greengrass V2 Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PosixUser" : String,
  "SystemResourceLimits" : SystemResourceLimits,
  "WindowsUser" : String
}

```

### YAML

```yaml

  PosixUser: String
  SystemResourceLimits:
    SystemResourceLimits
  WindowsUser: String

```

## Properties

`PosixUser`

The POSIX system user and (optional) group to use to run this component. Specify the user
and group separated by a colon ( `:`) in the following format:
`user:group`. The group is optional. If you don't specify a group, the AWS IoT Greengrass Core software uses the primary user for the group.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SystemResourceLimits`

The system resource limits to apply to this component's process on the core device.
AWS IoT Greengrass supports this feature only on Linux core devices.

If you omit this parameter, the AWS IoT Greengrass Core software uses the default system
resource limits that you configure on the AWS IoT Greengrass nucleus component. For more
information, see [Configure system resource limits for components](../../../greengrass/v2/developerguide/configure-greengrass-core-v2.md#configure-component-system-resource-limits).

_Required_: No

_Type_: [SystemResourceLimits](aws-properties-greengrassv2-deployment-systemresourcelimits.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`WindowsUser`

The Windows user to use to run this component on Windows core devices. The user must exist
on each Windows core device, and its name and password must be in the LocalSystem account's
Credentials Manager instance.

If you omit this parameter, the AWS IoT Greengrass Core software uses the default Windows
user that you configure on the AWS IoT Greengrass nucleus component. For more information,
see [Configure the user and group that run components](../../../greengrass/v2/developerguide/configure-greengrass-core-v2.md#configure-component-user).

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ComponentDeploymentSpecification

DeploymentComponentUpdatePolicy

All content copied from https://docs.aws.amazon.com/.
