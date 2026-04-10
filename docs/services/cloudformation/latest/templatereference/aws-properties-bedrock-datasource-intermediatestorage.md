This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataSource IntermediateStorage

A location for storing content from data sources temporarily as it is processed by
custom components in the ingestion pipeline.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3Location" : S3Location
}

```

### YAML

```yaml

  S3Location:
    S3Location

```

## Properties

`S3Location`

An S3 bucket path.

_Required_: Yes

_Type_: [S3Location](aws-properties-bedrock-datasource-s3location.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HierarchicalChunkingLevelConfiguration

ParsingConfiguration

All content copied from https://docs.aws.amazon.com/.
