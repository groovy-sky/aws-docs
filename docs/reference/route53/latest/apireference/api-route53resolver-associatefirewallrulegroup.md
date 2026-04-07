# AssociateFirewallRuleGroup

Associates a [FirewallRuleGroup](https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_FirewallRuleGroup.html) with a VPC, to provide DNS filtering for the VPC.

## Request Syntax

```nohighlight

{
   "CreatorRequestId": "string",
   "FirewallRuleGroupId": "string",
   "MutationProtection": "string",
   "Name": "string",
   "Priority": number,
   "Tags": [
      {
         "Key": "string",
         "Value": "string"
      }
   ],
   "VpcId": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[CreatorRequestId](#API_route53resolver_AssociateFirewallRuleGroup_RequestSyntax)**

A unique string that identifies the request and that allows failed requests to be
retried without the risk of running the operation twice. `CreatorRequestId`
can be any unique string, for example, a date/time stamp.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: Yes

**[FirewallRuleGroupId](#API_route53resolver_AssociateFirewallRuleGroup_RequestSyntax)**

The unique identifier of the firewall rule group.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: Yes

**[MutationProtection](#API_route53resolver_AssociateFirewallRuleGroup_RequestSyntax)**

If enabled, this setting disallows modification or removal of the association, to help prevent against accidentally altering DNS firewall protections.
When you create the association, the default setting is `DISABLED`.

Type: String

Valid Values: `ENABLED | DISABLED`

Required: No

**[Name](#API_route53resolver_AssociateFirewallRuleGroup_RequestSyntax)**

A name that lets you identify the association, to manage and use it.

Type: String

Length Constraints: Maximum length of 64.

Pattern: `(?!^[0-9]+$)([a-zA-Z0-9\-_' ']+)`

Required: Yes

**[Priority](#API_route53resolver_AssociateFirewallRuleGroup_RequestSyntax)**

The setting that determines the processing order of the rule group among the rule
groups that you associate with the specified VPC. DNS Firewall filters VPC traffic
starting from the rule group with the lowest numeric priority setting.

You must specify a unique priority for each rule group that you associate with a single VPC.
To make it easier to insert rule groups later, leave space between the numbers, for example, use 101, 200, and so on. You
can change the priority setting for a rule group association after you create it.

The allowed values for `Priority` are between 100 and 9900.

Type: Integer

Required: Yes

**[Tags](#API_route53resolver_AssociateFirewallRuleGroup_RequestSyntax)**

A list of the tag keys and values that you want to associate with the rule group association.

Type: Array of [Tag](api-route53resolver-tag.md) objects

Array Members: Maximum number of 200 items.

Required: No

**[VpcId](#API_route53resolver_AssociateFirewallRuleGroup_RequestSyntax)**

The unique identifier of the VPC that you want to associate with the rule group.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: Yes

## Response Syntax

```nohighlight

{
   "FirewallRuleGroupAssociation": {
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
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[FirewallRuleGroupAssociation](#API_route53resolver_AssociateFirewallRuleGroup_ResponseSyntax)**

The association that you just created. The association has an ID that you can use to
identify it in other requests, like update and delete.

Type: [FirewallRuleGroupAssociation](https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_FirewallRuleGroupAssociation.html) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

The current account doesn't have the IAM permissions required to perform the specified Resolver operation.

This error can also be thrown when a customer has reached the 5120 character limit for a
resource policy for CloudWatch Logs.

HTTP Status Code: 400

**ConflictException**

The requested state transition isn't valid. For example, you can't delete a firewall
domain list if it is in the process of being deleted, or you can't import domains into a
domain list that is in the process of being deleted.

HTTP Status Code: 400

**InternalServiceErrorException**

We encountered an unknown error. Try again in a few minutes.

HTTP Status Code: 400

**LimitExceededException**

The request caused one or more limits to be exceeded.

**ResourceType**

For a `LimitExceededException` error, the type of resource that exceeded the current limit.

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/route53resolver-2018-04-01/AssociateFirewallRuleGroup)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/route53resolver-2018-04-01/AssociateFirewallRuleGroup)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/route53resolver-2018-04-01/AssociateFirewallRuleGroup)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/route53resolver-2018-04-01/AssociateFirewallRuleGroup)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/route53resolver-2018-04-01/AssociateFirewallRuleGroup)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/route53resolver-2018-04-01/AssociateFirewallRuleGroup)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/route53resolver-2018-04-01/AssociateFirewallRuleGroup)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/route53resolver-2018-04-01/AssociateFirewallRuleGroup)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/route53resolver-2018-04-01/AssociateFirewallRuleGroup)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/route53resolver-2018-04-01/AssociateFirewallRuleGroup)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon Route 53 Resolver

AssociateResolverEndpointIpAddress
