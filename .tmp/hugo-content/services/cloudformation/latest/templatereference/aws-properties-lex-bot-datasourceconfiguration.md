This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot DataSourceConfiguration

Contains details about the configuration of the knowledge store used for the `AMAZON.QnAIntent`. You must have already created the knowledge store and indexed the documents within it.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BedrockKnowledgeStoreConfiguration" : BedrockKnowledgeStoreConfiguration,
  "KendraConfiguration" : QnAKendraConfiguration,
  "OpensearchConfiguration" : OpensearchConfiguration
}

```

### YAML

```yaml

  BedrockKnowledgeStoreConfiguration:
    BedrockKnowledgeStoreConfiguration
  KendraConfiguration:
    QnAKendraConfiguration
  OpensearchConfiguration:
    OpensearchConfiguration

```

## Properties

`BedrockKnowledgeStoreConfiguration`

Contains details about the configuration of the Amazon Bedrock knowledge base used for the `AMAZON.QnAIntent`. To set up a knowledge base, follow the steps at [Building a knowledge base](../../../bedrock/latest/userguide/knowledge-base.md).

_Required_: No

_Type_: [BedrockKnowledgeStoreConfiguration](aws-properties-lex-bot-bedrockknowledgestoreconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KendraConfiguration`

Contains details about the configuration of the Amazon Kendra index used for the `AMAZON.QnAIntent`. To create a Amazon Kendra index, follow the steps at [Creating an index](../../../kendra/latest/dg/create-index.md).

_Required_: No

_Type_: [QnAKendraConfiguration](aws-properties-lex-bot-qnakendraconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OpensearchConfiguration`

Contains details about the configuration of the Amazon OpenSearch Service database used for the `AMAZON.QnAIntent`. To create a domain, follow the steps at [Creating and managing Amazon OpenSearch Service domains](../../../opensearch-service/latest/developerguide/createupdatedomains.md).

_Required_: No

_Type_: [OpensearchConfiguration](aws-properties-lex-bot-opensearchconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataPrivacy

DeepgramSpeechModelConfig

All content copied from https://docs.aws.amazon.com/.
