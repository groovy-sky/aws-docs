This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkFirewall::Firewall

Use the firewall to provide stateful, managed, network firewall and intrusion detection and prevention filtering for your VPCs in Amazon VPC.

The firewall defines the configuration settings for an AWS Network Firewall firewall. The settings include the firewall policy, the subnets in your VPC to use for the firewall endpoints, and any tags that are attached to the firewall AWS resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::NetworkFirewall::Firewall",
  "Properties" : {
      "AvailabilityZoneChangeProtection" : Boolean,
      "AvailabilityZoneMappings" : [ AvailabilityZoneMapping, ... ],
      "DeleteProtection" : Boolean,
      "Description" : String,
      "EnabledAnalysisTypes" : [ String, ... ],
      "FirewallName" : String,
      "FirewallPolicyArn" : String,
      "FirewallPolicyChangeProtection" : Boolean,
      "SubnetChangeProtection" : Boolean,
      "SubnetMappings" : [ SubnetMapping, ... ],
      "Tags" : [ Tag, ... ],
      "TransitGatewayId" : String,
      "VpcId" : String
    }
}

```

### YAML

```yaml

Type: AWS::NetworkFirewall::Firewall
Properties:
  AvailabilityZoneChangeProtection: Boolean
  AvailabilityZoneMappings:
    - AvailabilityZoneMapping
  DeleteProtection: Boolean
  Description: String
  EnabledAnalysisTypes:
    - String
  FirewallName: String
  FirewallPolicyArn: String
  FirewallPolicyChangeProtection: Boolean
  SubnetChangeProtection: Boolean
  SubnetMappings:
    - SubnetMapping
  Tags:
    - Tag
  TransitGatewayId: String
  VpcId: String

```

## Properties

`AvailabilityZoneChangeProtection`

A setting indicating whether the firewall is protected against changes to its Availability Zone configuration. When set to `TRUE`, you must first disable this protection before adding or removing Availability Zones.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AvailabilityZoneMappings`

The Availability Zones where the firewall endpoints are created for a transit gateway-attached firewall. Each mapping specifies an Availability Zone where the firewall processes traffic.

_Required_: No

_Type_: Array of [AvailabilityZoneMapping](aws-properties-networkfirewall-firewall-availabilityzonemapping.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeleteProtection`

A flag indicating whether it is possible to delete the firewall. A setting of `TRUE` indicates
that the firewall is protected against deletion. Use this setting to protect against
accidentally deleting a firewall that is in use. When you create a firewall, the operation initializes this flag to `TRUE`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the firewall.

_Required_: No

_Type_: String

_Pattern_: `^.*$`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnabledAnalysisTypes`

An optional setting indicating the specific traffic analysis types to enable on the firewall.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FirewallName`

The descriptive name of the firewall. You can't change the name of a firewall after you create it.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9-]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FirewallPolicyArn`

The Amazon Resource Name (ARN) of the firewall policy.

The relationship of firewall to firewall policy is many to one. Each firewall requires
one firewall policy association, and you can use the same firewall policy for multiple
firewalls.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws.*$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FirewallPolicyChangeProtection`

A setting indicating whether the firewall is protected against a change to the firewall policy association.
Use this setting to protect against
accidentally modifying the firewall policy for a firewall that is in use. When you create a firewall, the operation initializes this setting to `TRUE`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetChangeProtection`

A setting indicating whether the firewall is protected against changes to the subnet associations.
Use this setting to protect against
accidentally modifying the subnet associations for a firewall that is in use. When you create a firewall, the operation initializes this setting to `TRUE`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetMappings`

The primary public subnets that Network Firewall is using for the firewall. Network Firewall creates a firewall endpoint in each subnet. Create a subnet mapping for each Availability Zone where you want to use the firewall.

These subnets are all defined for a single, primary VPC, and each must belong to a different Availability Zone. Each of these subnets establishes the availability of the firewall in its Availability Zone.

In addition to these subnets, you can define other endpoints for the firewall in `VpcEndpointAssociation` resources. You can define these additional endpoints for any VPC, and for any of the Availability Zones where the firewall resource already has a subnet mapping. VPC endpoint associations give you the ability to protect multiple VPCs using a single firewall, and to define multiple firewall endpoints for a VPC in a single Availability Zone.

_Required_: No

_Type_: Array of [SubnetMapping](aws-properties-networkfirewall-firewall-subnetmapping.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-networkfirewall-firewall-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TransitGatewayId`

The unique identifier of the transit gateway associated with this firewall. This field is only present for transit gateway-attached firewalls.

_Required_: No

_Type_: String

_Pattern_: `^tgw-[0-9a-z]+$`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcId`

The unique identifier of the VPC where the firewall is in use. You can't change the VPC of a firewall after you create the firewall.

_Required_: No

_Type_: String

_Pattern_: `^vpc-[0-9a-f]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the firewall. For example:

`{ "Ref": "arn:aws:network-firewall:us-east-1:012345678901:firewall/myFirewallName" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`EndpointIds`

The unique IDs of the firewall endpoints for all of the subnets that you attached to the firewall. The subnets are not listed in any particular order. For example: `["us-west-2c:vpce-111122223333", "us-west-2a:vpce-987654321098", "us-west-2b:vpce-012345678901"]`.

`FirewallArn`

The Amazon Resource Name (ARN) of the firewall.

`FirewallId`

The name of the firewallresource.

`TransitGatewayAttachmentId`

The unique identifier of the transit gateway attachment associated with this firewall. This field is only present for transit gateway-attached firewalls.

## Examples

### Create a firewall

The following shows example firewall specifications.

#### JSON

```json

"SampleFirewall": {
    "Type": "AWS::NetworkFirewall::Firewall",
    "Properties": {
        "FirewallName": "SampleFirewallName",
        "FirewallPolicyArn": {
            "Ref": "SampleFirewallPolicy"
        },
        "VpcId": {
            "Ref": "SampleVPC"
        },
        "SubnetMappings": [
            {
                "SubnetId": {
                    "Ref": "SampleSubnet1"
                }
            },
            {
                "SubnetId": {
                    "Ref": "SampleSubnet2"
                }
            }
        ],
        "Description": "Firewall description goes here",
        "Tags": [
            {
                "Key": "Foo",
                "Value": "Bar"
            }
        ]
    }
```

#### YAML

```yaml

SampleFirewall:
  Type: AWS::NetworkFirewall::Firewall
  Properties:
    FirewallName: SampleFirewallName
    FirewallPolicyArn: !Ref SampleFirewallPolicy
    VpcId: !Ref SampleVPC
    SubnetMappings:
      - SubnetId: !Ref SampleSubnet1
      - SubnetId: !Ref SampleSubnet2
    Description: Firewall description goes here
    Tags:
      - Key: Foo
                Value: Bar

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Network Firewall

AvailabilityZoneMapping

All content copied from https://docs.aws.amazon.com/.
