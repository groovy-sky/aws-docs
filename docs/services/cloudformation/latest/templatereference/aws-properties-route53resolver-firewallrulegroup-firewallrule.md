This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53Resolver::FirewallRuleGroup FirewallRule

A single firewall rule in a rule group.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Action" : String,
  "BlockOverrideDnsType" : String,
  "BlockOverrideDomain" : String,
  "BlockOverrideTtl" : Integer,
  "BlockResponse" : String,
  "ConfidenceThreshold" : String,
  "DnsThreatProtection" : String,
  "FirewallDomainListId" : String,
  "FirewallDomainRedirectionAction" : String,
  "FirewallThreatProtectionId" : String,
  "Priority" : Integer,
  "Qtype" : String
}

```

### YAML

```yaml

  Action: String
  BlockOverrideDnsType: String
  BlockOverrideDomain: String
  BlockOverrideTtl: Integer
  BlockResponse: String
  ConfidenceThreshold: String
  DnsThreatProtection: String
  FirewallDomainListId: String
  FirewallDomainRedirectionAction: String
  FirewallThreatProtectionId: String
  Priority: Integer
  Qtype: String

```

## Properties

`Action`

The action that DNS Firewall should take on a DNS query when it matches one of the domains in the rule's domain list,
or a threat in a DNS Firewall Advvanced rule:

- `ALLOW` \- Permit the request to go through. Not available for DNS Firewall Advanced rules.

- `ALERT` \- Permit the request to go through but send an alert to the logs.

- `BLOCK` \- Disallow the request. If this is specified,then `BlockResponse` must also be specified.

if `BlockResponse` is `OVERRIDE`, then all of the
following `OVERRIDE` attributes must be specified:

- `BlockOverrideDnsType`

- `BlockOverrideDomain`

- `BlockOverrideTtl`

_Required_: Yes

_Type_: String

_Allowed values_: `ALLOW | BLOCK | ALERT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BlockOverrideDnsType`

The DNS record's type. This determines the format of the record value that you provided in `BlockOverrideDomain`. Used for the rule action `BLOCK` with a `BlockResponse` setting of `OVERRIDE`.

_Required_: No

_Type_: String

_Allowed values_: `CNAME`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BlockOverrideDomain`

The custom DNS record to send back in response to the query. Used for the rule action `BLOCK` with a `BlockResponse` setting of `OVERRIDE`.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BlockOverrideTtl`

The recommended amount of time, in seconds, for the DNS resolver or web browser to cache the provided override record. Used for the rule action `BLOCK` with a `BlockResponse` setting of `OVERRIDE`.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `604800`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BlockResponse`

The way that you want DNS Firewall to block the request. Used for the rule action setting `BLOCK`.

- `NODATA` \- Respond indicating that the query was successful, but no response is available for it.

- `NXDOMAIN` \- Respond indicating that the domain name that's in the query doesn't exist.

- `OVERRIDE` \- Provide a custom override in the response. This option requires custom handling details in the rule's `BlockOverride*` settings.

_Required_: No

_Type_: String

_Allowed values_: `NODATA | NXDOMAIN | OVERRIDE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConfidenceThreshold`

The confidence threshold for DNS Firewall Advanced. You must provide this value when you create a DNS Firewall Advanced rule. The confidence
level values mean:

- `LOW`: Provides the highest detection rate for threats, but also increases false positives.

- `MEDIUM`: Provides a balance between detecting threats and false positives.

- `HIGH`: Detects only the most well corroborated threats with a low rate of false positives.

_Required_: No

_Type_: String

_Allowed values_: `LOW | MEDIUM | HIGH`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DnsThreatProtection`

The type of the DNS Firewall Advanced rule. Valid values are:

- `DGA`: Domain generation algorithms detection. DGAs are used by attackers to generate a large number of domains
to to launch malware attacks.

- `DNS_TUNNELING`: DNS tunneling detection. DNS tunneling is used by attackers to exfiltrate data from the client by using the DNS tunnel without
making a network connection to the client.

_Required_: No

_Type_: String

_Allowed values_: `DGA | DNS_TUNNELING | DICTIONARY_DGA`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FirewallDomainListId`

The ID of the domain list that's used in the rule.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FirewallDomainRedirectionAction`

How you want the the rule to evaluate DNS redirection in the DNS redirection chain, such as CNAME, or DNAME.

`Inspect_Redirection_Domain `(Default) inspects all domains in the redirection chain. The individual domains in the redirection chain must be
added to the domain list.

`Trust_Redirection_Domain ` inspects only the first domain in the redirection chain. You don't need to add the subsequent domains in the domain in the redirection list to
the domain list.

_Required_: No

_Type_: String

_Allowed values_: `INSPECT_REDIRECTION_DOMAIN | TRUST_REDIRECTION_DOMAIN`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FirewallThreatProtectionId`

ID of the DNS Firewall Advanced rule.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Priority`

The priority of the rule in the rule group. This value must be unique within the rule group. DNS Firewall processes the rules in a rule group by order of priority, starting from the lowest setting.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Qtype`

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

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `16`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Route53Resolver::FirewallRuleGroup

Tag
