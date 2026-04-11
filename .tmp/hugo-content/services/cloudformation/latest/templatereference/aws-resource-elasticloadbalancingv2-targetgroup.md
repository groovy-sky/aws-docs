This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticLoadBalancingV2::TargetGroup

Specifies a target group for an Application Load Balancer, a Network Load Balancer, or a
Gateway Load Balancer.

Before you register a Lambda function as a target, you must create a
`AWS::Lambda::Permission` resource that grants the Elastic Load Balancing
service principal permission to invoke the Lambda function.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ElasticLoadBalancingV2::TargetGroup",
  "Properties" : {
      "HealthCheckEnabled" : Boolean,
      "HealthCheckIntervalSeconds" : Integer,
      "HealthCheckPath" : String,
      "HealthCheckPort" : String,
      "HealthCheckProtocol" : String,
      "HealthCheckTimeoutSeconds" : Integer,
      "HealthyThresholdCount" : Integer,
      "IpAddressType" : String,
      "Matcher" : Matcher,
      "Name" : String,
      "Port" : Integer,
      "Protocol" : String,
      "ProtocolVersion" : String,
      "Tags" : [ Tag, ... ],
      "TargetControlPort" : Integer,
      "TargetGroupAttributes" : [ TargetGroupAttribute, ... ],
      "Targets" : [ TargetDescription, ... ],
      "TargetType" : String,
      "UnhealthyThresholdCount" : Integer,
      "VpcId" : String
    }
}

```

### YAML

```yaml

Type: AWS::ElasticLoadBalancingV2::TargetGroup
Properties:
  HealthCheckEnabled: Boolean
  HealthCheckIntervalSeconds: Integer
  HealthCheckPath: String
  HealthCheckPort: String
  HealthCheckProtocol: String
  HealthCheckTimeoutSeconds: Integer
  HealthyThresholdCount: Integer
  IpAddressType: String
  Matcher:
    Matcher
  Name: String
  Port: Integer
  Protocol: String
  ProtocolVersion: String
  Tags:
    - Tag
  TargetControlPort: Integer
  TargetGroupAttributes:
    - TargetGroupAttribute
  Targets:
    - TargetDescription
  TargetType: String
  UnhealthyThresholdCount: Integer
  VpcId: String

```

## Properties

`HealthCheckEnabled`

Indicates whether health checks are enabled. If the target type is `lambda`,
health checks are disabled by default but can be enabled. If the target type is
`instance`, `ip`, or `alb`, health checks are always
enabled and can't be disabled.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HealthCheckIntervalSeconds`

The approximate amount of time, in seconds, between health checks of an individual target. The range is 5-300.
If the target group protocol is TCP, TLS, UDP, TCP\_UDP, QUIC, TCP\_QUIC, HTTP or HTTPS, the default is 30 seconds.
If the target group protocol is GENEVE, the default is 10 seconds.
If the target type is `lambda`, the default is 35 seconds.

_Required_: No

_Type_: Integer

_Minimum_: `5`

_Maximum_: `300`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HealthCheckPath`

\[HTTP/HTTPS health checks\] The destination for health checks on the targets.

\[HTTP1 or HTTP2 protocol version\] The ping path. The default is /.

\[GRPC protocol version\] The path of a custom health check method with the format
/package.service/method. The default is /AWS.ALB/healthcheck.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HealthCheckPort`

The port the load balancer uses when performing health checks on targets. If the protocol
is HTTP, HTTPS, TCP, TLS, UDP, TCP\_UDP, QUIC, or TCP\_QUIC the default is `traffic-port`, which is
the port on which each target receives traffic from the load balancer. If the protocol is
GENEVE, the default is port 80.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HealthCheckProtocol`

The protocol the load balancer uses when performing health checks on targets. For
Application Load Balancers, the default is HTTP. For Network Load Balancers and Gateway Load
Balancers, the default is TCP. The TCP protocol is not supported for health checks if the
protocol of the target group is HTTP or HTTPS. The GENEVE, TLS, UDP, TCP\_UDP, QUIC, and TCP\_QUIC protocols are
not supported for health checks.

_Required_: No

_Type_: String

_Allowed values_: `HTTP | HTTPS | TCP | TLS | UDP | TCP_UDP | GENEVE | QUIC | TCP_QUIC`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HealthCheckTimeoutSeconds`

The amount of time, in seconds, during which no response from a target means a failed
health check. The range is 2–120 seconds. For target groups with a protocol of HTTP, the
default is 6 seconds. For target groups with a protocol of TCP, TLS or HTTPS, the default
is 10 seconds. For target groups with a protocol of GENEVE, the default is 5 seconds. If
the target type is `lambda`, the default is 30 seconds.

