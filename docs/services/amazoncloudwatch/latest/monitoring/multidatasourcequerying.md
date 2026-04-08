# Query metrics from other data sources

You can use CloudWatch to query, visualize, and create alarms for metrics from other data sources.
To do so, you connect CloudWatch to the other data sources. This gives you a single, consolidated
monitoring experience within the CloudWatch console. You can have a unified view of your
infrastructure and application metrics regardless of where the data is stored, helping you
identify and resolve issues faster.

After you connect to a data source using a CloudWatch wizard, CloudWatch creates an AWS CloudFormation stack that
deploys and configures an AWS Lambda function. This Lambda function runs on demand every time you
query the data source. The CloudWatch query builder shows you in real time a list of elements that can
be queried, such as metrics, tables, fields, or labels. As you make choices, the query builder
pre-populates a query in the native language of the selected source.

CloudWatch provides guided wizards for you to connect to the following data sources. For these
data sources, you provide basic information to identify the data source and credentials. You can
also manually create connectors to other data sources by creating your own Lambda
functions.

- Amazon OpenSearch Service– Derive metrics from your OpenSearch Service logs and traces.

- Amazon Managed Service for Prometheus– Query these metrics using PromQL.

- Amazon RDS for MySQL – Use SQL to convert data stored in your Amazon RDS tables into metrics.

- Amazon RDS for PostgreSQL– Use SQL to convert data stored in your Amazon RDS tables into metrics.

- Amazon S3 CSV files– Display metrics data from a CSV file stored in an Amazon S3 bucket.

- Microsoft Azure Monitor– Query metrics from your Microsoft Azure Monitor
account.

- Prometheus– Query these metrics using PromQL.

After you create connectors to data sources, see [Creating a graph of metrics from another data source](graph-a-metric.md#create-metric-graph-multidatasource) for information about graphing a metric
from a data source. For information about setting an alarm on a metric from a data source, see
[Create an alarm based on a connected data source](create-multisource-alarm.md).

###### Topics

- [Managing access to data sources](cloudwatch-multidatasources-permissions.md)

- [Connect to a prebuilt data source with a wizard](cloudwatch-multidatasources-connect.md)

- [Create a custom connector to a data source](cloudwatch-multidatasources-connect-custom.md)

- [Use your custom data source](cloudwatch-multidatasources-custom-use.md)

- [Delete a connector to a data source](cloudwatch-multidatasources-delete.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Add related information to custom telemetry

Managing access to data sources

All content copied from https://docs.aws.amazon.com/.
