This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot TextInputSpecification

Specifies the text input specifications.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "StartTimeoutMs" : Integer
}

```

### YAML

```yaml

  StartTimeoutMs: Integer

```

## Properties

`StartTimeoutMs`

Time for which a bot waits before re-prompting a customer for text input.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TestBotAliasSettings

TextLogDestination

All content copied from https://docs.aws.amazon.com/.
