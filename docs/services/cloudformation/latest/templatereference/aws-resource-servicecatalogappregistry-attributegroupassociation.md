This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ServiceCatalogAppRegistry::AttributeGroupAssociation

Associates an attribute group with an application to augment the application's metadata with the group's attributes.
This feature enables applications to be described with user-defined details that are machine-readable, such as third-party integrations.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ServiceCatalogAppRegistry::AttributeGroupAssociation",
  "Properties" : {
      "Application" : String,
      "AttributeGroup" : String
    }
}

```

### YAML

```yaml

Type: AWS::ServiceCatalogAppRegistry::AttributeGroupAssociation
Properties:
  Application: String
  AttributeGroup: String

```

## Properties

`Application`

The name or ID of the application.

_Required_: Yes

_Type_: String

_Pattern_: `\w+|[a-z0-9]{12}`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AttributeGroup`

The name or ID of the attribute group which holds the attributes that describe the application.

_Required_: Yes

_Type_: String

_Pattern_: `\w+|[a-z0-9]{12}`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the application Id.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ApplicationArn`

The Amazon resource name (ARN) of the application that was augmented with attributes.

`AttributeGroupArn`

The Amazon resource name (ARN) of the attribute group which contains the application's new attributes.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ServiceCatalogAppRegistry::AttributeGroup

AWS::ServiceCatalogAppRegistry::ResourceAssociation
