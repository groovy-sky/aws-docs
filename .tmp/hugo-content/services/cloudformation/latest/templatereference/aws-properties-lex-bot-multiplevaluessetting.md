This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot MultipleValuesSetting

Indicates whether a slot can return multiple values.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllowMultipleValues" : Boolean
}

```

### YAML

```yaml

  AllowMultipleValues: Boolean

```

## Properties

`AllowMultipleValues`

Indicates whether a slot can return multiple values. When
`true`, the slot may return more than one value in a
response. When `false`, the slot returns only a single
value.

Multi-value slots are only available in the en-US locale. If you set
this value to `true` in any other locale, Amazon Lex throws a
`ValidationException`.

If the `allowMutlipleValues` is not set, the default
value is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MessageGroup

NluImprovementSpecification

All content copied from https://docs.aws.amazon.com/.
