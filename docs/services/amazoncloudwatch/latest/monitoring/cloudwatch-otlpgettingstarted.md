# Getting started

To get started with OpenTelemetry in CloudWatch, you can use the pre-packaged OpenTelemetry setup that is available with the
CloudWatch agent along with the AWS Distro for OpenTelemetry SDKs. This gives you the most integrated monitoring experience in CloudWatch.

###### Note

Make sure Transaction Search is enabled before you use the OTLP Endpoint for traces.

Alternatively, you have the flexibility to use the OpenTelemetry Collector or your own custom OpenTelemetry Collector to directly send telemetry
to the OTLP endpoint. You can use the AWS Distro for OpenTelemetry to go collector-less and to send telemetry directly to the OTLP endpoint. Make an informed choice based on the feature support:

FeatureOpenTelemetry CollectorCustom OpenTelemetry CollectorAWS Distro for OpenTelemetry

CloudWatch application signals (Application performance metrics, service discovery, and application map)

Yes

Yes

Yes

Search and analyze spans and trace summaries

Yes

Yes

Yes

Search and analyze logs summaries

Yes

Yes

Yes

Application performance monitoring telemetry enrichment with AWS infrastructure attributes that your application is hosted in.

No

Yes

Yes

Runtime metrics correlated with your application. For example, JVM metrics

No

Yes

No

AWS Support

Data received by AWS

Data received by AWS

Data received by AWS

Telemetry supported

Logs, Metrics, Traces

Logs, Traces, Metrics

Metrics, Traces

###### Topics

- [OpenTelemetry Collector](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-OTLPSimplesetup.html)

- [Build your own custom OpenTelemetry Collector](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-OTLPAdvancedsetup.html)

- [Exporting collector-less telemetry using AWS Distro for OpenTelemetry (ADOT) SDK](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-OTLP-UsingADOT.html)

- [Enabling vended metrics in PromQL](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-OTelEnrichment.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

OTLP Endpoints

OpenTelemetry Collector
