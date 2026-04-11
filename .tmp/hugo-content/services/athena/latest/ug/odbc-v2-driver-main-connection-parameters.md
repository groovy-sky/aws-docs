# Main ODBC 2.x connection parameters

The following sections describe each of the main connection parameters.

## Data source name

Specifies the name of your data source.

**Connection string name****Parameter type****Default value****Connection string example**DSNOptional for DSN-less connection types`none``DSN=AmazonAthenaOdbcUsWest1;`

## Description

Contains description of your data source.

**Connection string name****Parameter type****Default value****Connection string example**DescriptionOptional`none``Description=Connection to Amazon Athena us-west-1;`

## Catalog

Specifies the data catalog name. For more information about catalogs, see [DataCatalog](../../../../reference/athena/latest/apireference/api-datacatalog.md) in the
Amazon Athena API Reference.

**Connection string name****Parameter type****Default value****Connection string example**CatalogOptional`AwsDataCatalog``Catalog=AwsDataCatalog;`

## Region

Specifies the AWS Region. For information about AWS Regions, see [Regions and\
Availability Zones](https://aws.amazon.com/about-aws/global-infrastructure/regions_az).

**Connection string name****Parameter type****Default value****Connection string example**AwsRegionMandatory`none``AwsRegion` =us-west-1;

## Database

Specifies the database name. For more information about databases, see [Database](../../../../reference/athena/latest/apireference/api-database.md) in the _Amazon Athena API Reference_.

**Connection string name****Parameter type****Default value****Connection string example**SchemaOptional`default``Schema=default;`

## Workgroup

Specifies the workgroup name. For more information about workgroups, see [WorkGroup](../../../../reference/athena/latest/apireference/api-workgroup.md) in the
_Amazon Athena API Reference_.

**Connection string name****Parameter type****Default value****Connection string example**WorkgroupOptional`primary``Workgroup=primary;`

## Output location

Specifies the location in Amazon S3 where query results are stored. For more information
about output location, see [ResultConfiguration](../../../../reference/athena/latest/apireference/api-resultconfiguration.md) in the _Amazon Athena API Reference_.

**Connection string name****Parameter type****Default value****Connection string example**S3OutputLocationMandatory`none``S3OutputLocation=s3://amzn-s3-demo-bucket/;`

## Encryption options

**Dialog parameter name**: Encryption options

Specifies encryption option. For more information about encryption options, see [EncryptionConfiguration](../../../../reference/athena/latest/apireference/api-encryptionconfiguration.md)
in the _Amazon Athena API Reference_.

**Connection string name****Parameter type****Default value****Possible values****Connection string example**S3OutputEncOptionOptional`none``NOT_SET`, `SSE_S3`, `SSE_KMS`,
`CSE_KMS``S3OutputEncOption=SSE_S3;`

## KMS key

Specifies a KMS key for encryption. For more information about encryption
configuration for KMS Keys, see [EncryptionConfiguration](../../../../reference/athena/latest/apireference/api-encryptionconfiguration.md)
in the _Amazon Athena API Reference_.

**Connection string name****Parameter type****Default value****Connection string example**S3OutputEncKMSKeyOptional`none``S3OutputEncKMSKey=your_key;`

## Connection test

ODBC Data Source Administrator provides a **Test** option that you
can use to test your ODBC 2.x connection to Amazon Athena. For steps, see [Configuring a data source name on Windows](odbc-v2-driver-getting-started-windows.md#odbc-v2-driver-configuring-dsn-on-windows). When you test a
connection, the ODBC driver calls the [GetWorkGroup](../../../../reference/athena/latest/apireference/api-getworkgroup.md) Athena API action. The call uses the authentication type and
corresponding credentials provider that you specified to retrieve the credentials. There
is no charge for the connection test when you use the ODBC 2.x driver. The test does not
generate query results in your Amazon S3 bucket.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ODBC 2.x connection parameters

Authentication

All content copied from https://docs.aws.amazon.com/.
