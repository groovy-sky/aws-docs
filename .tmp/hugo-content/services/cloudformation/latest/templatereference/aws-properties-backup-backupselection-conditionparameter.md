This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Backup::BackupSelection ConditionParameter

Includes information about tags you define to assign tagged resources to a backup
plan.

Include the prefix `aws:ResourceTag` in your tags.
For example, `"aws:ResourceTag/TagKey1": "Value1"`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConditionKey" : String,
  "ConditionValue" : String
}

```

### YAML

```yaml

  ConditionKey: String
  ConditionValue: String

```

## Properties

`ConditionKey`

The key in a key-value pair. For example, in the tag `Department:
         Accounting`, `Department` is the key.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ConditionValue`

The value in a key-value pair. For example, in the tag `Department:
            Accounting`, `Accounting` is the value.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BackupSelectionResourceType

ConditionResourceType

All content copied from https://docs.aws.amazon.com/.
