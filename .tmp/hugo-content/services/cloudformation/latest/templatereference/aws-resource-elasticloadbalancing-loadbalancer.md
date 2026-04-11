This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticLoadBalancing::LoadBalancer

Specifies a Classic Load Balancer.

If this resource has a public IP address and is also in a VPC that is defined in the same template, you must use
the [DependsOn attribute](../userguide/aws-attribute-dependson.md)
to declare a dependency on the VPC-gateway attachment.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ElasticLoadBalancing::LoadBalancer",
  "Properties" : {
      "AccessLoggingPolicy" : AccessLoggingPolicy,
      "AppCookieStickinessPolicy" : [ AppCookieStickinessPolicy, ... ],
      "AvailabilityZones" : [ String, ... ],
      "ConnectionDrainingPolicy" : ConnectionDrainingPolicy,
      "ConnectionSettings" : ConnectionSettings,
      "CrossZone" : Boolean,
      "HealthCheck" : HealthCheck,
      "Instances" : [ String, ... ],
      "LBCookieStickinessPolicy" : [ LBCookieStickinessPolicy, ... ],
      "Listeners" : [ Listeners, ... ],
      "LoadBalancerName" : String,
      "Policies" : [ Policies, ... ],
      "Scheme" : String,
      "SecurityGroups" : [ String, ... ],
      "Subnets" : [ String, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ElasticLoadBalancing::LoadBalancer
Properties:
  AccessLoggingPolicy:
    AccessLoggingPolicy
  AppCookieStickinessPolicy:
    - AppCookieStickinessPolicy
  AvailabilityZones:
    - String
  ConnectionDrainingPolicy:
    ConnectionDrainingPolicy
  ConnectionSettings:
    ConnectionSettings
  CrossZone: Boolean
  HealthCheck:
    HealthCheck
  Instances:
    - String
  LBCookieStickinessPolicy:
    - LBCookieStickinessPolicy
  Listeners:
    - Listeners
  LoadBalancerName: String
  Policies:
    - Policies
  Scheme: String
  SecurityGroups:
    - String
  Subnets:
    - String
  Tags:
    - Tag

```

## Properties

`AccessLoggingPolicy`

Information about where and how access logs are stored for the load balancer.

_Required_: No

_Type_: [AccessLoggingPolicy](aws-properties-elasticloadbalancing-loadbalancer-accessloggingpolicy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AppCookieStickinessPolicy`

Information about a policy for application-controlled session stickiness.

_Required_: No

_Type_: [Array](aws-properties-elasticloadbalancing-loadbalancer-appcookiestickinesspolicy.md) of [AppCookieStickinessPolicy](aws-properties-elasticloadbalancing-loadbalancer-appcookiestickinesspolicy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AvailabilityZones`

The Availability Zones for a load balancer in a default VPC. For a load balancer in a nondefault VPC, specify `Subnets` instead.

Update requires replacement if you did not previously specify an Availability Zone or if you are removing all Availability Zones.
Otherwise, update requires no interruption.

_Required_: No

_Type_: Array of String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`ConnectionDrainingPolicy`

If enabled, the load balancer allows existing requests to complete before the load balancer shifts traffic away from a deregistered or unhealthy instance.

For more information, see [Configure connection draining](../../../elasticloadbalancing/latest/classic/config-conn-drain.md)
in the _User Guide for Classic Load Balancers_.

_Required_: No

_Type_: [ConnectionDrainingPolicy](aws-properties-elasticloadbalancing-loadbalancer-connectiondrainingpolicy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConnectionSettings`

If enabled, the load balancer allows the connections to remain idle (no data is sent over the connection) for the specified duration.

By default, Elastic Load Balancing maintains a 60-second idle connection timeout for both front-end and back-end connections of your load balancer.
For more information, see [Configure idle connection timeout](../../../elasticloadbalancing/latest/classic/config-idle-timeout.md)
in the _User Guide for Classic Load Balancers_.

_Required_: No

_Type_: [ConnectionSettings](aws-properties-elasticloadbalancing-loadbalancer-connectionsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CrossZone`

If enabled, the load balancer routes the request traffic evenly across all instances regardless of the Availability Zones.

For more information, see [Configure cross-zone load balancing](../../../elasticloadbalancing/latest/classic/enable-disable-crosszone-lb.md)
in the _User Guide for Classic Load Balancers_.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HealthCheck`

The health check settings to use when evaluating the health of your EC2 instances.

Update requires replacement if you did not previously specify health check settings or if you are removing the health check settings.
Otherwise, update requires no interruption.

_Required_: No

_Type_: [HealthCheck](aws-properties-elasticloadbalancing-loadbalancer-healthcheck.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Instances`

The IDs of the instances for the load balancer.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LBCookieStickinessPolicy`

Information about a policy for duration-based session stickiness.

_Required_: No

_Type_: [Array](aws-properties-elasticloadbalancing-loadbalancer-lbcookiestickinesspolicy.md) of [LBCookieStickinessPolicy](aws-properties-elasticloadbalancing-loadbalancer-lbcookiestickinesspolicy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Listeners`

The listeners for the load balancer. You can specify at most one listener per port.

If you update the properties for a listener, AWS CloudFormation deletes the existing listener and
creates a new one with the specified properties. While the new listener is being created, clients
cannot connect to the load balancer.

_Required_: Yes

_Type_: [Array](aws-properties-elasticloadbalancing-loadbalancer-listeners.md) of [Listeners](aws-properties-elasticloadbalancing-loadbalancer-listeners.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LoadBalancerName`

The name of the load balancer. This name must be unique within your set of load balancers for the region.

If you don't specify a name, AWS CloudFormation generates a unique physical ID for the load balancer.
For more information, see [Name Type](../userguide/aws-properties-name.md).
If you specify a name, you cannot perform updates that require replacement of this resource, but you can perform
other updates. To replace the resource, specify a new name.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Policies`

The policies defined for your Classic Load Balancer. Specify only back-end server policies.

_Required_: No

_Type_: [Array](aws-properties-elasticloadbalancing-loadbalancer-policies.md) of [Policies](aws-properties-elasticloadbalancing-loadbalancer-policies.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Scheme`

The type of load balancer. Valid only for load balancers in a VPC.

If `Scheme` is `internet-facing`, the load balancer
has a public DNS name that resolves to a public IP address.

If `Scheme` is `internal`, the load balancer has a public
DNS name that resolves to a private IP address.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SecurityGroups`

The security groups for the load balancer. Valid only for load balancers in a VPC.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Subnets`

The IDs of the subnets for the load balancer. You can specify at most one subnet per Availability Zone.

Update requires replacement if you did not previously specify a subnet or if you are removing all subnets.
Otherwise, update requires no interruption. To update to a different subnet in the current Availability Zone,
you must first update to a subnet in a different Availability Zone, then update to the new subnet in the original
Availability Zone.

_Required_: No

_Type_: Array of String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Tags`

The tags associated with a load balancer.

_Required_: No

_Type_: Array of [Tag](aws-properties-elasticloadbalancing-loadbalancer-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the load balancer.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CanonicalHostedZoneName`

The name of the Route 53 hosted zone that is associated with the load balancer. Internal-facing load balancers don't use this value,
use `DNSName` instead.

`CanonicalHostedZoneNameID`

The ID of the Route 53 hosted zone name that is associated with the load balancer.

`DNSName`

The DNS name for the load balancer.

`SourceSecurityGroup.GroupName`

The name of the security group that you can use as part of your inbound rules for your load balancer's back-end instances.

`SourceSecurityGroup.OwnerAlias`

The owner of the source security group.

## Examples

- [Classic Load Balancer in a default VPC](#aws-resource-elasticloadbalancing-loadbalancer--examples--Classic_Load_Balancer_in_a_default_VPC)

- [Classic Load Balancer in a nondefault VPC](#aws-resource-elasticloadbalancing-loadbalancer--examples--Classic_Load_Balancer_in_a_nondefault_VPC)

### Classic Load Balancer in a default VPC

The following example specifies a Classic Load Balancer with a secure listener in a default VPC.

#### JSON

```json

"MyLoadBalancer" : {
    "Type": "AWS::ElasticLoadBalancing::LoadBalancer",
    "Properties": {
        "AvailabilityZones": [ "us-east-2a" ],
        "CrossZone": "true",
        "Listeners": [{
            "InstancePort": "80",
            "InstanceProtocol": "HTTP",
            "LoadBalancerPort": "443",
            "Protocol": "HTTPS",
            "PolicyNames": [ "My-SSLNegotiation-Policy" ],
            "SSLCertificateId": "arn:aws:iam::123456789012:server-certificate/my-server-certificate"
        }],
        "HealthCheck": {
            "Target": "HTTP:80/",
            "HealthyThreshold": "2",
            "UnhealthyThreshold": "3",
            "Interval": "10",
            "Timeout": "5"
        },
        "Policies": [{
            "PolicyName": "My-SSLNegotiation-Policy",
            "PolicyType": "SSLNegotiationPolicyType",
            "Attributes": [{
                "Name": "Reference-Security-Policy",
                "Value": "ELBSecurityPolicy-TLS-1-2-2017-01"
            }]
        }]
    }
}
```

#### YAML

```yaml

MyLoadBalancer:
    Type: AWS::ElasticLoadBalancing::LoadBalancer
    Properties:
      AvailabilityZones:
      - "us-east-2a"
      CrossZone: true
      Listeners:
      - InstancePort: '80'
        InstanceProtocol: HTTP
        LoadBalancerPort: '443'
        Protocol: HTTPS
        PolicyNames:
        - My-SSLNegotiation-Policy
        SSLCertificateId: arn:aws:iam::123456789012:server-certificate/my-server-certificate
      HealthCheck:
        Target: HTTP:80/
        HealthyThreshold: '2'
        UnhealthyThreshold: '3'
        Interval: '10'
        Timeout: '5'
      Policies:
      - PolicyName: My-SSLNegotiation-Policy
        PolicyType: SSLNegotiationPolicyType
        Attributes:
        - Name: Reference-Security-Policy
          Value: ELBSecurityPolicy-TLS-1-2-2017-01
```

### Classic Load Balancer in a nondefault VPC

The following example specifies a Classic Load Balancer with an HTTP listener and a VPC with
one subnet and an internet gateway.

#### JSON

```json

{
    "Resources": {
        "myLoadBalancer": {
            "Type": "AWS::ElasticLoadBalancing::LoadBalancer",
            "Properties": {
                "Subnets": [
                    {
                        "Ref": "mySubnet"
                    }
                ],
                "Listeners": [
                    {
                        "LoadBalancerPort": "80",
                        "InstancePort": "80",
                        "Protocol": "HTTP"
                    }
                ]
            }
        },
        "myVPC": {
            "Type": "AWS::EC2::VPC",
            "Properties": {
                "CidrBlock": "10.0.0.0/16",
                "EnableDnsSupport": true,
                "EnableDnsHostnames": true
            }
        },
        "mySubnet": {
            "Type": "AWS::EC2::Subnet",
            "Properties": {
                "VpcId": {
                    "Ref": "myVPC"
                },
                "CidrBlock": "10.0.0.0/24",
                "AvailabilityZone": {
                    "Fn::Select": [
                        1,
                        {
                            "Fn::GetAZs": ""
                        }
                    ]
                }
            }
        },
        "myInternetGateway": {
            "Type": "AWS::EC2::InternetGateway"
        },
        "myGatewayAttachment": {
            "Type": "AWS::EC2::VPCGatewayAttachment",
            "Properties": {
                "VpcId": {
                    "Ref": "myVPC"
                },
                "InternetGatewayId": {
                    "Ref": "myInternetGateway"
                }
            }
        },
        "myRouteTable": {
            "Type": "AWS::EC2::RouteTable",
            "Properties": {
                "VpcId": {
                    "Ref": "myVPC"
                }
            }
        },
        "myRoute": {
            "Type": "AWS::EC2::Route",
            "Properties": {
                "DestinationCidrBlock": "0.0.0.0/0",
                "GatewayId": {
                    "Ref": "myInternetGateway"
                },
                "RouteTableId": {
                    "Ref": "myRouteTable"
                }
            }
        },
        "myRouteTableAssociation": {
            "Type": "AWS::EC2::SubnetRouteTableAssociation",
            "Properties": {
                "RouteTableId": {
                    "Ref": "myRouteTable"
                },
                "SubnetId": {
                    "Ref": "mySubnet"
                }
            }
        }
    }
}
```

#### YAML

```yaml

Resources:
  myLoadBalancer:
    Type: AWS::ElasticLoadBalancing::LoadBalancer
    Properties:
      Subnets:
      - !Ref mySubnet
      Listeners:
      - LoadBalancerPort: '80'
        InstancePort: '80'
        Protocol: HTTP
  myVPC:
    Type: AWS::EC2::VPC
    Properties:
      CidrBlock: 10.0.0.0/16
      EnableDnsSupport: true
      EnableDnsHostnames: true
  mySubnet:
    Type: AWS::EC2::Subnet
    Properties:
      VpcId: !Ref myVPC
      CidrBlock: 10.0.0.0/24
      AvailabilityZone: !Select [ 1, !GetAZs ]
  myInternetGateway:
    Type: AWS::EC2::InternetGateway
  myGatewayAttachment:
    Type: AWS::EC2::VPCGatewayAttachment
    Properties:
      VpcId: !Ref myVPC
      InternetGatewayId: !Ref myInternetGateway
  myRouteTable:
    Type: AWS::EC2::RouteTable
    Properties:
      VpcId: !Ref myVPC
  myRoute:
    Type: AWS::EC2::Route
    Properties:
      DestinationCidrBlock: 0.0.0.0/0
      GatewayId: !Ref myInternetGateway
      RouteTableId: !Ref myRouteTable
  myRouteTableAssociation:
    Type: AWS::EC2::SubnetRouteTableAssociation
    Properties:
      RouteTableId: !Ref myRouteTable
      SubnetId: !Ref mySubnet
```

## See also

- [Elastic Load Balancing Template Snippets](../userguide/quickref-elb.md)

- [CreateLoadBalancer](../../../../reference/elasticloadbalancing/2012-06-01/apireference/api-createloadbalancer.md) in the
_Elastic Load Balancing API Reference (version 2012-06-01)_

- [ModifyLoadBalancerAttributes](../../../../reference/elasticloadbalancing/2012-06-01/apireference/api-modifyloadbalancerattributes.md)
in the _Elastic Load Balancing API Reference (version 2012-06-01)_

- [ConfigureHealthCheck](../../../../reference/elasticloadbalancing/2012-06-01/apireference/api-configurehealthcheck.md) in the
_Elastic Load Balancing API Reference (version 2012-06-01)_

- [User Guide for Classic Load Balancers](../../../elasticloadbalancing/latest/classic.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Elastic Load Balancing

AccessLoggingPolicy

All content copied from https://docs.aws.amazon.com/.
