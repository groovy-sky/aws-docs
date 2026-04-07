# ListFirewallRules

Retrieves the firewall rules that you have defined for the specified firewall rule group. DNS Firewall uses the rules in a rule group to filter DNS network traffic for a VPC.

A single call might return only a partial list of the rules. For information, see `MaxResults`.

## Request Syntax

```nohighlight

{
   "Action": "string",
   "FirewallRuleGroupId": "string",
   "MaxResults": number,
   "NextToken": "string",
   "Priority": number
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[Action](#API_route53resolver_ListFirewallRules_RequestSyntax)**

Optional additional filter for the rules to retrieve.

The action that DNS Firewall should take on a DNS query when it matches one of the domains in the rule's domain list, or a threat in a DNS Firewall Advanced rule:

- `ALLOW` \- Permit the request to go through. Not availabe for DNS Firewall Advanced rules.

- `ALERT` \- Permit the request to go through but send an alert to the logs.

- `BLOCK` \- Disallow the request. If this is specified, additional handling details are provided in the rule's `BlockResponse` setting.

Type: String

Valid Values: `ALLOW | BLOCK | ALERT`

Required: No

**[FirewallRuleGroupId](#API_route53resolver_ListFirewallRules_RequestSyntax)**

The unique identifier of the firewall rule group that you want to retrieve the rules for.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: Yes

**[MaxResults](#API_route53resolver_ListFirewallRules_RequestSyntax)**

The maximum number of objects that you want Resolver to return for this request. If more
objects are available, in the response, Resolver provides a
`NextToken` value that you can use in a subsequent call to get the next batch of objects.

If you don't specify a value for `MaxResults`, Resolver returns up to 100 objects.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 100.

Required: No

**[NextToken](#API_route53resolver_ListFirewallRules_RequestSyntax)**

For the first call to this list request, omit this value.

When you request a list of objects, Resolver returns at most the number of objects
specified in `MaxResults`. If more objects are available for retrieval,
Resolver returns a `NextToken` value in the response. To retrieve the next
batch of objects, use the token that was returned for the prior request in your next request.

Type: String

Required: No

**[Priority](#API_route53resolver_ListFirewallRules_RequestSyntax)**

Optional additional filter for the rules to retrieve.

The setting that determines the processing order of the rules in a rule group. DNS Firewall
processes the rules in a rule group by order of priority, starting from the lowest setting.

Type: Integer

Required: No

## Response Syntax

```nohighlight

{
   "FirewallRules": [
      {
         "Action": "string",
         "BlockOverrideDnsType": "string",
         "BlockOverrideDomain": "string",
         "BlockOverrideTtl": number,
         "BlockResponse": "string",
         "ConfidenceThreshold": "string",
         "CreationTime": "string",
         "CreatorRequestId": "string",
         "DnsThreatProtection": "string",
         "FirewallDomainListId": "string",
         "FirewallDomainRedirectionAction": "string",
         "FirewallRuleGroupId": "string",
         "FirewallThreatProtectionId": "string",
         "ModificationTime": "string",
         "Name": "string",
         "Priority": number,
         "Qtype": "string"
      }
   ],
   "NextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[FirewallRules](#API_route53resolver_ListFirewallRules_ResponseSyntax)**

A list of the rules that you have defined.

This might be a partial list of the firewall rules that you've defined. For information,
see `MaxResults`.

Type: Array of [FirewallRule](https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_FirewallRule.html) objects

**[NextToken](#API_route53resolver_ListFirewallRules_ResponseSyntax)**

If objects are still available for retrieval, Resolver returns this token in the response.
To retrieve the next batch of objects, provide this token in your next request.

Type: String

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/route53resolver-2018-04-01/ListFirewallRules)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/route53resolver-2018-04-01/ListFirewallRules)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/route53resolver-2018-04-01/ListFirewallRules)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/route53resolver-2018-04-01/ListFirewallRules)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/route53resolver-2018-04-01/ListFirewallRules)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/route53resolver-2018-04-01/ListFirewallRules)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/route53resolver-2018-04-01/ListFirewallRules)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/route53resolver-2018-04-01/ListFirewallRules)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/route53resolver-2018-04-01/ListFirewallRules)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/route53resolver-2018-04-01/ListFirewallRules)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ListFirewallRuleGroups

ListOutpostResolvers
