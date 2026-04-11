This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::CloudFrontOriginAccessIdentity CloudFrontOriginAccessIdentityConfig

Origin access identity configuration. Send a `GET` request to the
`/CloudFront API version/CloudFront/identity ID/config`
resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Comment" : String
}

```

### YAML

```yaml

  Comment: String

```

## Properties

`Comment`

A comment to describe the origin access identity. The comment cannot be longer than
128 characters.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [CloudFrontOriginAccessIdentityConfig](../../../../reference/cloudfront/latest/apireference/api-cloudfrontoriginaccessidentityconfig.md) in the _Amazon CloudFront API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::CloudFront::CloudFrontOriginAccessIdentity

AWS::CloudFront::ConnectionFunction

All content copied from https://docs.aws.amazon.com/.
