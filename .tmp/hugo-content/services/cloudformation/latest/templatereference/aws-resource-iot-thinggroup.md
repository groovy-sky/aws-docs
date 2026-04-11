This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::ThingGroup

Creates a new thing group. A dynamic thing group is created if the resource template
contains the `QueryString` attribute. A dynamic thing group will not contain the
`ParentGroupName` attribute. A static thing group and dynamic thing group
can't be converted to each other via the addition or removal of the
`QueryString` attribute.

###### Note

This is a control plane operation. See [Authorization](../../../iot/latest/developerguide/iot-authorization.md) for
information about authorizing control plane actions.

Requires permission to access the [CreateThingGroup](../../../service-authorization/latest/reference/list-awsiot.md#awsiot-actions-as-permissions) action.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoT::ThingGroup",
  "Properties" : {
      "ParentGroupName" : String,
      "QueryString" : String,
      "Tags" : [ Tag, ... ],
      "ThingGroupName" : String,
      "ThingGroupProperties" : ThingGroupProperties
    }
}

```

### YAML

```yaml

Type: AWS::IoT::ThingGroup
Properties:
  ParentGroupName: String
  QueryString:
    String
  Tags:
    - Tag
  ThingGroupName: String
  ThingGroupProperties:
    ThingGroupProperties

```

## Properties

`ParentGroupName`

The parent thing group name.

A Dynamic Thing Group does not have `parentGroupName` defined.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9:_-]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`QueryString`

The dynamic thing group search query string.

The `queryString` attribute _is_ required for
`CreateDynamicThingGroup`. The `queryString` attribute
_is not_ required for `CreateThingGroup`.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Metadata which can be used to manage the thing group or dynamic thing group.

_Required_: No

_Type_: Array of [Tag](aws-properties-iot-thinggroup-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThingGroupName`

The thing group name.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9:_-]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ThingGroupProperties`

Thing group properties.

_Required_: No

_Type_: [ThingGroupProperties](aws-properties-iot-thinggroup-thinggroupproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the thing group id.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The thing group ARN.

`Id`

The thing group ID.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AttributePayload

AttributePayload

All content copied from https://docs.aws.amazon.com/.
