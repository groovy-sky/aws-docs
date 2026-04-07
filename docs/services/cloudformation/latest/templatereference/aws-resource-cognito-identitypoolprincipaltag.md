This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::IdentityPoolPrincipalTag

A list of the identity pool principal tag assignments for attributes for access
control.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Cognito::IdentityPoolPrincipalTag",
  "Properties" : {
      "IdentityPoolId" : String,
      "IdentityProviderName" : String,
      "PrincipalTags" : Json,
      "UseDefaults" : Boolean
    }
}

```

### YAML

```yaml

Type: AWS::Cognito::IdentityPoolPrincipalTag
Properties:
  IdentityPoolId: String
  IdentityProviderName: String
  PrincipalTags: Json
  UseDefaults: Boolean

```

## Properties

`IdentityPoolId`

The identity pool that you want to associate with this principal tag map.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IdentityProviderName`

The identity pool identity provider (IdP) that you want to associate with this
principal tag map.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PrincipalTags`

A JSON-formatted list of user claims and the principal tags that you want to associate
with them. When Amazon Cognito requests credentials, it sets the value of the principal
tag to the value of the user's claim.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseDefaults`

Use a default set of mappings between claims and tags for this provider, instead of a
custom map.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the principal tag primary ID, like
`us-east-1:1cf667a2-49a6-454b-9e45-23199EXAMPLE|graph.facebook.com`.

## Examples

### Creating a new principal tag attribute map for an identity pool

The following example maps the claim `aud` to principal tag
`app_id` and the claim `sub` to `user_id`
in the identity pool `Example_pool.`

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: Cognito Principal Tags Example
Resources:
  ExamplePool:
    Type: AWS::Cognito::IdentityPool
    Properties:
      IdentityPoolName: example_pool
      AllowUnauthenticatedIdentities: true
      SupportedLoginProviders:
        graph.facebook.com: abcdExampleClientId
  PrincipalTags:
    Type: AWS::Cognito::IdentityPoolPrincipalTag
    Properties:
      IdentityPoolId: !Ref ExamplePool
      IdentityProviderName: graph.facebook.com
      PrincipalTags:
        app_id: aud
        user_id: sub
      UseDefaults: false
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::Cognito::IdentityPoolRoleAttachment
