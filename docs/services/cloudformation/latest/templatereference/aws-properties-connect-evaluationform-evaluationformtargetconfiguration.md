This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::EvaluationForm EvaluationFormTargetConfiguration

Configuration that specifies the target for an evaluation form.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContactInteractionType" : String
}

```

### YAML

```yaml

  ContactInteractionType: String

```

## Properties

`ContactInteractionType`

The contact interaction type for this evaluation form.

_Required_: Yes

_Type_: String

_Allowed values_: `AGENT | AUTOMATED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

EvaluationFormSingleSelectQuestionProperties

EvaluationFormTextQuestionAutomation
