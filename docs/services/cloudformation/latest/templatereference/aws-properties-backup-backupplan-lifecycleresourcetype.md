This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Backup::BackupPlan LifecycleResourceType

Specifies an object containing an array of `Transition` objects that
determine how long in days before a recovery point transitions to cold storage or is
deleted.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DeleteAfterDays" : Number,
  "MoveToColdStorageAfterDays" : Number,
  "OptInToArchiveForSupportedResources" : Boolean
}

```

### YAML

```yaml

  DeleteAfterDays: Number
  MoveToColdStorageAfterDays: Number
  OptInToArchiveForSupportedResources: Boolean

```

## Properties

`DeleteAfterDays`

The number of days after creation that a recovery point is deleted. This value must be
at least 90 days after the number of days specified in `MoveToColdStorageAfterDays`.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MoveToColdStorageAfterDays`

The number of days after creation that a recovery point is moved to cold
storage.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OptInToArchiveForSupportedResources`

If the value is true, your backup plan transitions supported resources to
archive (cold) storage tier in accordance with your lifecycle settings.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IndexActionsResourceType

ScanActionResourceType

All content copied from https://docs.aws.amazon.com/.
