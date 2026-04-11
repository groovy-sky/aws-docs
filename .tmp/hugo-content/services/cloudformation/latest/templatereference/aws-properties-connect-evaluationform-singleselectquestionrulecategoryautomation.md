This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::EvaluationForm SingleSelectQuestionRuleCategoryAutomation

Information about the automation option based on a rule category for a single select
question.

_Length Constraints_: Minimum length of 1. Maximum length of
50.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Category" : String,
  "Condition" : String,
  "OptionRefId" : String
}

```

### YAML

```yaml

  Category: String
  Condition: String
  OptionRefId: String

```

## Properties

`Category`

The category name, as defined in Rules.

_Minimum_: 1

_Maximum_: 50

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Condition`

The condition to apply for the automation option. If the condition is PRESENT, then
the option is applied when the contact data includes the category. Similarly, if the
condition is NOT\_PRESENT, then the option is applied when the contact data does not
include the category.

_Allowed values_: `PRESENT` \|
`NOT_PRESENT`

_Maximum_: 50

_Required_: Yes

_Type_: String

_Allowed values_: `PRESENT | NOT_PRESENT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OptionRefId`

The identifier of the answer option. An identifier must be unique within the
question.

_Length Constraints_: Minimum length of 1. Maximum length of
40.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9._-]{1,40}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ScoringStrategy

Tag

All content copied from https://docs.aws.amazon.com/.
