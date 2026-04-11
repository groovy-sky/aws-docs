This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Backup::BackupSelection ConditionResourceType

Specifies an object that contains an array of triplets made up of a condition type (such
as `STRINGEQUALS`), a key, and a value. Conditions are used to filter resources
in a selection that is assigned to a backup plan.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConditionKey" : String,
  "ConditionType" : String,
  "ConditionValue" : String
}

```

### YAML

```yaml

  ConditionKey: String
  ConditionType: String
  ConditionValue: String

```

## Properties

`ConditionKey`

The key in a key-value pair. For example, in `"Department": "accounting"`,
`"Department"` is the key.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ConditionType`

An operation, such as `STRINGEQUALS`, that is applied to a key-value pair
used to filter resources in a selection.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ConditionValue`

The value in a key-value pair. For example, in `"Department": "accounting"`,
`"accounting"` is the value.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConditionParameter

Conditions

All content copied from https://docs.aws.amazon.com/.
