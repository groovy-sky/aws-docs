This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Redshift::ScheduledAction ResizeClusterMessage

Describes a resize cluster operation. For example, a scheduled action to run the `ResizeCluster` API operation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Classic" : Boolean,
  "ClusterIdentifier" : String,
  "ClusterType" : String,
  "NodeType" : String,
  "NumberOfNodes" : Integer
}

```

### YAML

```yaml

  Classic: Boolean
  ClusterIdentifier: String
  ClusterType: String
  NodeType: String
  NumberOfNodes: Integer

```

## Properties

`Classic`

A boolean value indicating whether the resize operation is using the classic resize
process. If you don't provide this parameter or set the value to
`false`, the resize type is elastic.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClusterIdentifier`

The unique identifier for the cluster to resize.

_Required_: Yes

_Type_: String

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClusterType`

The new cluster type for the specified cluster.

_Required_: No

_Type_: String

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NodeType`

The new node type for the nodes you are adding. If not specified, the cluster's current node type is used.

_Required_: No

_Type_: String

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumberOfNodes`

The new number of nodes for the cluster. If not specified, the cluster's current number of nodes is used.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PauseClusterMessage

ResumeClusterMessage

All content copied from https://docs.aws.amazon.com/.
