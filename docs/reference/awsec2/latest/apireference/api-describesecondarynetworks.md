# DescribeSecondaryNetworks

Describes one or more secondary networks.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request, and provides an error response. If you have the required permissions, the error response is `DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Filter.N**

The filters. The following are the possible values:

- `ipv4-cidr-block-association.association-id` \- The association ID for an IPv4 CIDR block associated with the secondary network.

- `ipv4-cidr-block-association.cidr-block` \- An IPv4 CIDR block associated with the secondary network.

- `ipv4-cidr-block-association.state` \- The state of an IPv4 CIDR block associated with the secondary network.

- `owner-id` \- The ID of the AWS account that owns the secondary network.

- `secondary-network-id` \- The ID of the secondary network.

- `secondary-network-arn` \- The ARN of the secondary network.

- `state` \- The state of the secondary network ( `create-in-progress` \| `create-complete` \| `create-failed` \| `delete-in-progress` \| `delete-complete` \| `delete-failed`).

- `tag`:<key> - The key/value combination of a tag assigned to the resource. Use the tag key in the filter name and the tag value as the filter value. For example, to find all resources that have a tag with the key `Owner` and the value `TeamA`, specify `tag:Owner` for the filter name and `TeamA` for the filter value.

- `tag-key` \- The key of a tag assigned to the resource. Use this filter to find all resources assigned a tag with a specific key, regardless of the tag value.

- `type` \- The type of the secondary network ( `rdma`).

Type: Array of [Filter](api-filter.md) objects

Required: No

**MaxResults**

The maximum number of results to return with a single call. To retrieve the remaining results, make another call with the returned `nextToken` value.

Type: Integer

Valid Range: Minimum value of 5. Maximum value of 1000.

Required: No

**NextToken**

The token for the next page of results.

Type: String

Required: No

**SecondaryNetworkId.N**

The IDs of the secondary networks.

Type: Array of strings

Required: No

## Response Elements

The following elements are returned by the service.

**nextToken**

The token to use to retrieve the next page of results. This value is `null` when there are no more results to return.

Type: String

**requestId**

The ID of the request.

Type: String

**secondaryNetworkSet**

Information about the secondary networks.

Type: Array of [SecondaryNetwork](api-secondarynetwork.md) objects

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example 1

This example describes the specified secondary networks.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeSecondaryNetworks
   &SecondaryNetworkId.1=sn-0123456789abcdef0
   &SecondaryNetworkId.2=sn-0987654321fedcba0
   &AUTHPARAMS
```

#### Sample Response

```

<DescribeSecondaryNetworksResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
   <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
   <secondaryNetworkSet>
      <item>
         <secondaryNetworkId>sn-0123456789abcdef0</secondaryNetworkId>
         <secondaryNetworkArn>arn:aws:ec2:us-east-2:123456789012:secondary-network/sn-0123456789abcdef0</secondaryNetworkArn>
         <ownerId>123456789012</ownerId>
         <type>rdma</type>
         <state>create-complete</state>
         <ipv4CidrBlockAssociations>
            <item>
               <associationId>sn-cidr-assoc-56789901234abcdef0</associationId>
               <cidrBlock>10.0.0.0/16</cidrBlock>
               <state>associated</state>
            </item>
         </ipv4CidrBlockAssociations>
         <tagSet>
            <item>
               <key>Name</key>
               <value>Prod Secondary Network</value>
            </item>
         </tagSet>
      </item>
      <item>
         <secondaryNetworkId>sn-0987654321fedcba0</secondaryNetworkId>
         <secondaryNetworkArn>arn:aws:ec2:us-east-2:123456789012:secondary-network/sn-0987654321fedcba0</secondaryNetworkArn>
         <ownerId>123456789012</ownerId>
         <type>rdma</type>
         <state>create-complete</state>
         <ipv4CidrBlockAssociations>
            <item>
               <associationId>sn-cidr-assoc-09876543210fedcba0</associationId>
               <cidrBlock>10.1.0.0/16</cidrBlock>
               <state>associated</state>
            </item>
         </ipv4CidrBlockAssociations>
         <tagSet>
            <item>
               <key>Name</key>
               <value>Dev Secondary Network</value>
            </item>
         </tagSet>
      </item>
   </secondaryNetworkSet>
</DescribeSecondaryNetworksResponse>
```

### Example 2

This example uses filters to describe any secondary network you own that has a tag with the key Environment and the value Production and whose state is create-complete.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeSecondaryNetworks
   &Filter.1.Name=tag:Environment
   &Filter.1.Value.1=Production
   &Filter.2.Name=state
   &Filter.2.Value.1=create-complete
   &AUTHPARAMS
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DescribeSecondaryNetworks)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DescribeSecondaryNetworks)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describesecondarynetworks.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describesecondarynetworks.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describesecondarynetworks.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describesecondarynetworks.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describesecondarynetworks.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describesecondarynetworks.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DescribeSecondaryNetworks)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describesecondarynetworks.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeSecondaryInterfaces

DescribeSecondarySubnets
