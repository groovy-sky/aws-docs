This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GreengrassV2::Deployment SystemResourceLimits

Contains information about system resource limits that the software
applies to a component's processes.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Cpus" : Number,
  "Memory" : Integer
}

```

### YAML

```yaml

  Cpus: Number
  Memory: Integer

```

## Properties

`Cpus`

The maximum amount of CPU time that a component's processes can use on the core device. A
core device's total CPU time is equivalent to the device's number of CPU cores. For example,
on a core device with 4 CPU cores, you can set this value to 2 to limit the component's
processes to 50 percent usage of each CPU core. On a device with 1 CPU core, you can set this
value to 0.25 to limit the component's processes to 25 percent usage of the CPU. If you set
this value to a number greater than the number of CPU cores, the AWS IoT Greengrass Core
software doesn't limit the component's CPU usage.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Memory`

The maximum amount of RAM, expressed in kilobytes, that a component's processes can use on
the core device. For more information, see [Configure system resource limits for components](../../../greengrass/v2/developerguide/configure-greengrass-core-v2.md#configure-component-system-resource-limits).

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `9223372036854771712`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IoTJobTimeoutConfig

Next

All content copied from https://docs.aws.amazon.com/.
