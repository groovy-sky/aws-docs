This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::UserProfile

Creates a user profile. A user profile represents a single user within a domain, and is the main way to
reference a "person" for the purposes of sharing, reporting, and other user-oriented features. This entity is created
when a user onboards to Amazon SageMaker Studio. If an administrator invites a person by email or imports them from
IAM Identity Center, a user profile is automatically created. A user profile is the primary holder of settings for an
individual user and has a reference to the user's private Amazon Elastic File System (EFS) home directory.

###### Note

If you're using IAM Identity Center authentication, a user in IAM Identity Center, or a group in IAM Identity Center containing that user, must be assigned to the Amazon SageMaker Studio application from the IAM Identity Center
Console to create a user profile. For more information about application assignment, see [Assign user access](../../../singlesignon/latest/userguide/assignuserstoapp.md). After assignment is complete,
a user profile can be created for that user in IAM Identity Center with AWS CloudFormation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SageMaker::UserProfile",
  "Properties" : {
      "DomainId" : String,
      "SingleSignOnUserIdentifier" : String,
      "SingleSignOnUserValue" : String,
      "Tags" : [ Tag, ... ],
      "UserProfileName" : String,
      "UserSettings" : UserSettings
    }
}

```

### YAML

```yaml

Type: AWS::SageMaker::UserProfile
Properties:
  DomainId: String
  SingleSignOnUserIdentifier: String
  SingleSignOnUserValue: String
  Tags:
    - Tag
  UserProfileName: String
  UserSettings:
    UserSettings

```

## Properties

`DomainId`

The domain ID.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SingleSignOnUserIdentifier`

A specifier for the type of value specified in SingleSignOnUserValue. Currently, the only supported value is
"UserName". If the Domain's AuthMode is IAM Identity Center, this field is required. If the Domain's AuthMode is not
IAM Identity Center, this field cannot be specified.

_Required_: No

_Type_: String

_Pattern_: `UserName`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SingleSignOnUserValue`

The username of the associated AWS Single Sign-On User for this UserProfile. If the Domain's
AuthMode is IAM Identity Center, this field is required, and must match a valid username of a user in your directory.
If the Domain's AuthMode is not IAM Identity Center, this field cannot be specified.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of key-value pairs to apply to this resource.

Tags that you specify for the User Profile are also added to all apps that the User Profile launches.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-sagemaker-userprofile-tag.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UserProfileName`

The user profile name.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UserSettings`

A collection of settings that apply to users of Amazon SageMaker Studio.

_Required_: No

_Type_: [UserSettings](aws-properties-sagemaker-userprofile-usersettings.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Domain ID and the user profile name, such as `d-xxxxxxxxxxxx` and
`my-user-profile`, respectively.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`UserProfileArn`

The Amazon Resource Name (ARN) of the user profile, such as
`arn:aws:sagemaker:region:account-id:user-profile/domain-id/user-profile-name`.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AppLifecycleManagement

All content copied from https://docs.aws.amazon.com/.
