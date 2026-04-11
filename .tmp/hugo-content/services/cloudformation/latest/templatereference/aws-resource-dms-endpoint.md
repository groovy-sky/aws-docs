This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DMS::Endpoint

The `AWS::DMS::Endpoint` resource specifies an AWS DMS endpoint.

Currently, AWS CloudFormation supports all AWS DMS endpoint types.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DMS::Endpoint",
  "Properties" : {
      "CertificateArn" : String,
      "DatabaseName" : String,
      "DocDbSettings" : DocDbSettings,
      "DynamoDbSettings" : DynamoDbSettings,
      "ElasticsearchSettings" : ElasticsearchSettings,
      "EndpointIdentifier" : String,
      "EndpointType" : String,
      "EngineName" : String,
      "ExtraConnectionAttributes" : String,
      "GcpMySQLSettings" : GcpMySQLSettings,
      "IbmDb2Settings" : IbmDb2Settings,
      "KafkaSettings" : KafkaSettings,
      "KinesisSettings" : KinesisSettings,
      "KmsKeyId" : String,
      "MicrosoftSqlServerSettings" : MicrosoftSqlServerSettings,
      "MongoDbSettings" : MongoDbSettings,
      "MySqlSettings" : MySqlSettings,
      "NeptuneSettings" : NeptuneSettings,
      "OracleSettings" : OracleSettings,
      "Password" : String,
      "Port" : Integer,
      "PostgreSqlSettings" : PostgreSqlSettings,
      "RedisSettings" : RedisSettings,
      "RedshiftSettings" : RedshiftSettings,
      "ResourceIdentifier" : String,
      "S3Settings" : S3Settings,
      "ServerName" : String,
      "SslMode" : String,
      "SybaseSettings" : SybaseSettings,
      "Tags" : [ Tag, ... ],
      "Username" : String
    }
}

```

### YAML

```yaml

Type: AWS::DMS::Endpoint
Properties:
  CertificateArn: String
  DatabaseName: String
  DocDbSettings:
    DocDbSettings
  DynamoDbSettings:
    DynamoDbSettings
  ElasticsearchSettings:
    ElasticsearchSettings
  EndpointIdentifier: String
  EndpointType: String
  EngineName: String
  ExtraConnectionAttributes: String
  GcpMySQLSettings:
    GcpMySQLSettings
  IbmDb2Settings:
    IbmDb2Settings
  KafkaSettings:
    KafkaSettings
  KinesisSettings:
    KinesisSettings
  KmsKeyId: String
  MicrosoftSqlServerSettings:
    MicrosoftSqlServerSettings
  MongoDbSettings:
    MongoDbSettings
  MySqlSettings:
    MySqlSettings
  NeptuneSettings:
    NeptuneSettings
  OracleSettings:
    OracleSettings
  Password: String
  Port: Integer
  PostgreSqlSettings:
    PostgreSqlSettings
  RedisSettings:
    RedisSettings
  RedshiftSettings:
    RedshiftSettings
  ResourceIdentifier: String
  S3Settings:
    S3Settings
  ServerName: String
  SslMode: String
  SybaseSettings:
    SybaseSettings
  Tags:
    - Tag
  Username: String

```

## Properties

`CertificateArn`

The Amazon Resource Name (ARN) for the certificate.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatabaseName`

The name of the endpoint database. For a MySQL source or target endpoint, don't specify `DatabaseName`.
To migrate to a specific database, use this setting and `targetDbType`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DocDbSettings`

Settings in JSON format for the source and target DocumentDB endpoint.
For more information about other available settings, see
[Using extra connections attributes with Amazon DocumentDB as a source](../../../dms/latest/userguide/chap-source-documentdb.md#CHAP_Source.DocumentDB.ECAs) and
[Using Amazon DocumentDB as a target for AWS Database Migration Service](../../../dms/latest/userguide/chap-target-documentdb.md)
in the _AWS Database Migration Service User Guide_.

_Required_: No

_Type_: [DocDbSettings](aws-properties-dms-endpoint-docdbsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DynamoDbSettings`

