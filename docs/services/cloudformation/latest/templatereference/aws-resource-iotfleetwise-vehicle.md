This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTFleetWise::Vehicle

Creates a vehicle, which is an instance of a vehicle model (model manifest). Vehicles
created from the same vehicle model consist of the same signals inherited from the
vehicle model.

###### Note

If you have an existing AWS IoT thing, you can use AWS IoT FleetWise to create a
vehicle and collect data from your thing.

For more information, see [Vehicles](../../../iot-fleetwise/latest/developerguide/vehicles.md) in the _AWS IoT FleetWise Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoTFleetWise::Vehicle",
  "Properties" : {
      "AssociationBehavior" : String,
      "Attributes" : {Key: Value, ...},
      "DecoderManifestArn" : String,
      "ModelManifestArn" : String,
      "Name" : String,
      "StateTemplates" : [ StateTemplateAssociation, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IoTFleetWise::Vehicle
Properties:
  AssociationBehavior: String
  Attributes:
    Key: Value
  DecoderManifestArn: String
  ModelManifestArn: String
  Name: String
  StateTemplates:
    - StateTemplateAssociation
  Tags:
    - Tag

```

## Properties

`AssociationBehavior`

An option to create a new AWS IoT thing when creating a vehicle, or to validate an
existing thing as a vehicle.

_Required_: No

_Type_: String

_Allowed values_: `CreateIotThing | ValidateIotThingExists`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Attributes`

Static information about a vehicle in a key-value pair. For example: `"engine
                Type"` : `"v6"`

_Required_: No

_Type_: Object of String

_Pattern_: `^[a-zA-Z0-9_.-]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DecoderManifestArn`

The Amazon Resource Name (ARN) of a decoder manifest associated with the vehicle to create.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelManifestArn`

The Amazon Resource Name (ARN) of the vehicle model (model manifest) to create the vehicle from.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The unique ID of the vehicle.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z\d\-_:]+$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StateTemplates`

Associate state templates to track the state of the vehicle. State templates determine which signal updates the vehicle sends to the cloud.

_Required_: No

_Type_: Array of [StateTemplateAssociation](aws-properties-iotfleetwise-vehicle-statetemplateassociation.md)

_Minimum_: `0`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Metadata which can be used to manage the vehicle.

_Required_: No

_Type_: Array of [Tag](aws-properties-iotfleetwise-vehicle-tag.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the vehicle.

`CreationTime`

The time the vehicle was created in seconds since epoch (January 1, 1970 at midnight UTC time).

`LastModificationTime`

The time the vehicle was last updated in seconds since epoch (January 1, 1970 at midnight UTC time).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

PeriodicStateTemplateUpdateStrategy

All content copied from https://docs.aws.amazon.com/.
