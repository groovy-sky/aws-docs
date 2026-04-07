This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DirectConnect::TransitVirtualInterface BgpPeer

Information about a BGP peer.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AddressFamily" : String,
  "AmazonAddress" : String,
  "Asn" : String,
  "AuthKey" : String,
  "BgpPeerId" : String,
  "CustomerAddress" : String
}

```

### YAML

```yaml

  AddressFamily: String
  AmazonAddress: String
  Asn: String
  AuthKey: String
  BgpPeerId: String
  CustomerAddress: String

```

## Properties

`AddressFamily`

The address family for the BGP peer.

_Required_: Yes

_Type_: String

_Pattern_: `^(ipv4)|(ipv6)$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AmazonAddress`

The IP address assigned to the Amazon interface.

_Required_: No

_Type_: String

_Pattern_: `^[0-9a-fA-F:.]+/[0-9]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Asn`

The autonomous system number (ASN). The valid range is from 1 to 4294967294 for Border Gateway Protocol (BGP) configuration. If you provide a number greater than the maximum, an error is returned.

This is configured as a string to support long ASNs. For more details about long ASN support, see [Long ASN support in Direct Connect](https://docs.aws.amazon.com/directconnect/latest/UserGuide/long-asn-support.html) in the _Direct Connect User Guide_.

_Required_: Yes

_Type_: String

_Pattern_: `^[1-9][0-9]*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuthKey`

The authentication key for BGP configuration. This string has a minimum length of 6 characters and and a maximun lenth of 80 characters.

_Required_: No

_Type_: String

_Pattern_: ``^[A-Za-z0-9\\!"#$%&'()*+,\-./:;<=>?@\[\]\^_`{|}~]{6,80}$``

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BgpPeerId`

The ID of the BGP peer.

_Required_: No

_Type_: String

_Pattern_: `^dxpeer-[a-z0-9]{8}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomerAddress`

The IP address assigned to the customer interface.

_Required_: No

_Type_: String

_Pattern_: `^[0-9a-fA-F:.]+/[0-9]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::DirectConnect::TransitVirtualInterface

Tag
