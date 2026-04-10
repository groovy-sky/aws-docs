This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::PlacementGroup

Specifies a placement group in which to launch instances. The strategy of the placement
group determines how the instances are organized within the group.

A `cluster` placement group is a logical grouping of instances within a
single Availability Zone that benefit from low network latency, high network throughput. A
`spread` placement group places instances on distinct hardware. A
`partition` placement group places groups of instances in different
partitions, where instances in one partition do not share the same hardware with instances
in another partition.

For more information, see [Placement Groups](../../../ec2/latest/userguide/placement-groups.md) in the
_Amazon EC2 User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::PlacementGroup",
  "Properties" : {
      "PartitionCount" : Integer,
      "SpreadLevel" : String,
      "Strategy" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::EC2::PlacementGroup
Properties:
  PartitionCount: Integer
  SpreadLevel: String
  Strategy: String
  Tags:
    - Tag

```

## Properties

`PartitionCount`

The number of partitions. Valid only when **Strategy** is
set to `partition`.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SpreadLevel`

Determines how placement groups spread instances.

- Host – You can use `host` only with Outpost placement
groups.

- Rack – No usage restrictions.

_Required_: No

_Type_: String

_Allowed values_: `host | rack`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Strategy`

The placement strategy.

_Required_: No

_Type_: String

_Allowed values_: `cluster | spread | partition`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags to apply to the new placement group.

_Required_: No

_Type_: Array of [Tag](aws-properties-ec2-placementgroup-tag.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the placement group.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`GroupName`

The name of the placement group.

## Examples

### Create a placement group

The following example declares a placement group with a cluster placement
strategy.

#### JSON

```json

"PlacementGroup" : {
   "Type" : "AWS::EC2::PlacementGroup",
   "Properties" : {
            "Strategy" : "cluster"
   }
}
```

#### YAML

```yaml

PlacementGroup:
  Type: AWS::EC2::PlacementGroup
   Properties:
     Strategy: cluster
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::EC2::NetworkPerformanceMetricSubscription

Tag

All content copied from https://docs.aws.amazon.com/.
