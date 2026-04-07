This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodePipeline::Webhook WebhookAuthConfiguration

The authentication applied to incoming webhook trigger requests.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllowedIPRange" : String,
  "SecretToken" : String
}

```

### YAML

```yaml

  AllowedIPRange: String
  SecretToken: String

```

## Properties

`AllowedIPRange`

The property used to configure acceptance of webhooks in an IP address range. For
IP, only the `AllowedIPRange` property must be set. This property must be set
to a valid CIDR range.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretToken`

The property used to configure GitHub authentication. For GITHUB\_HMAC, only the
`SecretToken` property must be set.

###### Important

When creating CodePipeline webhooks, do not use your own credentials or
reuse the same secret token across multiple webhooks. For optimal security, generate
a unique secret token for each webhook you create. The secret token is an arbitrary
string that you provide, which GitHub uses to compute and sign the webhook payloads
sent to CodePipeline, for protecting the integrity and authenticity of the
webhook payloads. Using your own credentials or reusing the same token across
multiple webhooks can lead to security vulnerabilities.

###### Note

If a secret token was provided, it will be redacted in the response.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::CodePipeline::Webhook

WebhookFilterRule
