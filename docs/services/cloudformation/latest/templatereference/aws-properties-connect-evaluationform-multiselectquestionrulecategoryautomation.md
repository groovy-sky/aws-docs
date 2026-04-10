This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::EvaluationForm MultiSelectQuestionRuleCategoryAutomation

Automation rule for multi-select questions based on rule categories.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Category" : String,
  "Condition" : String,
  "OptionRefIds" : [ String, ... ]
}

```

### YAML

```yaml

  Category: String
  Condition: String
  OptionRefIds:
    - String

```

## Properties

`Category`

The category name for this automation rule.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Condition`

The condition for this automation rule.

_Required_: Yes

_Type_: String

_Allowed values_: `PRESENT | NOT_PRESENT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OptionRefIds`

Reference IDs of options for this automation rule.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EvaluationReviewNotificationRecipientValue

NumericQuestionPropertyValueAutomation

All content copied from https://docs.aws.amazon.com/.
