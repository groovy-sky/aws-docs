# CreateIpamResourceDiscovery

Creates an IPAM resource discovery. A resource discovery is an IPAM component that enables IPAM to manage and monitor resources that belong to the owning account.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ClientToken**

A client token for the IPAM resource discovery.

Type: String

Required: No

**Description**

A description for the IPAM resource discovery.

Type: String

Required: No

**DryRun**

A check for whether you have the required permissions for the action without actually making the request
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**OperatingRegion.N**

Operating Regions for the IPAM resource discovery. Operating Regions are AWS Regions where the IPAM is allowed to manage IP address CIDRs. IPAM only discovers and monitors resources in the AWS Regions you select as operating Regions.

Type: Array of [AddIpamOperatingRegion](api-addipamoperatingregion.md) objects

Array Members: Minimum number of 0 items. Maximum number of 50 items.

Required: No

**TagSpecification.N**

Tag specifications for the IPAM resource discovery.

Type: Array of [TagSpecification](api-tagspecification.md) objects

Required: No

## Response Elements

The following elements are returned by the service.

**ipamResourceDiscovery**

An IPAM resource discovery.

Type: [IpamResourceDiscovery](api-ipamresourcediscovery.md) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/CreateIpamResourceDiscovery)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/CreateIpamResourceDiscovery)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/createipamresourcediscovery.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/createipamresourcediscovery.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/createipamresourcediscovery.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/createipamresourcediscovery.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/createipamresourcediscovery.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/createipamresourcediscovery.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/CreateIpamResourceDiscovery)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/createipamresourcediscovery.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateIpamPrefixListResolverTarget

CreateIpamScope
