# CreateIpamExternalResourceVerificationToken

Create a verification token.

A verification token is an AWS-generated random value that you can use to prove ownership of an external resource. For example, you can use a verification token to validate that you control a public IP address range when you bring an IP address range to AWS (BYOIP).

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ClientToken**

A unique, case-sensitive identifier that you provide to ensure the idempotency of the request. For more information, see [Ensuring idempotency](../../../../services/ec2/latest/devguide/ec2-api-idempotency.md).

Type: String

Required: No

**DryRun**

A check for whether you have the required permissions for the action without actually making the request
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**IpamId**

The ID of the IPAM that will create the token.

Type: String

Required: Yes

**TagSpecification.N**

Token tags.

Type: Array of [TagSpecification](api-tagspecification.md) objects

Required: No

## Response Elements

The following elements are returned by the service.

**ipamExternalResourceVerificationToken**

The verification token.

Type: [IpamExternalResourceVerificationToken](api-ipamexternalresourceverificationtoken.md) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/CreateIpamExternalResourceVerificationToken)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/CreateIpamExternalResourceVerificationToken)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/createipamexternalresourceverificationtoken.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/createipamexternalresourceverificationtoken.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/createipamexternalresourceverificationtoken.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/createipamexternalresourceverificationtoken.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/createipamexternalresourceverificationtoken.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/createipamexternalresourceverificationtoken.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/CreateIpamExternalResourceVerificationToken)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/createipamexternalresourceverificationtoken.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateIpam

CreateIpamPolicy
