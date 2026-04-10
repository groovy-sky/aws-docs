This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticLoadBalancingV2::Listener TargetGroupStickinessConfig

Information about the target group stickiness for a rule.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DurationSeconds" : Integer,
  "Enabled" : Boolean
}

```

### YAML

```yaml

  DurationSeconds: Integer
  Enabled: Boolean

```

## Properties

`DurationSeconds`

\[Application Load Balancers\] The time period, in seconds, during which requests from a
client should be routed to the same target group. The range is 1-604800 seconds (7 days). You
must specify this value when enabling target group stickiness.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

Indicates whether target group stickiness is enabled.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RedirectConfig

TargetGroupTuple

All content copied from https://docs.aws.amazon.com/.
