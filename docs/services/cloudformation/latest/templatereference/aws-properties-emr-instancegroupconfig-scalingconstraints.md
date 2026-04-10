This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMR::InstanceGroupConfig ScalingConstraints

`ScalingConstraints` is a subproperty of the `AutoScalingPolicy` property type. `ScalingConstraints` defines the upper and lower EC2 instance limits for an automatic scaling policy. Automatic scaling activities triggered by automatic scaling rules will not cause an instance group to grow above or shrink below these limits.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxCapacity" : Integer,
  "MinCapacity" : Integer
}

```

### YAML

```yaml

  MaxCapacity: Integer
  MinCapacity: Integer

```

## Properties

`MaxCapacity`

The upper boundary of Amazon EC2 instances in an instance group beyond which
scaling activities are not allowed to grow. Scale-out activities will not add instances
beyond this boundary.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinCapacity`

The lower boundary of Amazon EC2 instances in an instance group below which
scaling activities are not allowed to shrink. Scale-in activities will not terminate
instances below this boundary.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ScalingAction

ScalingRule

All content copied from https://docs.aws.amazon.com/.
