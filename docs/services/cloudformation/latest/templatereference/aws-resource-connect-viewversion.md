This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::ViewVersion

Creates a version for the specified customer-managed view within the specified
instance.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Connect::ViewVersion",
  "Properties" : {
      "VersionDescription" : String,
      "ViewArn" : String,
      "ViewContentSha256" : String
    }
}

```

### YAML

```yaml

Type: AWS::Connect::ViewVersion
Properties:
  VersionDescription: String
  ViewArn: String
  ViewContentSha256: String

```

## Properties

`VersionDescription`

The description of the view version.

_Required_: No

_Type_: String

_Pattern_: `^([\p{L}\p{N}_.:\/=+\-@,]+[\p{L}\p{Z}\p{N}_.:\/=+\-@,]*)$`

_Minimum_: `1`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ViewArn`

The unqualified Amazon Resource Name (ARN) of the view.

For example:

```

arn:<partition>:connect:<region>:<accountId>:instance/00000000-0000-0000-0000-000000000000/view/00000000-0000-0000-0000-000000000000
```

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/view/[-:a-zA-Z0-9]*$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ViewContentSha256`

Indicates the checksum value of the latest published view content.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9]{64}$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Version`

Current version of the view.

`ViewVersionArn`

The Amazon Resource Name (ARN) of the view version.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::Connect::Workspace
