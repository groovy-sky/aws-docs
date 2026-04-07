# RouteServerBgpOptionsRequest

The BGP configuration options requested for a route server peer.

## Contents

**PeerAsn**

The Border Gateway Protocol (BGP) Autonomous System Number (ASN) for the appliance. Valid values are from 1 to 4294967295. We recommend using a private ASN in the 64512–65534 (16-bit ASN) or 4200000000–4294967294 (32-bit ASN) range.

Type: Long

Required: Yes

**PeerLivenessDetection**

The requested liveness detection protocol for the BGP peer.

- `bgp-keepalive`: The standard BGP keep alive mechanism ( [RFC4271](https://www.rfc-editor.org/rfc/rfc4271)) that is stable but may take longer to fail-over in cases of network impact or router failure.

- `bfd`: An additional Bidirectional Forwarding Detection (BFD) protocol ( [RFC5880](https://www.rfc-editor.org/rfc/rfc5880)) that enables fast failover by using more sensitive liveness detection.

Defaults to `bgp-keepalive`.

Type: String

Valid Values: `bfd | bgp-keepalive`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/routeserverbgpoptionsrequest.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/routeserverbgpoptionsrequest.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/routeserverbgpoptionsrequest.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

RouteServerBgpOptions

RouteServerBgpStatus
