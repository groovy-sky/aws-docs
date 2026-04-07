This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Guardrail WordPolicyConfig

Contains details about the word policy to configured for the guardrail.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ManagedWordListsConfig" : [ ManagedWordsConfig, ... ],
  "WordsConfig" : [ WordConfig, ... ]
}

```

### YAML

```yaml

  ManagedWordListsConfig:
    - ManagedWordsConfig
  WordsConfig:
    - WordConfig

```

## Properties

`ManagedWordListsConfig`

A list of managed words to configure for the guardrail.

_Required_: No

_Type_: Array of [ManagedWordsConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-guardrail-managedwordsconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WordsConfig`

A list of words to configure for the guardrail.

_Required_: No

_Type_: Array of [WordConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-guardrail-wordconfig.html)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

WordConfig

AWS::Bedrock::GuardrailVersion
