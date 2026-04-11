This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ServiceCatalogAppRegistry::AttributeGroup

Creates a new attribute group as a container for user-defined attributes. This feature
enables users to have full control over their cloud application's metadata in a rich
machine-readable format to facilitate integration with automated workflows and third-party
tools.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ServiceCatalogAppRegistry::AttributeGroup",
  "Properties" : {
      "Attributes" : Json,
      "Description" : String,
      "Name" : String,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::ServiceCatalogAppRegistry::AttributeGroup
Properties:
  Attributes: Json
  Description: String
  Name: String
  Tags:
    Key: Value

```

## Properties

`Attributes`

A nested object
in a JSON or YAML template
that supports arbitrary definitions.
Represents the attributes
in an attribute group
that describes an application and its components.

_Required_: Yes

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the attribute group that the user provides.

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the attribute group.

_Required_: Yes

_Type_: String

_Pattern_: `\w+`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Key-value pairs you can use to associate with the attribute group.

_Required_: No

_Type_: Object of String

_Pattern_: `^[a-zA-Z+-=._:/]+$`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the application Id.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon resource name (ARN)
that specifies the attribute group
across services.

`Id`

The globally unique attribute group identifier
of the attribute group.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ServiceCatalogAppRegistry::Application

AWS::ServiceCatalogAppRegistry::AttributeGroupAssociation

All content copied from https://docs.aws.amazon.com/.
