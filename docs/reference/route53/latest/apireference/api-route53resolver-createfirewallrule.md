# CreateFirewallRule

Creates a single DNS Firewall rule in the specified rule group, using the specified domain list.

## Request Syntax

```nohighlight

{
   "Action": "string",
   "BlockOverrideDnsType": "string",
   "BlockOverrideDomain": "string",
   "BlockOverrideTtl": number,
   "BlockResponse": "string",
   "ConfidenceThreshold": "string",
   "CreatorRequestId": "string",
   "DnsThreatProtection": "string",
   "FirewallDomainListId": "string",
   "FirewallDomainRedirectionAction": "string",
   "FirewallRuleGroupId": "string",
   "Name": "string",
   "Priority": number,
   "Qtype": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[Action](#API_route53resolver_CreateFirewallRule_RequestSyntax)**

The action that DNS Firewall should take on a DNS query when it matches one of the domains in the rule's domain list, or a threat in a DNS Firewall Advanced rule:

- `ALLOW` \- Permit the request to go through. Not available for DNS Firewall Advanced rules.

- `ALERT` \- Permit the request and send metrics and logs to Cloud Watch.

- `BLOCK` \- Disallow the request. This option requires additional details in the rule's `BlockResponse`.

Type: String

Valid Values: `ALLOW | BLOCK | ALERT`

Required: Yes

**[BlockOverrideDnsType](#API_route53resolver_CreateFirewallRule_RequestSyntax)**

The DNS record's type. This determines the format of the record value that you provided in `BlockOverrideDomain`. Used for the rule action `BLOCK` with a `BlockResponse` setting of `OVERRIDE`.

This setting is required if the `BlockResponse` setting is `OVERRIDE`.

Type: String

Valid Values: `CNAME`

Required: No

**[BlockOverrideDomain](#API_route53resolver_CreateFirewallRule_RequestSyntax)**

The custom DNS record to send back in response to the query. Used for the rule action `BLOCK` with a `BlockResponse` setting of `OVERRIDE`.

This setting is required if the `BlockResponse` setting is `OVERRIDE`.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: No

**[BlockOverrideTtl](#API_route53resolver_CreateFirewallRule_RequestSyntax)**

The recommended amount of time, in seconds, for the DNS resolver or web browser to cache the provided override record. Used for the rule action `BLOCK` with a `BlockResponse` setting of `OVERRIDE`.

This setting is required if the `BlockResponse` setting is `OVERRIDE`.

Type: Integer

Valid Range: Minimum value of 0. Maximum value of 604800.

Required: No

**[BlockResponse](#API_route53resolver_CreateFirewallRule_RequestSyntax)**

The way that you want DNS Firewall to block the request, used with the rule action
setting `BLOCK`.

- `NODATA` \- Respond indicating that the query was successful, but no response is available for it.

- `NXDOMAIN` \- Respond indicating that the domain name that's in the query doesn't exist.

- `OVERRIDE` \- Provide a custom override in the response. This option requires custom handling details in the rule's `BlockOverride*` settings.

This setting is required if the rule action setting is `BLOCK`.

Type: String

Valid Values: `NODATA | NXDOMAIN | OVERRIDE`

Required: No

**[ConfidenceThreshold](#API_route53resolver_CreateFirewallRule_RequestSyntax)**

The confidence threshold for DNS Firewall Advanced. You must provide this value when you create a DNS Firewall Advanced rule. The confidence
level values mean:

- `LOW`: Provides the highest detection rate for threats, but also increases false positives.

- `MEDIUM`: Provides a balance between detecting threats and false positives.

- `HIGH`: Detects only the most well corroborated threats with a low rate of false positives.

Type: String

Valid Values: `LOW | MEDIUM | HIGH`

Required: No

**[CreatorRequestId](#API_route53resolver_CreateFirewallRule_RequestSyntax)**

A unique string that identifies the request and that allows you to retry failed requests
without the risk of running the operation twice. `CreatorRequestId` can be
any unique string, for example, a date/time stamp.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: Yes

**[DnsThreatProtection](#API_route53resolver_CreateFirewallRule_RequestSyntax)**

Use to create a DNS Firewall Advanced rule.

Type: String

Valid Values: `DGA | DNS_TUNNELING`

Required: No

**[FirewallDomainListId](#API_route53resolver_CreateFirewallRule_RequestSyntax)**

The ID of the domain list that you want to use in the rule. Can't be used together with `DnsThreatProtecton`.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: No

**[FirewallDomainRedirectionAction](#API_route53resolver_CreateFirewallRule_RequestSyntax)**

How you want the the rule to evaluate DNS redirection in the DNS redirection chain, such as CNAME or DNAME.

`INSPECT_REDIRECTION_DOMAIN`: (Default) inspects all domains in the redirection chain. The individual domains in the redirection chain must be
added to the domain list.

`TRUST_REDIRECTION_DOMAIN`: Inspects only the first domain in the redirection chain. You don't need to add the subsequent domains in the domain in the redirection list to
the domain list.

Type: String

Valid Values: `INSPECT_REDIRECTION_DOMAIN | TRUST_REDIRECTION_DOMAIN`

Required: No

**[FirewallRuleGroupId](#API_route53resolver_CreateFirewallRule_RequestSyntax)**

The unique identifier of the firewall rule group where you want to create the rule.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: Yes

**[Name](#API_route53resolver_CreateFirewallRule_RequestSyntax)**

A name that lets you identify the rule in the rule group.

Type: String

Length Constraints: Maximum length of 64.

Pattern: `(?!^[0-9]+$)([a-zA-Z0-9\-_' ']+)`

Required: Yes

**[Priority](#API_route53resolver_CreateFirewallRule_RequestSyntax)**

The setting that determines the processing order of the rule in the rule group. DNS Firewall
processes the rules in a rule group by order of priority, starting from the lowest setting.

You must specify a unique priority for each rule in a rule group.
To make it easier to insert rules later, leave space between the numbers, for example, use 100, 200, and so on. You
can change the priority setting for the rules in a rule group at any time.

Type: Integer

Required: Yes

**[Qtype](#API_route53resolver_CreateFirewallRule_RequestSyntax)**

The DNS query type you want the rule to evaluate. Allowed values are;

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

**[FirewallRule](#API_route53resolver_CreateFirewallRule_ResponseSyntax)**

The firewall rule that you just created.

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/route53resolver-2018-04-01/CreateFirewallRule)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/route53resolver-2018-04-01/CreateFirewallRule)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/route53resolver-2018-04-01/CreateFirewallRule)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/route53resolver-2018-04-01/CreateFirewallRule)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/route53resolver-2018-04-01/CreateFirewallRule)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/route53resolver-2018-04-01/CreateFirewallRule)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/route53resolver-2018-04-01/CreateFirewallRule)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/route53resolver-2018-04-01/CreateFirewallRule)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/route53resolver-2018-04-01/CreateFirewallRule)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/route53resolver-2018-04-01/CreateFirewallRule)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CreateFirewallDomainList

CreateFirewallRuleGroup
