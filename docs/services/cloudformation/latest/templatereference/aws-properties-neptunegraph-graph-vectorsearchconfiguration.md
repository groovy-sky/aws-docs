This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NeptuneGraph::Graph VectorSearchConfiguration

The vector-search configuration for the graph, which specifies the vector dimension
to use in the vector index, if any.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "VectorSearchDimension" : Integer
}

```

### YAML

```yaml

  VectorSearchDimension: Integer

```

## Properties

`VectorSearchDimension`

The number of dimensions.

_Required_: Yes

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::NeptuneGraph::PrivateGraphEndpoint
