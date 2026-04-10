This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::EvaluationForm AutomaticFailConfiguration

Information about automatic fail configuration for an evaluation form.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TargetSection" : String
}

```

### YAML

```yaml

  TargetSection: String

```

## Properties

`TargetSection`

The referenceId of the target section for auto failure.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9._-]{1,40}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AutoEvaluationConfiguration

EvaluationFormBaseItem

All content copied from https://docs.aws.amazon.com/.
