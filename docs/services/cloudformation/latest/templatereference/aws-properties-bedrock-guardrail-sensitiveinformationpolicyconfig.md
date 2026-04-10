This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Guardrail SensitiveInformationPolicyConfig

Contains details about PII entities and regular expressions to configure for the guardrail.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PiiEntitiesConfig" : [ PiiEntityConfig, ... ],
  "RegexesConfig" : [ RegexConfig, ... ]
}

```

### YAML

```yaml

  PiiEntitiesConfig:
    - PiiEntityConfig
  RegexesConfig:
    - RegexConfig

```

## Properties

`PiiEntitiesConfig`

A list of PII entities to configure to the guardrail.

_Required_: No

_Type_: Array of [PiiEntityConfig](aws-properties-bedrock-guardrail-piientityconfig.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RegexesConfig`

A list of regular expressions to configure to the guardrail.

_Required_: No

_Type_: Array of [RegexConfig](aws-properties-bedrock-guardrail-regexconfig.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RegexConfig

Tag

All content copied from https://docs.aws.amazon.com/.
