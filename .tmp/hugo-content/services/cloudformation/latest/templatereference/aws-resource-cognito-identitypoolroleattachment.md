This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::IdentityPoolRoleAttachment

The `AWS::Cognito::IdentityPoolRoleAttachment` resource manages the role
configuration for an Amazon Cognito identity pool.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Cognito::IdentityPoolRoleAttachment",
  "Properties" : {
      "IdentityPoolId" : String,
      "RoleMappings" : RoleMapping,
      "Roles" : String
    }
}

```

### YAML

```yaml

Type: AWS::Cognito::IdentityPoolRoleAttachment
Properties:
  IdentityPoolId: String
  RoleMappings:
    RoleMapping
  Roles: String

```

## Properties

`IdentityPoolId`

An identity pool ID in the format `REGION:GUID`.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RoleMappings`

How users for a specific identity provider are mapped to roles. This is a string to
the `RoleMapping` object map. The string identifies the identity provider.
For example: `graph.facebook.com` or
`cognito-idp.us-east-1.amazonaws.com/us-east-1_abcdefghi:app_client_id`.

If the `IdentityProvider` field isn't provided in this object, the string
is used as the identity provider name.

For more information, see the [RoleMapping property](../userguide/aws-properties-cognito-identitypoolroleattachment-rolemapping.md).

_Required_: No

_Type_: [RoleMapping](aws-properties-cognito-identitypoolroleattachment-rolemapping.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Roles`

The map of the roles associated with this pool. For a given role, the key is either
"authenticated" or "unauthenticated". The value is the role ARN.

_Required_: No

_Type_: String

_Pattern_: `^.+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `IdentityPoolId`, such as
`us-east-2:0d01f4d7-1305-4408-b437-12345EXAMPLE`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

The resource ID.

## Examples

### Setting the roles for an identity pool

The following example sets roles for an identity pool. It sets “authenticated”
and “unauthenticated” roles and maps two identity providers to them. The first
identity provider is “graph.facebook.com”. The second is using a reference to
set the identity provider name.

#### JSON

```json

{
   "IdentityPoolRoleAttachment":{
      "Type":"AWS::Cognito::IdentityPoolRoleAttachment",
      "Properties":{
         "IdentityPoolId":{
            "Ref":"IdentityPool"
         },
         "Roles":{
            "authenticated":{
               "Fn::GetAtt":[
                  "AuthenticatedRole",
                  "Arn"
               ]
            },
            "unauthenticated":{
               "Fn::GetAtt":[
                  "UnAuthenticatedRole",
                  "Arn"
               ]
            }
         },
         "RoleMappings":{
            "graph.facebook.com":{
               "IdentityProvider":"graph.facebook.com",
               "AmbiguousRoleResolution":"Deny",
               "Type":"Rules",
               "RulesConfiguration":{
                  "Rules":[
                     {
                        "Claim":"sub",
                        "MatchType":"Equals",
                        "RoleARN":{
                           "Fn::GetAtt":[
                              "AuthenticatedRole",
                              "Arn"
                           ]
                        },
                        "Value":"goodvalue"
                     }
                  ]
               }
            },
            "userpool1":{
               "IdentityProvider":{
                  "Ref":"CognitoUserPool"
               },
               "AmbiguousRoleResolution":"Deny",
               "Type":"Rules",
               "RulesConfiguration":{
                  "Rules":[
                     {
                        "Claim":"sub",
                        "MatchType":"Equals",
                        "RoleARN":{
                           "Fn::GetAtt":[
                              "AuthenticatedRole",
                              "Arn"
                           ]
                        },
                        "Value":"goodvalue"
                     }
                  ]
               }
            }
         }
      }
   }
}
```

#### YAML

```yaml

IdentityPoolRoleAttachment:
  Type: AWS::Cognito::IdentityPoolRoleAttachment
  Properties:
    IdentityPoolId: !Ref IdentityPool
    Roles:
      "authenticated": !GetAtt AuthenticatedRole.Arn
      "unauthenticated": !GetAtt UnAuthenticatedRole.Arn
    RoleMappings:
      "graph.facebook.com":
        IdentityProvider: "graph.facebook.com"
        AmbiguousRoleResolution: Deny
        Type: Rules
        RulesConfiguration:
          Rules:
            - Claim: "sub"
              MatchType: "Equals"
              RoleARN: !GetAtt AuthenticatedRole.Arn
              Value: "goodvalue"
      "userpool1":
        IdentityProvider: !Ref CognitoUserPool
        AmbiguousRoleResolution: Deny
        Type: Rules
        RulesConfiguration:
          Rules:
            - Claim: "sub"
              MatchType: "Equals"
              RoleARN: !GetAtt AuthenticatedRole.Arn
              Value: "goodvalue"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Cognito::IdentityPoolPrincipalTag

MappingRule

All content copied from https://docs.aws.amazon.com/.
