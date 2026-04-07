# ReplaceNetworkAclEntry

Replaces an entry (rule) in a network ACL. For more information, see [Network ACLs](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-network-acls.html) in the
_Amazon VPC User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**CidrBlock**

The IPv4 network range to allow or deny, in CIDR notation (for example
`172.16.0.0/24`).

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Egress**

Indicates whether to replace the egress rule.

Default: If no value is specified, we replace the ingress rule.

Type: Boolean

Required: Yes

**Icmp**

ICMP protocol: The ICMP or ICMPv6 type and code. Required if specifying protocol
1 (ICMP) or protocol 58 (ICMPv6) with an IPv6 CIDR block.

Type: [IcmpTypeCode](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_IcmpTypeCode.html) object

Required: No

**Ipv6CidrBlock**

The IPv6 network range to allow or deny, in CIDR notation (for example
`2001:bd8:1234:1a00::/64`).

Type: String

Required: No

**NetworkAclId**

The ID of the ACL.

Type: String

Required: Yes

**PortRange**

TCP or UDP protocols: The range of ports the rule applies to.
Required if specifying protocol 6 (TCP) or 17 (UDP).

Type: [PortRange](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_PortRange.html) object

Required: No

**Protocol**

The protocol number. A value of "-1" means all protocols. If you specify "-1" or a
protocol number other than "6" (TCP), "17" (UDP), or "1" (ICMP), traffic on all ports is
allowed, regardless of any ports or ICMP types or codes that you specify. If you specify
protocol "58" (ICMPv6) and specify an IPv4 CIDR block, traffic for all ICMP types and
codes allowed, regardless of any that you specify. If you specify protocol "58" (ICMPv6)
and specify an IPv6 CIDR block, you must specify an ICMP type and code.

Type: String

Required: Yes

**RuleAction**

Indicates whether to allow or deny the traffic that matches the rule.

Type: String

Valid Values: `allow | deny`

Required: Yes

**RuleNumber**

The rule number of the entry to replace.

Type: Integer

Required: Yes

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**return**

Is `true` if the request succeeds, and an error otherwise.

Type: Boolean

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example replaces the egress entry numbered `110` in the specified network ACL.
The new rule denies egress traffic destined for any IPv4 address ( `0.0.0.0/0`) on TCP port 139.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=ReplaceNetworkAclEntry
&NetworkAclId=acl-2cb85d45
&RuleNumber=110
&Protocol="6"
&RuleAction=deny
&Egress=true
&CidrBlock=0.0.0.0/0
&PortRange.From=139
&PortRange.To=139
&AUTHPARAMS
```

#### Sample Response

```

<ReplaceNetworkAclEntryResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
   <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
   <return>true</return>
</ReplaceNetworkAclEntryResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/ReplaceNetworkAclEntry)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/ReplaceNetworkAclEntry)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/ReplaceNetworkAclEntry)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/ReplaceNetworkAclEntry)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/ReplaceNetworkAclEntry)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/ReplaceNetworkAclEntry)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/ReplaceNetworkAclEntry)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/ReplaceNetworkAclEntry)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/ReplaceNetworkAclEntry)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/ReplaceNetworkAclEntry)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ReplaceNetworkAclAssociation

ReplaceRoute
