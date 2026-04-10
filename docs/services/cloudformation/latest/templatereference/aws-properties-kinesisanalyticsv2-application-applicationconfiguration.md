This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalyticsV2::Application ApplicationConfiguration

Specifies the creation parameters for a Managed Service for Apache Flink application.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApplicationCodeConfiguration" : ApplicationCodeConfiguration,
  "ApplicationEncryptionConfiguration" : ApplicationEncryptionConfiguration,
  "ApplicationSnapshotConfiguration" : ApplicationSnapshotConfiguration,
  "ApplicationSystemRollbackConfiguration" : ApplicationSystemRollbackConfiguration,
  "EnvironmentProperties" : EnvironmentProperties,
  "FlinkApplicationConfiguration" : FlinkApplicationConfiguration,
  "SqlApplicationConfiguration" : SqlApplicationConfiguration,
  "VpcConfigurations" : [ VpcConfiguration, ... ],
  "ZeppelinApplicationConfiguration" : ZeppelinApplicationConfiguration
}

```

### YAML

```yaml

  ApplicationCodeConfiguration:
    ApplicationCodeConfiguration
  ApplicationEncryptionConfiguration:
    ApplicationEncryptionConfiguration
  ApplicationSnapshotConfiguration:
    ApplicationSnapshotConfiguration
  ApplicationSystemRollbackConfiguration:
    ApplicationSystemRollbackConfiguration
  EnvironmentProperties:
    EnvironmentProperties
  FlinkApplicationConfiguration:
    FlinkApplicationConfiguration
  SqlApplicationConfiguration:
    SqlApplicationConfiguration
  VpcConfigurations:
    - VpcConfiguration
  ZeppelinApplicationConfiguration:
    ZeppelinApplicationConfiguration

```

## Properties

`ApplicationCodeConfiguration`

The code location and type parameters for a Managed Service for Apache Flink application.

_Required_: Conditional

_Type_: [ApplicationCodeConfiguration](aws-properties-kinesisanalyticsv2-application-applicationcodeconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ApplicationEncryptionConfiguration`

The configuration to manage encryption at rest.

_Required_: No

_Type_: [ApplicationEncryptionConfiguration](aws-properties-kinesisanalyticsv2-application-applicationencryptionconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ApplicationSnapshotConfiguration`

Describes whether snapshots are enabled for a Managed Service for Apache Flink application.

_Required_: No

_Type_: [ApplicationSnapshotConfiguration](aws-properties-kinesisanalyticsv2-application-applicationsnapshotconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ApplicationSystemRollbackConfiguration`

Describes whether system rollbacks are enabled for a Managed Service for Apache Flink application.

_Required_: No

_Type_: [ApplicationSystemRollbackConfiguration](aws-properties-kinesisanalyticsv2-application-applicationsystemrollbackconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnvironmentProperties`

Describes execution properties for a Managed Service for Apache Flink application.

_Required_: No

_Type_: [EnvironmentProperties](aws-properties-kinesisanalyticsv2-application-environmentproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FlinkApplicationConfiguration`

The creation and update parameters for a Managed Service for Apache Flink application.

_Required_: No

_Type_: [FlinkApplicationConfiguration](aws-properties-kinesisanalyticsv2-application-flinkapplicationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SqlApplicationConfiguration`

The creation and update parameters for a SQL-based Kinesis Data Analytics application.

_Required_: No

_Type_: [SqlApplicationConfiguration](aws-properties-kinesisanalyticsv2-application-sqlapplicationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcConfigurations`

The array of descriptions of VPC configurations available to the application.

_Required_: No

_Type_: Array of [VpcConfiguration](aws-properties-kinesisanalyticsv2-application-vpcconfiguration.md)

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ZeppelinApplicationConfiguration`

The configuration parameters for a Kinesis Data Analytics Studio notebook.

_Required_: No

_Type_: [ZeppelinApplicationConfiguration](aws-properties-kinesisanalyticsv2-application-zeppelinapplicationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [ApplicationConfiguration](../../../managed-flink/latest/apiv2/api-applicationconfiguration.md) in the _Amazon Kinesis Data_
_Analytics API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ApplicationCodeConfiguration

ApplicationEncryptionConfiguration

All content copied from https://docs.aws.amazon.com/.
