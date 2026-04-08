# ModifyIpamPolicyAllocationRules

Modifies the allocation rules in an IPAM policy.

An IPAM policy is a set of rules that define how public IPv4 addresses from IPAM pools are allocated to AWS resources. Each rule maps an AWS service to IPAM pools that the service will use to get IP addresses. A single policy can have multiple rules and be applied to multiple AWS Regions. If the IPAM pool run out of addresses then the services fallback to Amazon-provided IP addresses. A policy can be applied to an individual AWS account or an entity within AWS Organizations.

Allocation rules are optional configurations within an IPAM policy that map AWS resource types to specific IPAM pools. If no rules are defined, the resource types default to using Amazon-provided IP addresses.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**AllocationRule.N**

The new allocation rules to apply to the IPAM policy.

Allocation rules are optional configurations within an IPAM policy that map AWS resource types to specific IPAM pools. If no rules are defined, the resource types default to using Amazon-provided IP addresses.

Type: Array of [IpamPolicyAllocationRuleRequest](api-ipampolicyallocationrulerequest.md) objects

Required: No

**DryRun**

A check for whether you have the required permissions for the action without actually making the request
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**IpamPolicyId**

The ID of the IPAM policy whose allocation rules you want to modify.

Type: String

Required: Yes

**Locale**

The locale for which to modify the allocation rules.

Type: String

Required: Yes

**ResourceType**

The resource type for which to modify the allocation rules.

The AWS service or resource type that can use IP addresses through IPAM policies. Supported services and resource types include:

- Elastic IP addresses

Type: String

Valid Values: `alb | eip | rds | rnat`

Required: Yes

## Response Elements

The following elements are returned by the service.

**ipamPolicyDocument**

The modified IPAM policy containing the updated allocation rules.

Type: [IpamPolicyDocument](api-ipampolicydocument.md) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/modifyipampolicyallocationrules.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/modifyipampolicyallocationrules.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/modifyipampolicyallocationrules.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/modifyipampolicyallocationrules.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/modifyipampolicyallocationrules.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/modifyipampolicyallocationrules.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/modifyipampolicyallocationrules.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/modifyipampolicyallocationrules.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/modifyipampolicyallocationrules.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/modifyipampolicyallocationrules.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ModifyIpam

ModifyIpamPool
