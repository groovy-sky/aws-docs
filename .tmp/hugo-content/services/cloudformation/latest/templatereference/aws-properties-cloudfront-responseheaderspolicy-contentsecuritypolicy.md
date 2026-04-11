This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::ResponseHeadersPolicy ContentSecurityPolicy

The policy directives and their values that CloudFront includes as values for the
`Content-Security-Policy` HTTP response header.

For more information about the `Content-Security-Policy` HTTP response
header, see [Content-Security-Policy](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy) in the MDN Web Docs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContentSecurityPolicy" : String,
  "Override" : Boolean
}

```

### YAML

```yaml

  ContentSecurityPolicy: String
  Override: Boolean

```

## Properties

`ContentSecurityPolicy`

The policy directives and their values that CloudFront includes as values for the
`Content-Security-Policy` HTTP response header.

_Required_: Yes

_Type_: [String](aws-properties-cloudfront-responseheaderspolicy-contentsecuritypolicy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Override`

A Boolean that determines whether CloudFront overrides the
`Content-Security-Policy` HTTP response header received from the origin
with the one specified in this response headers policy.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AccessControlExposeHeaders

ContentTypeOptions

All content copied from https://docs.aws.amazon.com/.
