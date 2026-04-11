This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::JobDefinition NodeRangeProperty

This is an object that represents the properties of the node range for a multi-node parallel
job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConsumableResourceProperties" : ConsumableResourceProperties,
  "Container" : MultiNodeContainerProperties,
  "EcsProperties" : MultiNodeEcsProperties,
  "EksProperties" : EksProperties,
  "InstanceTypes" : [ String, ... ],
  "TargetNodes" : String
}

```

### YAML

```yaml

  ConsumableResourceProperties:
    ConsumableResourceProperties
  Container:
    MultiNodeContainerProperties
  EcsProperties:
    MultiNodeEcsProperties
  EksProperties:
    EksProperties
  InstanceTypes:
    - String
  TargetNodes: String

```

## Properties

`ConsumableResourceProperties`

Contains a list of consumable resources required by a job.

_Required_: No

_Type_: [ConsumableResourceProperties](aws-properties-batch-jobdefinition-consumableresourceproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Container`

The container details for the node range.

_Required_: No

_Type_: [MultiNodeContainerProperties](aws-properties-batch-jobdefinition-multinodecontainerproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EcsProperties`

This is an object that represents the properties of the node range for a multi-node parallel
job.

_Required_: No

_Type_: [MultiNodeEcsProperties](aws-properties-batch-jobdefinition-multinodeecsproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EksProperties`

This is an object that represents the properties of the node range for a multi-node parallel job.

_Required_: No

_Type_: [EksProperties](aws-properties-batch-jobdefinition-eksproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceTypes`

The instance types of the underlying host infrastructure of a multi-node parallel
job.

###### Note

This parameter isn't applicable to jobs that are running on Fargate resources.

In addition, this list object is currently limited to one element.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetNodes`

The range of nodes, using node index values. A range of `0:3` indicates nodes
with index values of `0` through `3`. If the starting range value is
omitted ( `:n`), then `0` is used to start the range. If the ending range
value is omitted ( `n:`), then the highest possible node index is used to end the
range. Your accumulative node ranges must account for all nodes ( `0:n`). You can nest
node ranges (for example, `0:10` and `4:5`). In this case, the
`4:5` range properties override the `0:10` properties.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Multi-node Parallel\
Jobs](../../../batch/latest/userguide/multi-node-parallel-jobs.md) in the _AWS Batch User Guide_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NodeProperties

RepositoryCredentials

All content copied from https://docs.aws.amazon.com/.
