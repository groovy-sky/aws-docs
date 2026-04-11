# DnsOptionsSpecification

Describes the DNS options for an endpoint.

## Contents

**DnsRecordIpType**

The DNS records created for the endpoint.

Type: String

Valid Values: `ipv4 | dualstack | ipv6 | service-defined`

Required: No

**PrivateDnsOnlyForInboundResolverEndpoint**

Indicates whether to enable private DNS only for inbound endpoints. This option is
available only for services that support both gateway and interface endpoints. It routes
traffic that originates from the VPC to the gateway endpoint and traffic that originates
from on-premises to the interface endpoint.

Type: Boolean

Required: No

**PrivateDnsPreference**

The preference for which private domains have a private hosted zone created for and associated with the specified VPC. Only supported when private DNS is enabled and when the VPC endpoint type is ServiceNetwork or Resource.

- `ALL_DOMAINS` \- VPC Lattice provisions private hosted zones for all custom domain names.

- `VERIFIED_DOMAINS_ONLY` \- VPC Lattice provisions a private hosted zone only if custom domain name has been verified by the provider.

- `VERIFIED_DOMAINS_AND_SPECIFIED_DOMAINS` \- VPC Lattice provisions private hosted zones for all verified custom domain names and other domain names that the resource consumer specifies. The resource consumer specifies the domain names in the PrivateDnsSpecifiedDomains parameter.

- `SPECIFIED_DOMAINS_ONLY` \- VPC Lattice provisions a private hosted zone for domain names specified by the resource consumer. The resource consumer specifies the domain names in the PrivateDnsSpecifiedDomains parameter.

Type: String

Required: No

**PrivateDnsSpecifiedDomain.N**

Indicates which of the private domains to create private hosted zones for and associate with the specified VPC. Only supported when private DNS is enabled and the private DNS preference is verified-domains-and-specified-domains or specified-domains-only.

Type: Array of strings

Array Members: Minimum number of 0 items. Maximum number of 10 items.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/dnsoptionsspecification.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/dnsoptionsspecification.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/dnsoptionsspecification.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DnsOptions

DnsServersOptionsModifyStructure
