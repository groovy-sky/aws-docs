This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53GlobalResolver::DnsView

Creates a DNS view within a Route 53 Global Resolver. A DNS view models end users, user groups, networks,
and devices, and serves as a parent resource that holds configurations controlling access,
authorization, DNS firewall rules, and forwarding rules.

###### Important

Route 53 Global Resolver is a global service that supports resolvers in multiple AWS Regions but you must specify the
US East (Ohio) Region to create, update, or otherwise work with Route 53 Global Resolver resources. That is, for example,
specify
`--region us-east-2`
on AWS CLI commands.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Route53GlobalResolver::DnsView",
  "Properties" : {
      "ClientToken" : String,
      "Description" : String,
      "DnssecValidation" : String,
      "EdnsClientSubnet" : String,
      "FirewallRulesFailOpen" : String,
      "GlobalResolverId" : String,
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Route53GlobalResolver::DnsView
Properties:
  ClientToken: String
  Description: String
  DnssecValidation: String
  EdnsClientSubnet: String
  FirewallRulesFailOpen: String
  GlobalResolverId: String
  Name: String
  Tags:
    - Tag

```

## Properties

`ClientToken`

A unique string that identifies the request and ensures idempotency.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

An optional description for the DNS view.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DnssecValidation`

Whether to enable DNSSEC validation for DNS queries in this DNS view. When enabled, the
resolver verifies the authenticity and integrity of DNS responses from public name servers
for DNSSEC-signed domains.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EdnsClientSubnet`

Whether to enable EDNS Client Subnet injection for DNS queries in this DNS view. When
enabled, client subnet information is forwarded to provide more accurate geographic-based
DNS responses.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FirewallRulesFailOpen`

Determines the behavior when Route 53 Global Resolver cannot apply DNS firewall rules due to service
impairment. When enabled, DNS queries are allowed through; when disabled, queries are
blocked.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GlobalResolverId`

The ID of the Route 53 Global Resolver to associate with this DNS view.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

A descriptive name for the DNS view.

_Required_: Yes

_Type_: String

_Pattern_: `(?!^[0-9]+$)([a-zA-Z0-9-_' ']+)`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Tags to associate with the DNS view.

_Required_: No

_Type_: Array of [Tag](aws-properties-route53globalresolver-dnsview-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) of the DNS view.

`CreatedAt`

The date and time when the DNS view was created.

`DnsViewId`

The unique identifier for the DNS view.

`Status`

The operational status of the DNS view.

`UpdatedAt`

The date and time when the DNS view was last updated.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
