# Use AWS Glue Data Catalog to connect to your data

Athena uses the AWS Glue Data Catalog to store metadata such as table and column names for your data
stored in Amazon S3. This metadata information becomes the databases, tables, and views that you
see in the Athena query editor.

When using Athena with the AWS Glue Data Catalog, you can use AWS Glue to create databases and tables
(schema) to be queried in Athena, or you can use Athena to create schema and then use them in
AWS Glue and related services.

To define schema information for AWS Glue, you can use a form in the Athena console, use the
query editor in Athena, or create an AWS Glue crawler in the AWS Glue console. AWS Glue crawlers
automatically infer database and table schema from your data in Amazon S3. Using a form offers
more customization. Writing your own `CREATE TABLE` statements requires more
effort, but offers the most control. For more information, see [CREATE TABLE](https://docs.aws.amazon.com/athena/latest/ug/create-table.html).

## Additional Resources

- For more information about the AWS Glue Data Catalog, see [Data Catalog and crawlers in\
AWS Glue](https://docs.aws.amazon.com/glue/latest/dg/catalog-and-crawler.html) in the _AWS Glue Developer Guide_.

- For an illustrative article showing how to use AWS Glue and Athena to process XML
data, see [Process and analyze highly nested and large XML files using AWS Glue and\
Amazon Athena](https://aws.amazon.com/blogs/big-data/process-and-analyze-highly-nested-and-large-xml-files-using-aws-glue-and-amazon-athena) in the AWS Big Data Blog.

- Separate charges apply to AWS Glue. For more information, see [AWS Glue pricing](https://aws.amazon.com/glue/pricing).

###### Topics

- [Register and use data catalogs in Athena](https://docs.aws.amazon.com/athena/latest/ug/gdc-register.html)

- [Register a Data Catalog from another account](https://docs.aws.amazon.com/athena/latest/ug/data-sources-glue-cross-account.html)

- [Control access to data catalogs with IAM policies](https://docs.aws.amazon.com/athena/latest/ug/datacatalogs-iam-policy.html)

- [Use a form in the Athena console to add an AWS Glue table](https://docs.aws.amazon.com/athena/latest/ug/data-sources-glue-manual-table.html)

- [Use a crawler to add a table](https://docs.aws.amazon.com/athena/latest/ug/schema-crawlers.html)

- [Optimize queries with AWS Glue partition indexing and filtering](https://docs.aws.amazon.com/athena/latest/ug/glue-best-practices-partition-index.html)

- [Use the AWS CLI to recreate an AWS Glue database and its tables](https://docs.aws.amazon.com/athena/latest/ug/glue-recreate-db-and-tables-cli.html)

- [Create tables for ETL jobs](https://docs.aws.amazon.com/athena/latest/ug/schema-classifier.html)

- [Work with CSV data in AWS Glue](https://docs.aws.amazon.com/athena/latest/ug/schema-csv.html)

- [Work with geospatial data in AWS Glue](https://docs.aws.amazon.com/athena/latest/ug/schema-geospatial.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Connect to data sources

Register and use data catalogs
