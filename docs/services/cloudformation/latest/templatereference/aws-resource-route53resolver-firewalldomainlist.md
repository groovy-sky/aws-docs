This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53Resolver::FirewallDomainList

High-level information about a list of firewall domains for use in a [AWS::Route53Resolver::FirewallRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53resolver-firewallrulegroup-rule.html). This is returned by [GetFirewallDomainList](https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_GetFirewallDomainList.html).

To retrieve the domains that are defined for this domain list, call
[ListFirewallDomains](https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_ListFirewallDomains.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Route53Resolver::FirewallDomainList",
  "Properties" : {
      "DomainFileUrl" : String,
      "Domains" : [ String, ... ],
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Route53Resolver::FirewallDomainList
Properties:
  DomainFileUrl: String
  Domains:
    - String
  Name: String
  Tags:
    - Tag

```

## Properties

`DomainFileUrl`

The fully qualified URL or URI of the file stored in Amazon Simple Storage Service
(Amazon S3) that contains the list of domains to import.

The file must be in an S3 bucket that's in the same Region
as your DNS Firewall. The file must be a text file and must contain a single domain per line.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Domains`

A list of the domain lists that you have defined.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the domain list.

_Required_: No

_Type_: String

_Pattern_: `(?!^[0-9]+$)([a-zA-Z0-9\-_' ']+)`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A list of the tag keys and values that you want to associate with the domain list.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-route53resolver-firewalldomainlist-tag.html)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `FirewallDomainList` object.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the firewall domain list.

`CreationTime`

The date and time that the domain list was created, in Unix time format and Coordinated Universal Time (UTC).

`CreatorRequestId`

A unique string defined by you to identify the request. This allows you to
retry failed requests without the risk of running the operation twice. This can be any
unique string, for example, a timestamp.

`DomainCount`

The number of domain names that are specified in the domain list.

`Id`

The ID of the domain list.

`ManagedOwnerName`

The owner of the list, used only for lists that are not managed by you.
For example, the managed domain list
`AWSManagedDomainsMalwareDomainList` has
the managed owner name `Route 53 Resolver DNS Firewall`.

`ModificationTime`

The date and time that the domain list was last modified, in Unix time format and Coordinated Universal Time (UTC).

`Status`

The status of the domain list.

`StatusMessage`

Additional information about the status of the list, if available.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon Route 53 Resolver

Tag
