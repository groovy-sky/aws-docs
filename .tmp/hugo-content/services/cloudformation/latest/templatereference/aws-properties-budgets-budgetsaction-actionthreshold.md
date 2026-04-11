This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Budgets::BudgetsAction ActionThreshold

The trigger threshold of the action.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : String,
  "Value" : Number
}

```

### YAML

```yaml

  Type: String
  Value: Number

```

## Properties

`Type`

The type of threshold for a notification.

_Required_: Yes

_Type_: String

_Allowed values_: `PERCENTAGE | ABSOLUTE_VALUE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The threshold of a notification.

_Required_: Yes

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Budgets::BudgetsAction

Definition

All content copied from https://docs.aws.amazon.com/.
