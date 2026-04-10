This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QBusiness::Plugin OAuth2ClientCredentialConfiguration

Information about the OAuth 2.0 authentication credential/token used to configure a
plugin.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthorizationUrl" : String,
  "RoleArn" : String,
  "SecretArn" : String,
  "TokenUrl" : String
}

```

### YAML

```yaml

  AuthorizationUrl: String
  RoleArn: String
  SecretArn: String
  TokenUrl: String

```

## Properties

`AuthorizationUrl`

The redirect URL required by the OAuth 2.0 protocol for Amazon Q Business to
authenticate a plugin user through a third party authentication server.

_Required_: No

_Type_: String

_Pattern_: `^(https?|ftp|file)://([^\s]*)$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The ARN of an IAM role used by Amazon Q Business to access the OAuth 2.0
authentication credentials stored in a Secrets Manager secret.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:[a-z0-9-\.]{1,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}$`

_Minimum_: `0`

_Maximum_: `1284`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretArn`

The ARN of the Secrets Manager secret that stores the OAuth 2.0 credentials/token
used for plugin configuration.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:[a-z0-9-\.]{1,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}$`

_Minimum_: `0`

_Maximum_: `1284`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TokenUrl`

The URL required by the OAuth 2.0 protocol to exchange an end user authorization code
for an access token.

_Required_: No

_Type_: String

_Pattern_: `^(https?|ftp|file)://([^\s]*)$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomPluginConfiguration

PluginAuthConfiguration

All content copied from https://docs.aws.amazon.com/.
