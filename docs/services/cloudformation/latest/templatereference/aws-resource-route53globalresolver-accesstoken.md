This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53GlobalResolver::AccessToken

Creates an access token for a DNS view. Access tokens provide token-based authentication
for DNS-over-HTTPS (DoH) and DNS-over-TLS (DoT) connections to the Route 53 Global Resolver.

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
  "Type" : "AWS::Route53GlobalResolver::AccessToken",
  "Properties" : {
      "ClientToken" : String,
      "DnsViewId" : String,
      "ExpiresAt" : String,
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Route53GlobalResolver::AccessToken
Properties:
  ClientToken: String
  DnsViewId: String
  ExpiresAt: String
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

`DnsViewId`

The ID of the DNS view associated with the token.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ExpiresAt`

The date and time when the token expires.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the token.

_Required_: No

_Type_: String

_Pattern_: `(?!^[0-9]+$)([a-zA-Z0-9-_' ']+)`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of user-defined keys and optional values. These tags can be used for
categorization and organization.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-route53globalresolver-accesstoken-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`AccessTokenId`

The unique identifier for the access token.

`Arn`

The Amazon Resource Name (ARN) of the token.

`CreatedAt`

The date and time when the token was created.

`GlobalResolverId`

The ID of the global resolver associated with the token.

`Status`

The current status of the token.

`UpdatedAt`

The date and time when the token was last updated.

`Value`

The access token value. This token should be included in DoH and DoT requests for
authentication. Keep this value secure as it provides access to your Route 53 Global Resolver.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
