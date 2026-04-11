This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTSiteWise::Asset AssetProperty

Contains asset property information.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Alias" : String,
  "ExternalId" : String,
  "Id" : String,
  "LogicalId" : String,
  "NotificationState" : String,
  "Unit" : String
}

```

### YAML

```yaml

  Alias: String
  ExternalId: String
  Id: String
  LogicalId: String
  NotificationState: String
  Unit: String

```

## Properties

`Alias`

The alias that identifies the property, such as an OPC-UA server data stream path (for
example, `/company/windfarm/3/turbine/7/temperature`). For more information,
see [Mapping industrial\
data streams to asset properties](../../../iot-sitewise/latest/userguide/connect-data-streams.md) in the _AWS IoT SiteWise User_
_Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExternalId`

The external ID of the property. For more information, see [Using external IDs](../../../iot-sitewise/latest/userguide/object-ids.md#external-ids) in the _AWS IoT SiteWise User Guide_.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9_][a-zA-Z_\-0-9.:]*[a-zA-Z0-9_]+`

_Minimum_: `2`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Id`

The ID of the asset property.

###### Note

This is a return value and can't be set.

_Required_: No

_Type_: String

_Pattern_: `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`

_Minimum_: `36`

_Maximum_: `36`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogicalId`

The `LogicalID` of the asset property.

_Required_: No

_Type_: String

_Pattern_: `[^\u0000-\u001F\u007F]+`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NotificationState`

The MQTT notification state (enabled or disabled) for this asset property.
When the notification state is enabled, AWS IoT SiteWise publishes property value
updates to a unique MQTT topic. For more information, see [Interacting with other services](../../../iot-sitewise/latest/userguide/interact-with-other-services.md) in the _AWS IoT SiteWise User Guide_.

If you omit this parameter, the notification state is set to `DISABLED`.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Unit`

The unit (such as `Newtons` or `RPM`) of the asset property.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AssetHierarchy

Tag

All content copied from https://docs.aws.amazon.com/.
