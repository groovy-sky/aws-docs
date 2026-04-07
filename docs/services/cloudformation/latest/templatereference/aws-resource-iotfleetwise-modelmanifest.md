This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTFleetWise::ModelManifest

###### Important

AWS IoT FleetWise will no longer be open to new customers starting April 30, 2026.
If you would like to use AWS IoT FleetWise, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see [AWS IoT FleetWise availability change](https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/iotfleetwise-availability-change.html).

Creates a vehicle model (model manifest) that specifies signals (attributes,
branches, sensors, and actuators).

For more information, see [Vehicle models](https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/vehicle-models.html)
in the _AWS IoT FleetWise Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoTFleetWise::ModelManifest",
  "Properties" : {
      "Description" : String,
      "Name" : String,
      "Nodes" : [ String, ... ],
      "SignalCatalogArn" : String,
      "Status" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IoTFleetWise::ModelManifest
Properties:
  Description: String
  Name: String
  Nodes:
    - String
  SignalCatalogArn: String
  Status: String
  Tags:
    - Tag

```

## Properties

`Description`

A brief description of the vehicle model.

_Required_: No

_Type_: String

_Pattern_: `^[^\u0000-\u001F\u007F]+$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the vehicle model.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z\d\-_:]+$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Nodes`

A list of nodes, which are a general abstraction of signals.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SignalCatalogArn`

The Amazon Resource Name (ARN) of the signal catalog associated with the vehicle model.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

The state of the vehicle model. If the status is `ACTIVE`, the vehicle
model can't be edited. If the status is `DRAFT`, you can edit the vehicle
model.

_Required_: No

_Type_: String

_Allowed values_: `ACTIVE | DRAFT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Metadata that can be used to manage the vehicle model.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotfleetwise-modelmanifest-tag.html)

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

The Amazon Resource Name (ARN) of the vehicle model.

`CreationTime`

The time the vehicle model was created, in seconds since epoch (January 1, 1970 at
midnight UTC time).

`LastModificationTime`

The time the vehicle model was last updated, in seconds since epoch (January 1, 1970
at midnight UTC time).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
