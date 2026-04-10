This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot ObfuscationSetting

Determines whether Amazon Lex obscures slot values in conversation logs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ObfuscationSettingType" : String
}

```

### YAML

```yaml

  ObfuscationSettingType: String

```

## Properties

`ObfuscationSettingType`

Value that determines whether Amazon Lex obscures slot values in
conversation logs. The default is to obscure the values.

_Required_: Yes

_Type_: String

_Allowed values_: `None | DefaultObfuscation`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NluImprovementSpecification

OpensearchConfiguration

All content copied from https://docs.aws.amazon.com/.
