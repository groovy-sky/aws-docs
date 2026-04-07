This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53GlobalResolver::AccessSource

Creates an access source for a DNS view. Access sources define IP addresses or CIDR
ranges that are allowed to send DNS queries to the Route 53 Global Resolver, along with the permitted DNS
protocols.

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
  "Type" : "AWS::Route53GlobalResolver::AccessSource",
  "Properties" : {
      "Cidr" : String,
      "ClientToken" : String,
      "DnsViewId" : String,
      "IpAddressType" : String,
      "Name" : String,
      "Protocol" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Route53GlobalResolver::AccessSource
Properties:
  Cidr: String
  ClientToken: String
  DnsViewId: String
  IpAddressType: String
  Name: String
  Protocol: String
  Tags:
    - Tag

```

## Properties

`Cidr`

The CIDR block that defines the IP address range for the access source.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `42`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientToken`

A unique string that identifies the request and ensures idempotency.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DnsViewId`

The ID of the DNS view that the access source is associated with.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IpAddressType`

The IP address type of the access source.

_Required_: No

_Type_: String

_Allowed values_: `IPV4 | IPV6`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the access source.

_Required_: No

_Type_: String

_Pattern_: `(?!^[0-9]+$)([a-zA-Z0-9-_' ']+)`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Protocol`

The protocol used by the access source.

_Required_: Yes

_Type_: String

_Allowed values_: `DO53 | DOH | DOT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Tags to associate with the access source.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-route53globalresolver-accesssource-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`AccessSourceId`

The unique identifier for the access source.

`Arn`

The Amazon Resource Name (ARN) of the access source.

`CreatedAt`

The date and time when the access source was created.

`Status`

The current status of the access source.

`UpdatedAt`

The date and time when the access source was last updated.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon Route 53 Global Resolver

Tag
