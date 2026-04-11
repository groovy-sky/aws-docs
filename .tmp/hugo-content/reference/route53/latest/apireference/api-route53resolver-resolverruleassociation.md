# ResolverRuleAssociation

In the response to an
[AssociateResolverRule](api-route53resolver-associateresolverrule.md),
[DisassociateResolverRule](api-route53resolver-disassociateresolverrule.md),
or
[ListResolverRuleAssociations](api-route53resolver-listresolverruleassociations.md)
request, provides information about an association between a Resolver rule and a VPC.
The association determines which DNS queries that originate in the VPC are forwarded to your network.

## Contents

**Id**

The ID of the association between a Resolver rule and a VPC. Resolver assigns this value when you submit an
[AssociateResolverRule](api-route53resolver-associateresolverrule.md)
request.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: No

**Name**

The name of an association between a Resolver rule and a VPC.

The name can be up to 64 characters long and can contain letters (a-z, A-Z), numbers (0-9), hyphens (-), underscores (\_), and spaces. The name cannot consist of only numbers.

Type: String

Length Constraints: Maximum length of 64.

Pattern: `(?!^[0-9]+$)([a-zA-Z0-9\-_' ']+)`

Required: No

**ResolverRuleId**

The ID of the Resolver rule that you associated with the VPC that is specified by `VPCId`.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: No

**Status**

A code that specifies the current status of the association between a Resolver rule and a VPC.

Type: String

Valid Values: `CREATING | COMPLETE | DELETING | FAILED | OVERRIDDEN`

Required: No

**StatusMessage**

A detailed description of the status of the association between a Resolver rule and a VPC.

Type: String

Length Constraints: Maximum length of 255.

Required: No

**VPCId**

The ID of the VPC that you associated the Resolver rule with.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/route53resolver-2018-04-01/resolverruleassociation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53resolver-2018-04-01/resolverruleassociation.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53resolver-2018-04-01/resolverruleassociation.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ResolverRule

ResolverRuleConfig
