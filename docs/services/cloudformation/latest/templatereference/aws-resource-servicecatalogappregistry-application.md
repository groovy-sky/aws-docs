This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ServiceCatalogAppRegistry::Application

Represents a AWS Service Catalog AppRegistry application that is the top-level node in a hierarchy of related
cloud resource abstractions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ServiceCatalogAppRegistry::Application",
  "Properties" : {
      "Description" : String,
      "Name" : String,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::ServiceCatalogAppRegistry::Application
Properties:
  Description: String
  Name: String
  Tags:
    Key: Value

```

## Properties

`Description`

The description of the application.

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the application. The name must be unique in the region in which you are creating the application.

_Required_: Yes

_Type_: String

_Pattern_: `\w+`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Key-value pairs you can use to associate with the application.

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

`ApplicationName`

The name of the application. The name must be unique in the region in which you are creating the application.

`ApplicationTagKey`

The key of the AWS application tag, which is `awsApplication`.
Applications created before 11/13/2023 or applications without the `AppTag` linked resource group return no value.

`ApplicationTagValue`

The value of the AWS application tag, which is the identifier of an associated resource.
Applications created before 11/13/2023 or applications without the `AppTag` linked resource group return no value.

`Arn`

The Amazon resource name (ARN) that specifies the application across services.

`Id`

The identifier of the application.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS Service Catalog AppRegistry

AWS::ServiceCatalogAppRegistry::AttributeGroup
