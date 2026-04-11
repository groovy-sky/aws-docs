This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::TrafficMirrorTarget

Specifies a target for your Traffic Mirror session.

A Traffic Mirror target is the destination for mirrored traffic. The Traffic Mirror
source and the Traffic Mirror target (monitoring appliances) can be in the same VPC, or in
different VPCs connected via VPC peering or a transit gateway.

A Traffic Mirror target can be a network interface, a Network Load Balancer, or a Gateway Load Balancer endpoint.

To use the target in a Traffic Mirror session, use [AWS::EC2::TrafficMirrorSession](../userguide/aws-resource-ec2-trafficmirrorsession.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::TrafficMirrorTarget",
  "Properties" : {
      "Description" : String,
      "GatewayLoadBalancerEndpointId" : String,
      "NetworkInterfaceId" : String,
      "NetworkLoadBalancerArn" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::EC2::TrafficMirrorTarget
Properties:
  Description: String
  GatewayLoadBalancerEndpointId: String
  NetworkInterfaceId: String
  NetworkLoadBalancerArn: String
  Tags:
    - Tag

```

## Properties

`Description`

The description of the Traffic Mirror target.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`GatewayLoadBalancerEndpointId`

The ID of the Gateway Load Balancer endpoint.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NetworkInterfaceId`

The network interface ID that is associated with the target.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NetworkLoadBalancerArn`

The Amazon Resource Name (ARN) of the Network Load Balancer that is associated with the target.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags to assign to the Traffic Mirror target.

_Required_: No

_Type_: Array of [Tag](aws-properties-ec2-trafficmirrortarget-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the Traffic Mirror target.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

- [Create a traffic mirror target associated with a Network Load Balancer](#aws-resource-ec2-trafficmirrortarget--examples--Create_a_traffic_mirror_target_associated_with_a_Network_Load_Balancer)

- [Create a traffic mirror target associated with a network interface](#aws-resource-ec2-trafficmirrortarget--examples--Create_a_traffic_mirror_target_associated_with_a_network_interface)

### Create a traffic mirror target associated with a Network Load Balancer

This is a traffic mirror target associated with a Network Load Balancer.

#### JSON

```json

{
  "SampleNLBTrafficMirrorTarget": {
    "Type": "AWS::EC2::TrafficMirrorTarget",
    "Properties": {
      "Description": "Example traffic mirror target associated with a network load balancer",
      "NetworkLoadBalancerArn": "arn:aws:elasticloadbalancing:us-east-1:724145273726:loadbalancer/net/NLB/7cabvhEXAMPLE",
       "Tags": [
        {
          "Key": "Name",
          "Value": "SampleNLBTarget"
        }
      ]
    }
  }
}
```

#### YAML

```yaml

SampleNLBTrafficMirrorTarget:
  Type: "AWS::EC2::TrafficMirrorTarget"
  Properties:
    Description: "Example traffic mirror target associated with a network load balancer",
    NetworkLoadBalancerArn: "arn:aws:elasticloadbalancing:us-east-1:724145273726:loadbalancer/net/NLB/7cabvhEXAMPLE"
  Tags:
    - Key: "Name"
      Value: "SampleNLBTarget"

```

### Create a traffic mirror target associated with a network interface

This is a traffic mirror target associated with a network interface.

#### JSON

```json

{
  "SampleNetworkInterfaceTarget": {
    "Type": "AWS::EC2::TrafficMirrorTarget",
    "Properties": {
      "Description": "Example traffic mirror target associated with a network interface",
      "NetworkInterfaceId": "eni-070203a001EXAMPLE",
      "Tags": [
        {
          "Key": "Name",
          "Value": "SampleNetworkInterfaceTarget"
        }
      ]
    }
  }
}

```

#### YAML

```yaml

SampleNetworkInterfaceTarget:
  Type: "AWS::EC2::TrafficMirrorTarget"
  Properties:
    Description: "Example traffic mirror target associated with a network interface"
    NetworkInterfaceId: "eni-070203a001EXAMPLE"
    Tags:
    - Key: "Name"
      Value: "SampleNetworkInterfaceTarget"

```

## See also

- [Traffic mirror\
targets](../../../vpc/latest/mirroring/traffic-mirroring-targets.md) in _Traffic Mirroring_

- [CreateTrafficMirrorTarget](../../../../reference/awsec2/latest/apireference/api-createtrafficmirrortarget.md) in the _Amazon EC2 API_
_Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
