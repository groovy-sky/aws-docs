This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticLoadBalancingV2::Listener

Specifies a listener for an Application Load Balancer, Network Load Balancer, or
Gateway Load Balancer.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ElasticLoadBalancingV2::Listener",
  "Properties" : {
      "AlpnPolicy" : [ String, ... ],
      "Certificates" : [ Certificate, ... ],
      "DefaultActions" : [ Action, ... ],
      "ListenerAttributes" : [ ListenerAttribute, ... ],
      "LoadBalancerArn" : String,
      "MutualAuthentication" : MutualAuthentication,
      "Port" : Integer,
      "Protocol" : String,
      "SslPolicy" : String
    }
}

```

### YAML

```yaml

Type: AWS::ElasticLoadBalancingV2::Listener
Properties:
  AlpnPolicy:
    - String
  Certificates:
    - Certificate
  DefaultActions:
    - Action
  ListenerAttributes:
    - ListenerAttribute
  LoadBalancerArn: String
  MutualAuthentication:
    MutualAuthentication
  Port: Integer
  Protocol: String
  SslPolicy: String

```

## Properties

`AlpnPolicy`

\[TLS listener\] The name of the Application-Layer Protocol Negotiation (ALPN)
policy.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Certificates`

The default SSL server certificate for a secure listener. You must provide exactly one
certificate if the listener protocol is HTTPS or TLS.

For an HTTPS listener, update requires some interruptions. For a TLS listener, update requires no interruption.

To create a certificate list for a secure listener, use [AWS::ElasticLoadBalancingV2::ListenerCertificate](../userguide/aws-resource-elasticloadbalancingv2-listenercertificate.md).

_Required_: Conditional

