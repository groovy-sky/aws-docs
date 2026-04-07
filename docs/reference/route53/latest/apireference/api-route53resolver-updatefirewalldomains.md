# UpdateFirewallDomains

Updates the firewall domain list from an array of domain specifications.

## Request Syntax

```nohighlight

{
   "Domains": [ "string" ],
   "FirewallDomainListId": "string",
   "Operation": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[Domains](#API_route53resolver_UpdateFirewallDomains_RequestSyntax)**

A list of domains to use in the update operation.

###### Important

There is a limit of 1000 domains per request.

Each domain specification in your domain list must satisfy the following
requirements:

- It can optionally start with `*` (asterisk).

- With the exception of the optional starting asterisk, it must only contain
the following characters: `A-Z`, `a-z`,
`0-9`, `-` (hyphen).

- It must be from 1-255 characters in length.

Type: Array of strings

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: Yes

**[FirewallDomainListId](#API_route53resolver_UpdateFirewallDomains_RequestSyntax)**

The ID of the domain list whose domains you want to update.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: Yes

**[Operation](#API_route53resolver_UpdateFirewallDomains_RequestSyntax)**

What you want DNS Firewall to do with the domains that you are providing:

- `ADD` \- Add the domains to the ones that are already in the domain list.

- `REMOVE` \- Search the domain list for the domains and remove them from the list.

- `REPLACE` \- Update the domain list to exactly match the list that you are providing.

Type: String

Valid Values: `ADD | REMOVE | REPLACE`

Required: Yes

## Response Syntax

```nohighlight

{
   "Id": "string",
   "Name": "string",
   "Status": "string",
   "StatusMessage": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[Id](#API_route53resolver_UpdateFirewallDomains_ResponseSyntax)**

The ID of the firewall domain list that DNS Firewall just updated.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

**[Name](#API_route53resolver_UpdateFirewallDomains_ResponseSyntax)**

The name of the domain list.

Type: String

Length Constraints: Maximum length of 64.

Pattern: `(?!^[0-9]+$)([a-zA-Z0-9\-_' ']+)`

**[Status](#API_route53resolver_UpdateFirewallDomains_ResponseSyntax)**

Status of the `UpdateFirewallDomains` request.

Type: String

Valid Values: `COMPLETE | COMPLETE_IMPORT_FAILED | IMPORTING | DELETING | UPDATING`

**[StatusMessage](#API_route53resolver_UpdateFirewallDomains_ResponseSyntax)**

Additional information about the status of the list, if available.

Type: String

Length Constraints: Maximum length of 255.

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/route53resolver-2018-04-01/UpdateFirewallDomains)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/route53resolver-2018-04-01/UpdateFirewallDomains)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/route53resolver-2018-04-01/UpdateFirewallDomains)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/route53resolver-2018-04-01/UpdateFirewallDomains)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/route53resolver-2018-04-01/UpdateFirewallDomains)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/route53resolver-2018-04-01/UpdateFirewallDomains)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/route53resolver-2018-04-01/UpdateFirewallDomains)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/route53resolver-2018-04-01/UpdateFirewallDomains)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/route53resolver-2018-04-01/UpdateFirewallDomains)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/route53resolver-2018-04-01/UpdateFirewallDomains)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

UpdateFirewallConfig

UpdateFirewallRule
