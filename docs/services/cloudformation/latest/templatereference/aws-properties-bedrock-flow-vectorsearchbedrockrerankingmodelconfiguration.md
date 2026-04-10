This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Flow VectorSearchBedrockRerankingModelConfiguration

Configuration for the Amazon Bedrock foundation model used for reranking vector search results. This specifies which model to use and any additional parameters required by the model.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdditionalModelRequestFields" : Json,
  "ModelArn" : String
}

```

### YAML

```yaml

  AdditionalModelRequestFields: Json
  ModelArn: String

```

## Properties

`AdditionalModelRequestFields`

A list of additional fields to include in the model request during reranking. These fields provide extra context or configuration options specific to the selected foundation model.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelArn`

The Amazon Resource Name (ARN) of the foundation model to use for reranking. This model processes the query and search results to determine a more relevant ordering.

_Required_: Yes

_Type_: String

_Pattern_: `^(arn:aws(-[^:]+)?:bedrock:[a-z0-9-]{1,20}::foundation-model/(.*))?$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VectorSearchBedrockRerankingConfiguration

VectorSearchRerankingConfiguration

All content copied from https://docs.aws.amazon.com/.