_Type_: Array of [Certificate](aws-properties-elasticloadbalancingv2-listener-certificate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultActions`

The actions for the default rule. You cannot define a condition for a default
rule.

To create additional rules for an Application Load Balancer, use [AWS::ElasticLoadBalancingV2::ListenerRule](../userguide/aws-resource-elasticloadbalancingv2-listenerrule.md).

_Required_: Yes

_Type_: Array of [Action](aws-properties-elasticloadbalancingv2-listener-action.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ListenerAttributes`

The listener attributes. Attributes that you do not modify retain their current values.

_Required_: No

_Type_: Array of [ListenerAttribute](aws-properties-elasticloadbalancingv2-listener-listenerattribute.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LoadBalancerArn`

The Amazon Resource Name (ARN) of the load balancer.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MutualAuthentication`

The mutual authentication configuration information.

_Required_: No

_Type_: [MutualAuthentication](aws-properties-elasticloadbalancingv2-listener-mutualauthentication.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Port`

The port on which the load balancer is listening. You can't specify a port for a Gateway
Load Balancer.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Protocol`

The protocol for connections from clients to the load balancer. For Application Load
Balancers, the supported protocols are HTTP and HTTPS. For Network Load Balancers, the
supported protocols are TCP, TLS, UDP, TCP\_UDP, QUIC, and TCP\_QUIC. You can’t specify the UDP, TCP\_UDP, QUIC, or TCP\_QUIC
protocol if dual-stack mode is enabled. You can't specify a protocol for a Gateway Load
Balancer.

_Required_: No

_Type_: String

_Allowed values_: `HTTP | HTTPS | TCP | TLS | UDP | TCP_UDP | GENEVE | QUIC | TCP_QUIC`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SslPolicy`

\[HTTPS and TLS listeners\] The security policy that defines which protocols and ciphers are supported.
For more information, see [Security policies](../../../elasticloadbalancing/latest/application/describe-ssl-policies.md)
in the _Application Load Balancers Guide_ and [Security policies](../../../elasticloadbalancing/latest/network/describe-ssl-policies.md)
in the _Network Load Balancers Guide_.

\[HTTPS listeners\] Updating the security policy can result in interruptions if the load balancer is handling a high volume of traffic.
To decrease the possibility of an interruption if your load balancer is handling a high volume of traffic,
create an additional load balancer or request an LCU reservation.

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the listener.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ListenerArn`

The Amazon Resource Name (ARN) of the listener.

## Examples

After you create your load balancer using [AWS::ElasticLoadBalancingV2::LoadBalancer](../userguide/aws-resource-elasticloadbalancingv2-loadbalancer.md), you can add a listener.

- [Create an HTTP listener](#aws-resource-elasticloadbalancingv2-listener--examples--Create_an_HTTP_listener)

- [Create an HTTPS listener](#aws-resource-elasticloadbalancingv2-listener--examples--Create_an_HTTPS_listener)

- [Create a TCP listener](#aws-resource-elasticloadbalancingv2-listener--examples--Create_a_TCP_listener)

- [Create a TLS listener](#aws-resource-elasticloadbalancingv2-listener--examples--Create_a_TLS_listener)

- [Create a UDP listener](#aws-resource-elasticloadbalancingv2-listener--examples--Create_a_UDP_listener)

- [Create a QUIC listener](#aws-resource-elasticloadbalancingv2-listener--examples--Create_a_QUIC_listener)

### Create an HTTP listener

The following example creates an HTTP listener with a default action that redirects HTTP
requests on port 80 to HTTPS requests on port 443, retaining the original host name,
path, and query string.

#### YAML

```yaml

myHTTPlistener:
  Type: 'AWS::ElasticLoadBalancingV2::Listener'
  Properties:
    LoadBalancerArn: !Ref myLoadBalancer
    Protocol: HTTP
    Port: 80
    DefaultActions:
      - Type: redirect
        RedirectConfig:
          Protocol: HTTPS
          Port: 443
          Host: "#{host}"
          Path: "/#{path}"
          Query: "#{query}"
          StatusCode: "HTTP_301"
```

#### JSON

```json

{
    "myHTTPlistener": {
        "Type": "AWS::ElasticLoadBalancingV2::Listener",
        "Properties": {
            "DefaultActions": [
                {
                    "Type": "redirect",
                    "RedirectConfig": {
                        "Protocol": "HTTPS",
                        "Port": 443,
                        "Host": "#{host}",
                        "Path": "/#{path}",
                        "Query": "#{query}",
                        "StatusCode": "HTTP_301"
                    }
                }
            ],
            "LoadBalancerArn": {
                "Ref": "myLoadBalancer"
            },
            "Port": 80,
            "Protocol": "HTTP"
        }
    }
}
```

### Create an HTTPS listener

The following example creates an HTTPS listener with a default action that forwards
traffic to the specified target group. When you create a secure listener, you must
specify a security policy and a certificate. You can create the target group using
[AWS::ElasticLoadBalancingV2::TargetGroup](../userguide/aws-resource-elasticloadbalancingv2-targetgroup.md).

#### YAML

```yaml

myHTTPSListener:
  Type: 'AWS::ElasticLoadBalancingV2::Listener'
  Properties:
    LoadBalancerArn: !Ref myLoadBalancer
    Protocol: HTTPS
    Port: 443
    SslPolicy: "ELBSecurityPolicy-TLS13-1-2-2021-06"
    Certificates:
      - CertificateArn: "arn:aws:acm:us-west-2:123456789012:certificate/88ca7932-756c-46f1-a70d-03fa7EXAMPLE"
    DefaultActions:
      - Type: forward
        TargetGroupArn: !Ref myTargetGroup
```

#### JSON

```json

{
    "myHTTPSListener": {
        "Type": "AWS::ElasticLoadBalancingV2::Listener",
        "Properties": {
            "LoadBalancerArn": {
                "Ref": "myLoadBalancer"
            },
            "Protocol": "HTTPS",
            "Port": 443,
            "SslPolicy": "ELBSecurityPolicy-TLS13-1-2-2021-06",
            "Certificates": [
                {
                    "CertificateArn": "arn:aws:acm:us-west-2:123456789012:certificate/88ca7932-756c-46f1-a70d-03fa7EXAMPLE"
                }
            ],
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

### Create a TCP listener

The following example creates a TCP listener with a default action that forwards
traffic to the specified target group. You can create the target group using
[AWS::ElasticLoadBalancingV2::TargetGroup](../userguide/aws-resource-elasticloadbalancingv2-targetgroup.md).

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

### Create a TLS listener

The following example creates a TLS listener with a default action that forwards
traffic to the specified target group. When you create a secure listener, you must
specify a security policy and a certificate. You can create the target group using
[AWS::ElasticLoadBalancingV2::TargetGroup](../userguide/aws-resource-elasticloadbalancingv2-targetgroup.md).

#### YAML

```yaml

myTLSListener:
  Type: 'AWS::ElasticLoadBalancingV2::Listener'
  Properties:
    LoadBalancerArn: !Ref myLoadBalancer
    Protocol: TLS
    Port: 443
    SslPolicy: "ELBSecurityPolicy-TLS13-1-2-2021-06"
    Certificates:
      - CertificateArn: "arn:aws:acm:us-west-2:123456789012:certificate/88ca7932-756c-46f1-a70d-03fa7EXAMPLE"
    DefaultActions:
      - Type: forward
        TargetGroupArn: !Ref myTargetGroup
```

#### JSON

```json

{
    "myTLSListener": {
        "Type": "AWS::ElasticLoadBalancingV2::Listener",
        "Properties": {
            "LoadBalancerArn": {
                "Ref": "myLoadBalancer"
            },
            "Protocol": "TLS",
            "Port": 443,
            "SslPolicy": "ELBSecurityPolicy-TLS13-1-2-2021-06",
            "Certificates": [
                {
                    "CertificateArn": "arn:aws:acm:us-west-2:123456789012:certificate/88ca7932-756c-46f1-a70d-03fa7EXAMPLE"
                }
            ],
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

### Create a UDP listener

The following example creates a UDP listener with a default action that forwards
traffic to the specified target group. You can create the target group using
[AWS::ElasticLoadBalancingV2::TargetGroup](../userguide/aws-resource-elasticloadbalancingv2-targetgroup.md).

#### YAML

```yaml

myUDPListener:
  Type: 'AWS::ElasticLoadBalancingV2::Listener'
  Properties:
    LoadBalancerArn: !Ref myLoadBalancer
    Protocol: UDP
    Port: 53
    DefaultActions:
      - Type: forward
        TargetGroupArn: !Ref myTargetGroup
```

#### JSON

```json

{
    "myUDPListener": {
        "Type": "AWS::ElasticLoadBalancingV2::Listener",
        "Properties": {
            "LoadBalancerArn": {
                "Ref": "myLoadBalancer"
            },
            "Protocol": "UDP",
            "Port": 53,
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

### Create a QUIC listener

The following example creates a QUIC listener with a default action that forwards
traffic to the specified target group. You can create the target group using
[AWS::ElasticLoadBalancingV2::TargetGroup](../userguide/aws-resource-elasticloadbalancingv2-targetgroup.md).

#### YAML

```yaml

myQUICListener:
  Type: 'AWS::ElasticLoadBalancingV2::Listener'
  Properties:
    LoadBalancerArn: !Ref myLoadBalancer
    Protocol: QUIC
    Port: 443
    DefaultActions:
      - Type: forward
        TargetGroupArn: !Ref myTargetGroup
```

#### JSON

```json

{
    "myUDPListener": {
        "Type": "AWS::ElasticLoadBalancingV2::Listener",
        "Properties": {
            "LoadBalancerArn": {
                "Ref": "myLoadBalancer"
            },
            "Protocol": "QUIC",
            "Port": 443,
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

## See also

- [CreateListener](../../../../reference/elasticloadbalancing/latest/apireference/api-createlistener.md) in the _Elastic Load Balancing API Reference_
_(version 2015-12-01)_

- [Listeners](../../../elasticloadbalancing/latest/application/load-balancer-listeners.md) in the _User Guide for Application Load_
_Balancers_

- [Listeners](../../../elasticloadbalancing/latest/network/load-balancer-listeners.md) in the _User Guide for Network Load_
_Balancers_

- [Listeners](../../../elasticloadbalancing/latest/gateway/gateway-listeners.md) in the _User Guide for Gateway Load_
_Balancers_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Elastic Load Balancing V2

Action

All content copied from https://docs.aws.amazon.com/.
