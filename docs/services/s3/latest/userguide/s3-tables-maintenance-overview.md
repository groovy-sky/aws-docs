# S3 Tables maintenance

Amazon S3 automatically performs maintenance to enhance the performance of your tables in S3
table buckets. Maintenance is performed at the table bucket and individual table level and
includes the following:

**Table bucket-level maintenance:**

- **Unreferenced file removal** –
Cleans up orphaned files to optimize storage usage and costs.

**Table-level maintenance:**

- **File compaction** – Consolidates
small files to improve query performance and reduce storage
costs.

- **Snapshot management** – Controls
table version history and prevents excessive metadata growth.

These options are enabled by default. You can edit or disable these operations through
maintenance configuration files.

In addition to these options, you can also enable and configure record expiration settings
for tables. With this option, Amazon S3 automatically removes records from a table when the
records expire.

###### Topics

- [S3 Tables maintenance job status](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-maintenance-status.html)

- [Maintenance for table buckets](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-table-buckets-maintenance.html)

- [Maintenance for tables](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-maintenance.html)

- [Record expiration for tables](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-record-expiration.html)

- [Considerations and limitations for maintenance jobs](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-considerations.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Deleting a tag from a table bucket

S3 Tables maintenance job status
