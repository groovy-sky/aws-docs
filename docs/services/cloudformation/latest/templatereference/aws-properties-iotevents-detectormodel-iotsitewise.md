This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTEvents::DetectorModel IotSiteWise

Sends information about the detector model instance and the event that triggered the
action to a specified asset property in AWS IoT SiteWise.

You must use expressions for all parameters in `IotSiteWiseAction`. The
expressions accept literals, operators, functions, references, and substitutions
templates.

###### Examples

- For literal values, the expressions must contain single quotes. For example, the value
for the `propertyAlias` parameter can be
`'/company/windfarm/3/turbine/7/temperature'`.

- For references, you must specify either variables or input values. For example, the
value for the `assetId` parameter can be
`$input.TurbineInput.assetId1`.

- For a substitution template, you must use `${}`, and the template must be
in single quotes. A substitution template can also contain a combination of literals,
operators, functions, references, and substitution templates.

In the following example, the value for the `propertyAlias` parameter uses
a substitution template.

`'company/windfarm/${$input.TemperatureInput.sensorData.windfarmID}/turbine/
              ${$input.TemperatureInput.sensorData.turbineID}/temperature'`

You must specify either `propertyAlias` or both `assetId` and
`propertyId` to identify the target asset property in AWS IoT SiteWise.

For more information,
see [Expressions](../../../iotevents/latest/developerguide/iotevents-expressions.md)
in the _AWS IoT Events Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AssetId" : String,
  "EntryId" : String,
  "PropertyAlias" : String,
  "PropertyId" : String,
  "PropertyValue" : AssetPropertyValue
}

```

### YAML

```yaml

  AssetId: String
  EntryId: String
  PropertyAlias: String
  PropertyId: String
  PropertyValue:
    AssetPropertyValue

```

## Properties

`AssetId`

The ID of the asset that has the specified property.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EntryId`

A unique identifier for this entry. You can use the entry ID to track which data entry
causes an error in case of failure. The default is a new unique identifier.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PropertyAlias`

The alias of the asset property.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PropertyId`

The ID of the asset property.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PropertyValue`

The value to send to the asset property. This value contains timestamp, quality, and value
(TQV) information.

_Required_: Yes

_Type_: [AssetPropertyValue](aws-properties-iotevents-detectormodel-assetpropertyvalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IotEvents

IotTopicPublish

All content copied from https://docs.aws.amazon.com/.
