This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::EvaluationForm ScoringStrategy

A scoring strategy of the evaluation form.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Mode" : String,
  "Status" : String
}

```

### YAML

```yaml

  Mode: String
  Status: String

```

## Properties

`Mode`

The scoring mode of the evaluation form.

_Allowed values_: `QUESTION_ONLY` \|
`SECTION_ONLY`

_Required_: Yes

_Type_: String

_Allowed values_: `QUESTION_ONLY | SECTION_ONLY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

The scoring status of the evaluation form.

_Allowed values_: `ENABLED` \|
`DISABLED`

_Required_: Yes

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NumericQuestionPropertyValueAutomation

SingleSelectQuestionRuleCategoryAutomation

All content copied from https://docs.aws.amazon.com/.
