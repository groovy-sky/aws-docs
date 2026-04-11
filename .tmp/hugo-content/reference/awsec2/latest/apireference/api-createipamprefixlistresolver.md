# CreateIpamPrefixListResolver

Creates an IPAM prefix list resolver.

An IPAM prefix list resolver is a component that manages the synchronization between IPAM's CIDR selection rules and customer-managed prefix lists. It automates connectivity configurations by selecting CIDRs from IPAM's database based on your business logic and synchronizing them with prefix lists used in resources such as VPC route tables and security groups.

For more information about IPAM prefix list resolver, see [Automate prefix list updates with IPAM](../../../../services/vpc/latest/ipam/automate-prefix-list-updates.md) in the _Amazon VPC IPAM User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**AddressFamily**

The address family for the IPAM prefix list resolver. Valid values are `ipv4` and `ipv6`. You must create separate resolvers for IPv4 and IPv6 CIDRs as they cannot be mixed in the same resolver.

Type: String

Valid Values: `ipv4 | ipv6`

Required: Yes

**ClientToken**

A unique, case-sensitive identifier that you provide to ensure the idempotency of the request. For more information, see [Ensuring idempotency](../../../../services/ec2/latest/devguide/ec2-api-idempotency.md).

Type: String

Required: No

**Description**

A description for the IPAM prefix list resolver to help you identify its purpose and configuration.

Type: String

Required: No

**DryRun**

A check for whether you have the required permissions for the action without actually making the request
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**IpamId**

The ID of the IPAM that will serve as the source of the IP address database for CIDR selection. The IPAM must be in the Advanced tier to use this feature.

Type: String

Required: Yes

**Rule.N**

The CIDR selection rules for the resolver.

CIDR selection rules define the business logic for selecting CIDRs from IPAM. If a CIDR matches any of the rules, it will be included. If a rule has multiple conditions, the CIDR has to match every condition of that rule. You can create a prefix list resolver without any CIDR selection rules, but it will generate empty versions (containing no CIDRs) until you add rules.

Type: Array of [IpamPrefixListResolverRuleRequest](api-ipamprefixlistresolverrulerequest.md) objects

Required: No

**TagSpecification.N**

The tags to apply to the IPAM prefix list resolver during creation. Tags help you organize and manage your AWS resources.

Type: Array of [TagSpecification](api-tagspecification.md) objects

Required: No

## Response Elements

The following elements are returned by the service.

**ipamPrefixListResolver**

Information about the IPAM prefix list resolver that was created.

Type: [IpamPrefixListResolver](api-ipamprefixlistresolver.md) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/createipamprefixlistresolver.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/createipamprefixlistresolver.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/createipamprefixlistresolver.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/createipamprefixlistresolver.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/createipamprefixlistresolver.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/createipamprefixlistresolver.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/createipamprefixlistresolver.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/createipamprefixlistresolver.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/createipamprefixlistresolver.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/createipamprefixlistresolver.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateIpamPool

CreateIpamPrefixListResolverTarget
