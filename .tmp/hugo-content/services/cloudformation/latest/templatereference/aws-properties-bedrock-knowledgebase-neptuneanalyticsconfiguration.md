This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::KnowledgeBase NeptuneAnalyticsConfiguration

Contains details about the storage configuration of the knowledge base in
Amazon Neptune Analytics. For more information, see [Create a vector index \
in Amazon Neptune Analytics](../../../bedrock/latest/userguide/knowledge-base-setup-neptune.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FieldMapping" : NeptuneAnalyticsFieldMapping,
  "GraphArn" : String
}

```

### YAML

```yaml

  FieldMapping:
    NeptuneAnalyticsFieldMapping
  GraphArn: String

```

## Properties

`FieldMapping`

Contains the names of the fields to which to map information about the vector store.

_Required_: Yes

_Type_: [NeptuneAnalyticsFieldMapping](aws-properties-bedrock-knowledgebase-neptuneanalyticsfieldmapping.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`GraphArn`

The Amazon Resource Name (ARN) of the Neptune Analytics vector store.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws(|-cn|-us-gov):neptune-graph:[a-zA-Z0-9-]*:[0-9]{12}:graph\/g-[a-zA-Z0-9]{10}$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MongoDbAtlasFieldMapping

NeptuneAnalyticsFieldMapping

All content copied from https://docs.aws.amazon.com/.
