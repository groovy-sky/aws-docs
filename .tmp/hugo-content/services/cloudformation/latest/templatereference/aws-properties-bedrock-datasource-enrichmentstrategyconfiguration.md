This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataSource EnrichmentStrategyConfiguration

The strategy used for performing context enrichment.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Method" : String
}

```

### YAML

```yaml

  Method: String

```

## Properties

`Method`

The method used for the context enrichment strategy.

_Required_: Yes

_Type_: String

_Allowed values_: `CHUNK_ENTITY_EXTRACTION`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataSourceConfiguration

FixedSizeChunkingConfiguration

All content copied from https://docs.aws.amazon.com/.
