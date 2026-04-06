# Automating log analysis with scheduled queries

Scheduled queries enable you to automate the execution of CloudWatch Logs Insights queries on a regular
schedule. Instead of manually running queries to analyze your log data, you can configure
scheduled queries to run automatically and deliver results to destinations such as Amazon S3
buckets or Amazon EventBridge event buses. This automation is ideal for generating regular reports,
monitoring trends, or triggering downstream processes based on log analysis results.

Scheduled queries support all three query languages available in CloudWatch Logs Insights:

- [Logs Insights query\
language (Logs Insights QL)](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_AnalyzeLogData_LogsInsights.html)

- [OpenSearch Service Piped Processing Language\
(PPL)](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_AnalyzeLogData_PPL.html)

- [OpenSearch Service Structured Query Language\
(SQL)](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_AnalyzeLogData_SQL.html)

###### Contents

- [Understanding scheduled queries concepts](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/scheduled-queries-concepts.html)

- [Schedule expression reference](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/scheduled-queries-schedule-reference.html)

- [Best practices](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/scheduled-queries-best-practices.html)

- [Getting started with scheduled queries](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/scheduled-queries-getting-started.html)

- [Configuring S3 destinations for scheduled queries](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/scheduled-queries-s3-destination.html)

- [Troubleshooting scheduled queries](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/scheduled-queries-troubleshooting.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Generate a natural language summary from CloudWatch Logs Insights query results

Understanding scheduled queries concepts
