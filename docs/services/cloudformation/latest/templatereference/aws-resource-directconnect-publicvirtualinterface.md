This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DirectConnect::PublicVirtualInterface

Creates a public virtual interface. A virtual interface is the VLAN that transports Direct Connect traffic.
A public virtual interface supports sending traffic to public services of AWS such as Amazon S3.

When creating an IPv6 public virtual interface ( `addressFamily` is `ipv6`), leave the `customer`
and `amazon` address fields blank to use auto-assigned IPv6 space. Custom IPv6 addresses are not supported.

For more information, see [Direct Connect virtual interfaces](https://docs.aws.amazon.com/directconnect/latest/UserGuide/WorkingWithVirtualInterfaces.html) in the _Direct Connect User Guide_.

Hosted virtual interfaces are supported by the CloudFormation resource for public virtual interfaces. The CloudFormation stack account will own the virtual interface, allowing usage of a connection or LAG in another AWS account. The connection or LAG owner account must have a role allowing the stack account to allocate public virtual interfaces.

For more information about hosted virtual interfaces, see [Hosted Direct Connect virtual interfaces](https://docs.aws.amazon.com/directconnect/latest/UserGuide/hosted-vif.html) in the _Direct Connect User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DirectConnect::PublicVirtualInterface",
  "Properties" : {
      "AllocatePublicVirtualInterfaceRoleArn" : String,
      "BgpPeers" : [ BgpPeer, ... ],
      "ConnectionId" : String,
      "RouteFilterPrefixes" : [ String, ... ],
      "Tags" : [ Tag, ... ],
      "VirtualInterfaceName" : String,
      "Vlan" : Integer
    }
}

```

### YAML

```yaml

Type: AWS::DirectConnect::PublicVirtualInterface
Properties:
  AllocatePublicVirtualInterfaceRoleArn: String
  BgpPeers:
    - BgpPeer
  ConnectionId: String
  RouteFilterPrefixes:
    - String
  Tags:
    - Tag
  VirtualInterfaceName: String
  Vlan: Integer

```

## Properties

`AllocatePublicVirtualInterfaceRoleArn`

The Amazon Resource Name (ARN) of the role to allocate the public virtual interface. The role needs to be in the account which owns the connection or LAG, and must have `directconnect:AllocatePublicVirtualInterface` permissions. If there are tags on the virtual interface, `directconnect:TagResource` permissions are also required.

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

_Type_: Array of [BgpPeer](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-directconnect-publicvirtualinterface-bgppeer.html)

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

`RouteFilterPrefixes`

The routes to be advertised to the AWS network in this Region. Applies to public virtual interfaces.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags associated with the public virtual interface.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-directconnect-publicvirtualinterface-tag.html)

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

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the public virtual interface.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`VirtualInterfaceArn`

The Amazon Resource Name (ARN) of the public virtual interface.

`VirtualInterfaceId`

The ID of the public virtual interface.

## Examples

### Create a public virtual interface using resources in the same account

This example shows a basic public virtual interface setup using a pre-existing connection in the same account.

#### JSON

```json

{
  "Resources": {
    "myPublicVirtualInterface": {
      "Type": "AWS::DirectConnect::PublicVirtualInterface",
      "Properties": {
        "ConnectionId": "dxcon-fgsdqeuv",
        "VirtualInterfaceName": "cfn-publicvirtualinterface-example",
        "Vlan": 101,
        "RouteFilterPrefixes": [
          "50.0.0.0/30"
        ],
        "BgpPeers": [
          {
            "AddressFamily": "ipv4",
            "AmazonAddress": "50.0.0.1/30",
            "CustomerAddress": "50.0.0.2/30",
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
  myPublicVirtualInterface:
    Type: AWS::DirectConnect::PublicVirtualInterface
    Properties:
      ConnectionId: dxcon-fgsdqeuv
      VirtualInterfaceName: cfn-publicvirtualinterface-example
      Vlan: 101
      RouteFilterPrefixes:
      - 50.0.0.0/30
      BgpPeers:
      - AddressFamily: ipv4
        AmazonAddress: 50.0.0.1/30
        CustomerAddress: 50.0.0.2/30
        Asn: '65000'
        AuthKey: example-auth-key
      - AddressFamily: ipv6
        Asn: '65000'
      Tags:
      - Key: example-key
        Value: example-value
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

BgpPeer
