This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::EvaluationForm EvaluationFormItemEnablementSourceValue

An enablement expression source value.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RefId" : String,
  "Type" : String
}

```

### YAML

```yaml

  RefId: String
  Type: String

```

## Properties

`RefId`

A referenceId of the source value.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9._-]{1,40}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

A type of source item value.

_Required_: No

_Type_: String

_Allowed values_: `OPTION_REF_ID`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EvaluationFormItemEnablementSource

EvaluationFormLanguageConfiguration

All content copied from https://docs.aws.amazon.com/.
