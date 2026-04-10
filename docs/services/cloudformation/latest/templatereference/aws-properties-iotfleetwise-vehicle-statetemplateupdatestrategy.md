This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTFleetWise::Vehicle StateTemplateUpdateStrategy

The update strategy for the state template. Vehicles associated with the state template can stream telemetry data with either an `onChange` or `periodic` update strategy.

###### Important

AWS IoT FleetWise will no longer be open to new customers starting April 30, 2026.
If you would like to use AWS IoT FleetWise, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see [AWS IoT FleetWise availability change](../../../iot-fleetwise/latest/developerguide/iotfleetwise-availability-change.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OnChange" : Json,
  "Periodic" : PeriodicStateTemplateUpdateStrategy
}

```

### YAML

```yaml

  OnChange: Json
  Periodic:
    PeriodicStateTemplateUpdateStrategy

```

## Properties

`OnChange`

Property description not available.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Periodic`

Property description not available.

_Required_: No

_Type_: [PeriodicStateTemplateUpdateStrategy](aws-properties-iotfleetwise-vehicle-periodicstatetemplateupdatestrategy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StateTemplateAssociation

Tag

All content copied from https://docs.aws.amazon.com/.
