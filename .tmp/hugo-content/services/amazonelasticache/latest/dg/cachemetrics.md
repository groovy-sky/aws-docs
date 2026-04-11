# Monitoring use with CloudWatch Metrics

ElastiCache provides metrics that enable you to monitor your clusters. You can access
these metrics through CloudWatch. For more information on CloudWatch, see the [CloudWatch documentation.](https://aws.amazon.com/documentation/cloudwatch)

ElastiCache provides both host-level metrics (for example, CPU usage) and metrics that are
specific to the cache engine software (for example, cache gets and cache misses). These
metrics are measured and published for each Cache node in 60-second intervals.

###### Important

You should consider setting CloudWatch alarms on certain key metrics, so that you will be notified if your cluster's performance starts to degrade.
For more information, see [Which Metrics Should I Monitor?](cachemetrics-whichshouldimonitor.md) in this guide.

###### Topics

- [Host-Level Metrics](cachemetrics-hostlevel.md)

- [Metrics for Valkey and Redis OSS](cachemetrics-redis.md)

- [Metrics for Memcached](cachemetrics-memcached.md)

- [Which Metrics Should I Monitor?](cachemetrics-whichshouldimonitor.md)

- [Choosing Metric Statistics and Periods](cachemetrics-choosingstatisticsandperiods.md)

- [Monitoring CloudWatch Cluster and Node Metrics](cloudwatchmetrics.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Specifying log delivery using the AWS CLI

Host-Level Metrics

All content copied from https://docs.aws.amazon.com/.
