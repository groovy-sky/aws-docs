This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NeptuneGraph::Graph

The `AWS::NeptuneGraph::Graph` resource creates an
Amazon Neptune Analytics graph. Amazon Neptune Analytics
is a memory-optimized graph database engine for analytics. For more information, see
[Amazon Neptune Analytics](https://docs.aws.amazon.com/neptune-analytics/latest/userguide/what-is-neptune-analytics.html).

You can use `AWS::NeptuneGraph::Graph.DeletionProtection` to help guard against unintended
deletion of your graph.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::NeptuneGraph::Graph",
  "Properties" : {
      "DeletionProtection" : Boolean,
      "GraphName" : String,
      "ProvisionedMemory" : Integer,
      "PublicConnectivity" : Boolean,
      "ReplicaCount" : Integer,
      "Tags" : [ Tag, ... ],
      "VectorSearchConfiguration" : VectorSearchConfiguration
    }
}

```

### YAML

```yaml

Type: AWS::NeptuneGraph::Graph
Properties:
  DeletionProtection: Boolean
  GraphName: String
  ProvisionedMemory: Integer
  PublicConnectivity: Boolean
  ReplicaCount: Integer
  Tags:
    - Tag
  VectorSearchConfiguration:
    VectorSearchConfiguration

```

## Properties

`DeletionProtection`

A value that indicates whether the graph has deletion protection enabled.
The graph can't be deleted when deletion protection is enabled.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GraphName`

The graph name. For example: `my-graph-1`.

The name must contain from 1 to 63 letters, numbers, or hyphens, and its
first character must be a letter. It cannot end with a hyphen or contain two
consecutive hyphens.

If you don't specify a graph name, a unique graph name is generated for you using the prefix `graph-for`,
followed by a combination of `Stack Name` and a `UUID`.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-z][a-zA-Z0-9]*(-[a-zA-Z0-9]+)*$`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProvisionedMemory`

The provisioned memory-optimized Neptune Capacity Units (m-NCUs) to use for the graph.

Min = 16

_Required_: Yes

_Type_: Integer

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`PublicConnectivity`

Specifies whether or not the graph can be reachable over the internet. All access to graphs is IAM authenticated.

When the graph is publicly available, its domain name system (DNS) endpoint resolves to the public IP address
from the internet. When the graph isn't publicly available, you need to create a `PrivateGraphEndpoint` in a
given VPC to ensure the DNS name resolves to a private IP address that is reachable from the VPC.

Default: If not specified, the default value is false.

###### Note

If enabling public connectivity for the first time, there will be a delay while it is enabled.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReplicaCount`

The number of replicas in other AZs.

Default: If not specified, the default value is 1.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Adds metadata tags to the new graph.
These tags can also be used with cost allocation reporting, or used in a Condition statement in an IAM policy.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-neptunegraph-graph-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VectorSearchConfiguration`

Specifies the number of dimensions for vector embeddings that will be loaded into the graph.
The value is specified as `dimension=` value. Max = 65,535

_Required_: No

_Type_: [VectorSearchConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-neptunegraph-graph-vectorsearchconfiguration.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the GraphId.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Endpoint`

The connection endpoint for the graph. For example: `g-12a3bcdef4.us-east-1.neptune-graph.amazonaws.com`

`GraphArn`

The ARN of the graph. For example: `arn:aws:neptune-graph:us-east-1:111122223333:graph/g-12a3bcdef4`

`GraphId`

The ID of the graph. For example: `g-12a3bcdef4`

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon Neptune Analytics

Tag
