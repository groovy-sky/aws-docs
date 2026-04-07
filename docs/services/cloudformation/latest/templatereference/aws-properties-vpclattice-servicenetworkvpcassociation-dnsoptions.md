This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VpcLattice::ServiceNetworkVpcAssociation DnsOptions

The DNS configuration options.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PrivateDnsPreference" : String,
  "PrivateDnsSpecifiedDomains" : [ String, ... ]
}

```

### YAML

```yaml

  PrivateDnsPreference: String
  PrivateDnsSpecifiedDomains:
    - String

```

## Properties

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

AWS::VpcLattice::ServiceNetworkVpcAssociation

Tag
