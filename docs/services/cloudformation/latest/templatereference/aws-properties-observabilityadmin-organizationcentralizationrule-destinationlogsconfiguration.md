This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ObservabilityAdmin::OrganizationCentralizationRule DestinationLogsConfiguration

Configuration for centralization destination log groups, including encryption and backup
settings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BackupConfiguration" : LogsBackupConfiguration,
  "LogGroupNameConfiguration" : LogGroupNameConfiguration,
  "LogsEncryptionConfiguration" : LogsEncryptionConfiguration
}

```

### YAML

```yaml

  BackupConfiguration:
    LogsBackupConfiguration
  LogGroupNameConfiguration:
    LogGroupNameConfiguration
  LogsEncryptionConfiguration:
    LogsEncryptionConfiguration

```

## Properties

`BackupConfiguration`

Configuration defining the backup region and an optional KMS key for the backup
destination.

_Required_: No

_Type_: [LogsBackupConfiguration](aws-properties-observabilityadmin-organizationcentralizationrule-logsbackupconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogGroupNameConfiguration`

Configuration that specifies a naming pattern for destination log groups created during centralization.
The pattern supports static text and dynamic variables that are replaced with source attributes
when log groups are created.

_Required_: No

_Type_: [LogGroupNameConfiguration](aws-properties-observabilityadmin-organizationcentralizationrule-loggroupnameconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogsEncryptionConfiguration`

The encryption configuration for centralization destination log groups.

_Required_: No

_Type_: [LogsEncryptionConfiguration](aws-properties-observabilityadmin-organizationcentralizationrule-logsencryptionconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CentralizationRuleSource

LogGroupNameConfiguration

All content copied from https://docs.aws.amazon.com/.
