This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DMS::DataProvider MongoDbSettings

Provides information that defines a MongoDB endpoint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthMechanism" : String,
  "AuthSource" : String,
  "AuthType" : String,
  "CertificateArn" : String,
  "DatabaseName" : String,
  "Port" : Integer,
  "ServerName" : String,
  "SslMode" : String
}

```

### YAML

```yaml

  AuthMechanism: String
  AuthSource: String
  AuthType: String
  CertificateArn: String
  DatabaseName: String
  Port: Integer
  ServerName: String
  SslMode: String

```

## Properties

`AuthMechanism`

The authentication mechanism you use to access the MongoDB source endpoint.

For the default value, in MongoDB version 2.x, `"default"` is
`"mongodb_cr"`. For MongoDB version 3.x or later, `"default"` is
`"scram_sha_1"`. This setting isn't used when `AuthType` is
set to `"no"`.

_Required_: No

_Type_: String

_Allowed values_: `default | mongodb_cr | scram_sha_1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuthSource`

The MongoDB database name. This setting isn't used when `AuthType` is
set to `"no"`.

The default is `"admin"`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuthType`

The authentication type you use to access the MongoDB source endpoint.

When when set to `"no"`, user name and password parameters are not used and
can be empty.

_Required_: No

_Type_: String

_Allowed values_: `no | password`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CertificateArn`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatabaseName`

The database name on the MongoDB source endpoint.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Port`

The port value for the MongoDB source endpoint.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServerName`

The name of the server on the MongoDB source endpoint. For MongoDB Atlas, provide the
server name for any of the servers in the replication set.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SslMode`

Property description not available.

_Required_: No

_Type_: String

_Allowed values_: `none | require | verify-full`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MicrosoftSqlServerSettings

MySqlSettings

All content copied from https://docs.aws.amazon.com/.
