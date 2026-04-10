This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMR::Cluster PlacementType

`PlacementType` is a property of the `AWS::EMR::Cluster` resource. `PlacementType` determines the Amazon EC2 Availability Zone configuration of the cluster (job flow).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AvailabilityZone" : String
}

```

### YAML

```yaml

  AvailabilityZone: String

```

## Properties

`AvailabilityZone`

The Amazon EC2 Availability Zone for the cluster. `AvailabilityZone`
is used for uniform instance groups, while `AvailabilityZones` (plural) is used
for instance fleets.

_Required_: Yes

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `10280`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PlacementGroupConfig

ScalingAction

All content copied from https://docs.aws.amazon.com/.
