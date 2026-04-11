This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::RouteServerPeer BgpOptions

The BGP configuration options for this peer, including ASN (Autonomous System Number) and BFD (Bidrectional Forwarding Detection) settings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PeerAsn" : Integer,
  "PeerLivenessDetection" : String
}

```

### YAML

```yaml

  PeerAsn: Integer
  PeerLivenessDetection: String

```

## Properties

`PeerAsn`

The Border Gateway Protocol (BGP) Autonomous System Number (ASN) for the appliance. Valid values are from 1 to 4294967295. We recommend using a private ASN in the 64512–65534 (16-bit ASN) or 4200000000–4294967294 (32-bit ASN) range.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `4294967294`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PeerLivenessDetection`

The liveness detection protocol used for the BGP peer.

The requested liveness detection protocol for the BGP peer.

- `bgp-keepalive`: The standard BGP keep alive mechanism ( [RFC4271](https://www.rfc-editor.org/rfc/rfc4271)) that is stable but may take longer to fail-over in cases of network impact or router failure.

- `bfd`: An additional Bidirectional Forwarding Detection (BFD) protocol ( [RFC5880](https://www.rfc-editor.org/rfc/rfc5880)) that enables fast failover by using more sensitive liveness detection.

Defaults to `bgp-keepalive`.

_Required_: No

_Type_: String

_Allowed values_: `bfd | bgp-keepalive`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::EC2::RouteServerPeer

Tag

All content copied from https://docs.aws.amazon.com/.
