This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DMS::DataProvider PostgreSqlSettings

Provides information that defines a PostgreSQL endpoint.

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

Endpoint TCP port. The default is 5432.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServerName`

The host name of the endpoint database.

For an Amazon RDS PostgreSQL instance, this is the output of [DescribeDBInstances](../../../../reference/amazonrds/latest/apireference/api-describedbinstances.md), in the `Endpoint.Address` field.

For an Aurora PostgreSQL instance, this is the output of [DescribeDBClusters](../../../../reference/amazonrds/latest/apireference/api-describedbclusters.md), in the `Endpoint` field.

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

OracleSettings

RedshiftSettings

All content copied from https://docs.aws.amazon.com/.
