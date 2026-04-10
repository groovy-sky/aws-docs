This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lightsail::Instance Hardware

`Hardware` is a property of the [AWS::Lightsail::Instance](../userguide/aws-resource-lightsail-instance.md) resource. It describes the hardware properties for the
instance, such as the vCPU count, attached disks, and amount of RAM.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CpuCount" : Integer,
  "Disks" : [ Disk, ... ],
  "RamSizeInGb" : Integer
}

```

### YAML

```yaml

  CpuCount: Integer
  Disks:
    - Disk
  RamSizeInGb: Integer

```

## Properties

`CpuCount`

The number of vCPUs the instance has.

###### Note

The `CpuCount` property is read-only and should not be specified in a
create instance or update instance request.

_Required_: No

_Type_: Integer

_Update requires_: Updates are not supported.

`Disks`

The disks attached to the instance.

The instance restarts when performing an attach disk or detach disk request. This resets
the public IP address of your instance if a static IP isn't attached to it.

_Required_: No

_Type_: Array of [Disk](aws-properties-lightsail-instance-disk.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`RamSizeInGb`

The amount of RAM in GB on the instance (for example, `1.0`).

###### Note

The `RamSizeInGb` property is read-only and should not be specified in a
create instance or update instance request.

_Required_: No

_Type_: Integer

_Update requires_: Updates are not supported.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Disk

Location

All content copied from https://docs.aws.amazon.com/.
