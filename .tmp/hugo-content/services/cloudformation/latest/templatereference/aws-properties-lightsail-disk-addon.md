This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lightsail::Disk AddOn

`AddOn` is a property of the [AWS::Lightsail::Disk](../userguide/aws-resource-lightsail-disk.md) resource. It describes the add-ons for a disk.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AddOnType" : String,
  "AutoSnapshotAddOnRequest" : AutoSnapshotAddOn,
  "Status" : String
}

```

### YAML

```yaml

  AddOnType: String
  AutoSnapshotAddOnRequest:
    AutoSnapshotAddOn
  Status: String

```

## Properties

`AddOnType`

The add-on type (for example, `AutoSnapshot`).

###### Note

`AutoSnapshot` is the only add-on that can be enabled for a disk.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AutoSnapshotAddOnRequest`

The parameters for the automatic snapshot add-on, such as the daily time when an
automatic snapshot will be created.

_Required_: No

_Type_: [AutoSnapshotAddOn](aws-properties-lightsail-disk-autosnapshotaddon.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

The status of the add-on.

Valid Values: `Enabled` \| `Disabled`

_Required_: No

_Type_: String

_Allowed values_: `Enabling | Disabling | Enabled | Terminating | Terminated | Disabled | Failed`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Lightsail::Disk

AutoSnapshotAddOn

All content copied from https://docs.aws.amazon.com/.
