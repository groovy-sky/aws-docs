This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataSource CustomTransformationConfiguration

Settings for customizing steps in the data source content ingestion pipeline.

You can configure the data source to process documents with a Lambda function after
they are parsed and converted into chunks. When you add a post-chunking transformation,
the service stores chunked documents in an S3 bucket and invokes a Lambda function to process
them.

To process chunked documents with a Lambda function, define an S3 bucket path for input
and output objects, and a transformation that specifies the Lambda function to invoke. You can
use the Lambda function to customize how chunks are split, and the metadata for each chunk.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IntermediateStorage" : IntermediateStorage,
  "Transformations" : [ Transformation, ... ]
}

```

### YAML

```yaml

  IntermediateStorage:
    IntermediateStorage
  Transformations:
    - Transformation

```

## Properties

`IntermediateStorage`

An S3 bucket path for input and output objects.

_Required_: Yes

_Type_: [IntermediateStorage](aws-properties-bedrock-datasource-intermediatestorage.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Transformations`

A Lambda function that processes documents.

_Required_: Yes

_Type_: Array of [Transformation](aws-properties-bedrock-datasource-transformation.md)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CrawlFilterConfiguration

DataSourceConfiguration

All content copied from https://docs.aws.amazon.com/.
