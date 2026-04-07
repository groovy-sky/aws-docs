# DeleteFirewallRule

Deletes the specified firewall rule.

## Request Syntax

```nohighlight

{
   "FirewallDomainListId": "string",
   "FirewallRuleGroupId": "string",
   "FirewallThreatProtectionId": "string",
   "Qtype": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[FirewallDomainListId](#API_route53resolver_DeleteFirewallRule_RequestSyntax)**

The ID of the domain list that's used in the rule.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: No

**[FirewallRuleGroupId](#API_route53resolver_DeleteFirewallRule_RequestSyntax)**

The unique identifier of the firewall rule group that you want to delete the rule from.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: Yes

**[FirewallThreatProtectionId](#API_route53resolver_DeleteFirewallRule_RequestSyntax)**

The ID that is created for a DNS Firewall Advanced rule.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: No

**[Qtype](#API_route53resolver_DeleteFirewallRule_RequestSyntax)**

The DNS query type that the rule you are deleting evaluates. Allowed values are;

- A: Returns an IPv4 address.

- AAAA: Returns an Ipv6 address.

- CAA: Restricts CAs that can create SSL/TLS certifications for the domain.

- CNAME: Returns another domain name.

- DS: Record that identifies the DNSSEC signing key of a delegated zone.

- MX: Specifies mail servers.

- NAPTR: Regular-expression-based rewriting of domain names.

- NS: Authoritative name servers.

- PTR: Maps an IP address to a domain name.

- SOA: Start of authority record for the zone.

- SPF: Lists the servers authorized to send emails from a domain.

- SRV: Application specific values that identify servers.

- TXT: Verifies email senders and application-specific values.

- A query type you define by using the DNS type ID, for example 28 for AAAA. The values must be
defined as TYPENUMBER, where the
NUMBER can be 1-65334, for
example, TYPE28. For more information, see
[List of DNS record types](https://en.wikipedia.org/wiki/List_of_DNS_record_types).

Type: String

Length Constraints: Minimum length of 1. Maximum length of 16.

Required: No

## Response Syntax

```nohighlight

{
   "FirewallRule": {
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
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[FirewallRule](#API_route53resolver_DeleteFirewallRule_ResponseSyntax)**

The specification for the firewall rule that you just deleted.

Type: [FirewallRule](https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_FirewallRule.html) object

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/route53resolver-2018-04-01/DeleteFirewallRule)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/route53resolver-2018-04-01/DeleteFirewallRule)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/route53resolver-2018-04-01/DeleteFirewallRule)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/route53resolver-2018-04-01/DeleteFirewallRule)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/route53resolver-2018-04-01/DeleteFirewallRule)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/route53resolver-2018-04-01/DeleteFirewallRule)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/route53resolver-2018-04-01/DeleteFirewallRule)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/route53resolver-2018-04-01/DeleteFirewallRule)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/route53resolver-2018-04-01/DeleteFirewallRule)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/route53resolver-2018-04-01/DeleteFirewallRule)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DeleteFirewallDomainList

DeleteFirewallRuleGroup
