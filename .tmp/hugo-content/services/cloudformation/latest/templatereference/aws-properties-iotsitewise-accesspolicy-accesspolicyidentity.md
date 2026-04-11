This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTSiteWise::AccessPolicy AccessPolicyIdentity

The identity (IAM Identity Center user, IAM Identity Center group, or IAM user) to which this access policy
applies.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IamRole" : IamRole,
  "IamUser" : IamUser,
  "User" : User
}

```

### YAML

```yaml

  IamRole:
    IamRole
  IamUser:
    IamUser
  User:
    User

```

## Properties

`IamRole`

An IAM role identity.

_Required_: No

_Type_: [IamRole](aws-properties-iotsitewise-accesspolicy-iamrole.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IamUser`

An IAM user identity.

_Required_: No

_Type_: [IamUser](aws-properties-iotsitewise-accesspolicy-iamuser.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`User`

An IAM Identity Center user identity.

_Required_: No

_Type_: [User](aws-properties-iotsitewise-accesspolicy-user.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::IoTSiteWise::AccessPolicy

AccessPolicyResource

All content copied from https://docs.aws.amazon.com/.
