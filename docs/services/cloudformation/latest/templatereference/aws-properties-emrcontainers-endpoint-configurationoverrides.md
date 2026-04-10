This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMRContainers::Endpoint ConfigurationOverrides

A configuration specification to be used to override existing configurations.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApplicationConfiguration" : [ EMREKSConfiguration, ... ],
  "MonitoringConfiguration" : MonitoringConfiguration
}

```

### YAML

```yaml

  ApplicationConfiguration:
    - EMREKSConfiguration
  MonitoringConfiguration:
    MonitoringConfiguration

```

## Properties

`ApplicationConfiguration`

The configurations for the application running by the job run.

_Required_: No

_Type_: Array of [EMREKSConfiguration](aws-properties-emrcontainers-endpoint-emreksconfiguration.md)

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MonitoringConfiguration`

The configurations for monitoring.

_Required_: No

_Type_: [MonitoringConfiguration](aws-properties-emrcontainers-endpoint-monitoringconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CloudWatchMonitoringConfiguration

ContainerLogRotationConfiguration

All content copied from https://docs.aws.amazon.com/.
