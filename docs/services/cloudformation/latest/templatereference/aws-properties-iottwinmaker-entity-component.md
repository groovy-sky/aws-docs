This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTTwinMaker::Entity Component

The entity component.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ComponentName" : String,
  "ComponentTypeId" : String,
  "DefinedIn" : String,
  "Description" : String,
  "Properties" : {Key: Value, ...},
  "PropertyGroups" : {Key: Value, ...},
  "Status" : Status
}

```

### YAML

```yaml

  ComponentName: String
  ComponentTypeId: String
  DefinedIn: String
  Description: String
  Properties:
    Key: Value
  PropertyGroups:
    Key: Value
  Status:
    Status

```

## Properties

`ComponentName`

The name of the component.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z_\-0-9]+`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComponentTypeId`

The ID of the component type.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z_\-0-9]+`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefinedIn`

The name of the property definition set in the request.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the component.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Properties`

An object that maps strings to the properties to set in the component type.
Each string in the mapping must be unique to this object.

_Required_: No

_Type_: Object of [Property](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iottwinmaker-entity-property.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PropertyGroups`

An object that maps strings to the property groups in the component type. Each string in the mapping must be unique to this object.

_Required_: No

_Type_: Object of [PropertyGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iottwinmaker-entity-propertygroup.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

The status of the component.

_Required_: No

_Type_: [Status](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iottwinmaker-entity-status.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::IoTTwinMaker::Entity

CompositeComponent
