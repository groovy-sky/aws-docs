This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::AIAgent AssociationConfiguration

The configuration for an Amazon Q in Connect Assistant Association.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AssociationConfigurationData" : AssociationConfigurationData,
  "AssociationId" : String,
  "AssociationType" : String
}

```

### YAML

```yaml

  AssociationConfigurationData:
    AssociationConfigurationData
  AssociationId: String
  AssociationType: String

```

## Properties

`AssociationConfigurationData`

A typed union of the data of the configuration for an Amazon Q in Connect Assistant
Association.

_Required_: No

_Type_: [AssociationConfigurationData](aws-properties-wisdom-aiagent-associationconfigurationdata.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AssociationId`

The identifier of the association for this Association Configuration.

_Required_: No

_Type_: String

_Pattern_: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AssociationType`

The type of the association for this Association Configuration.

_Required_: No

_Type_: String

_Allowed values_: `KNOWLEDGE_BASE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AnswerRecommendationAIAgentConfiguration

AssociationConfigurationData

All content copied from https://docs.aws.amazon.com/.
