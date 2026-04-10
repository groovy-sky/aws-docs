This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::EvaluationForm EvaluationFormQuestionTypeProperties

Information about properties for a question in an evaluation form. The question type
properties must be either for a numeric question or a single select question.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MultiSelect" : EvaluationFormMultiSelectQuestionProperties,
  "Numeric" : EvaluationFormNumericQuestionProperties,
  "SingleSelect" : EvaluationFormSingleSelectQuestionProperties,
  "Text" : EvaluationFormTextQuestionProperties
}

```

### YAML

```yaml

  MultiSelect:
    EvaluationFormMultiSelectQuestionProperties
  Numeric:
    EvaluationFormNumericQuestionProperties
  SingleSelect:
    EvaluationFormSingleSelectQuestionProperties
  Text:
    EvaluationFormTextQuestionProperties

```

## Properties

`MultiSelect`

Properties for multi-select question types.

_Required_: No

_Type_: [EvaluationFormMultiSelectQuestionProperties](aws-properties-connect-evaluationform-evaluationformmultiselectquestionproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Numeric`

The properties of the numeric question.

_Required_: No

_Type_: [EvaluationFormNumericQuestionProperties](aws-properties-connect-evaluationform-evaluationformnumericquestionproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SingleSelect`

The properties of the numeric question.

_Required_: No

_Type_: [EvaluationFormSingleSelectQuestionProperties](aws-properties-connect-evaluationform-evaluationformsingleselectquestionproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Text`

The properties of the text question.

_Required_: No

_Type_: [EvaluationFormTextQuestionProperties](aws-properties-connect-evaluationform-evaluationformtextquestionproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EvaluationFormQuestionAutomationAnswerSource

EvaluationFormSection

All content copied from https://docs.aws.amazon.com/.
