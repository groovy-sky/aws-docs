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

_Type_: [ApplicationCodeConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisanalyticsv2-application-applicationcodeconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ApplicationEncryptionConfiguration`

The configuration to manage encryption at rest.

_Required_: No

_Type_: [ApplicationEncryptionConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisanalyticsv2-application-applicationencryptionconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ApplicationSnapshotConfiguration`

Describes whether snapshots are enabled for a Managed Service for Apache Flink application.

_Required_: No

_Type_: [ApplicationSnapshotConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisanalyticsv2-application-applicationsnapshotconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ApplicationSystemRollbackConfiguration`

Describes whether system rollbacks are enabled for a Managed Service for Apache Flink application.

_Required_: No

_Type_: [ApplicationSystemRollbackConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisanalyticsv2-application-applicationsystemrollbackconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnvironmentProperties`

Describes execution properties for a Managed Service for Apache Flink application.

_Required_: No

_Type_: [EnvironmentProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisanalyticsv2-application-environmentproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FlinkApplicationConfiguration`

The creation and update parameters for a Managed Service for Apache Flink application.

_Required_: No

_Type_: [FlinkApplicationConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisanalyticsv2-application-flinkapplicationconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SqlApplicationConfiguration`

The creation and update parameters for a SQL-based Kinesis Data Analytics application.

_Required_: No

_Type_: [SqlApplicationConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisanalyticsv2-application-sqlapplicationconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcConfigurations`

The array of descriptions of VPC configurations available to the application.

_Required_: No

_Type_: Array of [VpcConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisanalyticsv2-application-vpcconfiguration.html)

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ZeppelinApplicationConfiguration`

The configuration parameters for a Kinesis Data Analytics Studio notebook.

_Required_: No

_Type_: [ZeppelinApplicationConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisanalyticsv2-application-zeppelinapplicationconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [ApplicationConfiguration](https://docs.aws.amazon.com/managed-flink/latest/apiv2/API_ApplicationConfiguration.html) in the _Amazon Kinesis Data_
_Analytics API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ApplicationCodeConfiguration

ApplicationEncryptionConfiguration
