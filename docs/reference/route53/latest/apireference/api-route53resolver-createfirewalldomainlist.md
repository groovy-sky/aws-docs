# CreateFirewallDomainList

Creates an empty firewall domain list for use in DNS Firewall rules. You can populate the domains for the new list with a file, using [ImportFirewallDomains](api-route53resolver-importfirewalldomains.md), or with domain strings, using [UpdateFirewallDomains](api-route53resolver-updatefirewalldomains.md).

## Request Syntax

```nohighlight

{
   "CreatorRequestId": "string",
   "Name": "string",
   "Tags": [
      {
         "Key": "string",
         "Value": "string"
      }
   ]
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[CreatorRequestId](#API_route53resolver_CreateFirewallDomainList_RequestSyntax)**

A unique string that identifies the request and that allows you to retry failed requests
without the risk of running the operation twice. `CreatorRequestId` can be
any unique string, for example, a date/time stamp.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: Yes

**[Name](#API_route53resolver_CreateFirewallDomainList_RequestSyntax)**

A name that lets you identify the domain list to manage and use it.

Type: String

Length Constraints: Maximum length of 64.

Pattern: `(?!^[0-9]+$)([a-zA-Z0-9\-_' ']+)`

Required: Yes

**[Tags](#API_route53resolver_CreateFirewallDomainList_RequestSyntax)**

A list of the tag keys and values that you want to associate with the domain list.

Type: Array of [Tag](api-route53resolver-tag.md) objects

Array Members: Maximum number of 200 items.

Required: No

## Response Syntax

```nohighlight

{
   "FirewallDomainList": {
      "Arn": "string",
      "CreationTime": "string",
      "CreatorRequestId": "string",
      "DomainCount": number,
      "Id": "string",
      "ManagedOwnerName": "string",
      "ModificationTime": "string",
      "Name": "string",
      "Status": "string",
      "StatusMessage": "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[FirewallDomainList](#API_route53resolver_CreateFirewallDomainList_ResponseSyntax)**

The domain list that you just created.

Type: [FirewallDomainList](api-route53resolver-firewalldomainlist.md) object

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

**LimitExceededException**

The request caused one or more limits to be exceeded.

**ResourceType**

For a `LimitExceededException` error, the type of resource that exceeded the current limit.

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53resolver-2018-04-01/createfirewalldomainlist.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53resolver-2018-04-01/createfirewalldomainlist.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53resolver-2018-04-01/createfirewalldomainlist.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53resolver-2018-04-01/createfirewalldomainlist.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53resolver-2018-04-01/createfirewalldomainlist.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53resolver-2018-04-01/createfirewalldomainlist.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53resolver-2018-04-01/createfirewalldomainlist.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53resolver-2018-04-01/createfirewalldomainlist.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53resolver-2018-04-01/createfirewalldomainlist.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53resolver-2018-04-01/createfirewalldomainlist.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AssociateResolverRule

CreateFirewallRule
