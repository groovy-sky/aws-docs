This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::EvaluationForm EvaluationFormSingleSelectQuestionAutomation

Information about the automation configuration in single select questions. Automation
options are evaluated in order, and the first matched option is applied. If no
automation option matches, and there is a default option, then the default option is
applied.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AnswerSource" : EvaluationFormQuestionAutomationAnswerSource,
  "DefaultOptionRefId" : String,
  "Options" : [ EvaluationFormSingleSelectQuestionAutomationOption, ... ]
}

```

### YAML

```yaml

  AnswerSource:
    EvaluationFormQuestionAutomationAnswerSource
  DefaultOptionRefId: String
  Options:
    - EvaluationFormSingleSelectQuestionAutomationOption

```

## Properties

`AnswerSource`

Automation answer source.

_Required_: No

_Type_: [EvaluationFormQuestionAutomationAnswerSource](aws-properties-connect-evaluationform-evaluationformquestionautomationanswersource.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultOptionRefId`

The identifier of the default answer option, when none of the automation options match
the criteria.

_Length Constraints_: Minimum length of 1. Maximum length of
40.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9._-]{1,40}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Options`

The automation options of the single select question.

_Minimum_: 1

_Maximum_: 20

_Required_: Yes

_Type_: Array of [EvaluationFormSingleSelectQuestionAutomationOption](aws-properties-connect-evaluationform-evaluationformsingleselectquestionautomationoption.md)

_Minimum_: `1`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EvaluationFormSection

EvaluationFormSingleSelectQuestionAutomationOption

All content copied from https://docs.aws.amazon.com/.
