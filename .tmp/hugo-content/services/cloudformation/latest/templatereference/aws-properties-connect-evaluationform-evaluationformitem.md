This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::EvaluationForm EvaluationFormItem

Items that are part of the evaluation form. The total number of sections and questions
must not exceed 100 each. Questions must be contained in a section.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Question" : EvaluationFormQuestion,
  "Section" : EvaluationFormSection
}

```

### YAML

```yaml

  Question:
    EvaluationFormQuestion
  Section:
    EvaluationFormSection

```

## Properties

`Question`

The information of the question.

_Required_: No

_Type_: [EvaluationFormQuestion](aws-properties-connect-evaluationform-evaluationformquestion.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Section`

The information of the section.

_Required_: No

_Type_: [EvaluationFormSection](aws-properties-connect-evaluationform-evaluationformsection.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EvaluationFormBaseItem

EvaluationFormItemEnablementCondition

All content copied from https://docs.aws.amazon.com/.
