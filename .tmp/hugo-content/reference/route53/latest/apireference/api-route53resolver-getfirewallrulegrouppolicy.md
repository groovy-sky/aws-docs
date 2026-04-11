# GetFirewallRuleGroupPolicy

Returns the AWS Identity and Access Management (AWS IAM) policy for sharing the
specified rule group. You can use the policy to share the rule group using AWS Resource Access Manager (AWS RAM).

## Request Syntax

```nohighlight

{
   "Arn": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[Arn](#API_route53resolver_GetFirewallRuleGroupPolicy_RequestSyntax)**

The ARN (Amazon Resource Name) for the rule group.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: Yes

## Response Syntax

```nohighlight

{
   "FirewallRuleGroupPolicy": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[FirewallRuleGroupPolicy](#API_route53resolver_GetFirewallRuleGroupPolicy_ResponseSyntax)**

The AWS Identity and Access Management (AWS IAM) policy for sharing the specified rule
group. You can use the policy to share the rule group using AWS Resource Access Manager
(AWS RAM).

Type: String

Length Constraints: Maximum length of 30000.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

The current account doesn't have the IAM permissions required to perform the specified Resolver operation.

This error can also be thrown when a customer has reached the 5120 character limit for a
resource policy for CloudWatch Logs.

HTTP Status Code: 400

**InternalServiceErrorException**

We encountered an unknown error. Try again in a few minutes.

HTTP Status Code: 400

**ResourceNotFoundException**

The specified resource doesn't exist.

**ResourceType**

For a `ResourceNotFoundException` error, the type of resource that doesn't exist.

HTTP Status Code: 400

**ThrottlingException**

The request was throttled. Try again in a few minutes.

HTTP Status Code: 400

**ValidationException**

You have provided an invalid command. If you ran the `UpdateFirewallDomains` request. supported values are `ADD`,
`REMOVE`, or `REPLACE` a domain.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53resolver-2018-04-01/getfirewallrulegrouppolicy.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53resolver-2018-04-01/getfirewallrulegrouppolicy.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53resolver-2018-04-01/getfirewallrulegrouppolicy.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53resolver-2018-04-01/getfirewallrulegrouppolicy.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53resolver-2018-04-01/getfirewallrulegrouppolicy.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53resolver-2018-04-01/getfirewallrulegrouppolicy.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53resolver-2018-04-01/getfirewallrulegrouppolicy.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53resolver-2018-04-01/getfirewallrulegrouppolicy.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53resolver-2018-04-01/getfirewallrulegrouppolicy.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53resolver-2018-04-01/getfirewallrulegrouppolicy.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetFirewallRuleGroupAssociation

GetOutpostResolver
