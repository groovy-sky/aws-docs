This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GreengrassV2::ComponentVersion LambdaDeviceMount

Contains information about a device that Linux processes in a container can access.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AddGroupOwner" : Boolean,
  "Path" : String,
  "Permission" : String
}

```

### YAML

```yaml

  AddGroupOwner: Boolean
  Path: String
  Permission: String

```

## Properties

`AddGroupOwner`

Whether or not to add the component's system user as an owner of the device.

Default: `false`

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Path`

The mount path for the device in the file system.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Permission`

The permission to access the device: read/only ( `ro`) or read/write
( `rw`).

Default: `ro`

_Required_: No

_Type_: String

_Allowed values_: `ro | rw`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LambdaContainerParams

LambdaEventSource

All content copied from https://docs.aws.amazon.com/.
