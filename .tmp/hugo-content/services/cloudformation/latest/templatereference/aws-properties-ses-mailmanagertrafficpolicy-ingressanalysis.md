This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::MailManagerTrafficPolicy IngressAnalysis

The Add On ARN and its returned value that is evaluated in a policy statement's
conditional expression to either deny or block the incoming email.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Analyzer" : String,
  "ResultField" : String
}

```

### YAML

```yaml

  Analyzer: String
  ResultField: String

```

## Properties

`Analyzer`

The Amazon Resource Name (ARN) of an Add On.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9:_/+=,@.#-]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResultField`

The returned value from an Add On.

_Required_: Yes

_Type_: String

_Pattern_: `^(addon\.)?[\sa-zA-Z0-9_]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SES::MailManagerTrafficPolicy

IngressBooleanExpression

All content copied from https://docs.aws.amazon.com/.
