# DisableCapacityManager

Disables EC2 Capacity Manager for your account. This stops data ingestion and removes access to capacity analytics and optimization recommendations.
Previously collected data is retained but no new data will be processed.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ClientToken**

Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.
If you have the required permissions, the error response is `DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

## Response Elements

The following elements are returned by the service.

**capacityManagerStatus**

The current status of Capacity Manager after the disable operation.

Type: String

Valid Values: `enabled | disabled`

**organizationsAccess**

Indicates whether Organizations access is enabled. This will be `false` after disabling Capacity Manager.

Type: Boolean

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/disablecapacitymanager.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/disablecapacitymanager.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/disablecapacitymanager.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/disablecapacitymanager.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/disablecapacitymanager.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/disablecapacitymanager.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/disablecapacitymanager.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/disablecapacitymanager.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/disablecapacitymanager.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/disablecapacitymanager.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DisableAwsNetworkPerformanceMetricSubscription

DisableEbsEncryptionByDefault
