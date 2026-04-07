# GetEnabledIpamPolicy

Gets the enabled IPAM policy.

An IPAM policy is a set of rules that define how public IPv4 addresses from IPAM pools are allocated to AWS resources. Each rule maps an AWS service to IPAM pools that the service will use to get IP addresses. A single policy can have multiple rules and be applied to multiple AWS Regions. If the IPAM pool run out of addresses then the services fallback to Amazon-provided IP addresses. A policy can be applied to an individual AWS account or an entity within AWS Organizations.

## Request Parameters

For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

A check for whether you have the required permissions for the action without actually making the request
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

## Response Elements

The following elements are returned by the service.

**ipamPolicyEnabled**

Indicates whether the IPAM policy is enabled.

Type: Boolean

**ipamPolicyId**

The ID of the enabled IPAM policy.

Type: String

**managedBy**

The entity that manages the IPAM policy.

Type: String

Valid Values: `account | delegated-administrator-for-ipam`

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/GetEnabledIpamPolicy)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/GetEnabledIpamPolicy)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/getenabledipampolicy.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/getenabledipampolicy.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/getenabledipampolicy.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/getenabledipampolicy.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/getenabledipampolicy.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/getenabledipampolicy.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/GetEnabledIpamPolicy)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/getenabledipampolicy.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetEbsEncryptionByDefault

GetFlowLogsIntegrationTemplate
