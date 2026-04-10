This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::EvaluationForm EvaluationFormMultiSelectQuestionOption

An option for a multi-select question in an evaluation form.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RefId" : String,
  "Text" : String
}

```

### YAML

```yaml

  RefId: String
  Text: String

```

## Properties

`RefId`

Reference identifier for this option.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9._-]{1,40}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Text`

Display text for this option.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EvaluationFormMultiSelectQuestionAutomationOption

EvaluationFormMultiSelectQuestionProperties

All content copied from https://docs.aws.amazon.com/.
