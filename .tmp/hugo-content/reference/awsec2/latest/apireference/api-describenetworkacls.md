# DescribeNetworkAcls

Describes your network ACLs. The default is to describe all your network ACLs.
Alternatively, you can specify specific network ACL IDs or filter the results to
include only the network ACLs that match specific criteria.

For more information, see [Network ACLs](../../../../services/vpc/latest/userguide/vpc-network-acls.md) in the
_Amazon VPC User Guide_.

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

- `association.association-id` \- The ID of an association ID for the ACL.

- `association.network-acl-id` \- The ID of the network ACL involved in the association.

- `association.subnet-id` \- The ID of the subnet involved in the association.

- `default` \- Indicates whether the ACL is the default network ACL for the VPC.

- `entry.cidr` \- The IPv4 CIDR range specified in the entry.

- `entry.icmp.code` \- The ICMP code specified in the entry, if any.

- `entry.icmp.type` \- The ICMP type specified in the entry, if any.

- `entry.ipv6-cidr` \- The IPv6 CIDR range specified in the entry.

- `entry.port-range.from` \- The start of the port range specified in the entry.

- `entry.port-range.to` \- The end of the port range specified in the entry.

- `entry.protocol` \- The protocol specified in the entry ( `tcp` \| `udp` \| `icmp` or a protocol number).

- `entry.rule-action` \- Allows or denies the matching traffic ( `allow` \| `deny`).

- `entry.egress` \- A Boolean that indicates the type of rule. Specify `true`
for egress rules, or `false` for ingress rules.

- `entry.rule-number` \- The number of an entry (in other words, rule) in
the set of ACL entries.

- `network-acl-id` \- The ID of the network ACL.

- `owner-id` \- The ID of the AWS account that owns the network ACL.

- `tag` \- The key/value combination of a tag assigned to the resource. Use the tag key in the filter name and the tag value as the filter value.
For example, to find all resources that have a tag with the key `Owner` and the value `TeamA`, specify `tag:Owner` for the filter name and `TeamA` for the filter value.

- `tag-key` \- The key of a tag assigned to the resource. Use this filter to find all resources assigned a tag with a specific key, regardless of the tag value.

- `vpc-id` \- The ID of the VPC for the network ACL.

Type: Array of [Filter](api-filter.md) objects

Required: No

**MaxResults**

The maximum number of items to return for this request.
To get the next page of items, make another request with the token returned in the output.
For more information, see [Pagination](query-requests.md#api-pagination).

Type: Integer

Valid Range: Minimum value of 5. Maximum value of 1000.

Required: No

**NetworkAclId.N**

The IDs of the network ACLs.

Type: Array of strings

Required: No

**NextToken**

The token returned from a previous paginated request. Pagination continues from the end of the items returned by the previous request.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**networkAclSet**

Information about the network ACLs.

Type: Array of [NetworkAcl](api-networkacl.md) objects

**nextToken**

The token to include in another request to get the next page of items. This value is `null` when there are no more items to return.

Type: String

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example describes all your network ACLs.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeNetworkAcls
&AUTHPARAMS
```

#### Sample Response

```

<DescribeNetworkAclsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>be8171a0-2b2a-4a02-8b13-9c3436f2f02d</requestId>
    <networkAclSet>
        <item>
            <networkAclId>acl-0ea1f54ca7EXAMPLE</networkAclId>
            <vpcId>vpc-06e4ab6c6cEXAMPLE</vpcId>
            <ownerId>111122223333</ownerId>
            <default>true</default>
            <entrySet>
                <item>
                    <ruleNumber>100</ruleNumber>
                    <protocol>-1</protocol>
                    <ruleAction>allow</ruleAction>
                    <egress>true</egress>
                    <cidrBlock>0.0.0.0/0</cidrBlock>
                </item>
                <item>
                    <ruleNumber>32767</ruleNumber>
                    <protocol>-1</protocol>
                    <ruleAction>deny</ruleAction>
                    <egress>true</egress>
                    <cidrBlock>0.0.0.0/0</cidrBlock>
                </item>
                <item>
                    <ruleNumber>100</ruleNumber>
                    <protocol>-1</protocol>
                    <ruleAction>allow</ruleAction>
                    <egress>false</egress>
                    <cidrBlock>0.0.0.0/0</cidrBlock>
                </item>
                <item>
                    <ruleNumber>32767</ruleNumber>
                    <protocol>-1</protocol>
                    <ruleAction>deny</ruleAction>
                    <egress>false</egress>
                    <cidrBlock>0.0.0.0/0</cidrBlock>
                </item>
            </entrySet>
            <associationSet>
                <item>
                    <networkAclAssociationId>aclassoc-0c1679dc41EXAMPLE</networkAclAssociationId>
                    <networkAclId>acl-0ea1f54ca7EXAMPLE</networkAclId>
                    <subnetId>subnet-0931fc2fa5EXAMPLE</subnetId>
                </item>
            </associationSet>
            <tagSet/>
        </item>
        <item>
            <networkAclId>acl-09a47ac966EXAMPLE</networkAclId>
            <vpcId>vpc-06b7830650EXAMPLE</vpcId>
            <ownerId>111122223333</ownerId>
            <default>true</default>
            <entrySet>
                <item>
                    <ruleNumber>100</ruleNumber>
                    <protocol>-1</protocol>
                    <ruleAction>allow</ruleAction>
                    <egress>true</egress>
                    <cidrBlock>0.0.0.0/0</cidrBlock>
                </item>
                <item>
                    <ruleNumber>32767</ruleNumber>
                    <protocol>-1</protocol>
                    <ruleAction>deny</ruleAction>
                    <egress>true</egress>
                    <cidrBlock>0.0.0.0/0</cidrBlock>
                </item>
                <item>
                    <ruleNumber>100</ruleNumber>
                    <protocol>-1</protocol>
                    <ruleAction>allow</ruleAction>
                    <egress>false</egress>
                    <cidrBlock>0.0.0.0/0</cidrBlock>
                </item>
                <item>
                    <ruleNumber>32767</ruleNumber>
                    <protocol>-1</protocol>
                    <ruleAction>deny</ruleAction>
                    <egress>false</egress>
                    <cidrBlock>0.0.0.0/0</cidrBlock>
                </item>
            </entrySet>
            <associationSet/>
            <tagSet/>
        </item>
    </networkAclSet>
</DescribeNetworkAclsResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/describenetworkacls.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/describenetworkacls.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describenetworkacls.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describenetworkacls.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describenetworkacls.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describenetworkacls.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describenetworkacls.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describenetworkacls.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/describenetworkacls.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describenetworkacls.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeNatGateways

DescribeNetworkInsightsAccessScopeAnalyses
