# DescribeStaleSecurityGroups

Describes the stale security group rules for security groups referenced across a VPC
peering connection, transit gateway connection, or with a security group VPC
association. Rules are stale when they reference a deleted security group. Rules can
also be stale if they reference a security group in a peer VPC for which the VPC peering
connection has been deleted, across a transit gateway where the transit gateway has been
deleted (or [the transit\
gateway security group referencing feature](../../../../services/vpc/latest/tgw/tgw-vpc-attachments.md#vpc-attachment-security) has been disabled), or if a
security group VPC association has been disassociated.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**MaxResults**

The maximum number of items to return for this request. To get the next page of items,
make another request with the token returned in the output. For more information,
see [Pagination](query-requests.md#api-pagination).

Type: Integer

Valid Range: Minimum value of 5. Maximum value of 255.

Required: No

**NextToken**

The token returned from a previous paginated request. Pagination continues from the end of the items returned by the previous request.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**VpcId**

The ID of the VPC.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**nextToken**

The token to include in another request to get the next page of items. This value is `null` when there are no more items to return.

Type: String

**requestId**

The ID of the request.

Type: String

**staleSecurityGroupSet**

Information about the stale security groups.

Type: Array of [StaleSecurityGroup](api-stalesecuritygroup.md) objects

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example describes stale security group rules for vpc-11223344. The
response shows that sg-5fa68d3a in your account has a stale ingress SSH rule that
references sg-279ab042 in the peer VPC, and sg-fe6fba9a in your account has a stale egress
SSH rule that references sg-ef6fba8b in the peer VPC.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeStaleSecurityGroups
&VpcId=vpc-11223344
&AUTHPARAMS
```

#### Sample Response

```

<DescribeStaleSecurityGroupsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>ece1f9a0-b201-4eec-b74b-example</requestId>
        <staleSecurityGroupSet>
        <item>
            <staleIpPermissionsEgress>
                <item>
                    <fromPort>22</fromPort>
                    <toPort>22</toPort>
                    <groups>
                        <item>
                            <vpcId>vpc-7a20e51f</vpcId>
                            <groupId>sg-ef6fba8b</groupId>
                            <vpcPeeringConnectionId>pcx-b04deed9</vpcPeeringConnectionId>
                            <peeringStatus>active</peeringStatus>
                            <description>Access to pcx-b04deed9</description>
                        </item>
                    </groups>
                    <ipProtocol>tcp</ipProtocol>
                </item>
            </staleIpPermissionsEgress>
            <groupName>Sg-1</groupName>
            <vpcId>vpc-11223344</vpcId>
            <groupId>sg-fe6fba9a</groupId>
            <description>Sg-1 for peering</description>
            <staleIpPermissions/>
        </item>
        <item>
            <staleIpPermissionsEgress/>
            <groupName>Sg-2</groupName>
            <vpcId>vpc-11223344</vpcId>
            <groupId>sg-5fa68d3a</groupId>
            <description>Sg-2 for peering</description>
            <staleIpPermissions>
                <item>
                    <fromPort>22</fromPort>
                    <toPort>22</toPort>
                    <groups>
                        <item>
                            <vpcId>vpc-7a20e51f</vpcId>
                            <groupId>sg-279ab042</groupId>
                            <vpcPeeringConnectionId>pcx-b04deed9</vpcPeeringConnectionId>
                            <peeringStatus>active</peeringStatus>
                            <description>Access from pcx-b04deed9</description>
                        </item>
                    </groups>
                    <ipProtocol>tcp</ipProtocol>
                </item>
            </staleIpPermissions>
        </item>
    </staleSecurityGroupSet>
</DescribeStaleSecurityGroupsResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DescribeStaleSecurityGroups)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DescribeStaleSecurityGroups)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describestalesecuritygroups.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describestalesecuritygroups.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describestalesecuritygroups.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describestalesecuritygroups.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describestalesecuritygroups.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describestalesecuritygroups.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DescribeStaleSecurityGroups)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describestalesecuritygroups.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeSpotPriceHistory

DescribeStoreImageTasks
