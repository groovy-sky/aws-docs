This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DMS::DataProvider DocDbSettings

Provides information that defines a DocumentDB endpoint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CertificateArn" : String,
  "DatabaseName" : String,
  "Port" : Integer,
  "ServerName" : String,
  "SslMode" : String
}

```

### YAML

```yaml

  CertificateArn: String
  DatabaseName: String
  Port: Integer
  ServerName: String
  SslMode: String

```

## Properties

`CertificateArn`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatabaseName`

The database name on the DocumentDB source endpoint.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Port`

The port value for the DocumentDB source endpoint.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServerName`

The name of the server on the DocumentDB source endpoint.

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

AWS::DMS::DataProvider

IbmDb2LuwSettings

All content copied from https://docs.aws.amazon.com/.
