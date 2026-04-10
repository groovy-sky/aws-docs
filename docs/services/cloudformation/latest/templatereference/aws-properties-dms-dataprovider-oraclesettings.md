This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DMS::DataProvider OracleSettings

Provides information that defines an Oracle endpoint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AsmServer" : String,
  "CertificateArn" : String,
  "DatabaseName" : String,
  "Port" : Integer,
  "SecretsManagerOracleAsmAccessRoleArn" : String,
  "SecretsManagerOracleAsmSecretId" : String,
  "SecretsManagerSecurityDbEncryptionAccessRoleArn" : String,
  "SecretsManagerSecurityDbEncryptionSecretId" : String,
  "ServerName" : String,
  "SslMode" : String
}

```

### YAML

```yaml

  AsmServer: String
  CertificateArn: String
  DatabaseName: String
  Port: Integer
  SecretsManagerOracleAsmAccessRoleArn: String
  SecretsManagerOracleAsmSecretId: String
  SecretsManagerSecurityDbEncryptionAccessRoleArn: String
  SecretsManagerSecurityDbEncryptionSecretId: String
  ServerName: String
  SslMode: String

```

## Properties

`AsmServer`

For an Oracle source endpoint, your ASM server address. You can set this value from the
`asm_server` value. You set `asm_server` as part of the extra
connection attribute string to access an Oracle server with Binary Reader that uses ASM.
For more information, see [Configuration for change data capture (CDC) on an Oracle source\
database](../../../dms/latest/userguide/chap-source-oracle.md#dms/latest/userguide/CHAP_Source.Oracle.html#CHAP_Source.Oracle.CDC.Configuration).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

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

`SecretsManagerOracleAsmAccessRoleArn`

Required only if your Oracle endpoint uses Automatic Storage Management (ASM). The full
ARN of the IAM role that specifies AWS DMS as the trusted entity and grants the required
permissions to access the `SecretsManagerOracleAsmSecret`. This
`SecretsManagerOracleAsmSecret` has the secret value that allows access to
the Oracle ASM of the endpoint.

###### Note

You can specify one of two sets of values for these permissions. You can specify the
values for this setting and `SecretsManagerOracleAsmSecretId`. Or you can
specify clear-text values for `AsmUser`, `AsmPassword`, and
`AsmServerName`. You can't specify both. For more information on
creating this `SecretsManagerOracleAsmSecret` and the
`SecretsManagerOracleAsmAccessRoleArn` and
`SecretsManagerOracleAsmSecretId` required to access it, see [Using secrets to access AWS Database Migration Service resources](../../../dms/latest/userguide/chap-security.md#security-iam-secretsmanager) in the
_AWS Database Migration Service User Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretsManagerOracleAsmSecretId`

Required only if your Oracle endpoint uses Automatic Storage Management (ASM). The full
ARN, partial ARN, or friendly name of the `SecretsManagerOracleAsmSecret` that
contains the Oracle ASM connection details for the Oracle endpoint.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretsManagerSecurityDbEncryptionAccessRoleArn`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretsManagerSecurityDbEncryptionSecretId`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServerName`

Fully qualified domain name of the endpoint.

For an Amazon RDS Oracle instance, this is the output of [DescribeDBInstances](../../../../reference/amazonrds/latest/apireference/api-describedbinstances.md), in the `Endpoint.Address` field.

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

MySqlSettings

PostgreSqlSettings

All content copied from https://docs.aws.amazon.com/.
