This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OpenSearchServerless::Collection VectorOptions

Configuration options for vector search capabilities in an OpenSearch Serverless collection.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ServerlessVectorAcceleration" : String
}

```

### YAML

```yaml

  ServerlessVectorAcceleration: String

```

## Properties

`ServerlessVectorAcceleration`

Specifies whether serverless vector acceleration is enabled for the collection.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED | ALLOWED`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::OpenSearchServerless::CollectionGroup
