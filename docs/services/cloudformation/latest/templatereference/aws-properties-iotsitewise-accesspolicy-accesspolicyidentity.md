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

_Type_: [IamRole](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotsitewise-accesspolicy-iamrole.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IamUser`

An IAM user identity.

_Required_: No

_Type_: [IamUser](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotsitewise-accesspolicy-iamuser.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`User`

An IAM Identity Center user identity.

_Required_: No

_Type_: [User](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotsitewise-accesspolicy-user.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::IoTSiteWise::AccessPolicy

AccessPolicyResource
