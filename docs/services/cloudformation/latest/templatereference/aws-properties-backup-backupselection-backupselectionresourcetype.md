This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Backup::BackupSelection BackupSelectionResourceType

Specifies an object containing properties used to assign a set of resources to a backup
plan.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Conditions" : Conditions,
  "IamRoleArn" : String,
  "ListOfTags" : [ ConditionResourceType, ... ],
  "NotResources" : [ String, ... ],
  "Resources" : [ String, ... ],
  "SelectionName" : String
}

```

### YAML

```yaml

  Conditions:
    Conditions
  IamRoleArn: String
  ListOfTags:
    - ConditionResourceType
  NotResources:
    - String
  Resources:
    - String
  SelectionName: String

```

## Properties

`Conditions`

A list of conditions that you define to assign resources to your backup plans using
tags. For example, `"StringEquals": { "ConditionKey": "aws:ResourceTag/CreatedByCryo",
            "ConditionValue": "true" },`. Condition operators are case sensitive.

`Conditions` differs from `ListOfTags` as follows:

- When you specify more than one condition, you only assign the resources that match
ALL conditions (using AND logic).

- `Conditions` supports `StringEquals`,
`StringLike`, `StringNotEquals`, and
`StringNotLike`. `ListOfTags` only supports
`StringEquals`.

_Required_: No

_Type_: [Conditions](aws-properties-backup-backupselection-conditions.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IamRoleArn`

The ARN of the IAM role that AWS Backup uses to authenticate when backing up
the target resource; for example,
`arn:aws:iam::123456789012:role/S3Access`.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ListOfTags`

A list of conditions that you define to assign resources to your backup plans using
tags. For example, `"StringEquals": { "ConditionKey": "aws:ResourceTag/CreatedByCryo",
            "ConditionValue": "true" },`. Condition operators are case sensitive.

`ListOfTags` differs from `Conditions` as follows:

- When you specify more than one condition, you assign all resources that match AT
LEAST ONE condition (using OR logic).

- `ListOfTags` only supports `StringEquals`.
`Conditions` supports `StringEquals`,
`StringLike`, `StringNotEquals`, and
`StringNotLike`.

_Required_: No

_Type_: Array of [ConditionResourceType](aws-properties-backup-backupselection-conditionresourcetype.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NotResources`

A list of Amazon Resource Names (ARNs) to exclude from a backup plan. The maximum number
of ARNs is 500 without wildcards, or 30 ARNs with wildcards.

If you need to exclude many resources from a backup plan, consider a different resource
selection strategy, such as assigning only one or a few resource types or refining your
resource selection using tags.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Resources`

An array of strings that contain Amazon Resource Names (ARNs) of resources to assign to
a backup plan.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SelectionName`

The display name of a resource selection document.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Backup::BackupSelection

ConditionParameter

All content copied from https://docs.aws.amazon.com/.
