This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::Connection RedshiftPropertiesInput

The Amazon Redshift properties.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Credentials" : RedshiftCredentials,
  "DatabaseName" : String,
  "Host" : String,
  "LineageSync" : RedshiftLineageSyncConfigurationInput,
  "Port" : Number,
  "Storage" : RedshiftStorageProperties
}

```

### YAML

```yaml

  Credentials:
    RedshiftCredentials
  DatabaseName: String
  Host: String
  LineageSync:
    RedshiftLineageSyncConfigurationInput
  Port: Number
  Storage:
    RedshiftStorageProperties

```

## Properties

`Credentials`

The Amaon Redshift credentials.

_Required_: No

_Type_: [RedshiftCredentials](aws-properties-datazone-connection-redshiftcredentials.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatabaseName`

The Amazon Redshift database name.

_Required_: No

_Type_: String

_Pattern_: `^[a-z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Host`

The Amazon Redshift host.

_Required_: No

_Type_: String

_Pattern_: `^[\S]*$`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LineageSync`

The lineage sync of the Amazon Redshift.

_Required_: No

_Type_: [RedshiftLineageSyncConfigurationInput](aws-properties-datazone-connection-redshiftlineagesyncconfigurationinput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Port`

The Amaon Redshift port.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Storage`

The Amazon Redshift storage.

_Required_: No

_Type_: [RedshiftStorageProperties](aws-properties-datazone-connection-redshiftstorageproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RedshiftLineageSyncConfigurationInput

RedshiftStorageProperties

All content copied from https://docs.aws.amazon.com/.
