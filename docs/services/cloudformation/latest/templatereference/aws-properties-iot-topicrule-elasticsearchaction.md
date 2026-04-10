This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::TopicRule ElasticsearchAction

Describes an action that writes data to an Amazon OpenSearch Service
domain.

###### Note

The `Elasticsearch` action can only be used by existing rule actions. To create a
new rule action or to update an existing rule action, use the
`OpenSearch` rule action instead. For more information, see [OpenSearchAction](../../../../reference/iot/latest/apireference/api-opensearchaction.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Endpoint" : String,
  "Id" : String,
  "Index" : String,
  "RoleArn" : String,
  "Type" : String
}

```

### YAML

```yaml

  Endpoint: String
  Id: String
  Index: String
  RoleArn: String
  Type: String

```

## Properties

`Endpoint`

The endpoint of your OpenSearch domain.

_Required_: Yes

_Type_: String

_Pattern_: `https?://.*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Id`

The unique identifier for the document you are storing.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Index`

The index where you want to store your data.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The IAM role ARN that has access to OpenSearch.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of document you are storing.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DynamoDBv2Action

FirehoseAction

All content copied from https://docs.aws.amazon.com/.
