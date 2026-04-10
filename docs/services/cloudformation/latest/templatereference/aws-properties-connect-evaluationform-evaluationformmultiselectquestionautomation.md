This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::EvaluationForm EvaluationFormMultiSelectQuestionAutomation

Automation configuration for multi-select questions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AnswerSource" : EvaluationFormQuestionAutomationAnswerSource,
  "DefaultOptionRefIds" : [ String, ... ],
  "Options" : [ EvaluationFormMultiSelectQuestionAutomationOption, ... ]
}

```

### YAML

```yaml

  AnswerSource:
    EvaluationFormQuestionAutomationAnswerSource
  DefaultOptionRefIds:
    - String
  Options:
    - EvaluationFormMultiSelectQuestionAutomationOption

```

## Properties

`AnswerSource`

Property description not available.

_Required_: No

_Type_: [EvaluationFormQuestionAutomationAnswerSource](aws-properties-connect-evaluationform-evaluationformquestionautomationanswersource.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultOptionRefIds`

Reference IDs of default options.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Options`

Automation options for the multi-select question.

_Required_: Yes

_Type_: Array of [EvaluationFormMultiSelectQuestionAutomationOption](aws-properties-connect-evaluationform-evaluationformmultiselectquestionautomationoption.md)

_Minimum_: `1`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EvaluationFormLanguageConfiguration

EvaluationFormMultiSelectQuestionAutomationOption

All content copied from https://docs.aws.amazon.com/.
