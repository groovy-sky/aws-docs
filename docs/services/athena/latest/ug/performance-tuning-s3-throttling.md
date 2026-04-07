# Prevent Amazon S3 throttling

Throttling is the process of limiting the rate at which you use a service, an application,
or a system. In AWS, you can use throttling to prevent overuse of the Amazon S3 service and
increase the availability and responsiveness of Amazon S3 for all users. However, because
throttling limits the rate at which the data can be transferred to or from Amazon S3, it's
important to consider preventing your interactions from being throttled.

As pointed out in the [performance tuning](https://docs.aws.amazon.com/athena/latest/ug/performance-tuning.html) chapter, optimizations can
depend on your service level decisions, on how you structure your tables and data, and on
how you write your queries.

###### Topics

- [Reduce throttling at the service level](https://docs.aws.amazon.com/athena/latest/ug/performance-tuning-s3-throttling-reduce-throttling-at-the-service-level.html)

- [Optimize your tables](https://docs.aws.amazon.com/athena/latest/ug/performance-tuning-s3-throttling-optimizing-your-tables.html)

- [Optimize your queries](https://docs.aws.amazon.com/athena/latest/ug/performance-tuning-s3-throttling-optimizing-queries.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Use custom prefixes

Reduce throttling at the service level
