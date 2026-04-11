This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OpenSearchServerless::Index Method

Configuration for k-NN search method.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Engine" : String,
  "Name" : String,
  "Parameters" : Parameters,
  "SpaceType" : String
}

```

### YAML

```yaml

  Engine: String
  Name: String
  Parameters:
    Parameters
  SpaceType: String

```

## Properties

`Engine`

The k-NN search engine to use

_Required_: No

_Type_: String

_Allowed values_: `nmslib | faiss | lucene`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The algorithm name for k-NN search.

_Required_: Yes

_Type_: String

_Allowed values_: `hnsw | ivf`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Parameters`

Additional parameters for the k-NN algorithm.

_Required_: No

_Type_: [Parameters](aws-properties-opensearchserverless-index-parameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SpaceType`

The distance function used for k-NN search.

_Required_: No

_Type_: String

_Allowed values_: `l2 | l1 | linf | cosinesimil | innerproduct | hamming`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Mappings

Parameters

All content copied from https://docs.aws.amazon.com/.
