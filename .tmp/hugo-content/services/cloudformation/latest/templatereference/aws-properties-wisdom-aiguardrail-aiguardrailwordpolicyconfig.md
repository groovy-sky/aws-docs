This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::AIGuardrail AIGuardrailWordPolicyConfig

Word policy config for a guardrail.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ManagedWordListsConfig" : [ GuardrailManagedWordsConfig, ... ],
  "WordsConfig" : [ GuardrailWordConfig, ... ]
}

```

### YAML

```yaml

  ManagedWordListsConfig:
    - GuardrailManagedWordsConfig
  WordsConfig:
    - GuardrailWordConfig

```

## Properties

`ManagedWordListsConfig`

A config for the list of managed words.

_Required_: No

_Type_: Array of [GuardrailManagedWordsConfig](aws-properties-wisdom-aiguardrail-guardrailmanagedwordsconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WordsConfig`

List of custom word configurations.

_Required_: No

_Type_: Array of [GuardrailWordConfig](aws-properties-wisdom-aiguardrail-guardrailwordconfig.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AIGuardrailTopicPolicyConfig

GuardrailContentFilterConfig

All content copied from https://docs.aws.amazon.com/.
