This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataSource Transformation

A custom processing step for documents moving through a data source ingestion pipeline. To
process documents after they have been converted into chunks, set the step to apply to
`POST_CHUNKING`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "StepToApply" : String,
  "TransformationFunction" : TransformationFunction
}

```

### YAML

```yaml

  StepToApply: String
  TransformationFunction:
    TransformationFunction

```

## Properties

`StepToApply`

When the service applies the transformation.

_Required_: Yes

_Type_: String

_Allowed values_: `POST_CHUNKING`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TransformationFunction`

A Lambda function that processes documents.

_Required_: Yes

_Type_: [TransformationFunction](aws-properties-bedrock-datasource-transformationfunction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SharePointSourceConfiguration

TransformationFunction

All content copied from https://docs.aws.amazon.com/.
