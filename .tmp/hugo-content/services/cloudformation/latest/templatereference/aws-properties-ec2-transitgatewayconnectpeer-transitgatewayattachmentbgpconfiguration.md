This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::TransitGatewayConnectPeer TransitGatewayAttachmentBgpConfiguration

The BGP configuration information.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BgpStatus" : String,
  "PeerAddress" : String,
  "PeerAsn" : Number,
  "TransitGatewayAddress" : String,
  "TransitGatewayAsn" : Number
}

```

### YAML

```yaml

  BgpStatus: String
  PeerAddress: String
  PeerAsn: Number
  TransitGatewayAddress: String
  TransitGatewayAsn: Number

```

## Properties

`BgpStatus`

The BGP status.

_Required_: No

_Type_: String

_Allowed values_: `up | down`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PeerAddress`

The interior BGP peer IP address for the appliance.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PeerAsn`

The peer Autonomous System Number (ASN).

_Required_: No

_Type_: Number

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TransitGatewayAddress`

The interior BGP peer IP address for the transit gateway.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TransitGatewayAsn`

The transit gateway Autonomous System Number (ASN).

_Required_: No

_Type_: Number

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

TransitGatewayConnectPeerConfiguration

All content copied from https://docs.aws.amazon.com/.
