# Application performance monitoring (APM)

CloudWatch Application Signals provides application performance monitoring (APM) features such as pre-built, standardized dashboards for critical application metrics,
correlated trace spans, and a application map to enable you to visualize interactions between applications and their dependencies. You can also search and analyze
transaction spans and trace summaries to debug distributed application issues in a business context, for cases such as troubleshooting customer support tickets or
finding top impacted customers. You can also create Service Level Objectives (SLOs) to closely track the performance KPIs of critical operations in your application,
enabling you to easily identify and triage operations that do not meet your business KPIs.

See the following sections for an overview of these troubleshooting capabilities:

- [Monitor the operational health of your applications with Application Signals](services.md)

- [Searching and analyzing spans](cloudwatch-transaction-search-search-analyze-spans.md)

**Start collecting application metrics and traces**

Get the most integrated application performance monitoring experience by auto-instrumenting applications to easily collect telemetry, whether they are running
in [Amazon EKS clusters](cloudwatch-application-signals-enable-eks.md), [Amazon EC2](cloudwatch-application-signals-enable-ec2main.md),
[Amazon ECS](cloudwatch-application-signals-enable-ecsmain.md),
[Kubernetes](cloudwatch-application-signals-enable-kubernetesmain.md), [Lambda](cloudwatch-application-signals-enable-lambdamain.md), or
[on-premise](cloudwatch-application-signals-enable.md). Optionally,
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

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Using the resource health view

Application Signals

All content copied from https://docs.aws.amazon.com/.
