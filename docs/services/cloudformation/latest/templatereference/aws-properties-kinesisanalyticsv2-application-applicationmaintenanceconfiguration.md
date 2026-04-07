This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalyticsV2::Application ApplicationMaintenanceConfiguration

Specifies the maintenance configuration for a Amazon Managed Service for Apache Flink.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApplicationMaintenanceWindowStartTime" : String
}

```

### YAML

```yaml

  ApplicationMaintenanceWindowStartTime: String

```

## Properties

`ApplicationMaintenanceWindowStartTime`

The UTC timestamp of a day from which the eight-hour maintenance window will begin every day of the week. Maintenance of the application happens only during this eight-hour window.

_Required_: Yes

_Type_: String

_Pattern_: `^([01][0-9]|2[0-3]):[0-5][0-9]$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ApplicationEncryptionConfiguration

ApplicationRestoreConfiguration
