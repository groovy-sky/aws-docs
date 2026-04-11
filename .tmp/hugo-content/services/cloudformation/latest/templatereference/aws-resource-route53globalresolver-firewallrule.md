This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53GlobalResolver::FirewallRule

Creates a DNS firewall rule. Firewall rules define actions (ALLOW, BLOCK, or ALERT) to
take on DNS queries that match specified domain lists, managed domain lists, or advanced
threat protections.

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
  "Type" : "AWS::Route53GlobalResolver::FirewallRule",
  "Properties" : {
      "Action" : String,
      "BlockOverrideDnsType" : String,
      "BlockOverrideDomain" : String,
      "BlockOverrideTtl" : Integer,
      "BlockResponse" : String,
      "ClientToken" : String,
      "ConfidenceThreshold" : String,
      "Description" : String,
      "DnsAdvancedProtection" : String,
      "DnsViewId" : String,
      "FirewallDomainListId" : String,
      "Name" : String,
      "Priority" : Integer,
      "QType" : String
    }
}

```

### YAML

```yaml

Type: AWS::Route53GlobalResolver::FirewallRule
Properties:
  Action: String
  BlockOverrideDnsType: String
  BlockOverrideDomain: String
  BlockOverrideTtl: Integer
  BlockResponse: String
  ClientToken: String
  ConfidenceThreshold: String
  Description: String
  DnsAdvancedProtection: String
  DnsViewId: String
  FirewallDomainListId: String
  Name: String
  Priority: Integer
  QType: String

```

## Properties

`Action`

The action configured for the updated firewall rule.

_Required_: Yes

_Type_: String

_Allowed values_: `ALLOW | ALERT | BLOCK`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BlockOverrideDnsType`

The DNS record type configured for the updated firewall rule's custom response.

_Required_: No

_Type_: String

_Allowed values_: `CNAME`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BlockOverrideDomain`

The custom domain name configured for the updated firewall rule's BLOCK response.

_Required_: No

_Type_: String

_Pattern_: `\*?[-a-zA-Z0-9.]+`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BlockOverrideTtl`

The TTL value configured for the updated firewall rule's custom response.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `604800`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BlockResponse`

The type of block response configured for the updated firewall rule.

_Required_: No

_Type_: String

_Allowed values_: `NODATA | NXDOMAIN | OVERRIDE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientToken`

The unique string that identified the request and ensured idempotency.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ConfidenceThreshold`

The confidence threshold configured for the updated firewall rule's advanced threat
detection.

_Required_: No

_Type_: String

_Allowed values_: `LOW | MEDIUM | HIGH`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the updated firewall rule.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DnsAdvancedProtection`

Whether advanced DNS threat protection is enabled for the updated firewall rule.

_Required_: No

_Type_: String

_Allowed values_: `DGA | DNS_TUNNELING | DICTIONARY_DGA`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DnsViewId`

The ID of the DNS view associated with the updated firewall rule.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FirewallDomainListId`

The ID of the firewall domain list associated with the updated firewall rule.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the updated firewall rule.

_Required_: Yes

_Type_: String

_Pattern_: `(?!^[0-9]+$)([a-zA-Z0-9-_' ']+)`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Priority`

The priority of the updated firewall rule.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `10000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QType`

The DNS query type that the firewall rule should match.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `16`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

`CreatedAt`

The date and time when the firewall rule was originally created.

`FirewallRuleId`

The unique identifier of the firewall rule to update.

`QueryType`

The DNS query type that the updated firewall rule matches.

`Status`

The current status of the updated firewall rule.

`UpdatedAt`

The date and time when the firewall rule was last updated.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::Route53GlobalResolver::GlobalResolver

All content copied from https://docs.aws.amazon.com/.
