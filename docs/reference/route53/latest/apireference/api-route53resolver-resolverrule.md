# ResolverRule

For queries that originate in your VPC, detailed information about a Resolver rule, which specifies how to route DNS queries
out of the VPC. The `ResolverRule` parameter appears in the response to a
[CreateResolverRule](https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_CreateResolverRule.html),
[DeleteResolverRule](https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_DeleteResolverRule.html),
[GetResolverRule](https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_GetResolverRule.html),
[ListResolverRules](https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_ListResolverRules.html),
or
[UpdateResolverRule](https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_UpdateResolverRule.html) request.

## Contents

**Arn**

The ARN (Amazon Resource Name) for the Resolver rule specified by `Id`.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: No

**CreationTime**

The date and time that the Resolver rule was created, in Unix time format and Coordinated Universal Time (UTC).

Type: String

Length Constraints: Minimum length of 20. Maximum length of 40.

Required: No

**CreatorRequestId**

A unique string that you specified when you created the Resolver rule.
`CreatorRequestId` identifies the request and allows failed requests to
be retried without the risk of running the operation twice.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: No

**DelegationRecord**

DNS queries with delegation records that point to this domain name are forwarded to resolvers on your network.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: No

**DomainName**

DNS queries for this domain name are forwarded to the IP addresses that are specified in `TargetIps`. If a query matches
multiple Resolver rules (example.com and www.example.com), the query is routed using the Resolver rule that contains the most specific domain name
(www.example.com).

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: No

**Id**

The ID that Resolver assigned to the Resolver rule when you created it.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: No

**ModificationTime**

The date and time that the Resolver rule was last updated, in Unix time format and Coordinated Universal Time (UTC).

Type: String

Length Constraints: Minimum length of 20. Maximum length of 40.

Required: No

**Name**

The name for the Resolver rule, which you specified when you created the Resolver rule.

The name can be up to 64 characters long and can contain letters (a-z, A-Z), numbers (0-9), hyphens (-), underscores (\_), and spaces. The name cannot consist of only numbers.

Type: String

Length Constraints: Maximum length of 64.

Pattern: `(?!^[0-9]+$)([a-zA-Z0-9\-_' ']+)`

Required: No

**OwnerId**

When a rule is shared with another AWS account, the account ID of the account that the rule is shared with.

Type: String

Length Constraints: Minimum length of 12. Maximum length of 32.

Required: No

**ResolverEndpointId**

The ID of the endpoint that the rule is associated with.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: No

**RuleType**

When you want to forward DNS queries for specified domain name to resolvers on your network, specify `FORWARD` or `DELEGATE`.
If a query matches multiple Resolver rules (example.com and www.example.com), outbound DNS queries are routed using the Resolver rule that contains
the most specific domain name (www.example.com).

When you have a forwarding rule to forward DNS queries for a domain to your network and you want Resolver to process queries for
a subdomain of that domain, specify `SYSTEM`.

For example, to forward DNS queries for example.com to resolvers on your network, you create a rule and specify `FORWARD`
for `RuleType`. To then have Resolver process queries for apex.example.com, you create a rule and specify
`SYSTEM` for `RuleType`.

Currently, only Resolver can create rules that have a value of `RECURSIVE` for `RuleType`.

Type: String

Valid Values: `FORWARD | SYSTEM | RECURSIVE | DELEGATE`

Required: No

**ShareStatus**

Whether the rule is shared and, if so, whether the current account is sharing the rule with
another account, or another account is sharing the rule with the current account.

Type: String

Valid Values: `NOT_SHARED | SHARED_WITH_ME | SHARED_BY_ME`

Required: No

**Status**

A code that specifies the current status of the Resolver rule.

Type: String

Valid Values: `COMPLETE | DELETING | UPDATING | FAILED`

Required: No

**StatusMessage**

A detailed description of the status of a Resolver rule.

Type: String

Length Constraints: Maximum length of 255.

Required: No

**TargetIps**

An array that contains the IP addresses and ports that an outbound endpoint forwards DNS queries to. Typically,
these are the IP addresses of DNS resolvers on your network.

Type: Array of [TargetAddress](https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_TargetAddress.html) objects

Array Members: Minimum number of 1 item.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/route53resolver-2018-04-01/ResolverRule)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/route53resolver-2018-04-01/ResolverRule)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/route53resolver-2018-04-01/ResolverRule)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ResolverQueryLogConfigAssociation

ResolverRuleAssociation
