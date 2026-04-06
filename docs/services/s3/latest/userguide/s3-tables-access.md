# Accessing table data

There are multiple ways to access tables in Amazon S3 table buckets. You can integrate tables with AWS analytics services using AWS Glue Data Catalog, or access tables directly using the Amazon S3 Tables Iceberg REST endpoint or the Amazon S3 Tables Catalog for Apache Iceberg. The access method you use will depend on your catalog setup, governance model, and access control needs. The following is an overview of these access methods.

**AWS Glue Data Catalog integration**

This is the recommended access method for working with tables in S3 table buckets. This integration gives you a unified view of your data estate across multiple AWS analytics services through the AWS Glue Data Catalog. After integration, you can query tables using services such as Athena and Amazon Redshift. Access to tables is managed using IAM permissions. To access tables using this integration, the IAM identity you use needs access to your S3 Tables resources and actions, AWS Glue Data Catalog objects, and the query engine you're using. For more information, see [Access management for S3 Tables](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-setting-up.html).

**Direct access**

Use this method if you need to work with AWS Partner Network (APN) catalog implementations, custom catalog implementations, or if you only need to perform basic read/write operations on tables within a single table bucket. Access to tables is managed using IAM permissions. To access tables, the IAM identity you use needs access to your table resources and S3 Tables actions. For more information, see [Access management for S3 Tables](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-setting-up.html).

## Accessing tables through the AWS Glue Data Catalog integration

You can integrate S3 table buckets with AWS Glue Data Catalog to access tables from AWS analytics services, such as Amazon Athena, Amazon Redshift, and Quick. The integration populates the AWS Glue Data Catalog with your table resources, and federates access to those resources. For more information on integrating, see [Integrating Amazon S3 Tables with AWS analytics services](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-integrating-aws.html).

The following AWS analytics services can access tables through this integration:

- [Amazon Athena](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-integrating-athena.html)

- [Amazon Redshift](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-integrating-redshift.html)

- [Amazon EMR](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-integrating-emr.html)

- [Quick](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-integrating-quicksight.html)

- [Amazon Data Firehose](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-integrating-firehose.html)

- [AWS Glue\
ETL](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-integrating-glue.html)

- [Querying S3 Tables with SageMaker Unified Studio](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-integrating-sagemaker.html)

### Accessing tables using the AWS Glue Iceberg REST endpoint

Once your S3 table buckets are integrated with AWS Glue Data Catalog, you can also use the AWS Glue Iceberg REST endpoint to connect to S3 tables from third-party query engines that support Iceberg. For more information, see [Accessing Amazon S3 tables using the AWS Glue Iceberg REST endpoint](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-integrating-glue-endpoint.html).

We recommend using the AWS Glue Iceberg REST endpoint when you want to access tables from Spark, PyIceberg, or other Iceberg-compatible clients.

The following clients can access tables directly through the AWS Glue Iceberg REST endpoint:

- Any Iceberg client, including Spark, PyIceberg, and more.

## Accessing tables directly

You can access tables directly from open source query engines through methods that bridge S3 Tables management operations to your Apache Iceberg analytics applications. There are two direct access methods: the Amazon S3 Tables Iceberg REST endpoint or the Amazon S3 Tables Catalog for Apache Iceberg. The REST endpoint is recommended.

We recommend direct access if you access tables in self-managed catalog implementations, or only need to perform basic read/write operations on tables in a single table bucket. For other access scenarios, we recommend the AWS Glue Data Catalog integration.

Direct access to tables is managed through either IAM identity-based policies or resource-based policies attached to tables and table buckets.

### Accessing tables through the Amazon S3 Tables Iceberg REST endpoint

You can use the Amazon S3 Tables Iceberg REST endpoint to access your tables directly from any Iceberg REST compatible clients through HTTP endpoints, for more information, see [Accessing tables using the Amazon S3 Tables Iceberg REST endpoint](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-integrating-open-source.html).

The following AWS analytics services and query engines can access tables directly using the Amazon S3 Tables Iceberg REST endpoint:

###### Supported query engines

- Any Iceberg client, including Spark, PyIceberg, and more.

- [Amazon EMR](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-integrating-emr.html)

- [AWS Glue\
ETL](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-integrating-glue.html)

### Accessing tables directly through the Amazon S3 Tables Catalog for Apache Iceberg

You can also access tables directly from query engines like Apache Spark by using the S3 Tables client catalog, for more information, see [Accessing Amazon S3 tables with the Amazon S3 Tables Catalog for Apache Iceberg](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-client-catalog.html). However, S3 recommends using the Amazon S3 Tables Iceberg REST endpoint for direct access because it supports more applications, without requiring language or engine-specific code.

The following query engines can access tables directly using the client catalog:

- [Apache Spark](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-client-catalog.html#s3-tables-integrating-open-source-spark)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Deleting a tag from a table

S3 Tables integration overview
