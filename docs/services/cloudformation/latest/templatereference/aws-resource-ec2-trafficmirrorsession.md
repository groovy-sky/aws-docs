This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::TrafficMirrorSession

Creates a Traffic Mirror session.

A Traffic Mirror session actively copies packets from a Traffic Mirror source to a
Traffic Mirror target. Create a filter, and then assign it to the session to define a
subset of the traffic to mirror, for example all TCP traffic.

The Traffic Mirror source and the Traffic Mirror target (monitoring appliances) can be
in the same VPC, or in a different VPC connected via VPC peering or a transit gateway.

By default, no traffic is mirrored. Use [AWS::EC2::TrafficMirrorFilterRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html) to specify filter rules that specify the
traffic to mirror.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::TrafficMirrorSession",
  "Properties" : {
      "Description" : String,
      "NetworkInterfaceId" : String,
      "OwnerId" : String,
      "PacketLength" : Integer,
      "SessionNumber" : Integer,
      "Tags" : [ Tag, ... ],
      "TrafficMirrorFilterId" : String,
      "TrafficMirrorTargetId" : String,
      "VirtualNetworkId" : Integer
    }
}

```

### YAML

```yaml

Type: AWS::EC2::TrafficMirrorSession
Properties:
  Description: String
  NetworkInterfaceId: String
  OwnerId: String
  PacketLength: Integer
  SessionNumber: Integer
  Tags:
    - Tag
  TrafficMirrorFilterId: String
  TrafficMirrorTargetId: String
  VirtualNetworkId: Integer

```

## Properties

`Description`

The description of the Traffic Mirror session.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkInterfaceId`

The ID of the source network interface.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OwnerId`

The ID of the account that owns the Traffic Mirror session.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PacketLength`

The number of bytes in each packet to mirror. These are bytes after the VXLAN header. Do
not specify this parameter when you want to mirror the entire packet. To mirror a subset of
the packet, set this to the length (in bytes) that you want to mirror. For example, if you
set this value to 100, then the first 100 bytes that meet the filter criteria are copied to
the target.

If you do not want to mirror the entire packet, use the `PacketLength` parameter to specify the number of bytes in each packet to mirror.

For sessions with Network Load Balancer (NLB) Traffic Mirror targets the default `PacketLength` will be set to 8500. Valid values are 1-8500. Setting a `PacketLength` greater than 8500 will result in an error response.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SessionNumber`

The session number determines the order in which sessions are evaluated when an interface is used by multiple sessions. The first session with a matching filter is the one that mirrors the packets.

Valid values are 1-32766.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to assign to a Traffic Mirror session.

_Required_: No

_Type_: Array of [Tag](aws-properties-ec2-trafficmirrorsession-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrafficMirrorFilterId`

The ID of the Traffic Mirror filter.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrafficMirrorTargetId`

The ID of the Traffic Mirror target.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VirtualNetworkId`

The VXLAN ID for the Traffic Mirror session. For more information about the VXLAN
protocol, see [RFC 7348](https://datatracker.ietf.org/doc/html/rfc7348). If you do
not specify a `VirtualNetworkId`, an account-wide unique ID is chosen at
random.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the Traffic Mirror Session.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

### Create a traffic mirror session

This is a traffic mirror session that mirrors the first 100 bytes in each
packet.

#### JSON

```json

{
  "SampleTrafficMirrorSession": {
    "Type": "AWS::EC2::TrafficMirrorSession",
    "Properties": {
      "Description": "Example traffic mirror session",
      "NetworkInterfaceId": "eni-070203a001EXAMPLE",
      "TrafficMirrorTargetId": "tmt-5319fsEXAMPLE",
      "TrafficMirrorFilterId": "tmf-1tdbhqEXAMPLE",
      "SessionNumber": 1,
      "PacketLength": 100,
      "VirtualNetworkId": 1234,
      "Tags": [
        {
          "Key": "Name",
          "Value": "SampleSession"
        }
      ]
    }
  }
}

```

#### YAML

```yaml

SampleTrafficMirrorSession:
  Type: "AWS::EC2::TrafficMirrorSession"
  Properties:
    Description: "Example traffic mirror session"
    NetworkInterfaceId: "eni-070203a001EXAMPLE"
    TrafficMirrorTargetId: "tmt-5319fsEXAMPLE"
    TrafficMirrorFilterId: "tmf-1tdbhqEXAMPLE"
    SessionNumber: 1
    PacketLength: 100
    VirtualNetworkId: 1234
    Tags:
    - Key: "Name"
      Value: "SampleSession"

```

## See also

- [Traffic mirror\
sessions](https://docs.aws.amazon.com/vpc/latest/mirroring/traffic-mirroring-sessions.html) in _Traffic Mirroring_

- [CreateTrafficMirrorSession](../../../../reference/awsec2/latest/apireference/api-createtrafficmirrorsession.md) in the _Amazon EC2 API_
_Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TrafficMirrorPortRange

Tag
