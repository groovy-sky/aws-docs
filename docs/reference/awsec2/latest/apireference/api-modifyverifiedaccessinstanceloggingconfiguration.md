# ModifyVerifiedAccessInstanceLoggingConfiguration

Modifies the logging configuration for the specified AWS Verified Access instance.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**AccessLogs**

The configuration options for Verified Access instances.

Type: [VerifiedAccessLogOptions](api-verifiedaccesslogoptions.md) object

Required: Yes

**ClientToken**

A unique, case-sensitive token that you provide to ensure idempotency of your
modification request. For more information, see [Ensuring idempotency](../../../../services/ec2/latest/devguide/ec2-api-idempotency.md).

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**VerifiedAccessInstanceId**

The ID of the Verified Access instance.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**loggingConfiguration**

The logging configuration for the Verified Access instance.

Type: [VerifiedAccessInstanceLoggingConfiguration](api-verifiedaccessinstanceloggingconfiguration.md) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/ModifyVerifiedAccessInstanceLoggingConfiguration)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/ModifyVerifiedAccessInstanceLoggingConfiguration)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/modifyverifiedaccessinstanceloggingconfiguration.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/modifyverifiedaccessinstanceloggingconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/modifyverifiedaccessinstanceloggingconfiguration.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/modifyverifiedaccessinstanceloggingconfiguration.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/modifyverifiedaccessinstanceloggingconfiguration.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/modifyverifiedaccessinstanceloggingconfiguration.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/ModifyVerifiedAccessInstanceLoggingConfiguration)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/modifyverifiedaccessinstanceloggingconfiguration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ModifyVerifiedAccessInstance

ModifyVerifiedAccessTrustProvider
