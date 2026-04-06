# Query Amazon CloudFront logs

You can configure Amazon CloudFront CDN to export Web distribution access logs to Amazon Simple Storage Service. Use
these logs to explore users' surfing patterns across your web properties served by
CloudFront.

Before you begin querying the logs, enable Web distributions access log on your preferred
CloudFront distribution. For information, see [Access\
logs](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/AccessLogs.html) in the _Amazon CloudFront Developer Guide_. Make a note
of the Amazon S3 bucket in which you save these logs.

###### Topics

- [Create a table for CloudFront standard logs (legacy)](https://docs.aws.amazon.com/athena/latest/ug/create-cloudfront-table-standard-logs.html)

- [Create a table for CloudFront logs in Athena using manual partitioning with JSON](https://docs.aws.amazon.com/athena/latest/ug/create-cloudfront-table-manual-json.html)

- [Create a table for CloudFront logs in Athena using manual partitioning with Parquet](https://docs.aws.amazon.com/athena/latest/ug/create-cloudfront-table-manual-parquet.html)

- [Create a table for CloudFront logs in Athena using partition projection with JSON](https://docs.aws.amazon.com/athena/latest/ug/create-cloudfront-table-partition-json.html)

- [Create a table for CloudFront logs in Athena using partition projection with Parquet](https://docs.aws.amazon.com/athena/latest/ug/create-cloudfront-table-partition-parquet.html)

- [Create a table for CloudFront real-time logs](https://docs.aws.amazon.com/athena/latest/ug/create-cloudfront-table-real-time-logs.html)

- [Additional resources](https://docs.aws.amazon.com/athena/latest/ug/cloudfront-logs-additional-resources.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Elastic Load Balancing

Standard logs (legacy)
