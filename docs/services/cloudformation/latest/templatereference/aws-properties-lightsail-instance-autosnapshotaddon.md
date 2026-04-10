This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lightsail::Instance AutoSnapshotAddOn

`AutoSnapshotAddOn` is a property of the [AddOn](../userguide/aws-properties-lightsail-instance-addon.md) property. It describes the automatic snapshot add-on for an
instance.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SnapshotTimeOfDay" : String
}

```

### YAML

```yaml

  SnapshotTimeOfDay: String

```

## Properties

`SnapshotTimeOfDay`

The daily time when an automatic snapshot will be created.

Constraints:

- Must be in `HH:00` format, and in an hourly increment.

- Specified in Coordinated Universal Time (UTC).

- The snapshot will be automatically created between the time specified and up to 45
minutes after.

_Required_: No

_Type_: String

_Pattern_: `^[0-9]{2}:00$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AddOn

Disk

All content copied from https://docs.aws.amazon.com/.
