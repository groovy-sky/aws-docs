This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::KnowledgeBase S3VectorsConfiguration

Contains the storage configuration of the knowledge base for S3 vectors.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IndexArn" : String,
  "IndexName" : String,
  "VectorBucketArn" : String
}

```

### YAML

```yaml

  IndexArn: String
  IndexName: String
  VectorBucketArn: String

```

## Properties

`IndexArn`

The Amazon Resource Name (ARN) of the vector index used for the knowledge base. This ARN identifies the specific vector index resource within Amazon Bedrock.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IndexName`

The name of the vector index used for the knowledge base. This name identifies the vector index within the Amazon Bedrock service.

_Required_: No

_Type_: String

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VectorBucketArn`

The Amazon Resource Name (ARN) of the S3 bucket where vector embeddings are stored. This bucket contains the vector data used by the knowledge base.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3Location

SqlKnowledgeBaseConfiguration

All content copied from https://docs.aws.amazon.com/.
