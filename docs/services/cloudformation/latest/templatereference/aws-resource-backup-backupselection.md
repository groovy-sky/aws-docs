This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Backup::BackupSelection

Specifies a set of resources to assign to a backup plan.

For a sample CloudFormation template, see the [AWS Backup Developer Guide](../../../aws-backup/latest/devguide/assigning-resources.md#assigning-resources-cfn).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Backup::BackupSelection",
  "Properties" : {
      "BackupPlanId" : String,
      "BackupSelection" : BackupSelectionResourceType
    }
}

```

### YAML

```yaml

Type: AWS::Backup::BackupSelection
Properties:
  BackupPlanId: String
  BackupSelection:
    BackupSelectionResourceType

```

## Properties

`BackupPlanId`

Uniquely identifies a backup plan.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BackupSelection`

Specifies the body of a request to assign a set of resources to a backup plan.

It includes an array of resources, an optional array of patterns to exclude resources,
an optional role to provide access to the AWS service the resource belongs
to, and an optional array of tags used to identify a set of resources.

_Required_: Yes

_Type_: [BackupSelectionResourceType](aws-properties-backup-backupselection-backupselectionresourcetype.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns `BackupSelectionId`.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`BackupPlanId`

Uniquely identifies a backup plan.

`Id`

Uniquely identifies the backup selection.

`SelectionId`

Uniquely identifies a request to assign a set of resources
to a backup plan.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ScanSettingResourceType

BackupSelectionResourceType

All content copied from https://docs.aws.amazon.com/.
