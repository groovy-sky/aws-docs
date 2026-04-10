This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ODB::OdbPeeringConnection

Creates a peering connection between an ODB network and a VPC.

A peering connection enables private connectivity between the networks for application-tier communication.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ODB::OdbPeeringConnection",
  "Properties" : {
      "AdditionalPeerNetworkCidrs" : [ String, ... ],
      "DisplayName" : String,
      "OdbNetworkId" : String,
      "PeerNetworkId" : String,
      "PeerNetworkRouteTableIds" : [ String, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ODB::OdbPeeringConnection
Properties:
  AdditionalPeerNetworkCidrs:
    - String
  DisplayName: String
  OdbNetworkId: String
  PeerNetworkId: String
  PeerNetworkRouteTableIds:
    - String
  Tags:
    - Tag

```

## Properties

`AdditionalPeerNetworkCidrs`

Property description not available.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisplayName`

The display name of the ODB peering connection.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z_](?!.*--)[a-zA-Z0-9_-]*$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OdbNetworkId`

Property description not available.

_Required_: No

_Type_: String

_Pattern_: `^(arn:(?:aws|aws-cn|aws-us-gov|aws-iso-{0,1}[a-z]{0,1}):[a-z0-9-]+:[a-z0-9-]*:[0-9]+:[a-z0-9-]+/[a-zA-Z0-9_~.-]{6,64}|[a-zA-Z0-9_~.-]{6,64})$`

_Minimum_: `6`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PeerNetworkId`

Property description not available.

_Required_: No

_Type_: String

_Pattern_: `^(arn:(?:aws|aws-cn|aws-us-gov|aws-iso-{0,1}[a-z]{0,1}):[a-z0-9-]+:[a-z0-9-]*:[0-9]+:[a-z0-9-]+/[a-zA-Z0-9_~.-]{6,64}|[a-zA-Z0-9_~.-]{6,64})$`

_Minimum_: `6`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PeerNetworkRouteTableIds`

Property description not available.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](aws-properties-odb-odbpeeringconnection-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`OdbNetworkArn`

The Amazon Resource Name (ARN) of the ODB network that initiated the peering connection.

`OdbPeeringConnectionArn`

The Amazon Resource Name (ARN) of the ODB peering connection.

`OdbPeeringConnectionId`

The unique identifier of the ODB peering connection. A sample ID is `odbpcx-abcdefgh12345678`.

`PeerNetworkArn`

The Amazon Resource Name (ARN) of the peer network.

`PeerNetworkCidrs`

The CIDR blocks associated with the peering connection. These CIDR blocks define the IP address ranges that can communicate through the peering connection.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ZeroEtlAccess

Tag

All content copied from https://docs.aws.amazon.com/.
