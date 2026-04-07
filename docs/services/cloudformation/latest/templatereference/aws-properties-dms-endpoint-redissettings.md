This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DMS::Endpoint RedisSettings

Provides information that defines a Redis target endpoint. This
information includes the output format of records applied to the endpoint and details of
transaction and control table data information. For information about other available settings, see
[Specifying endpoint settings for Redis as a target](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Redis.html#CHAP_Target.Redis.EndpointSettings)
in the _AWS Database Migration Service User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthPassword" : String,
  "AuthType" : String,
  "AuthUserName" : String,
  "Port" : Number,
  "ServerName" : String,
  "SslCaCertificateArn" : String,
  "SslSecurityProtocol" : String
}

```

### YAML

```yaml

  AuthPassword: String
  AuthType: String
  AuthUserName: String
  Port: Number
  ServerName: String
  SslCaCertificateArn: String
  SslSecurityProtocol: String

```

## Properties

`AuthPassword`

The password provided with the `auth-role` and
`auth-token` options of the `AuthType` setting for a Redis
target endpoint.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuthType`

The type of authentication to perform when connecting to a Redis target. Options include
`none`, `auth-token`, and `auth-role`. The
`auth-token` option requires an `AuthPassword` value to be provided. The
`auth-role` option requires `AuthUserName` and `AuthPassword` values
to be provided.

_Required_: No

_Type_: String

_Allowed values_: `none | auth-role | auth-token`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuthUserName`

The user name provided with the `auth-role` option of the
`AuthType` setting for a Redis target endpoint.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Port`

Transmission Control Protocol (TCP) port for the endpoint.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServerName`

Fully qualified domain name of the endpoint.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SslCaCertificateArn`

The Amazon Resource Name (ARN) for the certificate authority (CA) that DMS uses to
connect to your Redis target endpoint.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SslSecurityProtocol`

The connection to a Redis target endpoint using Transport Layer Security (TLS). Valid
values include `plaintext` and `ssl-encryption`. The default is
`ssl-encryption`. The `ssl-encryption` option makes an encrypted
connection. Optionally, you can identify an Amazon Resource Name (ARN) for an SSL certificate authority (CA)
using the `SslCaCertificateArn ` setting. If an ARN isn't given for a CA, DMS
uses the Amazon root CA.

The `plaintext` option doesn't provide Transport Layer Security (TLS)
encryption for traffic between endpoint and database.

_Required_: No

_Type_: String

_Allowed values_: `plaintext | ssl-encryption`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PostgreSqlSettings

RedshiftSettings
