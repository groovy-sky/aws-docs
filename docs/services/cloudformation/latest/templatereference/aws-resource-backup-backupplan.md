This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Backup::BackupPlan

Contains an optional backup plan display name and an array of `BackupRule`
objects, each of which specifies a backup rule. Each rule in a backup plan is a separate
scheduled task and can back up a different selection of AWS
resources.

For a sample CloudFormation template, see the [AWS Backup Developer Guide](../../../aws-backup/latest/devguide/assigning-resources.md#assigning-resources-cfn).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Backup::BackupPlan",
  "Properties" : {
      "BackupPlan" : BackupPlanResourceType,
      "BackupPlanTags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::Backup::BackupPlan
Properties:
  BackupPlan:
    BackupPlanResourceType
  BackupPlanTags:
    Key: Value

```

## Properties

`BackupPlan`

Uniquely identifies the backup plan to be associated with the selection of
resources.

_Required_: Yes

_Type_: [BackupPlanResourceType](aws-properties-backup-backupplan-backupplanresourcetype.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BackupPlanTags`

The tags to assign to the backup plan.

_Required_: No

_Type_: Object of String

_Pattern_: `^.{1,128}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns `BackupPlanId`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`BackupPlanArn`

An Amazon Resource Name (ARN) that uniquely identifies a
backup plan; for example,
`arn:aws:backup:us-east-1:123456789012:plan:8F81F553-3A74-4A3F-B93D-B3360DC80C50`.

`BackupPlanId`

Uniquely identifies a backup plan.

`VersionId`

Unique, randomly generated, Unicode, UTF-8 encoded strings
that are at most 1,024 bytes long. Version Ids cannot be edited.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Backup

AdvancedBackupSettingResourceType

All content copied from https://docs.aws.amazon.com/.
