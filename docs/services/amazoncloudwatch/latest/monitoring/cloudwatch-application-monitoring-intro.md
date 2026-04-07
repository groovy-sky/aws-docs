# Application performance monitoring (APM)

CloudWatch Application Signals provides application performance monitoring (APM) features such as pre-built, standardized dashboards for critical application metrics,
correlated trace spans, and a application map to enable you to visualize interactions between applications and their dependencies. You can also search and analyze
transaction spans and trace summaries to debug distributed application issues in a business context, for cases such as troubleshooting customer support tickets or
finding top impacted customers. You can also create Service Level Objectives (SLOs) to closely track the performance KPIs of critical operations in your application,
enabling you to easily identify and triage operations that do not meet your business KPIs.

See the following sections for an overview of these troubleshooting capabilities:

- [Monitor the operational health of your applications with Application Signals](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Services.html)

- [Searching and analyzing spans](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Transaction-Search-search-analyze-spans.html)

**Start collecting application metrics and traces**

Get the most integrated application performance monitoring experience by auto-instrumenting applications to easily collect telemetry, whether they are running
in [Amazon EKS clusters](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Application-Signals-Enable-EKS.html), [Amazon EC2](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Application-Signals-Enable-EC2Main.html),
[Amazon ECS](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Application-Signals-Enable-ECSMain.html),
[Kubernetes](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Application-Signals-Enable-KubernetesMain.html), [Lambda](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Application-Signals-Enable-LambdaMain.html), or
[on-premise](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Application-Signals-Enable.html). Optionally,
you can also use [OpenTelemetry](cloudwatch-opentelemetry-sections.md) with Application Signals to collect telemetry.

###### Note

You must enable transaction search to get all APM features along with a new unified pricing for CloudWatch Application Signals, inclusive of X-Ray traces and application
transaction spans. For more information about pricing, see [Amazon CloudWatch Pricing](https://aws.amazon.com/cloudwatch/pricing).

###### Topics

- [Application Signals](cloudwatch-application-monitoring-sections.md)

- [Service level objectives (SLOs)](cloudwatch-servicelevelobjectives.md)

- [Transaction Search](cloudwatch-transaction-search.md)

- [Synthetic monitoring (canaries)](cloudwatch-synthetics-canaries.md)

- [CloudWatch RUM](cloudwatch-rum.md)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Using the resource health view

Application Signals
