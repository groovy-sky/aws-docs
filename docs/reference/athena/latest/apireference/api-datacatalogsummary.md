# DataCatalogSummary

The summary information for the data catalog, which includes its name and type.

## Contents

**CatalogName**

The name of the data catalog. The catalog name is unique for the AWS account and can use a maximum of 127 alphanumeric, underscore, at sign,
or hyphen characters. The remainder of the length constraint of 256 is reserved for use
by Athena.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Pattern: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

Required: No

**ConnectionType**

The type of connection for a `FEDERATED` data catalog (for example,
`REDSHIFT`, `MYSQL`, or `SQLSERVER`). For
information about individual connectors, see [Available data source\
connectors](../../../../services/athena/latest/ug/connectors-available.md).

Type: String

Valid Values: `DYNAMODB | MYSQL | POSTGRESQL | REDSHIFT | ORACLE | SYNAPSE | SQLSERVER | DB2 | OPENSEARCH | BIGQUERY | GOOGLECLOUDSTORAGE | HBASE | DOCUMENTDB | CMDB | TPCDS | TIMESTREAM | SAPHANA | SNOWFLAKE | DATALAKEGEN2 | DB2AS400`

Required: No

**Error**

Text of the error that occurred during data catalog creation or deletion.

Type: String

Required: No

**Status**

The status of the creation or deletion of the data catalog.

- The `LAMBDA`, `GLUE`, and `HIVE` data catalog
types are created synchronously. Their status is either
`CREATE_COMPLETE` or `CREATE_FAILED`.

- The `FEDERATED` data catalog type is created asynchronously.

Data catalog creation status:

- `CREATE_IN_PROGRESS`: Federated data catalog creation in
progress.

- `CREATE_COMPLETE`: Data catalog creation complete.

- `CREATE_FAILED`: Data catalog could not be created.

- `CREATE_FAILED_CLEANUP_IN_PROGRESS`: Federated data catalog
creation failed and is being removed.

- `CREATE_FAILED_CLEANUP_COMPLETE`: Federated data catalog creation
failed and was removed.

- `CREATE_FAILED_CLEANUP_FAILED`: Federated data catalog creation
failed but could not be removed.

Data catalog deletion status:

- `DELETE_IN_PROGRESS`: Federated data catalog deletion in
progress.

- `DELETE_COMPLETE`: Federated data catalog deleted.

- `DELETE_FAILED`: Federated data catalog could not be
deleted.

Type: String

Valid Values: `CREATE_IN_PROGRESS | CREATE_COMPLETE | CREATE_FAILED | CREATE_FAILED_CLEANUP_IN_PROGRESS | CREATE_FAILED_CLEANUP_COMPLETE | CREATE_FAILED_CLEANUP_FAILED | DELETE_IN_PROGRESS | DELETE_COMPLETE | DELETE_FAILED`

Required: No

**Type**

The data catalog type.

Type: String

Valid Values: `LAMBDA | GLUE | HIVE | FEDERATED`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/athena-2017-05-18/DataCatalogSummary)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/athena-2017-05-18/DataCatalogSummary)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/athena-2017-05-18/DataCatalogSummary)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DataCatalog

Datum
