This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AutoScaling::AutoScalingGroup TagProperty

A structure that specifies a tag for the `Tags` property of [AWS::AutoScaling::AutoScalingGroup](../userguide/aws-resource-autoscaling-autoscalinggroup.md) resource.

For more information, see [Tag Auto Scaling groups and\
instances](../../../autoscaling/ec2/userguide/autoscaling-tagging.md) in the _Amazon EC2 Auto Scaling User Guide_. You can
find a sample template snippet in the [Examples](../userguide/aws-resource-autoscaling-autoscalinggroup.md#aws-resource-autoscaling-autoscalinggroup--examples) section of the `AWS::AutoScaling::AutoScalingGroup`
resource.

CloudFormation adds the following tags to all Auto Scaling groups and associated
instances:

- aws:cloudformation:stack-name

- aws:cloudformation:stack-id

- aws:cloudformation:logical-id

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "PropagateAtLaunch" : Boolean,
  "Value" : String
}

```

### YAML

```yaml

  Key: String
  PropagateAtLaunch: Boolean
  Value: String

```

## Properties

`Key`

The tag key.

_Required_: Yes

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PropagateAtLaunch`

Set to `true` if you want CloudFormation to copy the tag to EC2 instances that
are launched as part of the Auto Scaling group. Set to `false` if you want the tag
attached only to the Auto Scaling group and not copied to any instances launched as part of
the Auto Scaling group.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The tag value.

_Required_: Yes

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RetentionTriggers

TotalLocalStorageGBRequest

All content copied from https://docs.aws.amazon.com/.
