This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTTwinMaker::Entity

Use the `AWS::IoTTwinMaker::Entity` resource to declare an entity.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoTTwinMaker::Entity",
  "Properties" : {
      "Components" : {Key: Value, ...},
      "CompositeComponents" : {Key: Value, ...},
      "Description" : String,
      "EntityId" : String,
      "EntityName" : String,
      "ParentEntityId" : String,
      "Tags" : {Key: Value, ...},
      "WorkspaceId" : String
    }
}

```

### YAML

```yaml

Type: AWS::IoTTwinMaker::Entity
Properties:
  Components:
    Key: Value
  CompositeComponents:
    Key: Value
  Description: String
  EntityId: String
  EntityName: String
  ParentEntityId: String
  Tags:
    Key: Value
  WorkspaceId: String

```

## Properties

`Components`

An object that maps strings to the components in the entity. Each string in the mapping must be unique to this object.

For information on the component object see the [component](https://docs.aws.amazon.com/iot-twinmaker/latest/apireference/API_ComponentResponse.html) API reference.

_Required_: No

_Type_: Object of [Component](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iottwinmaker-entity-component.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CompositeComponents`

Maps string to `compositeComponent` updates in the request.
Each key of the map represents the `componentPath` of the
`compositeComponent`.

_Required_: No

_Type_: Object of [CompositeComponent](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iottwinmaker-entity-compositecomponent.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the entity.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EntityId`

The ID of the entity.

_Required_: No

_Type_: String

_Pattern_: `[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}|^[a-zA-Z0-9][a-zA-Z_\-0-9.:]*[a-zA-Z0-9]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EntityName`

The entity name.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z_0-9-.][a-zA-Z_0-9-. ]*[a-zA-Z0-9]+`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParentEntityId`

The ID of the parent entity.

_Required_: No

_Type_: String

_Pattern_: `\$ROOT|^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}|^[a-zA-Z0-9][a-zA-Z_\-0-9.:]*[a-zA-Z0-9]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Metadata that you can use to manage the entity.

_Required_: No

_Type_: Object of String

_Pattern_: `^([\p{L}\p{Z}\p{N}_.:/=+\-@]*)$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WorkspaceId`

The ID of the workspace that contains the entity.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z_0-9][a-zA-Z_\-0-9]*[a-zA-Z0-9]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the workspace Id and the entity Id.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The entity ARN.

`CreationDateTime`

The date and time the entity was created.

`HasChildEntities`

A boolean value that specifies whether the entity has child entities or not.

`UpdateDateTime`

The date and time when the component type was last updated.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Status

Component
