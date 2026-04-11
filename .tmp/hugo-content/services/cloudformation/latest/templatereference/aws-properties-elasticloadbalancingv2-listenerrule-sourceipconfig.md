This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticLoadBalancingV2::ListenerRule SourceIpConfig

Information about a source IP condition.

You can use this condition to route based on the IP address of the source that connects to
the load balancer. If a client is behind a proxy, this is the IP address of the proxy not the
IP address of the client.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Values" : [ String, ... ]
}

```

### YAML

```yaml

  Values:
    - String

```

## Properties

`Values`

The source IP addresses, in CIDR format. You can use both IPv4 and IPv6
addresses. Wildcards are not supported.

If you specify multiple addresses, the condition is satisfied if the source IP address
of the request matches one of the CIDR blocks. This condition is not satisfied by the
addresses in the X-Forwarded-For header.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

The following example creates a listener rule with an action that forwards
traffic to the specified target group when the specified source IP condition is met.
You can create the target group using [AWS::ElasticLoadBalancingV2::TargetGroup](../userguide/aws-resource-elasticloadbalancingv2-targetgroup.md).

#### YAML

```yaml

myListenerRule:
  Type: 'AWS::ElasticLoadBalancingV2::ListenerRule'
  Properties:
    Actions:
     - Type: forward
       TargetGroupArn: !Ref myTargetGroup
    Conditions:
     - Field: source-ip
       SourceIpConfig:
        Values:
         - 192.168.0.0/16
```

#### JSON

```json

{
    "myListenerRule": {
        "Type": "AWS::ElasticLoadBalancingV2::ListenerRule",
        "Properties": {
            "Actions": [
                {
                    "Type": "forward",
                    "TargetGroupArn": {
                        "Ref": "myTargetGroup"
                    }
                }
            ],
            "Conditions": [
                {
                    "Field": "source-ip",
                    "SourceIpConfig": {
                        "Values": [
                            "192.168.0.0/24"
                        ]
                    }
                }
            ]
        }
    }
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RuleCondition

TargetGroupStickinessConfig

All content copied from https://docs.aws.amazon.com/.
