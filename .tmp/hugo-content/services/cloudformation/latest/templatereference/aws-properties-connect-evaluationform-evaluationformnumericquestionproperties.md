This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::EvaluationForm EvaluationFormNumericQuestionProperties

Information about properties for a numeric question in an evaluation form.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Automation" : EvaluationFormNumericQuestionAutomation,
  "MaxValue" : Integer,
  "MinValue" : Integer,
  "Options" : [ EvaluationFormNumericQuestionOption, ... ]
}

```

### YAML

```yaml

  Automation:
    EvaluationFormNumericQuestionAutomation
  MaxValue: Integer
  MinValue: Integer
  Options:
    - EvaluationFormNumericQuestionOption

```

## Properties

`Automation`

The automation properties of the numeric question.

_Required_: No

_Type_: [EvaluationFormNumericQuestionAutomation](aws-properties-connect-evaluationform-evaluationformnumericquestionautomation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxValue`

The maximum answer value.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinValue`

The minimum answer value.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Options`

The scoring options of the numeric question.

_Required_: No

_Type_: Array of [EvaluationFormNumericQuestionOption](aws-properties-connect-evaluationform-evaluationformnumericquestionoption.md)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EvaluationFormNumericQuestionOption

EvaluationFormQuestion

All content copied from https://docs.aws.amazon.com/.
