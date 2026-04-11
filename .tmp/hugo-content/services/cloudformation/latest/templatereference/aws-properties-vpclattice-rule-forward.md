This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VpcLattice::Rule Forward

The forward action. Traffic that matches the rule is forwarded to the specified target
groups.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TargetGroups" : [ WeightedTargetGroup, ... ]
}

```

### YAML

```yaml

  TargetGroups:
    - WeightedTargetGroup

```

## Properties

`TargetGroups`

The target groups. Traffic matching the rule is forwarded to the specified target groups.
With forward actions, you can assign a weight that controls the prioritization and selection of
each target group. This means that requests are distributed to individual target groups based on
their weights. For example, if two target groups have the same weight, each target group receives
half of the traffic.

The default value is 1. This means that if only one target group is provided, there is no
need to set the weight; 100% of the traffic goes to that target group.

_Required_: Yes

_Type_: Array of [WeightedTargetGroup](aws-properties-vpclattice-rule-weightedtargetgroup.md)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FixedResponse

HeaderMatch

All content copied from https://docs.aws.amazon.com/.
