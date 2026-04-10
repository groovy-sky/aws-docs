This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::EvaluationForm EvaluationFormMultiSelectQuestionProperties

Properties for a multi-select question in an evaluation form.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Automation" : EvaluationFormMultiSelectQuestionAutomation,
  "DisplayAs" : String,
  "Options" : [ EvaluationFormMultiSelectQuestionOption, ... ]
}

```

### YAML

```yaml

  Automation:
    EvaluationFormMultiSelectQuestionAutomation
  DisplayAs: String
  Options:
    - EvaluationFormMultiSelectQuestionOption

```

## Properties

`Automation`

Automation configuration for this multi-select question.

_Required_: No

_Type_: [EvaluationFormMultiSelectQuestionAutomation](aws-properties-connect-evaluationform-evaluationformmultiselectquestionautomation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisplayAs`

Display format for the multi-select question.

_Required_: No

_Type_: String

_Allowed values_: `DROPDOWN | CHECKBOX`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Options`

Options available for this multi-select question.

_Required_: Yes

_Type_: Array of [EvaluationFormMultiSelectQuestionOption](aws-properties-connect-evaluationform-evaluationformmultiselectquestionoption.md)

_Minimum_: `2`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EvaluationFormMultiSelectQuestionOption

EvaluationFormNumericQuestionAutomation

All content copied from https://docs.aws.amazon.com/.
