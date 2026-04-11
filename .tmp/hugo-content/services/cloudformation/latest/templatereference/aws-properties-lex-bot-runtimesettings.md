This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot RuntimeSettings

Contains specifications about the Amazon Lex runtime generative AI capabilities from Amazon Bedrock that you can turn on for your bot.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NluImprovementSpecification" : NluImprovementSpecification,
  "SlotResolutionImprovementSpecification" : SlotResolutionImprovementSpecification
}

```

### YAML

```yaml

  NluImprovementSpecification:
    NluImprovementSpecification
  SlotResolutionImprovementSpecification:
    SlotResolutionImprovementSpecification

```

## Properties

`NluImprovementSpecification`

Property description not available.

_Required_: No

_Type_: [NluImprovementSpecification](aws-properties-lex-bot-nluimprovementspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SlotResolutionImprovementSpecification`

Property description not available.

_Required_: No

_Type_: [SlotResolutionImprovementSpecification](aws-properties-lex-bot-slotresolutionimprovementspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResponseSpecification

S3BucketLogDestination

All content copied from https://docs.aws.amazon.com/.
