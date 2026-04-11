This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VpcLattice::Listener WeightedTargetGroup

Describes the weight of a target group.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TargetGroupIdentifier" : String,
  "Weight" : Integer
}

```

### YAML

```yaml

  TargetGroupIdentifier: String
  Weight: Integer

```

## Properties

`TargetGroupIdentifier`

The ID of the target group.

_Required_: Yes

_Type_: String

_Pattern_: `^((tg-[0-9a-z]{17})|(arn:[a-z0-9\-]+:vpc-lattice:[a-zA-Z0-9\-]+:\d{12}:targetgroup/tg-[0-9a-z]{17}))$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Weight`

Only required if you specify multiple target groups for a forward action. The weight
determines how requests are distributed to the target group. For example, if you specify two
target groups, each with a weight of 10, each target group receives half the requests. If you
specify two target groups, one with a weight of 10 and the other with a weight of 20, the target
group with a weight of 20 receives twice as many requests as the other target group. If there's
only one target group specified, then the default value is 100.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `999`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::VpcLattice::ResourceConfiguration

All content copied from https://docs.aws.amazon.com/.
