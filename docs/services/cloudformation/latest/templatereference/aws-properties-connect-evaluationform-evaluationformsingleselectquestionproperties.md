This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::EvaluationForm EvaluationFormSingleSelectQuestionProperties

Information about the options in single select questions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Automation" : EvaluationFormSingleSelectQuestionAutomation,
  "DisplayAs" : String,
  "Options" : [ EvaluationFormSingleSelectQuestionOption, ... ]
}

```

### YAML

```yaml

  Automation:
    EvaluationFormSingleSelectQuestionAutomation
  DisplayAs: String
  Options:
    - EvaluationFormSingleSelectQuestionOption

```

## Properties

`Automation`

The display mode of the single select question.

_Required_: No

_Type_: [EvaluationFormSingleSelectQuestionAutomation](aws-properties-connect-evaluationform-evaluationformsingleselectquestionautomation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisplayAs`

The display mode of the single select question.

_Allowed values_: `DROPDOWN` \| `RADIO`

_Required_: No

_Type_: String

_Allowed values_: `DROPDOWN | RADIO`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Options`

The answer options of the single select question.

_Minimum_: 2

_Maximum_: 256

_Required_: Yes

_Type_: Array of [EvaluationFormSingleSelectQuestionOption](aws-properties-connect-evaluationform-evaluationformsingleselectquestionoption.md)

_Minimum_: `2`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EvaluationFormSingleSelectQuestionOption

EvaluationFormTargetConfiguration

All content copied from https://docs.aws.amazon.com/.
