This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::TransitGateway

Specifies a transit gateway.

You can use a transit gateway to interconnect your virtual private clouds (VPC) and
on-premises networks. After the transit gateway enters the `available` state,
you can attach your VPCs and VPN connections to the transit gateway.

To attach your VPCs, use [AWS::EC2::TransitGatewayAttachment](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayattachment.html).

To attach a VPN connection, use [AWS::EC2::CustomerGateway](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-customer-gateway.html) to create a customer gateway and specify the ID of
the customer gateway and the ID of the transit gateway in a call to [AWS::EC2::VPNConnection](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-connection.html).

When you create a transit gateway, we create a default transit gateway route table and
use it as the default association route table and the default propagation route table. You
can use [AWS::EC2::TransitGatewayRouteTable](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayroutetable.html) to create additional transit gateway route
tables. If you disable automatic route propagation, we do not create a default transit
gateway route table. You can use [AWS::EC2::TransitGatewayRouteTablePropagation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayroutetablepropagation.html) to propagate routes from a
resource attachment to a transit gateway route table. If you disable automatic
associations, you can use [AWS::EC2::TransitGatewayRouteTableAssociation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayroutetableassociation.html) to associate a resource
attachment with a transit gateway route table.

To create a transit gateway with `EncryptionSupport` enabled through CloudFormation, you will need the `ec2:ModifyTransitGateway` Identity and Access Management (IAM) permission. For more information, see
`ModifyTransitGateway` in [Actions, resources, and condition keys for Amazon EC2](https://docs.aws.amazon.com/service-authorization/latest/reference/list_amazonec2.html#amazonec2-actions-as-) of the
_Identify and Access Management Service Authorization Reference_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::TransitGateway",
  "Properties" : {
      "AmazonSideAsn" : Integer,
      "AssociationDefaultRouteTableId" : String,
      "AutoAcceptSharedAttachments" : String,
      "DefaultRouteTableAssociation" : String,
      "DefaultRouteTablePropagation" : String,
      "Description" : String,
      "DnsSupport" : String,
      "EncryptionSupport" : String,
      "MulticastSupport" : String,
      "PropagationDefaultRouteTableId" : String,
      "SecurityGroupReferencingSupport" : String,
      "Tags" : [ Tag, ... ],
      "TransitGatewayCidrBlocks" : [ String, ... ],
      "VpnEcmpSupport" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::TransitGateway
Properties:
  AmazonSideAsn: Integer
  AssociationDefaultRouteTableId: String
  AutoAcceptSharedAttachments: String
  DefaultRouteTableAssociation: String
  DefaultRouteTablePropagation: String
  Description: String
  DnsSupport: String
  EncryptionSupport: String
  MulticastSupport: String
  PropagationDefaultRouteTableId: String
  SecurityGroupReferencingSupport: String
  Tags:
    - Tag
  TransitGatewayCidrBlocks:
    - String
  VpnEcmpSupport: String

```

## Properties

`AmazonSideAsn`

A private Autonomous System Number (ASN) for the Amazon side of a BGP session. The range
is 64512 to 65534 for 16-bit ASNs. The default is 64512.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AssociationDefaultRouteTableId`

The ID of the default association route table.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AutoAcceptSharedAttachments`

Enable or disable automatic acceptance of attachment requests. Disabled by default.

_Required_: No

_Type_: String

_Allowed values_: `enable | disable`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultRouteTableAssociation`

Enable or disable automatic association with the default association route table. Enabled by default. If `DefaultRouteTableAssociation` is set to enable, AWS Transit Gateway will create the default transit gateway route table.

_Required_: No

_Type_: String

_Allowed values_: `enable | disable`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultRouteTablePropagation`

Enable or disable automatic propagation of routes to the default propagation route table. Enabled by default. If `DefaultRouteTablePropagation` is set to enable, AWS Transit Gateway will create the default transit gateway route table.

_Required_: No

_Type_: String

_Allowed values_: `enable | disable`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the transit gateway.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DnsSupport`

Enable or disable DNS support. Enabled by default.

_Required_: No

_Type_: String

_Allowed values_: `enable | disable`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionSupport`

Enable or disable encryption support. Disabled by default.

_Required_: No

_Type_: String

_Allowed values_: `disable | enable`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MulticastSupport`

Indicates whether multicast is enabled on the transit gateway

_Required_: No

_Type_: String

_Allowed values_: `enable | disable`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PropagationDefaultRouteTableId`

The ID of the default propagation route table.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityGroupReferencingSupport`

Enables you to reference a security group across VPCs attached to a transit gateway (TGW). Use this option to simplify security group management and control of instance-to-instance traffic across VPCs that are connected by transit gateway. You can also use this option to migrate from VPC peering (which was the only option that supported security group referencing) to transit gateways (which now also support security group referencing). This option is disabled by default and there are no additional costs to use this feature.

For important information about this feature, see [Create a transit gateway](https://docs.aws.amazon.com/vpc/latest/tgw/tgw-transit-gateways.html#create-tgw) in the _AWS Transit Gateway Guide_.

_Required_: No

_Type_: String

_Allowed values_: `enable | disable`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags for the transit gateway.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-transitgateway-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TransitGatewayCidrBlocks`

The transit gateway CIDR blocks.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpnEcmpSupport`

Enable or disable Equal Cost Multipath Protocol support. Enabled by default.

_Required_: No

_Type_: String

_Allowed values_: `enable | disable`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the transit gateway.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`EncryptionSupportState`

The encryption support state of the transit gateway.

`Id`

The ID of the transit gateway.

## Examples

### Create a transit gateway

The following example declares a transit gateway.

#### JSON

```json

"myTransitGateway": {
   "Type": "AWS::EC2::TransitGateway",
   "Properties": {
      "AmazonSideAsn": 65000,
      "Description": "TGW Route Integration Test",
      "AutoAcceptSharedAttachments": "disable",
      "DefaultRouteTableAssociation": "enable",
      "DnsSupport": "enable",
      "EncryptionSupport": "disable",
      "VpnEcmpSupport": "enable",
      "Tags": [
         {
            "Key": "Application",
            "Value": {
               "Ref": "AWS::StackId"
            }
         }
      ]
   }
}
```

#### YAML

```yaml

  myTransitGateway:
    Type: "AWS::EC2::TransitGateway"
    Properties:
      AmazonSideAsn: 65000
      Description: "TGW Route Integration Test"
      AutoAcceptSharedAttachments: "disable"
      DefaultRouteTableAssociation: "enable"
      DnsSupport: "enable"
      EncryptionSupport: "disable"
      VpnEcmpSupport: "enable"
      Tags:
      - Key: Application
        Value: !Ref 'AWS::StackId'
```

## See also

- [CreateTransitGateway](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateTransitGateway.html) in the _Amazon EC2 API_
_Reference_

- [AWS::RAM::ResourceShare](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-resourceshare.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
