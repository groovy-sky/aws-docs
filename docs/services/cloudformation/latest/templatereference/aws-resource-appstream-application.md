This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppStream::Application

This resource creates an application. Applications store the details about how to launch applications on streaming instances. This is only supported for Elastic fleets.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppStream::Application",
  "Properties" : {
      "AppBlockArn" : String,
      "AttributesToDelete" : [ String, ... ],
      "Description" : String,
      "DisplayName" : String,
      "IconS3Location" : S3Location,
      "InstanceFamilies" : [ String, ... ],
      "LaunchParameters" : String,
      "LaunchPath" : String,
      "Name" : String,
      "Platforms" : [ String, ... ],
      "Tags" : [ Tag, ... ],
      "WorkingDirectory" : String
    }
}

```

### YAML

```yaml

Type: AWS::AppStream::Application
Properties:
  AppBlockArn: String
  AttributesToDelete:
    - String
  Description: String
  DisplayName: String
  IconS3Location:
    S3Location
  InstanceFamilies:
    - String
  LaunchParameters: String
  LaunchPath: String
  Name: String
  Platforms:
    - String
  Tags:
    - Tag
  WorkingDirectory: String

```

## Properties

`AppBlockArn`

The app block ARN with which the application should be associated.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws(?:\-cn|\-iso\-b|\-iso|\-us\-gov)?:[A-Za-z0-9][A-Za-z0-9_/.-]{0,62}:[A-Za-z0-9_/.-]{0,63}:[A-Za-z0-9_/.-]{0,63}:[A-Za-z0-9][A-Za-z0-9:_/+=,@.\\-]{0,1023}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AttributesToDelete`

A list of attributes to delete from an application.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the application.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisplayName`

The display name of the application. This name is visible to users in the application catalog.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IconS3Location`

The icon S3 location of the application.

_Required_: Yes

_Type_: [S3Location](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appstream-application-s3location.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceFamilies`

The instance families the application supports.

_Allowed Values_: `GENERAL_PURPOSE` \| `GRAPHICS_G4`

_Required_: Yes

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LaunchParameters`

The launch parameters of the application.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LaunchPath`

The launch path of the application.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the application. This name is visible to users when a name is not specified in the
DisplayName property.

_Pattern_: `^[a-zA-Z0-9][a-zA-Z0-9_.-]{0,100}$`

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Platforms`

The platforms the application supports.

_Allowed Values_: `WINDOWS_SERVER_2019` \| `AMAZON_LINUX2`

_Required_: Yes

_Type_: Array of String

_Maximum_: `4`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags of the application.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appstream-application-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WorkingDirectory`

The working directory of the application.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function,
`Ref` returns the `Arn` of the application, such as
`arn:aws:appstream:us-west-2:123456789123:application/abcdefg`.

For more information about using the `Ref` function, see [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN of the application.

`CreatedTime`

The time when the application was created.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VpcConfig

S3Location
