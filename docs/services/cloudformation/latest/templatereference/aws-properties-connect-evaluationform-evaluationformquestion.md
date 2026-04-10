This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::EvaluationForm EvaluationFormQuestion

Information about a question from an evaluation form.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enablement" : EvaluationFormItemEnablementConfiguration,
  "Instructions" : String,
  "NotApplicableEnabled" : Boolean,
  "QuestionType" : String,
  "QuestionTypeProperties" : EvaluationFormQuestionTypeProperties,
  "RefId" : String,
  "Title" : String,
  "Weight" : Number
}

```

### YAML

```yaml

  Enablement:
    EvaluationFormItemEnablementConfiguration
  Instructions: String
  NotApplicableEnabled: Boolean
  QuestionType: String
  QuestionTypeProperties:
    EvaluationFormQuestionTypeProperties
  RefId: String
  Title: String
  Weight: Number

```

## Properties

`Enablement`

A question conditional enablement.

_Required_: No

_Type_: [EvaluationFormItemEnablementConfiguration](aws-properties-connect-evaluationform-evaluationformitemenablementconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Instructions`

The instructions of the section.

_Length Constraints_: Minimum length of 0. Maximum length of
1024.

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NotApplicableEnabled`

The flag to enable not applicable answers to the question.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QuestionType`

The type of the question.

_Allowed values_: `NUMERIC` \| `SINGLESELECT`
\| `TEXT`

_Required_: Yes

_Type_: String

_Allowed values_: `NUMERIC | SINGLESELECT | TEXT | MULTISELECT | DATETIME`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QuestionTypeProperties`

The properties of the type of question. Text questions do not have to define question
type properties.

_Required_: No

_Type_: [EvaluationFormQuestionTypeProperties](aws-properties-connect-evaluationform-evaluationformquestiontypeproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RefId`

The identifier of the question. An identifier must be unique within the evaluation
form.

_Length Constraints_: Minimum length of 1. Maximum length of
40.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9._-]{1,40}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Title`

The title of the question.

_Length Constraints_: Minimum length of 1. Maximum length of
350.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `350`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Weight`

The scoring weight of the section.

_Minimum_: 0

_Maximum_: 100

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EvaluationFormNumericQuestionProperties

EvaluationFormQuestionAutomationAnswerSource

All content copied from https://docs.aws.amazon.com/.
