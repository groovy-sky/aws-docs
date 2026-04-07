# EnableIpamPolicy

Enables an IPAM policy.

An IPAM policy is a set of rules that define how public IPv4 addresses from IPAM pools are allocated to AWS resources. Each rule maps an AWS service to IPAM pools that the service will use to get IP addresses. A single policy can have multiple rules and be applied to multiple AWS Regions. If the IPAM pool run out of addresses then the services fallback to Amazon-provided IP addresses. A policy can be applied to an individual AWS account or an entity within AWS Organizations.

For more information, see [Define public IPv4 allocation strategy with IPAM policies](../../../../services/vpc/latest/ipam/define-public-ipv4-allocation-strategy-with-ipam-policies.md) in the _Amazon VPC IPAM User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

A check for whether you have the required permissions for the action without actually making the request
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**IpamPolicyId**

The ID of the IPAM policy to enable.

Type: String

Required: Yes

**OrganizationTargetId**

The ID of the AWS Organizations target for which to enable the IPAM policy. This parameter is required only when IPAM is integrated with AWS Organizations. When IPAM is not integrated with AWS Organizations, omit this parameter and the policy will apply to the current account.

A target can be an individual AWS account or an entity within an AWS Organization to which an IPAM policy can be applied.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**ipamPolicyId**

The ID of the IPAM policy that was enabled.

Type: String

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/EnableIpamPolicy)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/EnableIpamPolicy)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/enableipampolicy.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/enableipampolicy.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/enableipampolicy.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/enableipampolicy.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/enableipampolicy.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/enableipampolicy.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/EnableIpamPolicy)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/enableipampolicy.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

EnableIpamOrganizationAdminAccount

EnableReachabilityAnalyzerOrganizationSharing
