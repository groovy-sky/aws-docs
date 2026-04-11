This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::UserPoolUserToGroupAttachment

Adds a user to a group. A user who is in a group can present a preferred-role claim to
an identity pool, and populates a `cognito:groups` claim to their access and
identity tokens.

###### Note

Amazon Cognito evaluates AWS Identity and Access Management (IAM) policies in requests for this API operation. For
this operation, you must use IAM credentials to authorize requests, and you must
grant yourself the corresponding IAM permission in a policy.

###### Learn more

- [Signing AWS API Requests](../../../iam/latest/userguide/reference-aws-signing.md)

- [Using the Amazon Cognito user pools API and user pool endpoints](../../../cognito/latest/developerguide/user-pools-api-operations.md)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Cognito::UserPoolUserToGroupAttachment",
  "Properties" : {
      "GroupName" : String,
      "Username" : String,
      "UserPoolId" : String
    }
}

```

### YAML

```yaml

Type: AWS::Cognito::UserPoolUserToGroupAttachment
Properties:
  GroupName: String
  Username: String
  UserPoolId: String

```

## Properties

`GroupName`

The name of the group that you want to add your user to.

_Required_: Yes

_Type_: String

_Pattern_: `[\p{L}\p{M}\p{S}\p{N}\p{P}]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Username`

The user's username.

_Required_: Yes

_Type_: String

_Pattern_: `[\p{L}\p{M}\p{S}\p{N}\p{P}]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UserPoolId`

The ID of the user pool that contains the group that you want to add the user
to.

_Required_: Yes

_Type_: String

_Pattern_: `[\w-]+_[0-9a-zA-Z]+`

_Minimum_: `1`

_Maximum_: `55`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a generated ID, such as
`UserToGroupAttachment-YejJvzrEXAMPLE`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

### Adding users to a group

The following template adds the users "testuser1" and "testuser2" to the user
group "ExampleGroup" in the requested user pool.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "Add a user to a group in an Amazon Cognito user pool\n",
    "Resources": {
        "TestUser2ToExampleGroup": {
            "Properties": {
                "GroupName": "ExampleGroup",
                "Username": "testuser1",
                "UserPoolId": "us-west-2_EXAMPLE"
            },
            "Type": "AWS::Cognito::UserPoolUserToGroupAttachment"
        },
        "TestUserToExampleGroup": {
            "Properties": {
                "GroupName": "ExampleGroup",
                "Username": "testuser2",
                "UserPoolId": "us-west-2_EXAMPLE"
            },
            "Type": "AWS::Cognito::UserPoolUserToGroupAttachment"
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: "2010-09-09"

Description: |
  Add a user to a group in an Amazon Cognito user pool

Resources:
  TestUser2ToExampleGroup:
    Type: AWS::Cognito::UserPoolUserToGroupAttachment
    Properties:
      GroupName: ExampleGroup
      Username: testuser1
      UserPoolId: us-west-2_EXAMPLE

  TestUserToExampleGroup:
    Type: AWS::Cognito::UserPoolUserToGroupAttachment
    Properties:
      GroupName: ExampleGroup
      Username: testuser2
      UserPoolId: us-west-2_EXAMPLE
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AttributeType

Next

All content copied from https://docs.aws.amazon.com/.
