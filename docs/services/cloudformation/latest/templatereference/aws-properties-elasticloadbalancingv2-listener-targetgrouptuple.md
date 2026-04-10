This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticLoadBalancingV2::Listener TargetGroupTuple

Information about how traffic will be distributed between multiple target groups in a
forward rule.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TargetGroupArn" : String,
  "Weight" : Integer
}

```

### YAML

```yaml

  TargetGroupArn: String
  Weight: Integer

```

## Properties

`TargetGroupArn`

The Amazon Resource Name (ARN) of the target group.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Weight`

The weight. The range is 0 to 999.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TargetGroupStickinessConfig

AWS::ElasticLoadBalancingV2::ListenerCertificate

All content copied from https://docs.aws.amazon.com/.