Settings in JSON format for the target Amazon DynamoDB endpoint.
For information about other available settings, see
[Using object mapping to migrate data to DynamoDB](../../../dms/latest/userguide/chap-target-dynamodb.md#CHAP_Target.DynamoDB.ObjectMapping)
in the _AWS Database Migration Service User Guide_.

_Required_: No

_Type_: [DynamoDbSettings](aws-properties-dms-endpoint-dynamodbsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ElasticsearchSettings`

Settings in JSON format for the target OpenSearch endpoint. For more information
about the available settings, see
[Extra connection attributes when using OpenSearch as a target for AWS DMS](../../../dms/latest/userguide/chap-target-elasticsearch.md#CHAP_Target.Elasticsearch.Configuration)
in the _AWS Database Migration Service User Guide_.

_Required_: No

_Type_: [ElasticsearchSettings](aws-properties-dms-endpoint-elasticsearchsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EndpointIdentifier`

The database endpoint identifier. Identifiers must begin with a letter and must contain
only ASCII letters, digits, and hyphens. They can't end with a hyphen, or contain two
consecutive hyphens.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EndpointType`

The type of endpoint. Valid values are `source` and `target`.

_Required_: Yes

_Type_: String

_Allowed values_: `source | target`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EngineName`

The type of engine for the endpoint, depending on the `EndpointType` value.

_Valid values_: `mysql` \| `oracle` \|
`postgres` \| `mariadb` \| `aurora` \|
`aurora-postgresql` \| `opensearch` \| `redshift` \| `redshift-serverless` \| `s3` \|
`db2` \| `azuredb` \| `sybase` \| `dynamodb` \| `mongodb` \|
`kinesis` \| `kafka` \| `elasticsearch` \| `docdb` \|
`sqlserver` \| `neptune`

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExtraConnectionAttributes`

Additional attributes associated with the connection. Each attribute is specified as a
name-value pair associated by an equal sign (=). Multiple attributes are separated by a
semicolon (;) with no additional white space. For information on the attributes available
for connecting your source or target endpoint, see
[Working with AWS DMS Endpoints](../../../dms/latest/userguide/chap-endpoints.md)
in the _AWS Database Migration Service User Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GcpMySQLSettings`

Settings in JSON format for the source GCP MySQL endpoint. These settings are much the same as
the settings for any MySQL-compatible endpoint. For more information, see
[Extra connection attributes when using MySQL as a source for AWS DMS](../../../dms/latest/userguide/chap-source-mysql.md#CHAP_Source.MySQL.ConnectionAttrib)
in the _AWS Database Migration Service User Guide_.

_Required_: No

_Type_: [GcpMySQLSettings](aws-properties-dms-endpoint-gcpmysqlsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IbmDb2Settings`

Settings in JSON format for the source IBM Db2 LUW endpoint. For information about other available settings, see
[Extra connection attributes when using Db2 LUW as a source for AWS DMS](../../../dms/latest/userguide/chap-source-db2.md#CHAP_Source.DB2.ConnectionAttrib)
in the _AWS Database Migration Service User Guide_.

_Required_: No

_Type_: [IbmDb2Settings](aws-properties-dms-endpoint-ibmdb2settings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KafkaSettings`

Settings in JSON format for the target Apache Kafka endpoint.
For more information about other available settings, see
[Using object mapping to migrate data to a Kafka topic](../../../dms/latest/userguide/chap-target-kafka.md#CHAP_Target.Kafka.ObjectMapping)
in the _AWS Database Migration Service User Guide_.

_Required_: No

_Type_: [KafkaSettings](aws-properties-dms-endpoint-kafkasettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KinesisSettings`

Settings in JSON format for the target endpoint for Amazon Kinesis Data Streams.
For more information about other available settings, see
[Using object mapping to migrate data to a Kinesis data stream](../../../dms/latest/userguide/chap-target-kinesis.md#CHAP_Target.Kinesis.ObjectMapping)
in the _AWS Database Migration Service User Guide_.

_Required_: No

_Type_: [KinesisSettings](aws-properties-dms-endpoint-kinesissettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyId`

An AWS KMS key identifier that is used to encrypt the connection parameters for the endpoint.

If you don't specify a value for the `KmsKeyId` parameter, AWS DMS uses your default encryption key.

AWS KMS creates the default encryption key for your AWS account.
Your AWS account has a different default encryption key for each AWS Region.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MicrosoftSqlServerSettings`

Settings in JSON format for the source and target Microsoft SQL Server endpoint.
For information about other available settings, see
[Extra connection attributes when using SQL Server as a source for AWS DMS](../../../dms/latest/userguide/chap-source-sqlserver.md#CHAP_Source.SQLServer.ConnectionAttrib) and
[Extra connection attributes when using SQL Server as a target for AWS DMS](../../../dms/latest/userguide/chap-target-sqlserver.md#CHAP_Target.SQLServer.ConnectionAttrib)
in the _AWS Database Migration Service User Guide_.

_Required_: No

_Type_: [MicrosoftSqlServerSettings](aws-properties-dms-endpoint-microsoftsqlserversettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MongoDbSettings`

Settings in JSON format for the source MongoDB endpoint.
For more information about the available settings, see
[Using MongoDB as a target for AWS Database Migration Service](../../../dms/latest/userguide/chap-source-mongodb.md#CHAP_Source.MongoDB.Configuration)
in the _AWS Database Migration Service User Guide_.

_Required_: No

_Type_: [MongoDbSettings](aws-properties-dms-endpoint-mongodbsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MySqlSettings`

Settings in JSON format for the source and target MySQL endpoint. For information about
other available settings, see [Extra connection attributes when using MySQL as a source for AWS DMS](../../../dms/latest/userguide/chap-source-mysql.md#CHAP_Source.MySQL.ConnectionAttrib) and [Extra connection attributes when using a MySQL-compatible database as a target for\
AWS DMS](../../../dms/latest/userguide/chap-target-mysql.md#CHAP_Target.MySQL.ConnectionAttrib) in the _AWS Database Migration Service User Guide._

_Required_: No

_Type_: [MySqlSettings](aws-properties-dms-endpoint-mysqlsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NeptuneSettings`

Settings in JSON format for the target Amazon Neptune endpoint.
For more information about the available settings, see
[Specifying endpoint settings for Amazon Neptune as a target](../../../dms/latest/userguide/chap-target-neptune.md#CHAP_Target.Neptune.EndpointSettings)
in the _AWS Database Migration Service User Guide_.

_Required_: No

_Type_: [NeptuneSettings](aws-properties-dms-endpoint-neptunesettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OracleSettings`

Settings in JSON format for the source and target Oracle endpoint. For information about other available settings, see
[Extra connection attributes when using Oracle as a source for AWS DMS](../../../dms/latest/userguide/chap-source-oracle.md#CHAP_Source.Oracle.ConnectionAttrib) and
[Extra connection attributes when using Oracle as a target for AWS DMS](../../../dms/latest/userguide/chap-target-oracle.md#CHAP_Target.Oracle.ConnectionAttrib)
in the _AWS Database Migration Service User Guide_.

_Required_: No

_Type_: [OracleSettings](aws-properties-dms-endpoint-oraclesettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Password`

The password to be used to log in to the endpoint database.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Port`

The port used by the endpoint database.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PostgreSqlSettings`

Settings in JSON format for the source and target PostgreSQL endpoint.

For information about other available settings, see
[Extra connection attributes when using PostgreSQL as a source for AWS DMS](../../../dms/latest/userguide/chap-source-postgresql.md#CHAP_Source.PostgreSQL.ConnectionAttrib) and
[Extra connection attributes when using PostgreSQL as a target for AWS DMS](../../../dms/latest/userguide/chap-target-postgresql.md#CHAP_Target.PostgreSQL.ConnectionAttrib)
in the _AWS Database Migration Service User Guide_.

_Required_: No

_Type_: [PostgreSqlSettings](aws-properties-dms-endpoint-postgresqlsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RedisSettings`

Settings in JSON format for the target Redis endpoint.
For information about other available settings, see
[Specifying endpoint settings for Redis as a target](../../../dms/latest/userguide/chap-target-redis.md#CHAP_Target.Redis.EndpointSettings)
in the _AWS Database Migration Service User Guide_.

_Required_: No

_Type_: [RedisSettings](aws-properties-dms-endpoint-redissettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RedshiftSettings`

Settings in JSON format for the Amazon Redshift endpoint.

For more information about other available settings, see
[Extra connection attributes when using Amazon Redshift as a target for AWS DMS](../../../dms/latest/userguide/chap-target-redshift.md#CHAP_Target.Redshift.ConnectionAttrib)
in the _AWS Database Migration Service User Guide_.

_Required_: No

_Type_: [RedshiftSettings](aws-properties-dms-endpoint-redshiftsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceIdentifier`

A display name for the resource identifier at the end of the `EndpointArn`
response parameter that is returned in the created `Endpoint` object. The
value for this parameter can have up to 31 characters. It can contain only ASCII
letters, digits, and hyphen ('-'). Also, it can't end with a hyphen or contain two
consecutive hyphens, and can only begin with a letter, such as
`Example-App-ARN1`.

For example, this value might result in the `EndpointArn` value
`arn:aws:dms:eu-west-1:012345678901:rep:Example-App-ARN1`. If you don't
specify a `ResourceIdentifier` value, AWS DMS generates a default
identifier value for the end of `EndpointArn`.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3Settings`

Settings in JSON format for the source and target Amazon S3 endpoint. For more information about other available settings, see
[Extra connection attributes when using Amazon S3 as a source for AWS DMS](../../../dms/latest/userguide/chap-source-s3.md#CHAP_Source.S3.Configuring) and
[Extra connection attributes when using Amazon S3 as a target for AWS DMS](../../../dms/latest/userguide/chap-target-s3.md#CHAP_Target.S3.Configuring)
in the _AWS Database Migration Service User Guide_.

_Required_: No

_Type_: [S3Settings](aws-properties-dms-endpoint-s3settings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServerName`

The name of the server where the endpoint database resides.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SslMode`

The Secure Sockets Layer (SSL) mode to use for the SSL connection. The default is `none`.

###### Note

When `engine_name` is set to S3, the only allowed value is `none`.

_Required_: No

_Type_: String

_Allowed values_: `none | require | verify-ca | verify-full`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SybaseSettings`

Settings in JSON format for the source and target SAP ASE endpoint. For information about other available settings, see
[Extra connection attributes when using SAP ASE as a source for AWS DMS](../../../dms/latest/userguide/chap-source-sap.md#CHAP_Source.SAP.ConnectionAttrib) and
[Extra connection attributes when using SAP ASE as a target for AWS DMS](../../../dms/latest/userguide/chap-target-sap.md#CHAP_Target.SAP.ConnectionAttrib)
in the _AWS Database Migration Service User Guide_.

_Required_: No

_Type_: [SybaseSettings](aws-properties-dms-endpoint-sybasesettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

One or more tags to be assigned to the endpoint.

_Required_: No

_Type_: Array of [Tag](aws-properties-dms-endpoint-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Username`

The user name to be used to log in to the endpoint database.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the endpoint.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ExternalId`

A value that can be used for cross-account validation.

## Examples

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Resources": {
    "myBasicEndpoint": {
      "Type": "AWS::DMS::Endpoint",
      "Properties": {
        "EngineName": "mysql",
        "EndpointType": "source",
        "Username": "username",
        "Password": {
          "Ref": "PasswordParameter"
        },
        "ServerName": "source.db.amazon.com",
        "Port": 1234,
        "DatabaseName": "source-db"
      }
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: "Endpoint test"
Resources:
  BasicEndpoint:
    Type: AWS::DMS::Endpoint
    Properties:
      DatabaseName: my-db
      EndpointType: target
      EngineName: mysql
      Password: PasswordParameter
      Port: 1234
      ServerName: server.db.amazon.com
      Tags:
        - Key: type
          Value: new
      Username: username
```

## See also

- [CreateEndpoint](../../../../reference/dms/latest/apireference/api-createendpoint.md)
in the _AWS Database Migration Service API Reference_

- [Managing AWS resources as a single unit with AWS CloudFormation stacks](../userguide/stacks.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

DocDbSettings

All content copied from https://docs.aws.amazon.com/.
