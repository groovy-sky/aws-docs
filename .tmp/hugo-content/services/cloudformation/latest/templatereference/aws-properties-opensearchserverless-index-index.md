This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OpenSearchServerless::Index Index

An OpenSearch Serverless index resource

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Knn" : Boolean,
  "KnnAlgoParamEfSearch" : Integer,
  "RefreshInterval" : String
}

```

### YAML

```yaml

  Knn: Boolean
  KnnAlgoParamEfSearch: Integer
  RefreshInterval: String

```

## Properties

`Knn`

Enable or disable k-nearest neighbor search capability.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KnnAlgoParamEfSearch`

The size of the dynamic list for the nearest neighbors.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RefreshInterval`

How often to perform a refresh operation. For example, 1s or 5s.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::OpenSearchServerless::Index

IndexSettings

All content copied from https://docs.aws.amazon.com/.
