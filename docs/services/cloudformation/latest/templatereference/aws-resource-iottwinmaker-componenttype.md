This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTTwinMaker::ComponentType

Use the `AWS::IoTTwinMaker::ComponentType` resource to declare a component type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoTTwinMaker::ComponentType",
  "Properties" : {
      "ComponentTypeId" : String,
      "CompositeComponentTypes" : {Key: Value, ...},
      "Description" : String,
      "ExtendsFrom" : [ String, ... ],
      "Functions" : {Key: Value, ...},
      "IsSingleton" : Boolean,
      "PropertyDefinitions" : {Key: Value, ...},
      "PropertyGroups" : {Key: Value, ...},
      "Tags" : {Key: Value, ...},
      "WorkspaceId" : String
    }
}

```

### YAML

```yaml

Type: AWS::IoTTwinMaker::ComponentType
Properties:
  ComponentTypeId: String
  CompositeComponentTypes:
    Key: Value
  Description: String
  ExtendsFrom:
    - String
  Functions:
    Key: Value
  IsSingleton: Boolean
  PropertyDefinitions:
    Key: Value
  PropertyGroups:
    Key: Value
  Tags:
    Key: Value
  WorkspaceId: String

```

## Properties

`ComponentTypeId`

The ID of the component type.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z_\.\-0-9:]+`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CompositeComponentTypes`

Maps strings to `compositeComponentTypes` of the `componentType`.
`CompositeComponentType` is referenced by `componentTypeId`.

_Required_: No

_Type_: Object of [CompositeComponentType](aws-properties-iottwinmaker-componenttype-compositecomponenttype.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the component type.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExtendsFrom`

The name of the parent component type that this component type extends.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Functions`

An object that maps strings to the functions in the component type. Each string in the mapping must be unique to this object.

For information on the FunctionResponse object see the [FunctionResponse](../../../../reference/iot-twinmaker/latest/apireference/api-functionresponse.md) API reference.

_Required_: No

_Type_: Object of [Function](aws-properties-iottwinmaker-componenttype-function.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsSingleton`

A boolean value that specifies whether an entity can have more than one component of this type.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PropertyDefinitions`

An object that maps strings to the property definitions in the component type. Each string in the mapping must be unique to this object.

For information about the PropertyDefinitionResponse object, see the [PropertyDefinitionResponse](../../../../reference/iot-twinmaker/latest/apireference/api-propertydefinitionresponse.md) API reference.

_Required_: No

_Type_: Object of [PropertyDefinition](aws-properties-iottwinmaker-componenttype-propertydefinition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PropertyGroups`

An object that maps strings to the property groups in the component type. Each string in the mapping must be unique to this object.

_Required_: No

_Type_: Object of [PropertyGroup](aws-properties-iottwinmaker-componenttype-propertygroup.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The ComponentType tags.

_Required_: No

_Type_: Object of String

_Pattern_: `^([\p{L}\p{Z}\p{N}_.:/=+\-@]*)$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WorkspaceId`

The ID of the workspace that contains the component type.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z_0-9][a-zA-Z_\-0-9]*[a-zA-Z0-9]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the workspace Id and the ComponentType Id.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN of the component type.

`CreationDateTime`

The date and time when the component type was created.

`IsAbstract`

A boolean value that specifies whether the component type is abstract.

`IsSchemaInitialized`

A boolean value that specifies whether the component type has a schema initializer and that the schema initializer has run.

`UpdateDateTime`

The component type the update time.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS IoT TwinMaker

CompositeComponentType

All content copied from https://docs.aws.amazon.com/.
