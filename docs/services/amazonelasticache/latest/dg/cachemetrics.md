# Monitoring use with CloudWatch Metrics

ElastiCache provides metrics that enable you to monitor your clusters. You can access
these metrics through CloudWatch. For more information on CloudWatch, see the [CloudWatch documentation.](https://aws.amazon.com/documentation/cloudwatch)

ElastiCache provides both host-level metrics (for example, CPU usage) and metrics that are
specific to the cache engine software (for example, cache gets and cache misses). These
metrics are measured and published for each Cache node in 60-second intervals.

###### Important

You should consider setting CloudWatch alarms on certain key metrics, so that you will be notified if your cluster's performance starts to degrade.
For more information, see [Which Metrics Should I Monitor?](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/CacheMetrics.WhichShouldIMonitor.html) in this guide.

###### Topics

- [Host-Level Metrics](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/CacheMetrics.HostLevel.html)

- [Metrics for Valkey and Redis OSS](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/CacheMetrics.Redis.html)

- [Metrics for Memcached](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/CacheMetrics.Memcached.html)

- [Which Metrics Should I Monitor?](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/CacheMetrics.WhichShouldIMonitor.html)

- [Choosing Metric Statistics and Periods](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/CacheMetrics.ChoosingStatisticsAndPeriods.html)

- [Monitoring CloudWatch Cluster and Node Metrics](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/CloudWatchMetrics.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Specifying log delivery using the AWS CLI

Host-Level Metrics
