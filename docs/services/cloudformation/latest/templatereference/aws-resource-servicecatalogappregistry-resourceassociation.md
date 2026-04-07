This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ServiceCatalogAppRegistry::ResourceAssociation

Associates a resource
with an application.
Both the resource and the application can be specified either
by ID or name.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ServiceCatalogAppRegistry::ResourceAssociation",
  "Properties" : {
      "Application" : String,
      "Resource" : String,
      "ResourceType" : String
    }
}

```

### YAML

```yaml

Type: AWS::ServiceCatalogAppRegistry::ResourceAssociation
Properties:
  Application: String
  Resource: String
  ResourceType: String

```

## Properties

`Application`

The name or ID
of the application.

_Required_: Yes

_Type_: String

_Pattern_: `\w+|[a-z0-9]{12}`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Resource`

The name or ID
of the resource
of which the application
will be associated.

_Required_: Yes

_Type_: String

_Pattern_: `\w+|arn:aws[-a-z]*:cloudformation:[a-z]{2}(-gov)?-[a-z]+-\d:\d{12}:stack/[a-zA-Z][-A-Za-z0-9]{0,127}/[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResourceType`

The type
of resource
of which the application will be associated.

_Required_: Yes

_Type_: String

_Allowed values_: `CFN_STACK`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the application Id.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ApplicationArn`

The Amazon resource name (ARN)
that specifies the application.

`ResourceArn`

The Amazon resource name (ARN)
that specifies the resource.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ServiceCatalogAppRegistry::AttributeGroupAssociation

Next
