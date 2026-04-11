This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VpcLattice::ResourceConfiguration ResourceConfigurationDefinition

Identifies the resource configuration in one of the following ways:

- **Amazon Resource Name (ARN)** \- Supported resource-types
that are provisioned by AWS services, such as RDS databases, can be identified
by their ARN.

- **Domain name** \- Any domain name that is publicly
resolvable.

- **IP address** \- For IPv4 and IPv6, only IP addresses in the
VPC are supported.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ArnResource" : String,
  "DnsResource" : DnsResource,
  "IpResource" : String
}

```

### YAML

```yaml

  ArnResource: String
  DnsResource:
    DnsResource
  IpResource: String

```

## Properties

`ArnResource`

The Amazon Resource Name (ARN) of the resource configuration. For the ARN syntax and format, see [ARN format](../../../iam/latest/userguide/reference-arns.md#arns-syntax) in
the _AWS Identity and Access Management user guide_.

_Required_: No

_Type_: String

_Pattern_: `^arn.*`

_Maximum_: `1224`

_Update requires_: Updates are not supported.

`DnsResource`

The DNS name of the resource configuration.

_Required_: No

_Type_: [DnsResource](aws-properties-vpclattice-resourceconfiguration-dnsresource.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IpResource`

The IP address of the resource configuration.

_Required_: No

_Type_: String

_Minimum_: `4`

_Maximum_: `39`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DnsResource

Tag

All content copied from https://docs.aws.amazon.com/.
