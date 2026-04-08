# CreateIpam

Create an IPAM. Amazon VPC IP Address Manager (IPAM) is a VPC feature that you can use
to automate your IP address management workflows including assigning, tracking,
troubleshooting, and auditing IP addresses across AWS Regions and accounts
throughout your AWS Organization.

For more information, see [Create an IPAM](../../../../services/vpc/latest/ipam/create-ipam.md) in the _Amazon VPC IPAM User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ClientToken**

A unique, case-sensitive identifier that you provide to ensure the idempotency of the request. For more information, see [Ensuring idempotency](../../../../services/ec2/latest/devguide/ec2-api-idempotency.md).

Type: String

Required: No

**Description**

A description for the IPAM.

Type: String

Required: No

**DryRun**

A check for whether you have the required permissions for the action without actually making the request
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**EnablePrivateGua**

Enable this option to use your own GUA ranges as private IPv6 addresses. This option is disabled by default.

Type: Boolean

Required: No

**MeteredAccount**

A metered account is an AWS account that is charged for active IP addresses managed in IPAM. For more information, see [Enable cost distribution](../../../../services/vpc/latest/ipam/ipam-enable-cost-distro.md) in the _Amazon VPC IPAM User Guide_.

Possible values:

- `ipam-owner` (default): The AWS account which owns the IPAM is charged for all active IP addresses managed in IPAM.

- `resource-owner`: The AWS account that owns the IP address is charged for the active IP address.

Type: String

Valid Values: `ipam-owner | resource-owner`

Required: No

**OperatingRegion.N**

The operating Regions for the IPAM. Operating Regions are AWS Regions where the IPAM is allowed to manage IP address CIDRs. IPAM only discovers and monitors resources in the AWS Regions you select as operating Regions.

For more information about operating Regions, see [Create an IPAM](../../../../services/vpc/latest/ipam/create-ipam.md) in the _Amazon VPC IPAM User Guide_.

Type: Array of [AddIpamOperatingRegion](api-addipamoperatingregion.md) objects

Array Members: Minimum number of 0 items. Maximum number of 50 items.

Required: No

**TagSpecification.N**

The key/value combination of a tag assigned to the resource. Use the tag key in the filter name and the tag value as the filter value.
For example, to find all resources that have a tag with the key `Owner` and the value `TeamA`, specify `tag:Owner` for the filter name and `TeamA` for the filter value.

Type: Array of [TagSpecification](api-tagspecification.md) objects

Required: No

**Tier**

IPAM is offered in a Free Tier and an Advanced Tier. For more information about the features available in each tier and the costs associated with the tiers, see [Amazon VPC pricing > IPAM tab](http://aws.amazon.com/vpc/pricing).

Type: String

Valid Values: `free | advanced`

Required: No

## Response Elements

The following elements are returned by the service.

**ipam**

Information about the IPAM created.

Type: [Ipam](api-ipam.md) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/createipam.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/createipam.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/createipam.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/createipam.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/createipam.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/createipam.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/createipam.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/createipam.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/createipam.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/createipam.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateInterruptibleCapacityReservationAllocation

CreateIpamExternalResourceVerificationToken
