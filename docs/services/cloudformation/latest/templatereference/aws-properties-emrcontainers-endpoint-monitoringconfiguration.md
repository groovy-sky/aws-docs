This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMRContainers::Endpoint MonitoringConfiguration

Configuration setting for monitoring.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CloudWatchMonitoringConfiguration" : CloudWatchMonitoringConfiguration,
  "ContainerLogRotationConfiguration" : ContainerLogRotationConfiguration,
  "PersistentAppUI" : String,
  "S3MonitoringConfiguration" : S3MonitoringConfiguration
}

```

### YAML

```yaml

  CloudWatchMonitoringConfiguration:
    CloudWatchMonitoringConfiguration
  ContainerLogRotationConfiguration:
    ContainerLogRotationConfiguration
  PersistentAppUI: String
  S3MonitoringConfiguration:
    S3MonitoringConfiguration

```

## Properties

`CloudWatchMonitoringConfiguration`

Monitoring configurations for CloudWatch.

_Required_: No

_Type_: [CloudWatchMonitoringConfiguration](aws-properties-emrcontainers-endpoint-cloudwatchmonitoringconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ContainerLogRotationConfiguration`

Enable or disable container log rotation.

_Required_: No

_Type_: [ContainerLogRotationConfiguration](aws-properties-emrcontainers-endpoint-containerlogrotationconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PersistentAppUI`

Monitoring configurations for the persistent application UI.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3MonitoringConfiguration`

Amazon S3 configuration for monitoring log publishing.

_Required_: No

_Type_: [S3MonitoringConfiguration](aws-properties-emrcontainers-endpoint-s3monitoringconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EMREKSConfiguration

S3MonitoringConfiguration

All content copied from https://docs.aws.amazon.com/.
