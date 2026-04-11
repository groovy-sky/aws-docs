This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AutoScaling::WarmPool InstanceReusePolicy

A structure that specifies an instance reuse policy for the
`InstanceReusePolicy` property of the [AWS::AutoScaling::WarmPool](../userguide/aws-resource-autoscaling-warmpool.md) resource.

For more information, see [Warm pools for Amazon EC2\
Auto Scaling](../../../autoscaling/ec2/userguide/ec2-auto-scaling-warm-pools.md) in the _Amazon EC2 Auto Scaling User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ReuseOnScaleIn" : Boolean
}

```

### YAML

```yaml

  ReuseOnScaleIn: Boolean

```

## Properties

`ReuseOnScaleIn`

Specifies whether instances in the Auto Scaling group can be returned to the warm pool on
scale in.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::AutoScaling::WarmPool

Next

All content copied from https://docs.aws.amazon.com/.
