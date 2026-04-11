This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::IdNamespaceAssociation

Provides information to create the ID namespace association.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CleanRooms::IdNamespaceAssociation",
  "Properties" : {
      "Description" : String,
      "IdMappingConfig" : IdMappingConfig,
      "InputReferenceConfig" : IdNamespaceAssociationInputReferenceConfig,
      "MembershipIdentifier" : String,
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::CleanRooms::IdNamespaceAssociation
Properties:
  Description: String
  IdMappingConfig:
    IdMappingConfig
  InputReferenceConfig:
    IdNamespaceAssociationInputReferenceConfig
  MembershipIdentifier: String
  Name: String
  Tags:
    - Tag

```

## Properties

`Description`

The description of the ID namespace association.

_Required_: No

_Type_: String

_Pattern_: `^[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDBFF-\uDC00\uDFFF\t\r\n]*$`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IdMappingConfig`

The configuration settings for the ID mapping table.

_Required_: No

_Type_: [IdMappingConfig](aws-properties-cleanrooms-idnamespaceassociation-idmappingconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputReferenceConfig`

The input reference configuration for the ID namespace association.

_Required_: Yes

_Type_: [IdNamespaceAssociationInputReferenceConfig](aws-properties-cleanrooms-idnamespaceassociation-idnamespaceassociationinputreferenceconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MembershipIdentifier`

The unique identifier of the membership that contains the ID namespace association.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`

_Minimum_: `36`

_Maximum_: `36`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of this ID namespace association.

_Required_: Yes

_Type_: String

_Pattern_: `^(?!\s*$)[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDBFF-\uDC00\uDFFF\t]*$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An optional label that you can assign to a resource when you create it. Each tag
consists of a key and an optional value, both of which you define. When you use tagging,
you can also use tag-based access control in IAM policies to control access
to this resource.

_Required_: No

_Type_: Array of [Tag](aws-properties-cleanrooms-idnamespaceassociation-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name. For example:

`{"Ref": "MyIdNamespaceAssociation"}`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the ID namespace association.

`CollaborationArn`

The Amazon Resource Name (ARN) of the collaboration that contains this ID namespace association.

`CollaborationIdentifier`

The unique identifier of the collaboration that contains this ID namespace association.

`IdNamespaceAssociationIdentifier`

The unique identifier of the ID namespace association that you want to retrieve.

`MembershipArn`

The Amazon Resource Name (ARN) of the membership resource for this ID namespace association.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

IdMappingConfig

All content copied from https://docs.aws.amazon.com/.
