# DescribeVpcs

Describes your VPCs. The default is to describe all your VPCs.
Alternatively, you can specify specific VPC IDs or filter the results to
include only the VPCs that match specific criteria.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Filter.N**

The filters.

- `cidr` \- The primary IPv4 CIDR block of the VPC. The CIDR block you
specify must exactly match the VPC's CIDR block for information to be returned
for the VPC. Must contain the slash followed by one or two digits (for example,
`/28`).

- `cidr-block-association.cidr-block` \- An IPv4 CIDR block associated with the
VPC.

- `cidr-block-association.association-id` \- The association ID for
an IPv4 CIDR block associated with the VPC.

- `cidr-block-association.state` \- The state of an IPv4 CIDR block
associated with the VPC.

- `dhcp-options-id` \- The ID of a set of DHCP options.

- `ipv6-cidr-block-association.ipv6-cidr-block` \- An IPv6 CIDR
block associated with the VPC.

- `ipv6-cidr-block-association.ipv6-pool` \- The ID of the IPv6 address pool from which the IPv6 CIDR block is allocated.

- `ipv6-cidr-block-association.association-id` \- The association
ID for an IPv6 CIDR block associated with the VPC.

- `ipv6-cidr-block-association.state` \- The state of an IPv6 CIDR
block associated with the VPC.

- `is-default` \- Indicates whether the VPC is the default VPC.

- `owner-id` \- The ID of the AWS account that owns the VPC.

- `state` \- The state of the VPC ( `pending` \| `available`).

- `tag` \- The key/value combination of a tag assigned to the resource. Use the tag key in the filter name and the tag value as the filter value.
For example, to find all resources that have a tag with the key `Owner` and the value `TeamA`, specify `tag:Owner` for the filter name and `TeamA` for the filter value.

- `tag-key` \- The key of a tag assigned to the resource. Use this filter to find all resources assigned a tag with a specific key, regardless of the tag value.

- `vpc-id` \- The ID of the VPC.

Type: Array of [Filter](api-filter.md) objects

Required: No

**MaxResults**

The maximum number of items to return for this request.
To get the next page of items, make another request with the token returned in the output.
For more information, see [Pagination](query-requests.md#api-pagination).

Type: Integer

Valid Range: Minimum value of 5. Maximum value of 1000.

Required: No

**NextToken**

The token returned from a previous paginated request. Pagination continues from the end of the items returned by the previous request.

Type: String

Required: No

**VpcId.N**

The IDs of the VPCs.

Type: Array of strings

Required: No

## Response Elements

The following elements are returned by the service.

**nextToken**

The token to include in another request to get the next page of items. This value is `null` when there are no more items to return.

Type: String

**requestId**

The ID of the request.

Type: String

**vpcSet**

Information about the VPCs.

Type: Array of [Vpc](api-vpc.md) objects

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example 1

This example describes the specified VPC.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeVpcs
&VpcId.1=vpc-081ec835f3EXAMPLE
&vpcId.2=vpc-0ee975135dEXAMPLE
&VpcId.3=vpc-06e4ab6c6cEXAMPLE
&AUTHPARAMS
```

#### Sample Response

```

<DescribeVpcsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>8b67ac77-886c-4027-8f0e-d351f7fc9971</requestId>
    <vpcSet>
        <item>
            <vpcId>vpc-081ec835f3EXAMPLE</vpcId>
            <ownerId>111122223333</ownerId>
            <state>available</state>
            <cidrBlock>10.0.1.0/24</cidrBlock>
            <cidrBlockAssociationSet>
                <item>
                    <cidrBlock>10.0.1.0/24</cidrBlock>
                    <associationId>vpc-cidr-assoc-043f572c17EXAMPLE</associationId>
                    <cidrBlockState>
                        <state>associated</state>
                    </cidrBlockState>
                </item>
            </cidrBlockAssociationSet>
            <dhcpOptionsId>dopt-19edf471</dhcpOptionsId>
            <tagSet>
                <item>
                    <key>Name</key>
                    <value>MyVPC</value>
                </item>
            </tagSet>
            <instanceTenancy>default</instanceTenancy>
            <isDefault>false</isDefault>
        </item>
        <item>
            <vpcId>vpc-0ee975135dEXAMPLE</vpcId>
            <ownerId>111122223333</ownerId>
            <state>available</state>
            <cidrBlock>10.0.0.0/16</cidrBlock>
            <cidrBlockAssociationSet>
                <item>
                    <cidrBlock>10.0.0.0/16</cidrBlock>
                    <associationId>vpc-cidr-assoc-067c3a0a1fEXAMPLE</associationId>
                    <cidrBlockState>
                        <state>associated</state>
                    </cidrBlockState>
                </item>
            </cidrBlockAssociationSet>
            <dhcpOptionsId>dopt-f30d649a</dhcpOptionsId>
            <instanceTenancy>default</instanceTenancy>
            <isDefault>false</isDefault>
        </item>
        <item>
            <vpcId>vpc-06e4ab6c6cEXAMPLE</vpcId>
            <ownerId>123456789012</ownerId>
            <state>available</state>
            <cidrBlock>10.0.0.0/16</cidrBlock>
            <cidrBlockAssociationSet>
                <item>
                    <cidrBlock>10.0.0.0/16</cidrBlock>
                    <associationId>vpc-cidr-assoc-00b17b4eddEXAMPLE</associationId>
                    <cidrBlockState>
                        <state>associated</state>
                    </cidrBlockState>
                </item>
            </cidrBlockAssociationSet>
            <dhcpOptionsId>dopt-19edf471</dhcpOptionsId>
            <tagSet>
                <item>
                    <key>Name</key>
                    <value>Shared VPC</value>
                </item>
            </tagSet>
            <instanceTenancy>default</instanceTenancy>
            <isDefault>false</isDefault>
        </item>
    </vpcSet>
</DescribeVpcsResponse>
```

### Example 2

This example uses filters to describe any VPC you own that uses the set of DHCP
options with the ID `dopt-7a8b9c2d` or `dopt-2b2a3d3c` and whose
state is `available`.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeVpcs
&Filter.1.Name=dhcp-options-id
&Filter.1.Value.1=dopt-7a8b9c2d
&Filter.1.Value.2=dopt-2b2a3d3c
&Filter.2.Name=state
&Filter.2.Value.1=available
&AUTHPARAMS
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DescribeVpcs)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DescribeVpcs)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describevpcs.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describevpcs.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describevpcs.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describevpcs.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describevpcs.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describevpcs.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DescribeVpcs)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describevpcs.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeVpcPeeringConnections

DescribeVpnConcentrators
