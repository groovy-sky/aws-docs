This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::ThingType

Creates a new thing type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoT::ThingType",
  "Properties" : {
      "DeprecateThingType" : Boolean,
      "Tags" : [ Tag, ... ],
      "ThingTypeName" : String,
      "ThingTypeProperties" : ThingTypeProperties
    }
}

```

### YAML

```yaml

Type: AWS::IoT::ThingType
Properties:
  DeprecateThingType: Boolean
  Tags:
    - Tag
  ThingTypeName: String
  ThingTypeProperties:
    ThingTypeProperties

```

## Properties

`DeprecateThingType`

Deprecates a thing type. You can not associate new things with deprecated thing
type.

Requires permission to access the [DeprecateThingType](../../../service-authorization/latest/reference/list-awsiot.md#awsiot-actions-as-permissions) action.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Metadata which can be used to manage the thing type.

_Required_: No

_Type_: Array of [Tag](aws-properties-iot-thingtype-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThingTypeName`

The name of the thing type.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9:_-]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ThingTypeProperties`

The thing type properties for the thing type to create. It contains information about
the new thing type including a description, a list of searchable thing attribute names, and
a list of propagating attributes. After a thing type is created, you can only update `Mqtt5Configuration`.

_Required_: No

_Type_: [ThingTypeProperties](aws-properties-iot-thingtype-thingtypeproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the thing type id.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The thing type arn.

`Id`

The thing type id.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::IoT::ThingPrincipalAttachment

Mqtt5Configuration

All content copied from https://docs.aws.amazon.com/.
