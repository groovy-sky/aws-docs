This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataSource ContextEnrichmentConfiguration

Context enrichment configuration is used to provide additional context to the RAG application.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BedrockFoundationModelConfiguration" : BedrockFoundationModelContextEnrichmentConfiguration,
  "Type" : String
}

```

### YAML

```yaml

  BedrockFoundationModelConfiguration:
    BedrockFoundationModelContextEnrichmentConfiguration
  Type: String

```

## Properties

`BedrockFoundationModelConfiguration`

The configuration of the Amazon Bedrock foundation model used for context enrichment.

_Required_: No

_Type_: [BedrockFoundationModelContextEnrichmentConfiguration](aws-properties-bedrock-datasource-bedrockfoundationmodelcontextenrichmentconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The method used for context enrichment. It must be Amazon Bedrock foundation models.

_Required_: Yes

_Type_: String

_Allowed values_: `BEDROCK_FOUNDATION_MODEL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConfluenceSourceConfiguration

CrawlFilterConfiguration

All content copied from https://docs.aws.amazon.com/.
