This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53GlobalResolver::FirewallDomainList

Creates a firewall domain list. Domain lists are reusable sets of domain specifications
that you use in DNS firewall rules to allow, block, or alert on DNS queries to specific
domains.

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
  "Type" : "AWS::Route53GlobalResolver::FirewallDomainList",
  "Properties" : {
      "ClientToken" : String,
      "Description" : String,
      "DomainFileUrl" : String,
      "Domains" : [ String, ... ],
      "GlobalResolverId" : String,
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Route53GlobalResolver::FirewallDomainList
Properties:
  ClientToken: String
  Description: String
  DomainFileUrl: String
  Domains:
    - String
  GlobalResolverId: String
  Name: String
  Tags:
    - Tag

```

## Properties

`ClientToken`

A unique, case-sensitive identifier to ensure idempotency. This means that making the
same request multiple times with the same `clientToken` has the same result
every time.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

A description of the firewall domain list.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DomainFileUrl`

The fully qualified URL of the file in Amazon S3 that contains the list of domains to
import. The file should contain one domain per line.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Domains`

A list of the domains. You can add up to 1000 domains per request.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GlobalResolverId`

The ID of the global resolver that the firewall domain list is associated with.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the firewall domain list.

_Required_: Yes

_Type_: String

_Pattern_: `(?!^[0-9]+$)([a-zA-Z0-9-_' ']+)`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of user-defined keys and optional values. These tags can be used for
categorization and organization.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-route53globalresolver-firewalldomainlist-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) of the firewall domain list.

`CreatedAt`

The date and time when the firewall domain list was created.

`DomainCount`

Number of domains in the domain list.

`FirewallDomainListId`

ID of the domain list.

`Status`

The current status of the firewall domain list.

`StatusMessage`

Additional information about the status of the domain list.

`UpdatedAt`

The date and time when the firewall domain list was last updated.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
