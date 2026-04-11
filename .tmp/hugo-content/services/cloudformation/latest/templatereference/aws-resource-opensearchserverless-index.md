This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OpenSearchServerless::Index

An OpenSearch Serverless index resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::OpenSearchServerless::Index",
  "Properties" : {
      "CollectionEndpoint" : String,
      "IndexName" : String,
      "Mappings" : Mappings,
      "Settings" : IndexSettings
    }
}

```

### YAML

```yaml

Type: AWS::OpenSearchServerless::Index
Properties:
  CollectionEndpoint: String
  IndexName: String
  Mappings:
    Mappings
  Settings:
    IndexSettings

```

## Properties

`CollectionEndpoint`

The endpoint for the collection.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IndexName`

The name of the OpenSearch Serverless index.

_Required_: Yes

_Type_: String

_Pattern_: `^(?![_-])[a-z][a-z0-9_-]*$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Mappings`

Index mappings for the OpenSearch Serverless index.

_Required_: No

_Type_: [Mappings](aws-properties-opensearchserverless-index-mappings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Settings`

Index settings for the OpenSearch Serverless index.

_Required_: No

_Type_: [IndexSettings](aws-properties-opensearchserverless-index-indexsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

The index name and the collection endpoint in the following format: index name\|endpoint.
Here's an example:

1234TestIndex\|12345678TestEndPoint.us-east-1.aoss.amazonaws.com.

### Fn::GetAtt

`GetAtt` returns a value for a specified attribute of this type. For more information, see [Fn::GetAtt](../userguide/intrinsic-function-reference-getatt.md). The following are the available attributes and sample return values.

`Uuid`

The unique identifier for the index.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Index

All content copied from https://docs.aws.amazon.com/.
