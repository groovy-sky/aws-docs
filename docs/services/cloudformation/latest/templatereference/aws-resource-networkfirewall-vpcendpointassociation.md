This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkFirewall::VpcEndpointAssociation

A VPC endpoint association defines a single subnet to use for a firewall endpoint for a `Firewall`.
You can define VPC endpoint associations only in the Availability Zones that already have
a subnet mapping defined in the `Firewall` resource.

###### Note

You can retrieve the list of Availability Zones that are available for use by calling `DescribeFirewallMetadata`.

To manage firewall endpoints, first, in the `Firewall` specification, you specify a single VPC and one subnet
for each of the Availability Zones where you want to use the firewall. Then you can define additional endpoints as
VPC endpoint associations.

You can use VPC endpoint associations to expand the protections of the firewall as follows:

- **Protect multiple VPCs with a single firewall** \- You can use the firewall to protect other VPCs, either in your account or in accounts where the firewall is shared. You can only specify Availability Zones that already have a firewall endpoint defined in the `Firewall` subnet mappings.

- **Define multiple firewall endpoints for a VPC in an Availability Zone** \- You can create additional firewall endpoints for the VPC that you have defined in the firewall, in any Availability Zone that already has an endpoint defined in the `Firewall` subnet mappings. You can create multiple VPC endpoint associations for any other VPC where you use the firewall.

You can use AWS Resource Access Manager to share a `Firewall` that you own with other accounts, which gives them the ability to use the firewall
to create VPC endpoint associations. For information about sharing a firewall, see `PutResourcePolicy`
in this guide and see
[Sharing Network Firewall resources](https://docs.aws.amazon.com/network-firewall/latest/developerguide/sharing.html) in the _AWS Network Firewall Developer Guide_.

The status of the VPC endpoint association, which indicates whether it's ready to filter network traffic,
is provided in the corresponding VPC endpoint association status. You can retrieve both
the association and its status by calling `DescribeVpcEndpointAssociation`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::NetworkFirewall::VpcEndpointAssociation",
  "Properties" : {
      "Description" : String,
      "FirewallArn" : String,
      "SubnetMapping" : SubnetMapping,
      "Tags" : [ Tag, ... ],
      "VpcId" : String
    }
}

```

### YAML

```yaml

Type: AWS::NetworkFirewall::VpcEndpointAssociation
Properties:
  Description: String
  FirewallArn: String
  SubnetMapping:
    SubnetMapping
  Tags:
    - Tag
  VpcId: String

```

## Properties

`Description`

A description of the VPC endpoint association.

_Required_: No

_Type_: String

_Pattern_: `^.*$`

_Maximum_: `512`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FirewallArn`

The Amazon Resource Name (ARN) of the firewall.

_Required_: Yes

_Type_: String

_Pattern_: `^(arn:aws.*)$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SubnetMapping`

The ID for a subnet that's used in an association with a firewall. This is used in
`CreateFirewall`, `AssociateSubnets`, and `CreateVpcEndpointAssociation`. AWS Network Firewall
creates an instance of the associated firewall in each subnet that you specify, to filter
traffic in the subnet's Availability Zone.

_Required_: Yes

_Type_: [SubnetMapping](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-networkfirewall-vpcendpointassociation-subnetmapping.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The key:value pairs to associate with the resource.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-networkfirewall-vpcendpointassociation-tag.html)

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcId`

The unique identifier of the VPC for the endpoint association.

_Required_: Yes

_Type_: String

_Pattern_: `^vpc-[0-9a-f]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the firewall. For example:

`{ "Ref": "arn:aws:network-firewall:us-east-1:123456789012:vpc-endpoint-association/UUID" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`EndpointId`

The unique ID of the firewall endpoint for the subnet that you attached to the firewall.For example: "vpce-111122223333"

`VpcEndpointAssociationArn`

The Amazon Resource Name (ARN) of a VPC endpoint association.

`VpcEndpointAssociationId`

The unique identifier of the VPC endpoint association.

## Examples

### Create a VPC Endpoint Association

The following shows example VPC Endpoint Association specifications.

#### JSON

```json

"SampleVpcEndpointAssociation": {
    "Type": "AWS::NetworkFirewall::VpcEndpointAssociation",
    "Properties": {
        "Description": "VpcEndpointAssociation description goes here",
        "FirewallArn": {
            "Ref": "SampleFirewall"
        },
        "SubnetMapping": {
            "SubnetId": {
                "Ref": "SampleSubnet"
            }
        },
        "VpcId": {
            "Ref": "SampleVPC"
        },
        "Tags": [
            {
                "Key": "Foo",
                "Value": "Bar"
            }
        ]
    }
}
```

#### YAML

```yaml

SampleVpcEndpointAssociation:
  Type: AWS::NetworkFirewall::VpcEndpointAssociation
  Properties:
    FirewallArn: !Ref SampleFirewall
    VpcId: !Ref SampleVPC
    SubnetMapping:
      SubnetId: !Ref SampleSubnet
    Description: VpcEndpointAssociation description goes here
    Tags:
      - Key: Foo
        Value: Bar
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TLSInspectionConfiguration

SubnetMapping
