This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::PrivacyBudgetTemplate BudgetParameter

Individual budget parameter configuration that defines specific budget allocation settings for access budgets.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AutoRefresh" : String,
  "Budget" : Integer,
  "Type" : String
}

```

### YAML

```yaml

  AutoRefresh: String
  Budget: Integer
  Type: String

```

## Properties

`AutoRefresh`

Whether this individual budget parameter automatically refreshes when the budget period resets.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Budget`

The budget allocation amount for this specific parameter.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of budget parameter being configured.

_Required_: Yes

_Type_: String

_Allowed values_: `CALENDAR_DAY | CALENDAR_MONTH | CALENDAR_WEEK | LIFETIME`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::CleanRooms::PrivacyBudgetTemplate

Parameters

All content copied from https://docs.aws.amazon.com/.
