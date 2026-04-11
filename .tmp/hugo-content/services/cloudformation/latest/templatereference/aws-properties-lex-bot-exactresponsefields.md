This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot ExactResponseFields

Contains the names of the fields used for an exact response to the user.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AnswerField" : String,
  "QuestionField" : String
}

```

### YAML

```yaml

  AnswerField: String
  QuestionField: String

```

## Properties

`AnswerField`

The name of the field that contains the answer to the query made to the OpenSearch Service database.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `5000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QuestionField`

The name of the field that contains the query made to the OpenSearch Service database.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `5000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ErrorLogSettings

ExternalSourceSetting

All content copied from https://docs.aws.amazon.com/.
