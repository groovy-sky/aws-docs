This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppStream::AppBlockBuilder

Creates an app block builder.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppStream::AppBlockBuilder",
  "Properties" : {
      "AccessEndpoints" : [ AccessEndpoint, ... ],
      "AppBlockArns" : [ String, ... ],
      "Description" : String,
      "DisplayName" : String,
      "EnableDefaultInternetAccess" : Boolean,
      "IamRoleArn" : String,
      "InstanceType" : String,
      "Name" : String,
      "Platform" : String,
      "Tags" : [ Tag, ... ],
      "VpcConfig" : VpcConfig
    }
}

```

### YAML

```yaml

Type: AWS::AppStream::AppBlockBuilder
Properties:
  AccessEndpoints:
    - AccessEndpoint
  AppBlockArns:
    - String
  Description: String
  DisplayName: String
  EnableDefaultInternetAccess: Boolean
  IamRoleArn: String
  InstanceType: String
  Name: String
  Platform: String
  Tags:
    - Tag
  VpcConfig:
    VpcConfig

```

## Properties

`AccessEndpoints`

The access endpoints of the app block builder.

_Required_: No

_Type_: Array of [AccessEndpoint](aws-properties-appstream-appblockbuilder-accessendpoint.md)

_Minimum_: `1`

_Maximum_: `4`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AppBlockArns`

The ARN of the app block.

_Maximum_: `1`

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the app block builder.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisplayName`

The display name of the app block builder.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableDefaultInternetAccess`

Indicates whether default internet access is enabled for the app block builder.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IamRoleArn`

The ARN of the IAM role that is applied to the app block builder.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws(?:\-cn|\-iso\-b|\-iso|\-us\-gov)?:[A-Za-z0-9][A-Za-z0-9_/.-]{0,62}:[A-Za-z0-9_/.-]{0,63}:[A-Za-z0-9_/.-]{0,63}:[A-Za-z0-9][A-Za-z0-9:_/+=,@.\\-]{0,1023}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceType`

The instance type of the app block builder.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the app block builder.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Platform`

The platform of the app block builder.

_Allowed values_: `WINDOWS_SERVER_2019`

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags of the app block builder.

_Required_: No

_Type_: Array of [Tag](aws-properties-appstream-appblockbuilder-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcConfig`

The VPC configuration for the app block builder.

_Required_: Yes

_Type_: [VpcConfig](aws-properties-appstream-appblockbuilder-vpcconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function,
`Ref` returns the `Name` of the app block builder, such as
`abcdefAppBlockBuilder`.

For more information about using the `Ref` function, see [Ref](../userguide/intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN of the app block builder.

`CreatedTime`

The time when the app block builder was created.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TagItems

AccessEndpoint

All content copied from https://docs.aws.amazon.com/.
