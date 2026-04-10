This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::EvaluationForm EvaluationFormSection

Information about a section from an evaluation form. A section can contain sections
and/or questions. Evaluation forms can only contain sections and subsections (two level
nesting).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Instructions" : String,
  "Items" : [ EvaluationFormItem, ... ],
  "RefId" : String,
  "Title" : String,
  "Weight" : Number
}

```

### YAML

```yaml

  Instructions: String
  Items:
    - EvaluationFormItem
  RefId: String
  Title: String
  Weight: Number

```

## Properties

`Instructions`

The instructions of the section.

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Items`

The items of the section.

_Minimum_: 1

_Required_: No

_Type_: Array of [EvaluationFormItem](aws-properties-connect-evaluationform-evaluationformitem.md)

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RefId`

The identifier of the section. An identifier must be unique within the evaluation
form.

_Length Constraints_: Minimum length of 1. Maximum length of
40.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9._-]{1,40}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Title`

The title of the section.

_Length Constraints_: Minimum length of 1. Maximum length of
128.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

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

EvaluationFormQuestionTypeProperties

EvaluationFormSingleSelectQuestionAutomation

All content copied from https://docs.aws.amazon.com/.
