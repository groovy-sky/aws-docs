This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::UserPoolGroup

A user pool group. Contains details about the group and the way that it contributes to
IAM role decisions with identity pools. Identity pools can make decisions about the IAM
role to assign based on groups: users get credentials for the role associated with their
highest-priority group.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Cognito::UserPoolGroup",
  "Properties" : {
      "Description" : String,
      "GroupName" : String,
      "Precedence" : Integer,
      "RoleArn" : String,
      "UserPoolId" : String
    }
}

```

### YAML

```yaml

Type: AWS::Cognito::UserPoolGroup
Properties:
  Description: String
  GroupName: String
  Precedence: Integer
  RoleArn: String
  UserPoolId: String

```

## Properties

`Description`

A description of the group that you're creating.

_Required_: No

_Type_: String

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GroupName`

A name for the group. This name must be unique in your user pool.

_Required_: No

_Type_: String

_Pattern_: `[\p{L}\p{M}\p{S}\p{N}\p{P}]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Precedence`

A non-negative integer value that specifies the precedence of this group relative to
the other groups that a user can belong to in the user pool. Zero is the highest
precedence value. Groups with lower `Precedence` values take precedence over
groups with higher or null `Precedence` values. If a user belongs to two or
more groups, it is the group with the lowest precedence value whose role ARN is given in
the user's tokens for the `cognito:roles` and
`cognito:preferred_role` claims.

Two groups can have the same `Precedence` value. If this happens, neither
group takes precedence over the other. If two groups with the same
`Precedence` have the same role ARN, that role is used in the
`cognito:preferred_role` claim in tokens for users in each group. If the
two groups have different role ARNs, the `cognito:preferred_role` claim isn't
set in users' tokens.

The default `Precedence` value is null. The maximum `Precedence`
value is `2^31-1`.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The Amazon Resource Name (ARN) for the IAM role that you want to associate with the
group. A group role primarily declares a preferred role for the credentials that you get
from an identity pool. Amazon Cognito ID tokens have a `cognito:preferred_role` claim
that presents the highest-precedence group that a user belongs to. Both ID and access
tokens also contain a `cognito:groups` claim that list all the groups that a
user is a member of.

_Required_: No

_Type_: String

_Pattern_: `arn:[\w+=/,.@-]+:[\w+=/,.@-]+:([\w+=/,.@-]*)?:[0-9]+:[\w+=/,.@-]+(:[\w+=/,.@-]+)?(:[\w+=/,.@-]+)?`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserPoolId`

The ID of the user pool where you want to create a user group.

_Required_: Yes

_Type_: String

_Pattern_: `[\w-]+_[0-9a-zA-Z]+`

_Minimum_: `1`

_Maximum_: `55`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the user pool group. For example:
`Admins`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

### Creating a new user group in a user pool

The following example creates a group named "ExampleGroup" with an IAM role and a precedence of 1.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "Provision a new user group to an Amazon Cognito user pool\n",
    "Resources": {
        "ExampleGroup": {
            "Properties": {
                "Description": "My Example Group",
                "GroupName": "ExampleGroup",
                "Precedence": 1,
                "RoleArn": "arn:aws:iam::123456789012:role/my-test-cognito-role",
                "UserPoolId": "us-west-2_EXAMPLE"
            },
            "Type": "AWS::Cognito::UserPoolGroup"
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: "2010-09-09"

Description: |
  Provision a new user group to an Amazon Cognito user pool

Resources:
  ExampleGroup:
    Type: AWS::Cognito::UserPoolGroup
    Properties:
      Description: My Example Group
      GroupName: ExampleGroup
      Precedence: 1
      RoleArn: arn:aws:iam::123456789012:role/my-test-cognito-role
      UserPoolId: us-west-2_EXAMPLE
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CustomDomainConfigType

AWS::Cognito::UserPoolIdentityProvider
