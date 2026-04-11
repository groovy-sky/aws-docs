This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::View

Creates a customer-managed view in the published state within the specified
instance.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Connect::View",
  "Properties" : {
      "Actions" : [ String, ... ],
      "Description" : String,
      "InstanceArn" : String,
      "Name" : String,
      "Tags" : [ Tag, ... ],
      "Template" : Json
    }
}

```

### YAML

```yaml

Type: AWS::Connect::View
Properties:
  Actions:
    - String
  Description: String
  InstanceArn: String
  Name: String
  Tags:
    - Tag
  Template: Json

```

## Properties

`Actions`

A list of actions possible from the view.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `255 | 1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the view.

_Required_: No

_Type_: String

_Pattern_: `^([\p{L}\p{N}_.:\/=+\-@,]+[\p{L}\p{Z}\p{N}_.:\/=+\-@,]*)$`

_Minimum_: `0`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceArn`

The Amazon Resource Name (ARN) of the instance.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the view.

_Required_: Yes

_Type_: String

_Pattern_: `^([\p{L}\p{N}_.:\/=+\-@]+[\p{L}\p{Z}\p{N}_.:\/=+\-@]*)$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags associated with the view resource (not specific to view version).

_Required_: No

_Type_: Array of [Tag](aws-properties-connect-view-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Template`

The view template representing the structure of the view.

_Required_: Yes

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ViewArn`

The unqualified Amazon Resource Name (ARN) of the view.

For example:

```

arn:<partition>:connect:<region>:<accountId>:instance/00000000-0000-0000-0000-000000000000/view/00000000-0000-0000-0000-000000000000
```

`ViewContentSha256`

Indicates the checksum value of the latest published view content.

`ViewId`

The identifier of the view.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UserHierarchyStructure

Tag

All content copied from https://docs.aws.amazon.com/.
