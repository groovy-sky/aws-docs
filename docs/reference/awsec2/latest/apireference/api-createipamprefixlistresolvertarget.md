# CreateIpamPrefixListResolverTarget

Creates an IPAM prefix list resolver target.

An IPAM prefix list resolver target is an association between a specific customer-managed prefix list and an IPAM prefix list resolver. The target enables the resolver to synchronize CIDRs selected by its rules into the specified prefix list, which can then be referenced in AWS resources.

For more information about IPAM prefix list resolver, see [Automate prefix list updates with IPAM](../../../../services/vpc/latest/ipam/automate-prefix-list-updates.md) in the _Amazon VPC IPAM User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ClientToken**

A unique, case-sensitive identifier that you provide to ensure the idempotency of the request. For more information, see [Ensuring idempotency](../../../../services/ec2/latest/devguide/ec2-api-idempotency.md).

Type: String

Required: No

**DesiredVersion**

The specific version of the prefix list to target. If not specified, the resolver will target the latest version.

Type: Long

Required: No

**DryRun**

A check for whether you have the required permissions for the action without actually making the request
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**IpamPrefixListResolverId**

The ID of the IPAM prefix list resolver that will manage the synchronization of CIDRs to the target prefix list.

Type: String

Required: Yes

**PrefixListId**

The ID of the managed prefix list that will be synchronized with CIDRs selected by the IPAM prefix list resolver. This prefix list becomes an IPAM managed prefix list.

An IPAM-managed prefix list is a customer-managed prefix list that has been associated with an IPAM prefix list resolver target. When a prefix list becomes IPAM managed, its CIDRs are automatically synchronized based on the IPAM prefix list resolver's CIDR selection rules, and direct CIDR modifications are restricted.

Type: String

Required: Yes

**PrefixListRegion**

The AWS Region where the prefix list is located. This is required when referencing a prefix list in a different Region.

Type: String

Required: Yes

**TagSpecification.N**

The tags to apply to the IPAM prefix list resolver target during creation. Tags help you organize and manage your AWS resources.

Type: Array of [TagSpecification](api-tagspecification.md) objects

Required: No

**TrackLatestVersion**

Indicates whether the resolver target should automatically track the latest version of the prefix list. When enabled, the target will always synchronize with the most current version of the prefix list.

Choose this for automatic updates when you want your prefix lists to stay current with infrastructure changes without manual intervention.

Type: Boolean

Required: Yes

## Response Elements

The following elements are returned by the service.

**ipamPrefixListResolverTarget**

Information about the IPAM prefix list resolver target that was created.

Type: [IpamPrefixListResolverTarget](api-ipamprefixlistresolvertarget.md) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/CreateIpamPrefixListResolverTarget)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/CreateIpamPrefixListResolverTarget)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/createipamprefixlistresolvertarget.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/createipamprefixlistresolvertarget.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/createipamprefixlistresolvertarget.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/createipamprefixlistresolvertarget.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/createipamprefixlistresolvertarget.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/createipamprefixlistresolvertarget.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/CreateIpamPrefixListResolverTarget)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/createipamprefixlistresolvertarget.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateIpamPrefixListResolver

CreateIpamResourceDiscovery
