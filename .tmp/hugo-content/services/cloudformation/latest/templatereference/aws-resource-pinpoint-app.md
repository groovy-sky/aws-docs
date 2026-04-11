This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::App

An _app_ is an Amazon Pinpoint application, also referred to as a
_project_. An application is a collection of related settings,
customer information, segments, campaigns, and other types of Amazon Pinpoint
resources.

The App resource represents an Amazon Pinpoint application.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Pinpoint::App",
  "Properties" : {
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Pinpoint::App
Properties:
  Name: String
  Tags:
    - Tag

```

## Properties

`Name`

The display name of the application.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [`Tag`](aws-properties-resource-tags.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the unique identifier ( `ApplicationId`) for
the Amazon Pinpoint application.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the application.

`Id`

The unique identifier for the application. This identifier is displayed as the
**Project ID** on the Amazon Pinpoint
console.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Pinpoint::APNSVoipSandboxChannel

AWS::Pinpoint::ApplicationSettings

All content copied from https://docs.aws.amazon.com/.
