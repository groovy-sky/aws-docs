This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::EvaluationForm EvaluationFormSingleSelectQuestionOption

Information about the automation configuration in single select questions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AutomaticFail" : Boolean,
  "AutomaticFailConfiguration" : AutomaticFailConfiguration,
  "RefId" : String,
  "Score" : Integer,
  "Text" : String
}

```

### YAML

```yaml

  AutomaticFail: Boolean
  AutomaticFailConfiguration:
    AutomaticFailConfiguration
  RefId: String
  Score: Integer
  Text: String

```

## Properties

`AutomaticFail`

The flag to mark the option as automatic fail. If an automatic fail answer is
provided, the overall evaluation gets a score of 0.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AutomaticFailConfiguration`

Whether automatic fail is configured on a single select question.

_Required_: No

_Type_: [AutomaticFailConfiguration](aws-properties-connect-evaluationform-automaticfailconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RefId`

The identifier of the answer option. An identifier must be unique within the
question.

_Length Constraints_: Minimum length of 1. Maximum length of
40.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9._-]{1,40}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Score`

The score assigned to the answer option.

_Minimum_: 0

_Maximum_: 10

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Text`

The title of the answer option.

_Length Constraints_: Minimum length of 1. Maximum length of
128.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EvaluationFormSingleSelectQuestionAutomationOption

EvaluationFormSingleSelectQuestionProperties

All content copied from https://docs.aws.amazon.com/.
