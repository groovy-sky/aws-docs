This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSO::ApplicationAssignment

A structure that describes an assignment of a principal to an application.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SSO::ApplicationAssignment",
  "Properties" : {
      "ApplicationArn" : String,
      "PrincipalId" : String,
      "PrincipalType" : String
    }
}

```

### YAML

```yaml

Type: AWS::SSO::ApplicationAssignment
Properties:
  ApplicationArn: String
  PrincipalId: String
  PrincipalType: String

```

## Properties

`ApplicationArn`

The ARN of the application that has principals assigned.

_Required_: Yes

_Type_: String

_Pattern_: `arn:aws(-[a-z]{1,5}){0,3}:sso::\d{12}:application/(sso)?ins-[a-zA-Z0-9-.]{16}/apl-[a-zA-Z0-9]{16}`

_Minimum_: `10`

_Maximum_: `1224`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PrincipalId`

The unique identifier of the principal assigned to the application.

_Required_: Yes

_Type_: String

_Pattern_: `^([0-9a-f]{10}-|)[A-Fa-f0-9]{8}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{12}$`

_Minimum_: `1`

_Maximum_: `47`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PrincipalType`

The type of the principal assigned to the application.

_Required_: Yes

_Type_: String

_Allowed values_: `USER | GROUP`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a generated ID, combined by all fields with the delimiter
`|`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

### Creating a new application assignment for IAM Identity Center

The following example grants the user permission to access the example
application.

#### JSON

```json

"ApplicationAssignment": {
    "Type": "AWS::SSO::ApplicationAssignment",
    "Properties": {
        "ApplicationArn": "arn:aws:sso:::application/ssoins-exampleapplicationid",
        "PrincipalID": "user_id",
        "PrincipalType": "USER"
    }
}
```

#### YAML

```yaml

ApplicationAssignment:
    Type: AWS::SSO::ApplicationAssignment
    Properties:
        ApplicationArn: 'arn:aws:sso:::application/ssoins-exampleapplicationid'
        PrincipalID: 'user_id'
        PrincipalType: 'USER'

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::SSO::Assignment

All content copied from https://docs.aws.amazon.com/.
