This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticLoadBalancingV2::Listener ForwardConfig

Information for creating an action that distributes requests among multiple target
groups. Specify only when `Type` is `forward`.

If you specify both `ForwardConfig` and `TargetGroupArn`, you can
specify only one target group using `ForwardConfig` and it must be the same
target group specified in `TargetGroupArn`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TargetGroups" : [ TargetGroupTuple, ... ],
  "TargetGroupStickinessConfig" : TargetGroupStickinessConfig
}

```

### YAML

```yaml

  TargetGroups:
    - TargetGroupTuple
  TargetGroupStickinessConfig:
    TargetGroupStickinessConfig

```

## Properties

`TargetGroups`

Information about how traffic will be distributed between multiple target groups in a
forward rule.

_Required_: No

_Type_: Array of [TargetGroupTuple](aws-properties-elasticloadbalancingv2-listener-targetgrouptuple.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetGroupStickinessConfig`

Information about the target group stickiness for a rule.

_Required_: No

_Type_: [TargetGroupStickinessConfig](aws-properties-elasticloadbalancingv2-listener-targetgroupstickinessconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

The following example defines a listener with a default action that forwards
traffic to the specified target group. You can create the target group using [AWS::ElasticLoadBalancingV2::TargetGroup](../userguide/aws-resource-elasticloadbalancingv2-targetgroup.md).

#### YAML

```yaml

myTCPListener:
  Type: 'AWS::ElasticLoadBalancingV2::Listener'
  Properties:
    LoadBalancerArn: !Ref myLoadBalancer
    Protocol: TCP
    Port: 80
    DefaultActions:
      - Type: forward
        TargetGroupArn: !Ref myTargetGroup
```

#### JSON

```json

{
    "myTCPListener": {
        "Type": "AWS::ElasticLoadBalancingV2::Listener",
        "Properties": {
            "LoadBalancerArn": {
                "Ref": "myLoadBalancer"
            },
            "Protocol": "TCP",
            "Port": 80,
            "DefaultActions": [
                {
                    "Type": "forward",
                    "TargetGroupArn": {
                        "Ref": "myTargetGroup"
                    }
                }
            ]
        }
    }
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FixedResponseConfig

JwtValidationActionAdditionalClaim

All content copied from https://docs.aws.amazon.com/.
