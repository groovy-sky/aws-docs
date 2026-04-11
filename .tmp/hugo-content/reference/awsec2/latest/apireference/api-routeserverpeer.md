# RouteServerPeer

Describes a BGP peer configuration for a route server endpoint.

A route server peer is a session between a route server endpoint and the device deployed in AWS (such as a firewall appliance or other network security function running on an EC2 instance). The device must meet these requirements:

- Have an elastic network interface in the VPC

- Support BGP (Border Gateway Protocol)

- Can initiate BGP sessions

## Contents

**bfdStatus**

The current status of the BFD session with this peer.

Type: [RouteServerBfdStatus](api-routeserverbfdstatus.md) object

Required: No

**bgpOptions**

The BGP configuration options for this peer, including ASN (Autonomous System Number) and BFD (Bidrectional Forwarding Detection) settings.

Type: [RouteServerBgpOptions](api-routeserverbgpoptions.md) object

Required: No

**bgpStatus**

The current status of the BGP session with this peer.

Type: [RouteServerBgpStatus](api-routeserverbgpstatus.md) object

Required: No

**endpointEniAddress**

The IP address of the Elastic network interface for the route server endpoint.

Type: String

Required: No

**endpointEniId**

The ID of the Elastic network interface for the route server endpoint.

Type: String

Required: No

**failureReason**

The reason for any failure in peer creation or operation.

Type: String

Required: No

**peerAddress**

The IPv4 address of the peer device.

Type: String

Required: No

**routeServerEndpointId**

The ID of the route server endpoint associated with this peer.

Type: String

Required: No

**routeServerId**

The ID of the route server associated with this peer.

Type: String

Required: No

**routeServerPeerId**

The unique identifier of the route server peer.

Type: String

Required: No

**state**

The current state of the route server peer.

Type: String

Valid Values: `pending | available | deleting | deleted | failing | failed`

Required: No

**subnetId**

The ID of the subnet containing the route server peer.

Type: String

Required: No

**TagSet.N**

Any tags assigned to the route server peer.

Type: Array of [Tag](api-tag.md) objects

Required: No

**vpcId**

The ID of the VPC containing the route server peer.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/routeserverpeer.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/routeserverpeer.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/routeserverpeer.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

RouteServerEndpoint

RouteServerPropagation
