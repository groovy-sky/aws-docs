This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::Instance

_This is a preview release for Amazon Connect. It is subject to_
_change._

Initiates an Amazon Connect instance with all the supported channels enabled. It
does not attach any storage, such as Amazon Simple Storage Service (Amazon S3) or Amazon
Kinesis.

Amazon Connect enforces a limit on the total number of instances that you can
create or delete in 30 days. If you exceed this limit, you will get an error message
indicating there has been an excessive number of attempts at creating or deleting
instances. You must wait 30 days before you can restart creating and deleting instances
in your account.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Connect::Instance",
  "Properties" : {
      "Attributes" : Attributes,
      "DirectoryId" : String,
      "IdentityManagementType" : String,
      "InstanceAlias" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Connect::Instance
Properties:
  Attributes:
    Attributes
  DirectoryId: String
  IdentityManagementType: String
  InstanceAlias: String
  Tags:
    - Tag

```

## Properties

`Attributes`

A toggle for an individual feature at the instance level.

_Required_: Yes

_Type_: [Attributes](aws-properties-connect-instance-attributes.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DirectoryId`

The identifier for the directory.

_Required_: No

_Type_: String

_Pattern_: `^d-[0-9a-f]{10}$`

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IdentityManagementType`

The identity management type.

_Required_: Yes

_Type_: String

_Allowed values_: `SAML | CONNECT_MANAGED | EXISTING_DIRECTORY`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InstanceAlias`

The alias of instance. `InstanceAlias` is only required when
`IdentityManagementType` is `CONNECT_MANAGED` or
`SAML`. `InstanceAlias` is not required when
`IdentityManagementType` is `EXISTING_DIRECTORY`.

_Required_: No

_Type_: String

_Pattern_: `^(?!d-)([\da-zA-Z]+)([-]*[\da-zA-Z])*$`

_Minimum_: `1`

_Maximum_: `45`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of key-value pairs to apply to this resource.

_Required_: No

_Type_: Array of [Tag](aws-properties-connect-instance-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the Amazon Connect instance.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the instance.

`CreatedTime`

When the instance was created.

`Id`

The identifier of the Amazon Connect instance. You can find the instanceId in the
ARN of the instance.

`InstanceStatus`

The state of the instance.

`ServiceRole`

The service role of the instance.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Attributes

All content copied from https://docs.aws.amazon.com/.
