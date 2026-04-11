This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Backup::BackupSelection Conditions

Contains information about which resources to include or exclude from a backup plan
using their tags. Conditions are case sensitive.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "StringEquals" : [ ConditionParameter, ... ],
  "StringLike" : [ ConditionParameter, ... ],
  "StringNotEquals" : [ ConditionParameter, ... ],
  "StringNotLike" : [ ConditionParameter, ... ]
}

```

### YAML

```yaml

  StringEquals:
    - ConditionParameter
  StringLike:
    - ConditionParameter
  StringNotEquals:
    - ConditionParameter
  StringNotLike:
    - ConditionParameter

```

## Properties

`StringEquals`

Filters the values of your tagged resources for only those resources that you tagged
with the same value. Also called "exact matching."

_Required_: No

_Type_: Array of [ConditionParameter](aws-properties-backup-backupselection-conditionparameter.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StringLike`

Filters the values of your tagged resources for matching tag values with the use of a
wildcard character (\*) anywhere in the string. For example, "prod\*" or "\*rod\*" matches the
tag value "production".

_Required_: No

_Type_: Array of [ConditionParameter](aws-properties-backup-backupselection-conditionparameter.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StringNotEquals`

Filters the values of your tagged resources for only those resources that you tagged
that do not have the same value. Also called "negated matching."

_Required_: No

_Type_: Array of [ConditionParameter](aws-properties-backup-backupselection-conditionparameter.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StringNotLike`

Filters the values of your tagged resources for non-matching tag values with the use of
a wildcard character (\*) anywhere in the string.

_Required_: No

_Type_: Array of [ConditionParameter](aws-properties-backup-backupselection-conditionparameter.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConditionResourceType

AWS::Backup::BackupVault

All content copied from https://docs.aws.amazon.com/.
