# Query AWS service logs

This section includes several procedures for using Amazon Athena to query popular datasets,
such as AWS CloudTrail logs, Amazon CloudFront logs, Classic Load Balancer logs, Application Load Balancer
logs, Amazon VPC flow logs, and Network Load Balancer logs.

The tasks in this section use the Athena console, but you can also use other tools like the
[Athena JDBC driver](connect-with-jdbc.md), the [AWS CLI](https://docs.aws.amazon.com/cli/latest/reference/athena), or the [Amazon Athena API Reference](../../../../reference/athena/latest/apireference.md).

For information about using AWS CloudFormation to automatically create AWS service log tables,
partitions, and example queries in Athena, see [Automating AWS service logs table creation and querying them with Amazon Athena](https://aws.amazon.com/blogs/big-data/automating-aws-service-logs-table-creation-and-querying-them-with-amazon-athena)
in the AWS Big Data Blog. For information about using a Python library for AWS Glue to create
a common framework for processing AWS service logs and querying them in Athena, see [Easily\
query AWS service logs using Amazon Athena](https://aws.amazon.com/blogs/big-data/easily-query-aws-service-logs-using-amazon-athena).

The topics in this section assume that you have configured appropriate permissions to
access Athena and the Amazon S3 bucket where the data to query should reside. For more
information, see [Set up, administrative, and programmatic access](setting-up.md) and [Get started](getting-started.md).

###### Topics

- [Application Load Balancer](application-load-balancer-logs.md)

- [Elastic Load Balancing](elasticloadbalancer-classic-logs.md)

- [CloudFront](cloudfront-logs.md)

- [CloudTrail](cloudtrail-logs.md)

- [Amazon EMR](https://docs.aws.amazon.com/athena/latest/ug/emr-logs.html)

- [Global Accelerator](https://docs.aws.amazon.com/athena/latest/ug/querying-global-accelerator-flow-logs.html)

- [GuardDuty](https://docs.aws.amazon.com/athena/latest/ug/querying-guardduty.html)

- [Network Firewall](https://docs.aws.amazon.com/athena/latest/ug/querying-network-firewall-logs.html)

- [Network Load Balancer](https://docs.aws.amazon.com/athena/latest/ug/networkloadbalancer-classic-logs.html)

- [Route 53](https://docs.aws.amazon.com/athena/latest/ug/querying-r53-resolver-logs.html)

- [Amazon SES](https://docs.aws.amazon.com/athena/latest/ug/querying-ses-logs.html)

- [Amazon VPC](vpc-flow-logs.md)

- [AWS WAF](https://docs.aws.amazon.com/athena/latest/ug/waf-logs.html)

For information about querying Amazon S3 logs, see the following topics:

- [How do I analyze my Amazon S3 server access logs using Athena?](https://aws.amazon.com/premiumsupport/knowledge-center/analyze-logs-athena) in the AWS
Knowledge Center

- [Querying Amazon S3 access logs for requests using Amazon Athena](https://docs.aws.amazon.com/AmazonS3/latest/dev/using-s3-access-logs-to-identify-requests.html#querying-s3-access-logs-for-requests) in the
Amazon Simple Storage Service User Guide

- [Using AWS CloudTrail\
to identify Amazon S3 requests](https://docs.aws.amazon.com/AmazonS3/latest/dev/cloudtrail-request-identification.html) in the Amazon Simple Storage Service User Guide

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

List all columns

Application Load Balancer
