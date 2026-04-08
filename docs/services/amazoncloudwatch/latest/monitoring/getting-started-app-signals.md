# Supported instrumentation setups

You can enable CloudWatch Application Signals with different instrumentation setups.
This topic describes each of the setup methods and recommendations based on the method you choose.

## Use AWS Distro for OpenTelemetry with the CloudWatch Agent

The most integrated application performance monitoring(APM) experience in CloudWatch is delivered through the AWS Distro for OpenTelemetry (ADOT) SDKs and are used with the CloudWatch Agent to collect application metrics and traces.
This option works best if you want to get started with APM in CloudWatch quickly and also leverage out-of-the box integrations with features, such as Container Insights and CloudWatch Logs.
For more information, see [Enable Application Signals on Amazon EKS Clusters](cloudwatch-application-signals-enable-eks.md) and [Enable Application Signals on Amazon EC2, Amazon ECS, or Kubernates](cloudwatch-application-signals-enable-other.md).

## Use the OpenTelemetry SDK and Collector

This setup works for the following use cases:

1. You instrumented your application or plan with OpenTelemetry SDKs and currently are using OpenTelemetry Collector.

2. You're using languages, such as Erlang and Rust, that aren't supported by AWS Distro for OpenTelemetry (ADOT).

For more information, see [OpenTelemetry with CloudWatch](cloudwatch-opentelemetry-sections.md).

## Use the AWS X-Ray SDK and daemon

This option is best if you instrumented your application using X-Ray SDKs and haven't migrated ADOT SDKs or OpenTelemetry SDKs.

For more information, see [Transaction Search](cloudwatch-transaction-search.md).

## Feature comparison

FeatureADOT SDK + CloudWatch AgentOpen Telemetry SDK + OpenTelemetry CollectorX-Ray SDKsAWS SupportYesOnly for data sent to AWSYesNonstandard language supportNoYesNoContainer Insights integrationYesNoNoOut of the box logging with CloudWatch LogsYesNoNoOut of the box runtime metricsYesNoNoAlways gets metrics on 100% of trafficYesOnly at 100% sampling rateOnly at 100% sampling rate

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Supported systems

Enable Application Signals in your account

All content copied from https://docs.aws.amazon.com/.
