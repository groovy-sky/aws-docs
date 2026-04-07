This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticLoadBalancingV2::ListenerRule

Specifies a listener rule. The listener must be associated with an Application Load
Balancer. Each rule consists of a priority, one or more actions, and one or more
conditions.

For more information, see [Quotas for your Application Load Balancers](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-limits.html) in the
_User Guide for Application Load Balancers_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ElasticLoadBalancingV2::ListenerRule",
  "Properties" : {
      "Actions" : [ Action, ... ],
      "Conditions" : [ RuleCondition, ... ],
      "ListenerArn" : String,
      "Priority" : Integer,
      "Transforms" : [ Transform, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ElasticLoadBalancingV2::ListenerRule
Properties:
  Actions:
    - Action
  Conditions:
    - RuleCondition
  ListenerArn: String
  Priority: Integer
  Transforms:
    - Transform

```

## Properties

`Actions`

The actions.

The rule must include exactly one of the following types of actions:
`forward`, `fixed-response`, or `redirect`, and it must
be the last action to be performed. If the rule is for an HTTPS listener, it can also
optionally include an authentication action.

_Required_: Yes

_Type_: Array of [Action](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-elasticloadbalancingv2-listenerrule-action.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Conditions`

The conditions.

The rule can optionally include up to one of each of the following conditions:
`http-request-method`, `host-header`, `path-pattern`,
and `source-ip`. A rule can also optionally include one or more of each of the
following conditions: `http-header` and `query-string`.

_Required_: Yes

_Type_: Array of [RuleCondition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-elasticloadbalancingv2-listenerrule-rulecondition.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ListenerArn`

The Amazon Resource Name (ARN) of the listener.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Priority`

The rule priority. A listener can't have multiple rules with the same priority.

If you try to reorder rules by updating their priorities, do not specify a new priority
if an existing rule already uses this priority, as this can cause an error. If you need to
reuse a priority with a different rule, you must remove it as a priority first, and then
specify it in a subsequent update.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `50000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Transforms`

Property description not available.

_Required_: No

_Type_: Array of [Transform](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-elasticloadbalancingv2-listenerrule-transform.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the listener rule.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`IsDefault`

Indicates whether this is the default rule.

`RuleArn`

The Amazon Resource Name (ARN) of the rule.

## Examples

You can define rules for your listener in addition to the default action. Each rule must
specify an action and a condition.

- [Forward action with a host-header condition](#aws-resource-elasticloadbalancingv2-listenerrule--examples--Forward_action_with_a_host-header_condition)

- [Fixed-response action with a source-ip condition](#aws-resource-elasticloadbalancingv2-listenerrule--examples--Fixed-response_action_with_a_source-ip_condition)

- [Redirect action with an http-header condition](#aws-resource-elasticloadbalancingv2-listenerrule--examples--Redirect_action_with_an_http-header_condition)

### Forward action with a host-header condition

The following example creates a rule with a `forward` action and a `host-header` condition.

#### YAML

```yaml

myForwardListenerRule:
   Type: 'AWS::ElasticLoadBalancingV2::ListenerRule'
   Properties:
     ListenerArn: !Ref myListener
     Priority: 10
     Conditions:
       - Field: host-header
         Values:
           - example.com
           - www.example.com
     Actions:
       - Type: forward
         TargetGroupArn: !Ref myTargetGroup
```

#### JSON

```json

{
    "myForwardListenerRule": {
        "Type": "AWS::ElasticLoadBalancingV2::ListenerRule",
        "Properties": {
            "ListenerArn": {
                "Ref": "myListener"
            },
            "Priority": 10,
            "Conditions": [
                {
                    "Field": "host-header",
                    "Values": [
                        "example.com",
                        "www.example.com"
                    ]
                }
            ],
            "Actions": [
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

### Fixed-response action with a source-ip condition

The following example creates a rule with a `fixed-response` action and a `source-ip` condition.

#### YAML

```yaml

myFixedResponseListenerRule:
   Type: 'AWS::ElasticLoadBalancingV2::ListenerRule'
   Properties:
     ListenerArn: !Ref myListener
     Priority: 20
     Conditions:
       - Field: source-ip
         SourceIpConfig:
           Values:
             - 192.168.1.0/24
             - 10.0.0.0/16
     Actions:
       - Type: fixed-response
         FixedResponseConfig:
           StatusCode: 403
           ContentType: text/plain
           MessageBody: "Access denied"
```

#### JSON

```json

{
    "myFixedResponseListenerRule": {
        "Type": "AWS::ElasticLoadBalancingV2::ListenerRule",
        "Properties": {
            "ListenerArn": {
                "Ref": "myListener"
            },
            "Priority": 20,
            "Conditions": [
                {
                    "Field": "source-ip",
                    "SourceIpConfig": {
                        "Values": [
                            "192.168.1.0/24",
                            "10.0.0.0/16"
                        ]
                    }
                }
            ],
            "Actions": [
                {
                    "Type": "fixed-response",
                    "FixedResponseConfig": {
                        "StatusCode": 403,
                        "ContentType": "text/plain",
                        "MessageBody": "Access denied"
                    }
                }
            ]
        }
    }
}
```

### Redirect action with an http-header condition

The following example creates a rule with a `redirect` action and an `http-header` condition.

#### YAML

```yaml

myRedirectListenerRule:
   Type: 'AWS::ElasticLoadBalancingV2::ListenerRule'
   Properties:
     ListenerArn: !Ref myListener
     Priority: 30
     Conditions:
       - Field: http-header
         HttpHeaderConfig:
           HttpHeaderName: User-Agent
           Values:
             - "*Mobile*"
             - "*Android*"
             - "*iPhone*"
     Actions:
       - Type: redirect
         RedirectConfig:
           Host: m.example.com
           StatusCode: HTTP_302
```

#### JSON

```json

{
    "myRedirectListenerRule": {
        "Type": "AWS::ElasticLoadBalancingV2::ListenerRule",
        "Properties": {
            "ListenerArn": {
                "Ref": "myListener"
            },
            "Priority": 30,
            "Conditions": [
                {
                    "Field": "http-header",
                    "HttpHeaderConfig": {
                        "HttpHeaderName": "User-Agent",
                        "Values": [
                            "*Mobile*",
                            "*Android*",
                            "*iPhone*"
                        ]
                    }
                }
            ],
            "Actions": [
                {
                    "Type": "redirect",
                    "RedirectConfig": {
                        "Host": "m.example.com",
                        "StatusCode": "HTTP_302"
                    }
                }
            ]
        }
    }
}
```

## See also

- [CreateRule](https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_CreateRule.html) in the _Elastic Load Balancing API Reference_
_(version 2015-12-01)_

- [Listener Rules](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html#listener-rules) in the _User Guide for Application Load_
_Balancers_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Certificate

Action
