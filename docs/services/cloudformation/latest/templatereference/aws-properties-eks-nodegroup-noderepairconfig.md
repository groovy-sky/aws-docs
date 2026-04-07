This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EKS::Nodegroup NodeRepairConfig

The node auto repair configuration for the node group.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enabled" : Boolean,
  "MaxParallelNodesRepairedCount" : Integer,
  "MaxParallelNodesRepairedPercentage" : Integer,
  "MaxUnhealthyNodeThresholdCount" : Integer,
  "MaxUnhealthyNodeThresholdPercentage" : Integer,
  "NodeRepairConfigOverrides" : [ NodeRepairConfigOverrides, ... ]
}

```

### YAML

```yaml

  Enabled: Boolean
  MaxParallelNodesRepairedCount: Integer
  MaxParallelNodesRepairedPercentage: Integer
  MaxUnhealthyNodeThresholdCount: Integer
  MaxUnhealthyNodeThresholdPercentage: Integer
  NodeRepairConfigOverrides:
    - NodeRepairConfigOverrides

```

## Properties

`Enabled`

Specifies whether to enable node auto repair for the node group. Node auto repair is
disabled by default.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxParallelNodesRepairedCount`

Specify the maximum number of nodes that can be repaired concurrently or in parallel,
expressed as a count of unhealthy nodes. This gives you finer-grained control over the
pace of node replacements. When using this, you cannot also set
`maxParallelNodesRepairedPercentage` at the same time.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxParallelNodesRepairedPercentage`

Specify the maximum number of nodes that can be repaired concurrently or in parallel,
expressed as a percentage of unhealthy nodes. This gives you finer-grained control over the
pace of node replacements. When using this, you cannot also set
`maxParallelNodesRepairedCount` at the same time.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxUnhealthyNodeThresholdCount`

Specify a count threshold of unhealthy nodes, above which node auto
repair actions will stop. When using this, you cannot also set
`maxUnhealthyNodeThresholdPercentage` at the same time.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxUnhealthyNodeThresholdPercentage`

Specify a percentage threshold of unhealthy nodes, above which node auto
repair actions will stop. When using this, you cannot also set
`maxUnhealthyNodeThresholdCount` at the same time.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NodeRepairConfigOverrides`

Specify granular overrides for specific repair actions. These overrides control the
repair action and the repair delay time before a node is considered eligible for repair.
If you use this, you must specify all the values.

_Required_: No

_Type_: [Array](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-eks-nodegroup-noderepairconfigoverrides.html) of [NodeRepairConfigOverrides](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-eks-nodegroup-noderepairconfigoverrides.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

LaunchTemplateSpecification

NodeRepairConfigOverrides
