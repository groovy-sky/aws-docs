This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::Connection RedshiftLineageSyncConfigurationInput

The Amaon Redshift lineage sync configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enabled" : Boolean,
  "Schedule" : LineageSyncSchedule
}

```

### YAML

```yaml

  Enabled: Boolean
  Schedule:
    LineageSyncSchedule

```

## Properties

`Enabled`

Specifies whether the Amaon Redshift lineage sync configuration is enabled.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Schedule`

The schedule of the Amaon Redshift lineage sync configuration.

_Required_: No

_Type_: [LineageSyncSchedule](aws-properties-datazone-connection-lineagesyncschedule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RedshiftCredentials

RedshiftPropertiesInput

All content copied from https://docs.aws.amazon.com/.
