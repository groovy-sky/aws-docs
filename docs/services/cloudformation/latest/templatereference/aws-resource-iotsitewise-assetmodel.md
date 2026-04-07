This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTSiteWise::AssetModel

Creates an asset model from specified property and hierarchy definitions. You create
assets from asset models. With asset models, you can easily create assets of the same type
that have standardized definitions. Each asset created from a model inherits the asset model's
property and hierarchy definitions. For more information, see [Defining asset models](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/define-models.html) in the
_AWS IoT SiteWise User Guide_.

You can create three types of asset models, `ASSET_MODEL`,
`COMPONENT_MODEL`, or an `INTERFACE`.

- **ASSET\_MODEL** – (default) An asset model that
you can use to create assets. Can't be included as a component in another asset
model.

- **COMPONENT\_MODEL** – A reusable component that
you can include in the composite models of other asset models. You can't create
assets directly from this type of asset model.

- **INTERFACE** – An interface is a type of model
that defines a standard structure that can be applied to different asset models.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoTSiteWise::AssetModel",
  "Properties" : {
      "AssetModelCompositeModels" : [ AssetModelCompositeModel, ... ],
      "AssetModelDescription" : String,
      "AssetModelExternalId" : String,
      "AssetModelHierarchies" : [ AssetModelHierarchy, ... ],
      "AssetModelName" : String,
      "AssetModelProperties" : [ AssetModelProperty, ... ],
      "AssetModelType" : String,
      "EnforcedAssetModelInterfaceRelationships" : [ EnforcedAssetModelInterfaceRelationship, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IoTSiteWise::AssetModel
Properties:
  AssetModelCompositeModels:
    - AssetModelCompositeModel
  AssetModelDescription: String
  AssetModelExternalId: String
  AssetModelHierarchies:
    - AssetModelHierarchy
  AssetModelName: String
  AssetModelProperties:
    - AssetModelProperty
  AssetModelType: String
  EnforcedAssetModelInterfaceRelationships:
    - EnforcedAssetModelInterfaceRelationship
  Tags:
    - Tag

```

## Properties

`AssetModelCompositeModels`

The composite models that are part of this asset model. It groups properties
(such as attributes, measurements, transforms, and metrics) and child composite models that
model parts of your industrial equipment. Each composite model has a type that defines the
properties that the composite model supports. Use composite models to define alarms on this asset model.

###### Note

When creating custom composite models, you need to use [CreateAssetModelCompositeModel](https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_CreateAssetModelCompositeModel.html). For more information,
see [Creating custom composite models (Components)](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/create-custom-composite-models.html) in the
_AWS IoT SiteWise User Guide_.

_Required_: No

_Type_: Array of [AssetModelCompositeModel](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotsitewise-assetmodel-assetmodelcompositemodel.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AssetModelDescription`

A description for the asset model.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AssetModelExternalId`

The external ID of the asset model. For more information, see [Using external IDs](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/object-ids.html#external-ids) in the _AWS IoT SiteWise User Guide_.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9_][a-zA-Z_\-0-9.:]*[a-zA-Z0-9_]+`

_Minimum_: `2`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AssetModelHierarchies`

The hierarchy definitions of the asset model. Each hierarchy specifies an asset model
whose assets can be children of any other assets created from this asset model. For more
information, see [Asset hierarchies](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/asset-hierarchies.html) in the _AWS IoT SiteWise User Guide_.

You can specify up to 10 hierarchies per asset model. For more
information, see [Quotas](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/quotas.html) in the _AWS IoT SiteWise User Guide_.

_Required_: No

_Type_: Array of [AssetModelHierarchy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotsitewise-assetmodel-assetmodelhierarchy.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AssetModelName`

A unique name for the asset model.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AssetModelProperties`

The property definitions of the asset model. For more information, see
[Asset properties](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/asset-properties.html) in the _AWS IoT SiteWise User Guide_.

You can specify up to 200 properties per asset model. For more
information, see [Quotas](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/quotas.html) in the _AWS IoT SiteWise User Guide_.

_Required_: No

_Type_: Array of [AssetModelProperty](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotsitewise-assetmodel-assetmodelproperty.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AssetModelType`

The type of asset model.

- **ASSET\_MODEL** – (default) An asset model that you can use to create assets.
Can't be included as a component in another asset model.

- **COMPONENT\_MODEL** – A reusable component that you can include in the composite
models of other asset models. You can't create assets directly from this type of asset model.

- **INTERFACE** – An interface is a type of model
that defines a standard structure that can be applied to different asset models.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EnforcedAssetModelInterfaceRelationships`

Property description not available.

_Required_: No

_Type_: Array of [EnforcedAssetModelInterfaceRelationship](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotsitewise-assetmodel-enforcedassetmodelinterfacerelationship.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of key-value pairs that contain metadata for the asset. For more information, see
[Tagging your AWS IoT SiteWise\
resources](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/tag-resources.html) in the _AWS IoT SiteWise User Guide_.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotsitewise-assetmodel-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `AssetModelId`.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AssetModelId`

The ID of the asset model.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AssetModelCompositeModel
