This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::TransitGatewayConnectPeer TransitGatewayConnectPeerConfiguration

Describes the Connect peer details.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BgpConfigurations" : [ TransitGatewayAttachmentBgpConfiguration, ... ],
  "InsideCidrBlocks" : [ String, ... ],
  "PeerAddress" : String,
  "Protocol" : String,
  "TransitGatewayAddress" : String
}

```

### YAML

```yaml

  BgpConfigurations:
    - TransitGatewayAttachmentBgpConfiguration
  InsideCidrBlocks:
    - String
  PeerAddress: String
  Protocol: String
  TransitGatewayAddress: String

```

## Properties

`BgpConfigurations`

The BGP configuration details.

_Required_: No

_Type_: Array of [TransitGatewayAttachmentBgpConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-transitgatewayconnectpeer-transitgatewayattachmentbgpconfiguration.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InsideCidrBlocks`

The range of interior BGP peer IP addresses.

_Required_: Yes

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PeerAddress`

The Connect peer IP address on the appliance side of the tunnel.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Protocol`

The tunnel protocol.

_Required_: No

_Type_: String

_Allowed values_: `gre`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TransitGatewayAddress`

The Connect peer IP address on the transit gateway side of the tunnel.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TransitGatewayAttachmentBgpConfiguration

AWS::EC2::TransitGatewayMeteringPolicy
