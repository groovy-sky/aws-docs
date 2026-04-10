This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTFleetWise::Vehicle StateTemplateAssociation

The state template associated with a vehicle. State templates contain state properties, which are signals that belong to a signal catalog that is synchronized between the AWS IoT FleetWise Edge and the AWS Cloud.

###### Important

AWS IoT FleetWise will no longer be open to new customers starting April 30, 2026.
If you would like to use AWS IoT FleetWise, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see [AWS IoT FleetWise availability change](../../../iot-fleetwise/latest/developerguide/iotfleetwise-availability-change.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Identifier" : String,
  "StateTemplateUpdateStrategy" : StateTemplateUpdateStrategy
}

```

### YAML

```yaml

  Identifier: String
  StateTemplateUpdateStrategy:
    StateTemplateUpdateStrategy

```

## Properties

`Identifier`

The unique ID of the state template.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StateTemplateUpdateStrategy`

Property description not available.

_Required_: Yes

_Type_: [StateTemplateUpdateStrategy](aws-properties-iotfleetwise-vehicle-statetemplateupdatestrategy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PeriodicStateTemplateUpdateStrategy

StateTemplateUpdateStrategy

All content copied from https://docs.aws.amazon.com/.
