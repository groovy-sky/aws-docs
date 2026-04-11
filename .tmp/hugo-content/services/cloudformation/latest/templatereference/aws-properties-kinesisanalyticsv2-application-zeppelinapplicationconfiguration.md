This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalyticsV2::Application ZeppelinApplicationConfiguration

The configuration of a Kinesis Data Analytics Studio notebook.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CatalogConfiguration" : CatalogConfiguration,
  "CustomArtifactsConfiguration" : [ CustomArtifactConfiguration, ... ],
  "DeployAsApplicationConfiguration" : DeployAsApplicationConfiguration,
  "MonitoringConfiguration" : ZeppelinMonitoringConfiguration
}

```

### YAML

```yaml

  CatalogConfiguration:
    CatalogConfiguration
  CustomArtifactsConfiguration:
    - CustomArtifactConfiguration
  DeployAsApplicationConfiguration:
    DeployAsApplicationConfiguration
  MonitoringConfiguration:
    ZeppelinMonitoringConfiguration

```

## Properties

`CatalogConfiguration`

The Amazon Glue Data Catalog that you use in queries in a Kinesis Data Analytics
Studio notebook.

_Required_: No

_Type_: [CatalogConfiguration](aws-properties-kinesisanalyticsv2-application-catalogconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomArtifactsConfiguration`

A list of `CustomArtifactConfiguration` objects.

_Required_: No

_Type_: Array of [CustomArtifactConfiguration](aws-properties-kinesisanalyticsv2-application-customartifactconfiguration.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeployAsApplicationConfiguration`

The information required to deploy a Kinesis Data Analytics Studio notebook as an
application with durable state.

_Required_: No

_Type_: [DeployAsApplicationConfiguration](aws-properties-kinesisanalyticsv2-application-deployasapplicationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MonitoringConfiguration`

The monitoring configuration of a Kinesis Data Analytics Studio notebook.

_Required_: No

_Type_: [ZeppelinMonitoringConfiguration](aws-properties-kinesisanalyticsv2-application-zeppelinmonitoringconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VpcConfiguration

ZeppelinMonitoringConfiguration

All content copied from https://docs.aws.amazon.com/.
