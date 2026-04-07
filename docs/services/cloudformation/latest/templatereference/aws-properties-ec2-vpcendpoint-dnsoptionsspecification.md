This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::VPCEndpoint DnsOptionsSpecification

Describes the DNS options for an endpoint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DnsRecordIpType" : String,
  "PrivateDnsOnlyForInboundResolverEndpoint" : String,
  "PrivateDnsPreference" : String,
  "PrivateDnsSpecifiedDomains" : [ String, ... ]
}

```

### YAML

```yaml

  DnsRecordIpType: String
  PrivateDnsOnlyForInboundResolverEndpoint: String
  PrivateDnsPreference: String
  PrivateDnsSpecifiedDomains:
    - String

```

## Properties

`DnsRecordIpType`

The DNS records created for the endpoint.

_Required_: No

_Type_: String

_Allowed values_: `ipv4 | ipv6 | dualstack | service-defined | not-specified`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrivateDnsOnlyForInboundResolverEndpoint`

Indicates whether to enable private DNS only for inbound endpoints. This option is
available only for services that support both gateway and interface endpoints. It routes
traffic that originates from the VPC to the gateway endpoint and traffic that originates
from on-premises to the interface endpoint.

_Required_: No

_Type_: String

_Allowed values_: `OnlyInboundResolver | AllResolvers | NotSpecified`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrivateDnsPreference`

The preference for which private domains have a private hosted zone created for and associated with the specified VPC. Only supported when private DNS is enabled and when the VPC endpoint type is ServiceNetwork or Resource.

_Required_: No

_Type_: String

_Allowed values_: `VERIFIED_DOMAINS_ONLY | ALL_DOMAINS | VERIFIED_DOMAINS_AND_SPECIFIED_DOMAINS | SPECIFIED_DOMAINS_ONLY`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PrivateDnsSpecifiedDomains`

Indicates which of the private domains to create private hosted zones for and associate with the specified VPC. Only supported when private DNS is enabled and the private DNS preference is `VERIFIED_DOMAINS_AND_SPECIFIED_DOMAINS` or `SPECIFIED_DOMAINS_ONLY`.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `255 | 10`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::EC2::VPCEndpoint

Tag
