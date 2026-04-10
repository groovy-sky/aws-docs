This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::EvaluationForm EvaluationFormNumericQuestionAutomation

Information about the automation configuration in numeric questions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AnswerSource" : EvaluationFormQuestionAutomationAnswerSource,
  "PropertyValue" : NumericQuestionPropertyValueAutomation
}

```

### YAML

```yaml

  AnswerSource:
    EvaluationFormQuestionAutomationAnswerSource
  PropertyValue:
    NumericQuestionPropertyValueAutomation

```

## Properties

`AnswerSource`

A source of automation answer for numeric question.

_Required_: No

_Type_: [EvaluationFormQuestionAutomationAnswerSource](aws-properties-connect-evaluationform-evaluationformquestionautomationanswersource.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PropertyValue`

The property value of the automation.

_Required_: No

_Type_: [NumericQuestionPropertyValueAutomation](aws-properties-connect-evaluationform-numericquestionpropertyvalueautomation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EvaluationFormMultiSelectQuestionProperties

EvaluationFormNumericQuestionOption

All content copied from https://docs.aws.amazon.com/.
