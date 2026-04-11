# GetInstanceMetadataDefaults

Gets the default instance metadata service (IMDS) settings that are set at the account
level in the specified AWS  Region.

For more information, see [Order of precedence for instance metadata options](../../../../services/ec2/latest/userguide/configuring-instance-metadata-options.md#instance-metadata-options-order-of-precedence) in the
_Amazon EC2 User Guide_.

## Request Parameters

For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the operation, without actually making the
request, and provides an error response. If you have the required permissions, the error response is
`DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

## Response Elements

The following elements are returned by the service.

**accountLevel**

The account-level default IMDS settings.

Type: [InstanceMetadataDefaultsResponse](api-instancemetadatadefaultsresponse.md) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/getinstancemetadatadefaults.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/getinstancemetadatadefaults.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/getinstancemetadatadefaults.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/getinstancemetadatadefaults.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/getinstancemetadatadefaults.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/getinstancemetadatadefaults.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/getinstancemetadatadefaults.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/getinstancemetadatadefaults.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/getinstancemetadatadefaults.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/getinstancemetadatadefaults.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetImageBlockPublicAccessState

GetInstanceTpmEkPub
