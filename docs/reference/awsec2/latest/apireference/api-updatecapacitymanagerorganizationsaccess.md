# UpdateCapacityManagerOrganizationsAccess

Updates the Organizations access setting for EC2 Capacity Manager. This controls whether Capacity Manager can aggregate
data from all accounts in your AWS Organization or only from the current account.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ClientToken**

Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request, and provides an error response. If
you have the required permissions, the error response is `DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**OrganizationsAccess**

Specifies whether to enable or disable cross-account access for AWS Organizations. When enabled, Capacity Manager aggregates data from all accounts in your organization.

Type: Boolean

Required: Yes

## Response Elements

The following elements are returned by the service.

**capacityManagerStatus**

The current status of Capacity Manager after the update operation.

Type: String

Valid Values: `enabled | disabled`

**organizationsAccess**

The updated Organizations access setting indicating whether cross-account data aggregation is enabled.

Type: Boolean

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/UpdateCapacityManagerOrganizationsAccess)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/UpdateCapacityManagerOrganizationsAccess)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/updatecapacitymanagerorganizationsaccess.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/updatecapacitymanagerorganizationsaccess.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/updatecapacitymanagerorganizationsaccess.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/updatecapacitymanagerorganizationsaccess.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/updatecapacitymanagerorganizationsaccess.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/updatecapacitymanagerorganizationsaccess.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/UpdateCapacityManagerOrganizationsAccess)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/updatecapacitymanagerorganizationsaccess.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

UnmonitorInstances

UpdateInterruptibleCapacityReservationAllocation
