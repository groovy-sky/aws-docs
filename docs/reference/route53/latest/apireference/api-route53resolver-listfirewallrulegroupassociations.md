# ListFirewallRuleGroupAssociations

Retrieves the firewall rule group associations that you have defined. Each association enables DNS filtering for a VPC with one rule group.

A single call might return only a partial list of the associations. For information, see `MaxResults`.

## Request Syntax

```nohighlight

{
   "FirewallRuleGroupId": "string",
   "MaxResults": number,
   "NextToken": "string",
   "Priority": number,
   "Status": "string",
   "VpcId": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[FirewallRuleGroupId](#API_route53resolver_ListFirewallRuleGroupAssociations_RequestSyntax)**

The unique identifier of the firewall rule group that you want to retrieve the associations
for. Leave this blank to retrieve associations for any rule group.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: No

**[MaxResults](#API_route53resolver_ListFirewallRuleGroupAssociations_RequestSyntax)**

The maximum number of objects that you want Resolver to return for this request. If more
objects are available, in the response, Resolver provides a
`NextToken` value that you can use in a subsequent call to get the next batch of objects.

If you don't specify a value for `MaxResults`, Resolver returns up to 100 objects.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 100.

Required: No

**[NextToken](#API_route53resolver_ListFirewallRuleGroupAssociations_RequestSyntax)**

For the first call to this list request, omit this value.

When you request a list of objects, Resolver returns at most the number of objects
specified in `MaxResults`. If more objects are available for retrieval,
Resolver returns a `NextToken` value in the response. To retrieve the next
batch of objects, use the token that was returned for the prior request in your next request.

Type: String

Required: No

**[Priority](#API_route53resolver_ListFirewallRuleGroupAssociations_RequestSyntax)**

The setting that determines the processing order of the rule group among the rule
groups that are associated with a single VPC. DNS Firewall filters VPC traffic starting
from the rule group with the lowest numeric priority setting.

Type: Integer

Required: No

**[Status](#API_route53resolver_ListFirewallRuleGroupAssociations_RequestSyntax)**

The association `Status` setting that you want DNS Firewall to filter on for the list. If you don't specify this, then DNS Firewall returns all associations, regardless of status.

Type: String

Valid Values: `COMPLETE | DELETING | UPDATING`

Required: No

**[VpcId](#API_route53resolver_ListFirewallRuleGroupAssociations_RequestSyntax)**

The unique identifier of the VPC that you want to retrieve the associations
for. Leave this blank to retrieve associations for any VPC.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: No

## Response Syntax

```nohighlight

{
   "FirewallRuleGroupAssociations": [
      {
         "Arn": "string",
         "CreationTime": "string",
         "CreatorRequestId": "string",
         "FirewallRuleGroupId": "string",
         "Id": "string",
         "ManagedOwnerName": "string",
         "ModificationTime": "string",
         "MutationProtection": "string",
         "Name": "string",
         "Priority": number,
         "Status": "string",
         "StatusMessage": "string",
         "VpcId": "string"
      }
   ],
   "NextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[FirewallRuleGroupAssociations](#API_route53resolver_ListFirewallRuleGroupAssociations_ResponseSyntax)**

A list of your firewall rule group associations.

This might be a partial list of the associations that you have defined. For information, see `MaxResults`.

Type: Array of [FirewallRuleGroupAssociation](https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_FirewallRuleGroupAssociation.html) objects

**[NextToken](#API_route53resolver_ListFirewallRuleGroupAssociations_ResponseSyntax)**

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

**ThrottlingException**

The request was throttled. Try again in a few minutes.

HTTP Status Code: 400

**ValidationException**

You have provided an invalid command. If you ran the `UpdateFirewallDomains` request. supported values are `ADD`,
`REMOVE`, or `REPLACE` a domain.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/route53resolver-2018-04-01/ListFirewallRuleGroupAssociations)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/route53resolver-2018-04-01/ListFirewallRuleGroupAssociations)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/route53resolver-2018-04-01/ListFirewallRuleGroupAssociations)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/route53resolver-2018-04-01/ListFirewallRuleGroupAssociations)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/route53resolver-2018-04-01/ListFirewallRuleGroupAssociations)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/route53resolver-2018-04-01/ListFirewallRuleGroupAssociations)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/route53resolver-2018-04-01/ListFirewallRuleGroupAssociations)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/route53resolver-2018-04-01/ListFirewallRuleGroupAssociations)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/route53resolver-2018-04-01/ListFirewallRuleGroupAssociations)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/route53resolver-2018-04-01/ListFirewallRuleGroupAssociations)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ListFirewallDomains

ListFirewallRuleGroups
