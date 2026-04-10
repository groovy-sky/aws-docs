This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMRServerless::Application MonitoringConfiguration

The configuration setting for monitoring logs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CloudWatchLoggingConfiguration" : CloudWatchLoggingConfiguration,
  "ManagedPersistenceMonitoringConfiguration" : ManagedPersistenceMonitoringConfiguration,
  "PrometheusMonitoringConfiguration" : PrometheusMonitoringConfiguration,
  "S3MonitoringConfiguration" : S3MonitoringConfiguration
}

```

### YAML

```yaml

  CloudWatchLoggingConfiguration:
    CloudWatchLoggingConfiguration
  ManagedPersistenceMonitoringConfiguration:
    ManagedPersistenceMonitoringConfiguration
  PrometheusMonitoringConfiguration:
    PrometheusMonitoringConfiguration
  S3MonitoringConfiguration:
    S3MonitoringConfiguration

```

## Properties

`CloudWatchLoggingConfiguration`

The Amazon CloudWatch configuration for monitoring logs. You can configure your jobs
to send log information to CloudWatch.

_Required_: No

_Type_: [CloudWatchLoggingConfiguration](aws-properties-emrserverless-application-cloudwatchloggingconfiguration.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`ManagedPersistenceMonitoringConfiguration`

The managed log persistence configuration for a job run.

_Required_: No

_Type_: [ManagedPersistenceMonitoringConfiguration](aws-properties-emrserverless-application-managedpersistencemonitoringconfiguration.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`PrometheusMonitoringConfiguration`

The monitoring configuration object you can configure to send metrics to Amazon Managed Service for Prometheus for a job run.

_Required_: No

_Type_: [PrometheusMonitoringConfiguration](aws-properties-emrserverless-application-prometheusmonitoringconfiguration.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`S3MonitoringConfiguration`

The Amazon S3 configuration for monitoring log publishing.

_Required_: No

_Type_: [S3MonitoringConfiguration](aws-properties-emrserverless-application-s3monitoringconfiguration.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MaximumAllowedResources

NetworkConfiguration

All content copied from https://docs.aws.amazon.com/.
