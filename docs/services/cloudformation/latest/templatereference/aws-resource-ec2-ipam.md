This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::IPAM

IPAM is a VPC feature that you can use to automate your IP address management
workflows including assigning, tracking, troubleshooting, and auditing IP addresses
across AWS Regions and accounts throughout your AWS
Organization. For more information, see [What is IPAM?](https://docs.aws.amazon.com/vpc/latest/ipam/what-is-it-ipam.html) in the
_Amazon VPC IPAM User Guide_.

There are AWS Identity and Access Management (IAM) permissions required
to fully manage an IPAM in CloudFormation. For more information, see [Example\
policy](https://docs.aws.amazon.com/vpc/latest/ipam/iam-ipam-policy-examples.html) in the _Amazon VPC IPAM User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::IPAM",
  "Properties" : {
      "DefaultResourceDiscoveryOrganizationalUnitExclusions" : [ IpamOrganizationalUnitExclusion, ... ],
      "Description" : String,
      "EnablePrivateGua" : Boolean,
      "MeteredAccount" : String,
      "OperatingRegions" : [ IpamOperatingRegion, ... ],
      "Tags" : [ Tag, ... ],
      "Tier" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::IPAM
Properties:
  DefaultResourceDiscoveryOrganizationalUnitExclusions:
    - IpamOrganizationalUnitExclusion
  Description: String
  EnablePrivateGua: Boolean
  MeteredAccount: String
  OperatingRegions:
    - IpamOperatingRegion
  Tags:
    - Tag
  Tier: String

```

## Properties

`DefaultResourceDiscoveryOrganizationalUnitExclusions`

If your IPAM is integrated with AWS Organizations,
you can exclude an [organizational unit (OU)](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_getting-started_concepts.html#organizationalunit) from being managed by IPAM. When you exclude an OU, IPAM will not manage the IP addresses in accounts in that OU. For more information, see [Exclude organizational units from IPAM](https://docs.aws.amazon.com/vpc/latest/ipam/exclude-ous.html) in the _Amazon Virtual Private Cloud IP Address Manager User Guide_.

_Required_: No

_Type_: Array of [IpamOrganizationalUnitExclusion](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-ipam-ipamorganizationalunitexclusion.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description for the IPAM.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnablePrivateGua`

Enable this option to use your own GUA ranges as private IPv6 addresses. This option is disabled by default.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MeteredAccount`

A metered account is an AWS account that is charged for active IP addresses managed in IPAM. For more information, see [Enable cost distribution](https://docs.aws.amazon.com/vpc/latest/ipam/ipam-enable-cost-distro.html) in the _Amazon VPC IPAM User Guide_.

Possible values:

- `ipam-owner` (default): The AWS account which owns the IPAM is charged for all active IP addresses managed in IPAM.

- `resource-owner`: The AWS account that owns the IP address is charged for the active IP address.

_Required_: No

_Type_: String

_Allowed values_: `ipam-owner | resource-owner`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OperatingRegions`

The operating Regions for an IPAM. Operating Regions are AWS Regions where the IPAM is allowed to manage IP address CIDRs. IPAM only discovers and monitors resources in the AWS Regions you select as operating Regions.

For more information about operating Regions, see [Create an IPAM](https://docs.aws.amazon.com/vpc/latest/ipam/create-ipam.html) in the _Amazon VPC IPAM User Guide_.

_Required_: No

_Type_: Array of [IpamOperatingRegion](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-ipam-ipamoperatingregion.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The key/value combination of a tag assigned to the resource. Use the tag key in the filter name and the tag value as the filter value.
For example, to find all resources that have a tag with the key `Owner` and the value `TeamA`, specify `tag:Owner` for the filter name and `TeamA` for the filter value.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-ipam-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tier`

IPAM is offered in a Free Tier and an Advanced Tier. For more information about the
features available in each tier and the costs associated with the tiers, see the [VPC IPAM product pricing page](https://aws.amazon.com/vpc/pricing).

_Required_: No

_Type_: String

_Allowed values_: `free | advanced`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the IPAM ID.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN of the IPAM.

`DefaultResourceDiscoveryAssociationId`

The ID of the default resource discovery association.

`DefaultResourceDiscoveryId`

The ID of the default resource discovery.

`IpamId`

The ID of the IPAM.

`PrivateDefaultScopeId`

The ID of the default private scope.

`PublicDefaultScopeId`

The ID of the default public scope.

`ResourceDiscoveryAssociationCount`

The number of resource discovery associations.

`ScopeCount`

The number of scopes.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

IpamOperatingRegion