_Required_: No

_Type_: Integer

_Minimum_: `2`

_Maximum_: `120`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HealthyThresholdCount`

The number of consecutive health check successes required before considering a target healthy. The range is
2-10. If the target group protocol is TCP, TCP\_UDP, UDP, TLS, HTTP or HTTPS, the default is 5. For target groups
with a protocol of GENEVE, the default is 5. If the target type
is `lambda`, the default is 5.

_Required_: No

_Type_: Integer

_Minimum_: `2`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IpAddressType`

The IP address type. The default value is `ipv4`.

_Required_: No

_Type_: String

_Allowed values_: `ipv4 | ipv6`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Matcher`

\[HTTP/HTTPS health checks\] The HTTP or gRPC codes to use when checking for a successful
response from a target. For target groups with a protocol of TCP, TCP\_UDP, UDP, QUIC, TCP\_QUIC, or TLS the range
is 200-599. For target groups with a protocol of HTTP or HTTPS, the range is 200-499. For target
groups with a protocol of GENEVE, the range is 200-399.

_Required_: No

_Type_: [Matcher](aws-properties-elasticloadbalancingv2-targetgroup-matcher.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the target group.

This name must be unique per region per account, can have a maximum of 32 characters, must
contain only alphanumeric characters or hyphens, and must not begin or end with a
hyphen.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Port`

The port on which the targets receive traffic. This port is used unless you specify a port
override when registering the target. If the target is a Lambda function, this parameter does
not apply. If the protocol is GENEVE, the supported port is 6081.

_Required_: Conditional

_Type_: Integer

_Minimum_: `1`

_Maximum_: `65535`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Protocol`

The protocol to use for routing traffic to the targets. For Application Load Balancers,
the supported protocols are HTTP and HTTPS. For Network Load Balancers, the supported
protocols are TCP, TLS, UDP, TCP\_UDP, QUIC, or TCP\_QUIC. For Gateway Load Balancers, the supported protocol is
GENEVE. A TCP\_UDP listener must be associated with a TCP\_UDP target group. A TCP\_QUIC listener must be associated with a TCP\_QUIC target group. If the target is a
Lambda function, this parameter does not apply.

_Required_: Conditional

_Type_: String

_Allowed values_: `HTTP | HTTPS | TCP | TLS | UDP | TCP_UDP | GENEVE | QUIC | TCP_QUIC`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProtocolVersion`

\[HTTP/HTTPS protocol\] The protocol version. The possible values are `GRPC`,
`HTTP1`, and `HTTP2`.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags.

_Required_: No

