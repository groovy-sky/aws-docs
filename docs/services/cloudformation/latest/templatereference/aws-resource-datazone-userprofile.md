This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::UserProfile

The user type of the user for which the user profile is created.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DataZone::UserProfile",
  "Properties" : {
      "DomainIdentifier" : String,
      "Status" : String,
      "UserIdentifier" : String,
      "UserType" : String
    }
}

```

### YAML

```yaml

Type: AWS::DataZone::UserProfile
Properties:
  DomainIdentifier: String
  Status: String
  UserIdentifier: String
  UserType: String

```

## Properties

`DomainIdentifier`

The identifier of a Amazon DataZone domain in which a user profile exists.

_Required_: Yes

_Type_: String

_Pattern_: `^dzd[-_][a-zA-Z0-9_-]{1,36}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Status`

The status of the user profile.

_Required_: No

_Type_: String

_Allowed values_: `ASSIGNED | NOT_ASSIGNED | ACTIVATED | DEACTIVATED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserIdentifier`

The identifier of the user for which the user profile is created.

_Required_: Yes

_Type_: String

_Pattern_: `(^([0-9a-f]{10}-|)[A-Fa-f0-9]{8}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{12}$|^[a-zA-Z_0-9+=,.@-]+$|^arn:aws:iam::\d{12}:.+$)`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UserType`

The user type of the user for which the user profile is created.

_Required_: No

_Type_: String

_Allowed values_: `IAM_USER | IAM_ROLE | SSO_USER`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a string containing pipe-separated `DomainId` and
`UserProfileId` that uniquely identify the user profile. For example: `{ "Ref": "MyUserProfile"
    }` for the resource with the logical ID `MyUserProfile`, `Ref` returns
`DomainId|UserProfileId`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`DomainId`

The identifier of a Amazon DataZone domain in which a user profile exists.

`Id`

The ID of the user profile.

`Type`

The type of the user profile.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SubscriptionTargetForm

IamUserProfileDetails
