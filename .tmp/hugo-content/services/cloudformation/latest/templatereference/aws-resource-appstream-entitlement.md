This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppStream::Entitlement

Creates an entitlement to control access, based on user attributes, to specific
applications within a stack. Entitlements apply to SAML 2.0 federated user identities.
Amazon WorkSpaces Applications user pool and streaming URL users are entitled to all applications in
a stack. Entitlements don't apply to the desktop stream view application or to applications
managed by a dynamic app provider using the Dynamic Application Framework.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppStream::Entitlement",
  "Properties" : {
      "AppVisibility" : String,
      "Attributes" : [ Attribute, ... ],
      "Description" : String,
      "Name" : String,
      "StackName" : String
    }
}

```

### YAML

```yaml

Type: AWS::AppStream::Entitlement
Properties:
  AppVisibility: String
  Attributes:
    - Attribute
  Description: String
  Name: String
  StackName: String

```

## Properties

`AppVisibility`

Specifies whether to entitle all apps or only selected apps.

_Required_: Yes

_Type_: String

_Allowed values_: `ALL | ASSOCIATED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Attributes`

The attributes of the entitlement.

_Required_: Yes

_Type_: Array of [Attribute](aws-properties-appstream-entitlement-attribute.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the entitlement.

_Required_: No

_Type_: String

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the entitlement.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9][a-zA-Z0-9_.-]{0,100}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StackName`

The name of the stack.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9][a-zA-Z0-9_.-]{0,100}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function,
`Ref` returns the combination of the `StackName` and
`Name`, such as `abcdefStack|abcdefEntitlement`.

For more information about using the `Ref` function, see [Ref](../userguide/intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedTime`

The time when the entitlement was created.

`LastModifiedTime`

The time when the entitlement was last modified.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ServiceAccountCredentials

Attribute

All content copied from https://docs.aws.amazon.com/.