_Type_: Array of [Tag](aws-properties-elasticloadbalancingv2-targetgroup-tag.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetControlPort`

The port on which the target control agent and application load balancer exchange management traffic for the target optimizer feature.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetGroupAttributes`

The target group attributes. Attributes that you do not modify retain their current values.

_Required_: No

_Type_: Array of [TargetGroupAttribute](aws-properties-elasticloadbalancingv2-targetgroup-targetgroupattribute.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Targets`

The targets.

_Required_: No

_Type_: Array of [TargetDescription](aws-properties-elasticloadbalancingv2-targetgroup-targetdescription.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetType`

The type of target that you must specify when registering targets with this target group.
You can't specify targets for a target group using more than one target type.

- `instance` \- Register targets by instance ID. This is the default
value.

- `ip` \- Register targets by IP address. You can specify IP addresses from
the subnets of the virtual private cloud (VPC) for the target group, the RFC 1918 range
(10.0.0.0/8, 172.16.0.0/12, and 192.168.0.0/16), and the RFC 6598 range (100.64.0.0/10).
You can't specify publicly routable IP addresses.

- `lambda` \- Register a single Lambda function as a target.

- `alb` \- Register a single Application Load Balancer as a target.

_Required_: No

_Type_: String

_Allowed values_: `instance | ip | lambda | alb`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UnhealthyThresholdCount`

The number of consecutive health check failures required before considering a target unhealthy. The range is
2-10. If the target group protocol is TCP, TCP\_UDP, UDP, TLS, QUIC, TCP\_QUIC, HTTP or HTTPS, the default is 2. For target groups
with a protocol of GENEVE, the default is 2. If the target type
is `lambda`, the default is 5.

_Required_: No

_Type_: Integer

_Minimum_: `2`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcId`

The identifier of the virtual private cloud (VPC). If the target is a Lambda function,
this parameter does not apply. Otherwise, this parameter is required.

_Required_: Conditional

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the target group.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`LoadBalancerArns`

The Amazon Resource Name (ARN) of the load balancer that routes traffic to this target group.

`TargetGroupArn`

The Amazon Resource Name (ARN) of the target group.

`TargetGroupFullName`

The full name of the target group. For example, `targetgroup/my-target-group/cbf133c568e0d028`.

`TargetGroupName`

The name of the target group. For example, `my-target-group`.

## Examples

You can create target groups for use with your listener rules. For more information,
see [AWS::ElasticLoadBalancingV2::ListenerRule](../userguide/aws-resource-elasticloadbalancingv2-listenerrule.md).

- [Target group with instance targets](#aws-resource-elasticloadbalancingv2-targetgroup--examples--Target_group_with_instance_targets)

- [Target group with ip targets](#aws-resource-elasticloadbalancingv2-targetgroup--examples--Target_group_with_ip_targets)

- [Target group with QUIC targets](#aws-resource-elasticloadbalancingv2-targetgroup--examples--Target_group_with_QUIC_targets)

- [Target group with an Application Load Balancer target](#aws-resource-elasticloadbalancingv2-targetgroup--examples--Target_group_with_an_Application_Load_Balancer_target)

- [Target group with a Lambda target](#aws-resource-elasticloadbalancingv2-targetgroup--examples--Target_group_with_a_Lambda_target)

### Target group with instance targets

The following example creates a target group where the targets are EC2 instances.

#### YAML

```yaml

myTargetGroup:
    Type: AWS::ElasticLoadBalancingV2::TargetGroup
    Properties:
      Name: my-target-group
      Protocol: HTTP
      Port: 80
      TargetType: instance
      VpcId: !Ref myVPC
      Targets:
        - Id: !GetAtt Instance1.InstanceId
          Port: 80
        - Id: !GetAtt Instance2.InstanceId
          Port: 80
```

#### JSON

```json

{
    "myTargetGroup": {
        "Type": "AWS::ElasticLoadBalancingV2::TargetGroup",
        "Properties": {
            "Name": "my-target-group",
            "Protocol": "HTTP",
            "Port": 80,
            "TargetType": "instance",
            "VpcId": {
                "Ref": "myVPC"
            },
            "Targets": [
                {
                    "Id": {
                        "Fn::GetAtt": [
                            "Instance1",
                            "InstanceId"
                        ]
                    },
                    "Port": 80
                },
                {
                    "Id": {
                        "Fn::GetAtt": [
                            "Instance2",
                            "InstanceId"
                        ]
                    },
                    "Port": 80
                }
            ]
        }
    }
}
```

### Target group with ip targets

The following example creates a target group where the targets are IP addresses.

#### YAML

```yaml

myTargetGroup:
    Type: AWS::ElasticLoadBalancingV2::TargetGroup
    Properties:
      Name: my-target-group
      Protocol: TCP
      Port: 80
      TargetType: ip
      VpcId: !Ref myVPC
      Targets:
        - Id: !GetAtt Instance1.PrivateIp
          Port: 80
        - Id: !GetAtt Instance2.PrivateIp
          Port: 80
```

#### JSON

```json

{
    "myTargetGroup": {
        "Type": "AWS::ElasticLoadBalancingV2::TargetGroup",
        "Properties": {
            "Name": "my-target-group",
            "Protocol": "TCP",
            "Port": 80,
            "TargetType": "ip",
            "VpcId": {
                "Ref": "myVPC"
            },
            "Targets": [
                {
                    "Id": {
                        "Fn::GetAtt": [
                            "Instance1",
                            "PrivateIp"
                        ]
                    },
                    "Port": 80
                },
                {
                    "Id": {
                        "Fn::GetAtt": [
                            "Instance2",
                            "PrivateIp"
                        ]
                    },
                    "Port": 80
                }
            ]
        }
    }
}
```

### Target group with QUIC targets

The following example creates a target group where the targets are configured for QUIC traffic.

#### YAML

```yaml

myTargetGroup:
    Type: AWS::ElasticLoadBalancingV2::TargetGroup
    Properties:
      Name: my-target-group
      Protocol: QUIC
      Port: 443
      TargetType: ip
      VpcId: !Ref myVPC
      Targets:
        - Id: !GetAtt Instance1.PrivateIp
          Port: 80
          QuicServerId: 0xa1b2c3d4e5f65890
        - Id: !GetAtt Instance2.PrivateIp
          Port: 80
          QuicServerId: 0xa1b2c3d4e5f65899
```

#### JSON

```json

{
    "myTargetGroup": {
        "Type": "AWS::ElasticLoadBalancingV2::TargetGroup",
        "Properties": {
            "Name": "my-target-group",
            "Protocol": "QUIC",
            "Port": 443,
            "TargetType": "ip",
            "VpcId": {
                "Ref": "myVPC"
            },
            "Targets": [
                {
                    "Id": {
                        "Fn::GetAtt": [
                            "Instance1",
                            "PrivateIp"
                        ]
                    },
                    "Port": 80,
                    "QuicServerId": "0xa1b2c3d4e5f65890"
                },
                {
                    "Id": {
                        "Fn::GetAtt": [
                            "Instance2",
                            "PrivateIp"
                        ]
                    },
                    "Port": 80,
                    "QuicServerId": "0xa1b2c3d4e5f65899"
                }
            ]
        }
    }
}
```

### Target group with an Application Load Balancer target

The following example creates a target group where the target is an Application Load Balancer.
This target group must be used by a Network Load Balancer.

#### YAML

```yaml

myTargetGroup:
    Type: AWS::ElasticLoadBalancingV2::TargetGroup
    Properties:
      Name: my-target-group
      Protocol: TCP
      Port: 80
      TargetType: alb
      VpcId: !Ref albVPC
      Targets:
        - Id: !Ref myALB
          Port: 80
```

#### JSON

```json

{
    "myTargetGroup": {
        "Type": "AWS::ElasticLoadBalancingV2::TargetGroup",
        "Properties": {
            "Name": "my-target-group",
            "Protocol": "TCP",
            "Port": 80,
            "TargetType": "alb",
            "VpcId": {
                "Ref": "albVPC"
            },
            "Targets": [
                {
                    "Id": {
                        "Ref": "myALB"
                    },
                    "Port": 80
                }
            ]
        }
    }
}
```

### Target group with a Lambda target

The following example creates a target group where the target is a Lambda
function. This target group must be used by an Application Load Balancer.

#### YAML

```yaml

Resources:
  MyLambdaInvokePermission:
    Type: AWS::Lambda::Permission
    Properties:
      FunctionName: !GetAtt
        - MyLambdaFunction
        - Arn
      Action: 'lambda:InvokeFunction'
      Principal: elasticloadbalancing.amazonaws.com

  MyTargetGroup:
    Type: AWS::ElasticLoadBalancingV2::TargetGroup
    Properties:
      HealthCheckEnabled: false
      Name: MyTargets
      TargetType: lambda
      Targets:
      - Id: !GetAtt [ MyLambdaFunction, Arn ]

  MyLambdaFunction:
    Type: "AWS::Lambda::Function"
    Properties:
      Handler: "index.handler"
      Role: !GetAtt [ LambdaExecutionRole, Arn ]
      Code:
        ZipFile: !Sub |
          import json

          def handler(event, context):
            response = {
              "statusCode": 200,
              "statusDescription": "200 OK",
              "isBase64Encoded": False,
              "headers": {
                "Content-Type": "text/html; charset=utf-8"
              }
            }

            response['body'] = """<html>
            <head>
            <title>Hello World!</title>
            <style>
            html, body {
              margin: 0; padding: 0;
              font-family: arial; font-weight: 700; font-size: 3em;
              text-align: center;
            }
            </style>
            </head>
            <body>
            <p>Hello World from Lambda</p>
            </body>
            </html>"""
            return response
      Runtime: "runtime.version"
      Timeout: "25"

  LambdaExecutionRole:
    Type: "AWS::IAM::Role"
    Properties:
      AssumeRolePolicyDocument:
        Version: "2012-10-17"
        Statement:
          - Effect: Allow
            Principal:
              Service: lambda.amazonaws.com
            Action: "sts:AssumeRole"
```

## See also

- [CreateTargetGroup](../../../../reference/elasticloadbalancing/latest/apireference/api-createtargetgroup.md) in the _Elastic Load Balancing API_
_Reference (version 2015-12-01)_

- [Target groups](../../../elasticloadbalancing/latest/application/load-balancer-target-groups.md) in the _User Guide for Application Load_
_Balancers_

- [Target groups](../../../elasticloadbalancing/latest/network/load-balancer-target-groups.md) in the _User Guide for Network Load_
_Balancers_

- [Target groups](../../../elasticloadbalancing/latest/gateway/target-groups.md) in the _User Guide for Gateway Load_
_Balancers_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Matcher

All content copied from https://docs.aws.amazon.com/.
