This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OpenSearchServerless::CollectionGroup

Creates a collection group within OpenSearch Serverless. Collection groups let you manage OpenSearch Compute Units (OCUs) at a group level, with multiple collections sharing the group's capacity limits.

For more information, see [Managing collection groups](../../../opensearch-service/latest/developerguide/serverless-collection-groups.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::OpenSearchServerless::CollectionGroup",
  "Properties" : {
      "CapacityLimits" : CapacityLimits,
      "Description" : String,
      "Name" : String,
      "StandbyReplicas" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::OpenSearchServerless::CollectionGroup
Properties:
  CapacityLimits:
    CapacityLimits
  Description: String
  Name: String
  StandbyReplicas: String
  Tags:
    - Tag

```

## Properties

`CapacityLimits`

The maximum capacity limits for all OpenSearch Serverless collections, in OpenSearch Compute Units
(OCUs). These limits are used to scale your collections based on the current workload.
For more information, see [Managing\
capacity limits for Amazon OpenSearch Serverless](../../../opensearch-service/latest/developerguide/serverless-scaling.md).

_Required_: No

_Type_: [CapacityLimits](aws-properties-opensearchserverless-collectiongroup-capacitylimits.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the collection group.

_Required_: No

_Type_: String

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the collection group.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-z][a-z0-9-]{2,31}$`

_Minimum_: `3`

_Maximum_: `32`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StandbyReplicas`

Indicates whether standby replicas are used for the collection group. Can be either `ENABLED` or `DISABLED`.

_Required_: Yes

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of key-value pairs to apply to this resource.

_Required_: No

_Type_: Array of [Tag](aws-properties-opensearchserverless-collectiongroup-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the collection group ID. For more information about using the `Ref` function, see [Ref](../userguide/intrinsic-function-reference-ref.md).

### Fn::GetAtt

Details about a collection group.

`Arn`

The Amazon Resource Name (ARN) of the collection group.

`Id`

The unique identifier of the collection group.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VectorOptions

CapacityLimits

All content copied from https://docs.aws.amazon.com/.
