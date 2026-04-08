# ModifyIpamPrefixListResolverTarget

Modifies an IPAM prefix list resolver target. You can update version tracking settings and the desired version of the target prefix list.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ClientToken**

A unique, case-sensitive identifier that you provide to ensure the idempotency of the request. For more information, see [Ensuring idempotency](../../../../services/ec2/latest/devguide/ec2-api-idempotency.md).

Type: String

Required: No

**DesiredVersion**

The desired version of the prefix list to target. This allows you to pin the target to a specific version.

Type: Long

Required: No

**DryRun**

A check for whether you have the required permissions for the action without actually making the request
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**IpamPrefixListResolverTargetId**

The ID of the IPAM prefix list resolver target to modify.

Type: String

Required: Yes

**TrackLatestVersion**

Indicates whether the resolver target should automatically track the latest version of the prefix list. When enabled, the target will always synchronize with the most current version.

Choose this for automatic updates when you want your prefix lists to stay current with infrastructure changes without manual intervention.

Type: Boolean

Required: No

## Response Elements

The following elements are returned by the service.

**ipamPrefixListResolverTarget**

Information about the modified IPAM prefix list resolver target.

Type: [IpamPrefixListResolverTarget](api-ipamprefixlistresolvertarget.md) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/modifyipamprefixlistresolvertarget.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/modifyipamprefixlistresolvertarget.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/modifyipamprefixlistresolvertarget.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/modifyipamprefixlistresolvertarget.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/modifyipamprefixlistresolvertarget.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/modifyipamprefixlistresolvertarget.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/modifyipamprefixlistresolvertarget.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/modifyipamprefixlistresolvertarget.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/modifyipamprefixlistresolvertarget.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/modifyipamprefixlistresolvertarget.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ModifyIpamPrefixListResolver

ModifyIpamResourceCidr
