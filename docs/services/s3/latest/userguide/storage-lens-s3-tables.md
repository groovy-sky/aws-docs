# Working with S3 Storage Lens data in S3 Tables

Amazon S3 Storage Lens can export your storage analytics and insights to S3 Tables, enabling you to query your Storage Lens metrics using SQL with AWS analytics services like Amazon Athena, Amazon EMR, Amazon SageMaker Studio (SMStudio), and other AWS analytics tools. When you configure S3 Storage Lens to export to S3 Tables, your metrics are automatically stored in read-only Apache Iceberg tables in the AWS-managed `aws-s3` table bucket.

This integration provides structured data access for querying Storage Lens metrics using standard SQL, analytics integration with AWS analytics services, historical analysis capabilities, and cost optimization with no additional charges for exporting to AWS-managed S3 Tables.

###### Topics

- [Exporting S3 Storage Lens metrics to S3 Tables](https://docs.aws.amazon.com/AmazonS3/latest/userguide/storage-lens-s3-tables-export.html)

- [Table naming for S3 Storage Lens export to S3 Tables](https://docs.aws.amazon.com/AmazonS3/latest/userguide/storage-lens-s3-tables-naming.html)

- [Understanding S3 Storage Lens table schemas](https://docs.aws.amazon.com/AmazonS3/latest/userguide/storage-lens-s3-tables-schemas.html)

- [Permissions for S3 Storage Lens tables](https://docs.aws.amazon.com/AmazonS3/latest/userguide/storage-lens-s3-tables-permissions.html)

- [Querying S3 Storage Lens data with analytics tools](https://docs.aws.amazon.com/AmazonS3/latest/userguide/storage-lens-s3-tables-querying.html)

- [Using AI assistants with S3 Storage Lens tables](https://docs.aws.amazon.com/AmazonS3/latest/userguide/storage-lens-s3-tables-ai-tools.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

For performance

Exporting metrics to S3 Tables
