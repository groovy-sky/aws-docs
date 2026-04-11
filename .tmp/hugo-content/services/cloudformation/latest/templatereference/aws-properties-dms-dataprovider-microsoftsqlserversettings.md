This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DMS::DataProvider MicrosoftSqlServerSettings

Provides information that defines a Microsoft SQL Server endpoint.

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

Database name for the endpoint.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Port`

Endpoint TCP port.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServerName`

Fully qualified domain name of the endpoint. For an Amazon RDS SQL Server instance, this is
the output of [DescribeDBInstances](../../../../reference/amazonrds/latest/apireference/api-describedbinstances.md), in the `Endpoint.Address` field.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SslMode`

Property description not available.

_Required_: Yes

_Type_: String

_Allowed values_: `none | require | verify-ca | verify-full`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MariaDbSettings

MongoDbSettings

All content copied from https://docs.aws.amazon.com/.
