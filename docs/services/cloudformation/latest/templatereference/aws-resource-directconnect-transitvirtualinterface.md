This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DirectConnect::TransitVirtualInterface

Creates a transit virtual interface. A transit virtual interface should be used to access one or more transit gateways associated with Direct Connect gateways. A transit virtual interface enables the connection of multiple VPCs attached to a transit gateway to a Direct Connect gateway.

###### Important

If you associate your transit gateway with one or more Direct Connect gateways, the Autonomous System Number (ASN) used by the transit gateway and the Direct Connect gateway must be different. For example, if you use the default ASN 64512 for both your the transit gateway and Direct Connect gateway, the association request fails.

A jumbo MTU value must be either 1500 or 8500. No other values will be accepted. Setting
the MTU of a virtual interface to 8500 (jumbo frames) can cause an update to the underlying
physical connection if it wasn't updated to support jumbo frames. Updating the connection
disrupts network connectivity for all virtual interfaces associated with the connection for up
to 30 seconds. To check whether your connection supports jumbo frames, call [DescribeConnections](../../../../reference/directconnect/latest/apireference/api-describeconnections.md). To check whether your virtual
interface supports jumbo frames, call [DescribeVirtualInterfaces](../../../../reference/directconnect/latest/apireference/api-describevirtualinterfaces.md).

For more information, see [Direct Connect virtual interfaces](../../../directconnect/latest/userguide/workingwithvirtualinterfaces.md) in the _Direct Connect User Guide_.

Hosted virtual interfaces are supported by the CloudFormation resource for transit virtual interfaces. The CloudFormation stack account will own the virtual interface, allowing usage of a connection or LAG in another AWS account. The connection or LAG owner account must have a role allowing the stack account to allocate transit virtual interfaces.

For more information about hosted virtual interfaces, see [Hosted Direct Connect virtual interfaces](../../../directconnect/latest/userguide/hosted-vif.md) in the _Direct Connect User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DirectConnect::TransitVirtualInterface",
  "Properties" : {
      "AllocateTransitVirtualInterfaceRoleArn" : String,
      "BgpPeers" : [ BgpPeer, ... ],
      "ConnectionId" : String,
      "DirectConnectGatewayId" : String,
      "EnableSiteLink" : Boolean,
      "Mtu" : Integer,
      "Tags" : [ Tag, ... ],
      "VirtualInterfaceName" : String,
      "Vlan" : Integer
    }
}

```

### YAML

```yaml

Type: AWS::DirectConnect::TransitVirtualInterface
Properties:
  AllocateTransitVirtualInterfaceRoleArn: String
  BgpPeers:
    - BgpPeer
  ConnectionId: String
  DirectConnectGatewayId: String
  EnableSiteLink: Boolean
  Mtu: Integer
  Tags:
    - Tag
  VirtualInterfaceName: String
  Vlan: Integer

```

## Properties

`AllocateTransitVirtualInterfaceRoleArn`

The Amazon Resource Name (ARN) of the role to allocate the transit virtual interface. The role needs to be in the account which owns the connection or LAG, and must have `directconnect:AllocateTransitVirtualInterface` permissions. If there are tags on the virtual interface, `directconnect:TagResource` permissions are also required.

###### Note

This should only be used when creating hosted virtual interfaces.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[a-z-]*:iam::[0-9]{12}:role/.+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BgpPeers`

The BGP peers configured on this virtual interface.

Modifying the BGP peers on a virtual interface may cause interruptions.

_Required_: Yes

_Type_: Array of [BgpPeer](aws-properties-directconnect-transitvirtualinterface-bgppeer.md)

_Minimum_: `1`

_Maximum_: `2`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`ConnectionId`

The ID or ARN of the connection or LAG.

Connectivity over the virtual interface will be interrupted while associating to a new connection or LAG.

_Required_: Yes

_Type_: String

_Pattern_: `^((arn:aws[a-z-]*:directconnect:[a-z0-9-]+:[0-9]{12}:(dxcon/dxcon|dxlag/dxlag))|dx(con|lag))-[a-z0-9A-Z]{8,21}$`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`DirectConnectGatewayId`

The ID or ARN of the Direct Connect gateway.

_Required_: Yes

_Type_: String

_Pattern_: `^(arn:aws[a-z-]*:directconnect::[0-9]{12}:dx-gateway/)?[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EnableSiteLink`

Indicates whether to enable or disable SiteLink.

Connectivity over the virtual interface will be interrupted while enabling or disabling SiteLink.

_Required_: No

_Type_: Boolean

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Mtu`

The maximum transmission unit (MTU), in bytes. The supported values are 1500 and 8500. The default value is 1500.

Connectivity over the virtual interface will be interrupted while the MTU update is completed.

_Required_: No

_Type_: Integer

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Tags`

The tags associated with the transitive virtual interface.

_Required_: No

_Type_: Array of [Tag](aws-properties-directconnect-transitvirtualinterface-tag.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VirtualInterfaceName`

The name of the virtual interface assigned by the customer network. The name has a maximum of 100 characters. The following are valid characters: a-z, 0-9 and a hyphen (-).

_Required_: Yes

_Type_: String

_Pattern_: `^[\w \-_,\/]{1,100}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Vlan`

The ID of the VLAN.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Maximum_: `4095`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the transit virtual interface.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`VirtualInterfaceArn`

The Amazon Resource Name (ARN) of the transit virtual interface.

`VirtualInterfaceId`

The ID of the transit virtual interface.

## Examples

### Create a transit virtual interface using resources in the same account

This example shows a basic transit virtual interface setup using a pre-existing connection and Direct Connect gateway in the same account.

#### JSON

```json

{
  "Resources": {
    "myTransitVirtualInterface": {
      "Type": "AWS::DirectConnect::TransitVirtualInterface",
      "Properties": {
        "ConnectionId": "dxcon-fgsdqeuv",
        "DirectConnectGatewayId": "f07aff53-9814-41bd-8c09-aad589c88e87",
        "VirtualInterfaceName": "cfn-transitvirtualinterface-example",
        "Vlan": 101,
        "BgpPeers": [
          {
            "AddressFamily": "ipv4",
            "AmazonAddress": "192.168.1.1/30",
            "CustomerAddress": "192.168.1.2/30",
            "Asn": "65000",
            "AuthKey": "example-auth-key"
          },
          {
            "AddressFamily": "ipv6",
            "Asn": "65000"
          }
        ],
        "Tags": [
          {
            "Key": "example-key",
            "Value": "example-value"
          }
        ]
      }
    }
  }
}
```

#### YAML

```yaml

Resources:
  myTransitVirtualInterface:
    Type: AWS::DirectConnect::TransitVirtualInterface
    Properties:
      ConnectionId: dxcon-fgsdqeuv
      DirectConnectGatewayId: f07aff53-9814-41bd-8c09-aad589c88e87
      VirtualInterfaceName: cfn-transitvirtualinterface-example
      Vlan: 101
      BgpPeers:
      - AddressFamily: ipv4
        AmazonAddress: 192.168.1.1/30
        CustomerAddress: 192.168.1.2/30
        Asn: '65000'
        AuthKey: example-auth-key
      - AddressFamily: ipv6
        Asn: '65000'
      Tags:
      - Key: example-key
        Value: example-value
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

BgpPeer

All content copied from https://docs.aws.amazon.com/.
