This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::Workspace

Contains information about a workspace, which defines the user experience by mapping views to pages.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Connect::Workspace",
  "Properties" : {
      "Associations" : [ String, ... ],
      "Description" : String,
      "InstanceArn" : String,
      "Media" : [ MediaItem, ... ],
      "Name" : String,
      "Pages" : [ WorkspacePage, ... ],
      "Tags" : [ Tag, ... ],
      "Theme" : WorkspaceTheme,
      "Title" : String,
      "Visibility" : String
    }
}

```

### YAML

```yaml

Type: AWS::Connect::Workspace
Properties:
  Associations:
    - String
  Description: String
  InstanceArn: String
  Media:
    - MediaItem
  Name: String
  Pages:
    - WorkspacePage
  Tags:
    - Tag
  Theme:
    WorkspaceTheme
  Title: String
  Visibility: String

```

## Properties

`Associations`

Property description not available.

_Required_: No

_Type_: Array of String

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the workspace.

_Required_: No

_Type_: String

_Pattern_: `^[\P{C}
	]*$`

_Minimum_: `0`

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceArn`

The Amazon Resource Name (ARN) of the instance.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Media`

Property description not available.

_Required_: No

_Type_: Array of [MediaItem](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-workspace-mediaitem.html)

_Maximum_: `4`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the workspace.

_Required_: Yes

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Pages`

Property description not available.

_Required_: No

_Type_: Array of [WorkspacePage](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-workspace-workspacepage.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags used to organize, track, or control access for the workspace.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-workspace-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Theme`

The theme configuration for the workspace, including colors and styling.

_Required_: No

_Type_: [WorkspaceTheme](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-workspace-workspacetheme.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Title`

The title displayed for the workspace.

_Required_: No

_Type_: String

_Pattern_: `^[\P{C}]*$`

_Minimum_: `0`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Visibility`

Controls who can access the workspace. Valid values are: `ALL` (all users), `ASSIGNED`
(only assigned users and routing profiles), and `NONE` (not visible).

_Required_: No

_Type_: String

_Allowed values_: `ALL | ASSIGNED | NONE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

Property description not available.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Connect::ViewVersion

FontFamily
