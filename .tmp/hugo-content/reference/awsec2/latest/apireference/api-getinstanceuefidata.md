# GetInstanceUefiData

A binary representation of the UEFI variable store. Only non-volatile variables are
stored. This is a base64 encoded and zlib compressed binary value that must be properly
encoded.

When you use [register-image](../../../../services/cli/latest/reference/ec2/register-image.md) to create
an AMI, you can create an exact copy of your variable store by passing the UEFI data in
the `UefiData` parameter. You can modify the UEFI data by using the [python-uefivars tool](https://github.com/awslabs/python-uefivars) on
GitHub. You can use the tool to convert the UEFI data into a human-readable format
(JSON), which you can inspect and modify, and then convert back into the binary format
to use with register-image.

For more information, see [UEFI Secure Boot](../../../../services/ec2/latest/userguide/uefi-secure-boot.md) in the
_Amazon EC2 User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the operation, without actually making the
request, and provides an error response. If you have the required permissions, the error response is
`DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**InstanceId**

The ID of the instance from which to retrieve the UEFI data.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**instanceId**

The ID of the instance from which to retrieve the UEFI data.

Type: String

**requestId**

The ID of the request.

Type: String

**uefiData**

Base64 representation of the non-volatile UEFI variable store.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/getinstanceuefidata.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/getinstanceuefidata.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/getinstanceuefidata.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/getinstanceuefidata.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/getinstanceuefidata.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/getinstanceuefidata.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/getinstanceuefidata.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/getinstanceuefidata.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/getinstanceuefidata.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/getinstanceuefidata.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetInstanceTypesFromInstanceRequirements

GetIpamAddressHistory
