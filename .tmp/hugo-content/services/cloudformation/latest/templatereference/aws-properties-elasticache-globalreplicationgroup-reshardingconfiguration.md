This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElastiCache::GlobalReplicationGroup ReshardingConfiguration

A list of `PreferredAvailabilityZones` objects that specifies the
configuration of a node group in the resharded cluster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NodeGroupId" : String,
  "PreferredAvailabilityZones" : [ String, ... ]
}

```

### YAML

```yaml

  NodeGroupId: String
  PreferredAvailabilityZones:
    - String

```

## Properties

`NodeGroupId`

Either the ElastiCache supplied 4-digit id or a user supplied id for the
node group these configuration values apply to.

_Required_: No

_Type_: String

_Pattern_: `\d+`

_Minimum_: `1`

_Maximum_: `4`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PreferredAvailabilityZones`

A list of preferred availability zones for the nodes in this cluster.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RegionalConfiguration

AWS::ElastiCache::ParameterGroup

All content copied from https://docs.aws.amazon.com/.
