This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTSiteWise::AssetModel VariableValue

Identifies a property value used in an expression.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HierarchyExternalId" : String,
  "HierarchyId" : String,
  "HierarchyLogicalId" : String,
  "PropertyExternalId" : String,
  "PropertyId" : String,
  "PropertyLogicalId" : String,
  "PropertyPath" : [ PropertyPathDefinition, ... ]
}

```

### YAML

```yaml

  HierarchyExternalId: String
  HierarchyId: String
  HierarchyLogicalId: String
  PropertyExternalId: String
  PropertyId: String
  PropertyLogicalId: String
  PropertyPath:
    - PropertyPathDefinition

```

## Properties

`HierarchyExternalId`

The external ID of the hierarchy being referenced. For more information, see [Using external\
IDs](../../../iot-sitewise/latest/userguide/object-ids.md#external-ids) in the _AWS IoT SiteWise User Guide_.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9_][a-zA-Z_\-0-9.:]*[a-zA-Z0-9_]+`

_Minimum_: `2`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HierarchyId`

The ID of the hierarchy to query for the property ID. You can use the hierarchy's name
instead of the hierarchy's ID. If the hierarchy has an external ID, you can specify
`externalId:` followed by the external ID. For more information, see [Using external IDs](../../../iot-sitewise/latest/userguide/object-ids.md#external-ids) in the _AWS IoT SiteWise User Guide_.

You use a hierarchy ID instead of a model ID because you can have several hierarchies
using the same model and therefore the same `propertyId`. For example, you might
have separately grouped assets that come from the same asset model. For more information, see
[Asset hierarchies](../../../iot-sitewise/latest/userguide/asset-hierarchies.md) in the _AWS IoT SiteWise User Guide_.

_Required_: No

_Type_: String

_Pattern_: `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`

_Minimum_: `36`

_Maximum_: `36`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HierarchyLogicalId`

The `LogicalID` of the hierarchy to query for the
`PropertyLogicalID`.

You use a `hierarchyLogicalID` instead of a model ID because you can have
several hierarchies using the same model and therefore the same property. For example,
you might have separately grouped assets that come from the same asset model. For more
information, see [Defining relationships\
between asset models (hierarchies)](../../../iot-sitewise/latest/userguide/asset-hierarchies.md) in the _AWS IoT SiteWise User_
_Guide_.

_Required_: No

_Type_: String

_Pattern_: `[^\u0000-\u001F\u007F]+`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PropertyExternalId`

The external ID of the property being referenced. For more information, see [Using external\
IDs](../../../iot-sitewise/latest/userguide/object-ids.md#external-ids) in the _AWS IoT SiteWise User Guide_.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9_][a-zA-Z_\-0-9.:]*[a-zA-Z0-9_]+`

_Minimum_: `2`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PropertyId`

The ID of the property to use as the variable. You can use the property
`name` if it's from the same asset model. If the property has an external
ID, you can specify `externalId:` followed by the external ID. For more
information, see [Using external\
IDs](../../../iot-sitewise/latest/userguide/object-ids.md#external-ids) in the _AWS IoT SiteWise User Guide_.

###### Note

This is a return value and can't be set.

_Required_: No

_Type_: String

_Pattern_: `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`

_Minimum_: `36`

_Maximum_: `36`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PropertyLogicalId`

The `LogicalID` of the property that is being referenced.

_Required_: No

_Type_: String

_Pattern_: `[^\u0000-\u001F\u007F]+`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PropertyPath`

The path of the property. Each step of the path is the name of the step. See the
following example:

```

PropertyPath:
  Name: AssetModelName
  Name: Composite1
  Name: NestedComposite
```

_Required_: No

_Type_: Array of [PropertyPathDefinition](aws-properties-iotsitewise-assetmodel-propertypathdefinition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TumblingWindow

AWS::IoTSiteWise::ComputationModel

All content copied from https://docs.aws.amazon.com/.
