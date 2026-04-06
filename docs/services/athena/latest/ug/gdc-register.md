# Register and use data catalogs in Athena

Athena supports mounting and connecting to multiple data catalogs.

- You can mount Amazon Redshift data in the AWS Glue Data Catalog and query it from Athena without having
to copy or move data. For more information, see [Bringing\
Amazon Redshift data into the AWS Glue Data Catalog](https://docs.aws.amazon.com/lake-formation/latest/dg/managing-namespaces-datacatalog.html).

- Connect the AWS Glue Data Catalog to external data sources using AWS Glue connections, and
create federated catalogs to centrally manage permissions to the data with
fine-grained access control using Lake Formation. For more information, see [Register your connection as a Glue Data Catalog](https://docs.aws.amazon.com/athena/latest/ug/register-connection-as-gdc.html).

- Create catalogs from Amazon S3 table buckets, and use Lake Formation to centrally manage access
permissions and restrict user access to objects within the table bucket. For more
information, see [Working with Amazon S3 Tables and\
table buckets](../../../s3/latest/userguide/s3-tables.md) in the Amazon S3 User Guide.

###### Note

For any Glue catalog, you can only register a multi-level catalog like `123412341234:my_catalog/my_child`. You cannot register a single-level catalog like `123412341234:linkcontainer` or `my_catalog`. Single-level catalogs can only be queried by using the Glue data catalog directly in the Athena query. For more information, see [Query AWS Glue data catalogs in Athena](https://docs.aws.amazon.com/athena/latest/ug/gdc-register-query-the-data-source.html).

###### Topics

- [Register Redshift data catalogs in Athena](https://docs.aws.amazon.com/athena/latest/ug/gdc-register-rs.html)

- [Register federated catalogs in Athena](https://docs.aws.amazon.com/athena/latest/ug/gdc-register-federated.html)

- [Register S3 table bucket catalogs and query Tables from Athena](https://docs.aws.amazon.com/athena/latest/ug/gdc-register-s3-table-bucket-cat.html)

- [Query AWS Glue data catalogs in Athena](https://docs.aws.amazon.com/athena/latest/ug/gdc-register-query-the-data-source.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Use AWS Glue Data Catalog

Register Redshift data catalogs
