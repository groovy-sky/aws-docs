This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMR::StudioSessionMapping

The `AWS::EMR::StudioSessionMapping` resource is an Amazon EMR resource type that maps a user or group to the Amazon EMR Studio specified by `StudioId`, and
applies a session policy that defines Studio permissions for that user or group.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EMR::StudioSessionMapping",
  "Properties" : {
      "IdentityName" : String,
      "IdentityType" : String,
      "SessionPolicyArn" : String,
      "StudioId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EMR::StudioSessionMapping
Properties:
  IdentityName: String
  IdentityType: String
  SessionPolicyArn: String
  StudioId: String

```

## Properties

`IdentityName`

The name of the user or group. For more information, see [UserName](../../../../reference/singlesignon/latest/identitystoreapireference/api-user.md#singlesignon-Type-User-UserName) and [DisplayName](../../../../reference/singlesignon/latest/identitystoreapireference/api-group.md#singlesignon-Type-Group-DisplayName) in the _IAM Identity Center Identity Store API Reference_.

_Required_: Yes

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IdentityType`

Specifies whether the identity to map to the Amazon EMR Studio is a user or a
group.

_Required_: Yes

_Type_: String

_Allowed values_: `USER | GROUP`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SessionPolicyArn`

The Amazon Resource Name (ARN) for the session policy that will be applied to the user
or group. Session policies refine Studio user permissions without the need to use multiple
IAM user roles. For more information, see [Create an EMR Studio user role with session policies](../../../emr/latest/managementguide/emr-studio-user-role.md) in the _Amazon EMR Management Guide_.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws(-(cn|us-gov|iso-f|iso-e))?:iam::([0-9]{12})?:policy\/[^.]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StudioId`

The ID of the Amazon EMR Studio to which the user or group will be
mapped.

_Required_: Yes

_Type_: String

_Pattern_: `^es-[0-9A-Z]+`

_Minimum_: `4`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::EMR::WALWorkspace

All content copied from https://docs.aws.amazon.com/.
