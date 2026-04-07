This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMR::Cluster PlacementGroupConfig

Placement group configuration for an Amazon EMR cluster. The configuration
specifies the placement strategy that can be applied to instance roles during cluster
creation.

To use this configuration, consider attaching managed policy
AmazonElasticMapReducePlacementGroupPolicy to the Amazon EMR role.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InstanceRole" : String,
  "PlacementStrategy" : String
}

```

### YAML

```yaml

  InstanceRole: String
  PlacementStrategy: String

```

## Properties

`InstanceRole`

Role of the instance in the cluster.

Starting with Amazon EMR release 5.23.0, the only supported instance role is
`MASTER`.

_Required_: Yes

_Type_: String

_Allowed values_: `MASTER | CORE | TASK`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PlacementStrategy`

Amazon EC2 Placement Group strategy associated with instance role.

Starting with Amazon EMR release 5.23.0, the only supported placement strategy
is `SPREAD` for the `MASTER` instance role.

_Required_: No

_Type_: String

_Allowed values_: `SPREAD | PARTITION | CLUSTER | NONE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

OnDemandResizingSpecification

PlacementType
