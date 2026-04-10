This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::DataSource WebCrawlerBasicAuthentication

Provides the configuration information to connect to websites that require basic user
authentication.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Credentials" : String,
  "Host" : String,
  "Port" : Integer
}

```

### YAML

```yaml

  Credentials: String
  Host: String
  Port: Integer

```

## Properties

`Credentials`

The Amazon Resource Name (ARN) of an AWS Secrets Manager secret. You create a
secret to store your credentials in [AWS Secrets Manager](../../../secretsmanager/latest/userguide/intro.md)

You use a secret if basic authentication credentials are required to connect to a
website. The secret stores your credentials of user name and password.

_Required_: Yes

_Type_: String

_Pattern_: `arn:[a-z0-9-\.]{1,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}`

_Minimum_: `1`

_Maximum_: `1284`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Host`

The name of the website host you want to connect to using authentication
credentials.

For example, the host name of https://a.example.com/page1.html is
"a.example.com".

_Required_: Yes

_Type_: String

_Pattern_: `([^\s]*)`

_Minimum_: `1`

_Maximum_: `253`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Port`

The port number of the website host you want to connect to using authentication
credentials.

For example, the port for https://a.example.com/page1.html is 443, the standard port
for HTTPS.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WebCrawlerAuthenticationConfiguration

WebCrawlerConfiguration

All content copied from https://docs.aws.amazon.com/.
