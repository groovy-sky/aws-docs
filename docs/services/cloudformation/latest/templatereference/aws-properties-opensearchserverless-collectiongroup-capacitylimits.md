This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OpenSearchServerless::CollectionGroup CapacityLimits

The maximum capacity limits for all OpenSearch Serverless collections, in OpenSearch Compute Units
(OCUs). These limits are used to scale your collections based on the current workload.
For more information, see [Managing\
capacity limits for Amazon OpenSearch Serverless](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-scaling.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxIndexingCapacityInOcu" : Number,
  "MaxSearchCapacityInOcu" : Number,
  "MinIndexingCapacityInOcu" : Number,
  "MinSearchCapacityInOcu" : Number
}

```

### YAML

```yaml

  MaxIndexingCapacityInOcu: Number
  MaxSearchCapacityInOcu: Number
  MinIndexingCapacityInOcu: Number
  MinSearchCapacityInOcu: Number

```

## Properties

`MaxIndexingCapacityInOcu`

The maximum indexing capacity for collections.

_Required_: No

_Type_: Number

_Minimum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxSearchCapacityInOcu`

The maximum search capacity for collections.

_Required_: No

_Type_: Number

_Minimum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinIndexingCapacityInOcu`

The minimum indexing capacity for collections.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinSearchCapacityInOcu`

The minimum search capacity for collections.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::OpenSearchServerless::CollectionGroup

Tag
