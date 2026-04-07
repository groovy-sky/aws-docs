This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IAM::UserToGroupAddition

Adds the specified user to the specified group.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IAM::UserToGroupAddition",
  "Properties" : {
      "GroupName" : String,
      "Users" : [ String, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IAM::UserToGroupAddition
Properties:
  GroupName: String
  Users:
    - String

```

## Properties

`GroupName`

The name of the group to update.

This parameter allows (through its [regex pattern](http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric
characters with no spaces. You can also include any of the following characters: \_+=,.@-

_Required_: Yes

_Type_: String

_Pattern_: `[\w+=,.@-]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Users`

A list of the names of the users that you want to add to the group.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name.

For example:

`{ "Ref": "MyUserToGroupAddition" }`

For the `AWS::IAM::UserToGroupAddition` resource with the logical ID
`MyUserToGroupAddition`, `Ref` will return the AWS resource name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

The stable and unique string identifying the group. For more information about IDs, see
[IAM\
identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the _IAM User Guide_.

## See also

- To view `AWS::IAM::UserToGroupAddition` template example snippets,
see [Add Users to a Group](../userguide/quickref-iam.md#scenario-iam-addusertogroup).

- [AddUserToGroup](https://docs.aws.amazon.com/IAM/latest/APIReference/API_AddUserToGroup.html) in the _AWS Identity and Access Management API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::IAM::UserPolicy

AWS::IAM::VirtualMFADevice
