This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OpenSearchServerless::Index Parameters

Additional parameters for the k-NN algorithm.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EfConstruction" : Integer,
  "M" : Integer
}

```

### YAML

```yaml

  EfConstruction: Integer
  M: Integer

```

## Properties

`EfConstruction`

The size of the dynamic list used during k-NN graph creation.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`M`

Number of neighbors to consider during k-NN search.

_Required_: No

_Type_: Integer

_Minimum_: `2`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Method

PropertyMapping

All content copied from https://docs.aws.amazon.com/.
