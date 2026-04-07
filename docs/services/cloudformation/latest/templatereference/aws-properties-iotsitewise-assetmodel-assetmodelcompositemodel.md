This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTSiteWise::AssetModel AssetModelCompositeModel

Contains information about a composite model in an asset model. This object contains the
asset property definitions that you define in the composite model.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ComposedAssetModelId" : String,
  "CompositeModelProperties" : [ AssetModelProperty, ... ],
  "Description" : String,
  "ExternalId" : String,
  "Id" : String,
  "Name" : String,
  "ParentAssetModelCompositeModelExternalId" : String,
  "Path" : [ String, ... ],
  "Type" : String
}

```

### YAML

```yaml

  ComposedAssetModelId: String
  CompositeModelProperties:
    - AssetModelProperty
  Description: String
  ExternalId: String
  Id: String
  Name: String
  ParentAssetModelCompositeModelExternalId: String
  Path:
    - String
  Type: String

```

## Properties

`ComposedAssetModelId`

The ID of a component model which is reused to create this composite model.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CompositeModelProperties`

The asset property definitions for this composite model.

_Required_: No

_Type_: Array of [AssetModelProperty](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotsitewise-assetmodel-assetmodelproperty.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the composite model.

###### Note

If the composite model is a `component-model-based` composite model,
the description is inherited from the `COMPONENT_MODEL` asset model and
cannot be changed.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExternalId`

The external ID of a composite model on this asset model. For more information, see
[Using external\
IDs](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/object-ids.html#external-ids) in the _AWS IoT SiteWise User Guide_.

###### Note

One of `ExternalId` or `Path` must be specified.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9_][a-zA-Z_\-0-9.:]*[a-zA-Z0-9_]+`

_Minimum_: `2`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Id`

The ID of the asset model composite model.

###### Note

This is a return value and can't be set.

_Required_: No

_Type_: String

_Pattern_: `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`

_Minimum_: `36`

_Maximum_: `36`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the composite model.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParentAssetModelCompositeModelExternalId`

The external ID of the parent composite model. For more information, see [Using external\
IDs](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/object-ids.html#external-ids) in the _AWS IoT SiteWise User Guide_.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9_][a-zA-Z_\-0-9.:]*[a-zA-Z0-9_]+`

_Minimum_: `2`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Path`

The structured path to the property from the root of the asset using property names.
Path is used as the ID if the asset model is a derived composite model.

###### Note

One of `ExternalId` or `Path` must be specified.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of the composite model. For alarm composite models, this type is
`AWS/ALARM`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

- [Create an alarm model](#aws-properties-iotsitewise-assetmodel-assetmodelcompositemodel--examples--Create_an_alarm_model)

- [Create a component-model-based composite model](#aws-properties-iotsitewise-assetmodel-assetmodelcompositemodel--examples--Create_a_component-model-based_composite_model)

### Create an alarm model

You can modify the following example to create an alarm model.

###### Note

Replace `TestAlarmModel` with the name of your alarm
model.

#### YAML

```yaml

Resources: AssetModelWithAlarmCompositeModel: Type:
                AWS::IoTSiteWise::AssetModel Properties: AssetModelName:
                AssetModelWithValidAlarmCompositeModel AssetModelDescription:
                AssetModelWithValidAlarmCompositeModel AssetModelCompositeModels: - Description:
                compositeModel created by cfn Name: TestAlarmCompositeModel Type: AWS/ALARM
                CompositeModelProperties: - LogicalId: MyLogicalId_for_ALARM_TYPE_1 Name:
                AWS/ALARM_TYPE DataType: STRING Type: TypeName: Attribute Attribute: DefaultValue:
                IOT_EVENTS - LogicalId: MyLogicalId_for_ALARM_SOURCE_1 Name: AWS/ALARM_SOURCE
                DataType: STRING Type: TypeName: Attribute Attribute: DefaultValue: Fn::Sub:
                "arn:${AWS::Partition}:iotevents:${AWS::Region}:${AWS::AccountId}:alarmModel/TestAlarmModel"
                - LogicalId: MyLogicalId_for_ALARM_STATE_1 Name: AWS/ALARM_STATE DataType: STRUCT
                DataTypeSpec: AWS/ALARM_STATE Type: TypeName: Measurement
```

### Create a `component-model-based` composite model

You can modify the following example to create a
`component-model-based` composite model.

#### YAML

```yaml

Resources: MainAssetModel: Type: AWS::IoTSiteWise::AssetModel
                Properties: AssetModelExternalId: AssetModelName: AssetModelWithCustomComposites
                AssetModelDescription: Asset model with custom composite models
                AssetModelProperties: - Name: metric_property ExternalId:
                metric_property_external_id DataType: DOUBLE Type: TypeName: Transform Transform:
                Expression: abs(avgtemp) Variables: - Name: avgtemp Value: PropertyPath: - Name:
                AssetModelWithCustomComposites - Name: CompositeModel3 - Name:
                ComponentModel1CompositeModel - Name: ComponentModel2Property
                AssetModelCompositeModels: - ExternalId: CompositeModelExternalId1 Type: CUSTOM
                Name: CompositeModel1 CompositeModelProperties: - Name: CompositeModel1Property
                ExternalId: CompositeModel1PropertyExternalId DataType: DOUBLE Type: TypeName:
                Measurement - ExternalId: CompositeModelExternalId2 Type: CUSTOM Name:
                CompositeModel2 ParentAssetModelCompositeModelExternalId: CompositeModelExternalId1
                CompositeModelProperties: - Name: CompositeModel2Property ExternalId:
                CompositeModel2PropertyExternalId DataType: DOUBLE Type: TypeName: Measurement -
                ExternalId: CompositeModelExternalId3 Type: CUSTOM Name: CompositeModel3
                ComposedAssetModelId: !Ref ComponentModel1 - Path: - AssetModelWithCustomComposites
                - CompositeModel3 - ComponentModel1CompositeModel ExternalId:
                DerivedCompositeExternalId Type: CUSTOM Name: ComponentModel1CompositeModel
                CompositeModelProperties: - Name: ComponentModel2Property ExternalId:
                DerivedCompositePropertyExternalId DataType: DOUBLE Type: TypeName: Measurement
                ComponentModel1: Type: AWS::IoTSiteWise::AssetModel Properties:
                AssetModelExternalId: ComponentModel1ExternalId AssetModelType: COMPONENT_MODEL
                AssetModelName: ComponentModel1 AssetModelCompositeModels: - ExternalId:
                ComponentModel1CompositeModel ComposedAssetModelId: !Ref ComponentModel2 Name:
                ComponentModel1CompositeModel Type: CUSTOM ComponentModel2: Type:
                AWS::IoTSiteWise::AssetModel Properties: AssetModelExternalId:
                ComponentModel2ExternalId AssetModelType: COMPONENT_MODEL AssetModelName:
                ComponentModel2 AssetModelProperties: - Name: ComponentModel2Property ExternalId:
                ComponentModel2PropertyExternalId DataType: DOUBLE Type: TypeName: Measurement

```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::IoTSiteWise::AssetModel

AssetModelHierarchy
