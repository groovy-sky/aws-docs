# Configuring metadata tables

Amazon S3 Metadata accelerates data discovery by automatically capturing metadata for the objects
in your general purpose buckets and storing it in read-only, fully managed Apache
Iceberg tables that you can query. These read-only tables are called _metadata tables_. As objects are added to, updated, and removed from
your general purpose buckets, S3 Metadata automatically refreshes the corresponding metadata
tables to reflect the latest changes.

With S3 Metadata, you can easily find, store, and query metadata for your S3 objects, so
that you can quickly prepare data for use in business analytics, artificial intelligence and
machine learning (AI/ML) model training, and more.

To generate and store object metadata in AWS managed metadata tables, you create a metadata table
configuration for your general purpose bucket. Amazon S3 is designed to continuously update the metadata tables
to reflect the latest changes to your data as long as the configuration is active on the bucket.
Additionally, Amazon S3 continuously optimizes your metadata tables to help reduce storage costs and improve
analytics query performance.

To create a metadata table configuration, make sure that you have the necessary AWS Identity and Access Management (IAM)
permissions to create and manage metadata tables.

To monitor updates to your metadata table configuration, you can use AWS CloudTrail. For more information,
see [Amazon S3 bucket-level actions that are tracked by CloudTrail logging](cloudtrail-logging-s3-info.md#cloudtrail-bucket-level-tracking).

###### Topics

- [Setting up permissions for configuring metadata tables](metadata-tables-permissions.md)

- [Creating metadata table configurations](https://docs.aws.amazon.com/AmazonS3/latest/userguide/metadata-tables-create-configuration.html)

- [Controlling access to metadata tables](https://docs.aws.amazon.com/AmazonS3/latest/userguide/metadata-tables-access-control.html)

- [Expiring journal table records](metadata-tables-expire-journal-table-records.md)

- [Enabling or disabling live inventory tables](https://docs.aws.amazon.com/AmazonS3/latest/userguide/metadata-tables-enable-disable-inventory-tables.html)

- [Viewing metadata table configurations](https://docs.aws.amazon.com/AmazonS3/latest/userguide/metadata-tables-view-configuration.html)

- [Deleting metadata table configurations](https://docs.aws.amazon.com/AmazonS3/latest/userguide/metadata-tables-delete-configuration.html)

- [Deleting metadata tables](https://docs.aws.amazon.com/AmazonS3/latest/userguide/metadata-tables-delete-table.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Live inventory tables schema

Permissions for metadata tables
