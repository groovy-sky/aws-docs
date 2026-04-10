This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot DescriptiveBotBuilderSpecification

Contains specifications for the descriptive bot building feature.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BedrockModelSpecification" : BedrockModelSpecification,
  "Enabled" : Boolean
}

```

### YAML

```yaml

  BedrockModelSpecification:
    BedrockModelSpecification
  Enabled: Boolean

```

## Properties

`BedrockModelSpecification`

An object containing information about the Amazon Bedrock model used to interpret the prompt used in descriptive bot building.

_Required_: No

_Type_: [BedrockModelSpecification](aws-properties-lex-bot-bedrockmodelspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

Specifies whether the descriptive bot building feature is activated or not.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DefaultConditionalBranch

DialogAction

All content copied from https://docs.aws.amazon.com/.
