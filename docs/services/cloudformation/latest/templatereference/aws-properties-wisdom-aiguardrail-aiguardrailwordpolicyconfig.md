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

_Type_: Array of [GuardrailManagedWordsConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wisdom-aiguardrail-guardrailmanagedwordsconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WordsConfig`

List of custom word configurations.

_Required_: No

_Type_: Array of [GuardrailWordConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wisdom-aiguardrail-guardrailwordconfig.html)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AIGuardrailTopicPolicyConfig

GuardrailContentFilterConfig
