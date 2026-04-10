This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EKS::Nodegroup UpdateConfig

The update configuration for the node group.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxUnavailable" : Number,
  "MaxUnavailablePercentage" : Number,
  "UpdateStrategy" : String
}

```

### YAML

```yaml

  MaxUnavailable: Number
  MaxUnavailablePercentage: Number
  UpdateStrategy: String

```

## Properties

`MaxUnavailable`

The maximum number of nodes unavailable at once during a version update. Nodes are
updated in parallel. This value or `maxUnavailablePercentage` is required to
have a value.The maximum number is 100.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxUnavailablePercentage`

The maximum percentage of nodes unavailable during a version update. This percentage
of nodes are updated in parallel, up to 100 nodes at once. This value or
`maxUnavailable` is required to have a value.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UpdateStrategy`

The configuration for the behavior to follow during a node group version update of
this managed node group. You choose between two possible strategies for replacing nodes
during an [`UpdateNodegroupVersion`](../../../../reference/eks/latest/apireference/api-updatenodegroupversion.md) action.

An Amazon EKS managed node group updates by replacing nodes with new nodes of newer AMI
versions in parallel. The _update strategy_ changes the managed node
update behavior of the managed node group for each quantity. The
_default_ strategy has guardrails to protect you from
misconfiguration and launches the new instances first, before terminating the old
instances. The _minimal_ strategy removes the guardrails and
terminates the old instances before launching the new instances. This minimal strategy
is useful in scenarios where you are constrained to resources or costs (for example,
with hardware accelerators such as GPUs).

_Required_: No

_Type_: String

_Allowed values_: `DEFAULT | MINIMAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Taint

WarmPoolConfig

All content copied from https://docs.aws.amazon.com/.
