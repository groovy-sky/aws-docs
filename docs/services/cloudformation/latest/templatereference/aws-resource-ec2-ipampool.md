This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::IPAMPool

In IPAM, a pool is a collection of contiguous IP addresses CIDRs. Pools enable you to organize your IP addresses according to your routing and security needs. For example, if you have separate routing and security needs for development and production applications, you can create a pool for each.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::IPAMPool",
  "Properties" : {
      "AddressFamily" : String,
      "AllocationDefaultNetmaskLength" : Integer,
      "AllocationMaxNetmaskLength" : Integer,
      "AllocationMinNetmaskLength" : Integer,
      "AllocationResourceTags" : [ Tag, ... ],
      "AutoImport" : Boolean,
      "AwsService" : String,
      "Description" : String,
      "IpamScopeId" : String,
      "Locale" : String,
      "ProvisionedCidrs" : [ ProvisionedCidr, ... ],
      "PublicIpSource" : String,
      "PubliclyAdvertisable" : Boolean,
      "SourceIpamPoolId" : String,
      "SourceResource" : SourceResource,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::EC2::IPAMPool
Properties:
  AddressFamily: String
  AllocationDefaultNetmaskLength: Integer
  AllocationMaxNetmaskLength: Integer
  AllocationMinNetmaskLength: Integer
  AllocationResourceTags:
    - Tag
  AutoImport: Boolean
  AwsService: String
  Description: String
  IpamScopeId: String
  Locale: String
  ProvisionedCidrs:
    - ProvisionedCidr
  PublicIpSource: String
  PubliclyAdvertisable: Boolean
  SourceIpamPoolId: String
  SourceResource:
    SourceResource
  Tags:
    - Tag

```

## Properties

`AddressFamily`

The address family of the pool.

_Required_: Yes

_Type_: String

_Allowed values_: `ipv4 | ipv6`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AllocationDefaultNetmaskLength`

The default netmask length for allocations added to this pool. If, for example, the CIDR assigned to this pool is 10.0.0.0/8 and
you enter 16 here, new allocations will default to 10.0.0.0/16.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AllocationMaxNetmaskLength`

The maximum netmask length possible for CIDR allocations in this IPAM pool to be compliant. The maximum netmask length must be greater than the minimum netmask length. Possible netmask lengths for IPv4 addresses are 0 - 32. Possible netmask lengths for IPv6 addresses are 0 - 128.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AllocationMinNetmaskLength`

The minimum netmask length required for CIDR allocations in this IPAM pool to be compliant. The minimum netmask length must be less than the maximum netmask length. Possible netmask lengths for IPv4 addresses are 0 - 32. Possible netmask lengths for IPv6 addresses are 0 - 128.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AllocationResourceTags`

Tags that are required for resources that use CIDRs from this IPAM pool. Resources that do not have these tags will not be allowed to allocate space from the pool. If the resources have their tags changed after they have allocated space or if the allocation tagging requirements are changed on the pool, the resource may be marked as noncompliant.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-ipampool-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AutoImport`

If selected, IPAM will continuously look for resources within the CIDR range of this pool
and automatically import them as allocations into your IPAM. The CIDRs that will be allocated for
these resources must not already be allocated to other resources in order for the import to succeed. IPAM will import
a CIDR regardless of its compliance with the pool's allocation rules, so a resource might be imported and subsequently
marked as noncompliant. If IPAM discovers multiple CIDRs that overlap, IPAM will import the largest CIDR only. If IPAM
discovers multiple CIDRs with matching CIDRs, IPAM will randomly import one of them only.

A locale must be set on the pool for this feature to work.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AwsService`

Limits which service in AWS that the pool can be used in. "ec2", for example, allows users to use space for Elastic IP addresses and VPCs.

_Required_: No

_Type_: String

_Allowed values_: `ec2 | global-services`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

The description of the IPAM pool.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IpamScopeId`

The ID of the scope in which you would like to create the IPAM pool.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Locale`

The locale of the IPAM pool.

The locale for the pool should be one of the following:

- An AWS Region where you want this IPAM pool to be available for allocations.

- The network border group for an AWS Local Zone where you want this IPAM pool to be available for allocations ( [supported Local Zones](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-byoip.html#byoip-zone-avail)). This option is only available for IPAM IPv4 pools in the public scope.

If you choose an AWS Region for locale that has not been configured as an operating Region for the IPAM, you'll get an error.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProvisionedCidrs`

Information about the CIDRs provisioned to an IPAM pool.

_Required_: No

_Type_: Array of [ProvisionedCidr](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-ipampool-provisionedcidr.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PublicIpSource`

The IP address source for pools in the public scope. Only used for provisioning IP address CIDRs to pools in the public scope. Default is `BYOIP`. For more information, see [Create IPv6 pools](https://docs.aws.amazon.com/vpc/latest/ipam/intro-create-ipv6-pools.html) in the _Amazon VPC IPAM User Guide_.
By default, you can add only one Amazon-provided IPv6 CIDR block to a top-level IPv6 pool. For information on increasing the default limit, see [Quotas for your IPAM](https://docs.aws.amazon.com/vpc/latest/ipam/quotas-ipam.html) in the _Amazon VPC IPAM User Guide_.

_Required_: No

_Type_: String

_Allowed values_: `byoip | amazon`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PubliclyAdvertisable`

Determines if a pool is publicly advertisable. This option is not available for pools with AddressFamily set to `ipv4`.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceIpamPoolId`

The ID of the source IPAM pool. You can use this option to create an IPAM pool within an existing source pool.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceResource`

The resource used to provision CIDRs to a resource planning pool.

_Required_: No

_Type_: [SourceResource](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-ipampool-sourceresource.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The key/value combination of a tag assigned to the resource. Use the tag key in the filter name and the tag value as the filter value.
For example, to find all resources that have a tag with the key `Owner` and the value `TeamA`, specify `tag:Owner` for the filter name and `TeamA` for the filter value.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-ipampool-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the IPAM pool ID.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN of the IPAM pool.

`IpamArn`

The ARN of the IPAM.

`IpamPoolId`

The ID of the IPAM pool.

`IpamScopeArn`

The ARN of the scope of the IPAM pool.

`IpamScopeType`

The scope of the IPAM.

`PoolDepth`

The depth of pools in your IPAM pool. The pool depth quota is 10.

`State`

The state of the IPAM pool.

`StateMessage`

A message related to the failed creation of an IPAM pool.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::EC2::IPAMAllocation

ProvisionedCidr
