This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53GlobalResolver::GlobalResolver

Creates a new Route 53 Global Resolver instance. A Route 53 Global Resolver is a global, internet-accessible DNS resolver
that provides secure DNS resolution for both public and private domains through global
anycast IP addresses.

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
  "Type" : "AWS::Route53GlobalResolver::GlobalResolver",
  "Properties" : {
      "ClientToken" : String,
      "Description" : String,
      "IpAddressType" : String,
      "Name" : String,
      "ObservabilityRegion" : String,
      "Regions" : [ String, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Route53GlobalResolver::GlobalResolver
Properties:
  ClientToken: String
  Description: String
  IpAddressType: String
  Name: String
  ObservabilityRegion: String
  Regions:
    - String
  Tags:
    - Tag

```

## Properties

`ClientToken`

The unique string that identifies the request and ensures idempotency.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

A description of the global resolver.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IpAddressType`

The IP address type for the Route 53 Global Resolver. Valid values are IPV4 (default) or
DUAL\_STACK for both IPv4 and IPv6 support.

_Required_: No

_Type_: String

_Allowed values_: `IPV4 | DUAL_STACK`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the global resolver.

_Required_: Yes

_Type_: String

_Pattern_: `(?!^[0-9]+$)([a-zA-Z0-9-_' ']+)`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ObservabilityRegion`

The AWS Region where observability data is collected for the global resolver.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Regions`

The AWS Regions where the global resolver is deployed.

_Required_: Yes

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Tags to associate with the Route 53 Global Resolver. Tags are key-value pairs that help you organize and
identify your resources.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-route53globalresolver-globalresolver-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) of the global resolver.

`CreatedAt`

The date and time when the global resolver was created.

`DnsName`

The DNS name of the global resolver.

`GlobalResolverId`

The unique identifier for the Route 53 Global Resolver.

`IPv4Addresses`

The IPv4 addresses assigned to the global resolver.

`IPv6Addresses`

The global anycast IPv6 addresses associated with the Route 53 Global Resolver.
This field is only populated when ipAddressType is DUAL\_STACK.
DNS clients can send queries to these addresses from anywhere on the internet.

`Status`

The current status of the global resolver.

`UpdatedAt`

The date and time when the global resolver was last updated.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Route53GlobalResolver::FirewallRule

Tag
