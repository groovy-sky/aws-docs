# Basic connection parameters

The following sections describe the basic connection parameters for the JDBC 3.x
driver.

## Region

The AWS Region where queries will be run. For a list of regions, see [Amazon Athena endpoints and\
quotas](../../../../general/latest/gr/athena.md).

Parameter nameAliasParameter typeDefault valueRegionAwsRegion (deprecated)Mandatory (but if not provided, will be searched using the [DefaultAwsRegionProviderChain](https://sdk.amazonaws.com/java/api/latest/software/amazon/awssdk/regions/providers/DefaultAwsRegionProviderChain.html)) none

## Catalog

The catalog that contains the databases and the tables that will be accessed with the
driver. For information about catalogs, see [DataCatalog](../../../../reference/athena/latest/apireference/api-datacatalog.md).

Parameter nameAliasParameter typeDefault valueCatalognoneOptionalAwsDataCatalog

## Database

The database where queries will run. Tables that are not explicitly qualified with a
database name are resolved to this database. For information about databases, see [Database](../../../../reference/athena/latest/apireference/api-database.md).

Parameter nameAliasParameter typeDefault valueDatabaseSchemaOptionaldefault

## Workgroup

The workgroup in which queries will run. For information about workgroups, see [WorkGroup](../../../../reference/athena/latest/apireference/api-workgroup.md).

Parameter nameAliasParameter typeDefault valueWorkGroupnoneOptionalprimary

## Output location

The location in Amazon S3 where query results will be stored. For information about output
location, see [ResultConfiguration](../../../../reference/athena/latest/apireference/api-resultconfiguration.md).

Parameter nameAliasParameter typeDefault valueOutputLocationS3OutputLocation (deprecated)Mandatory (unless the workgroup specifies an output location)none

[Document Conventions](../../../../general/latest/gr/docconventions.md)

JDBC 3.x connection parameters

Advanced

All content copied from https://docs.aws.amazon.com/.
