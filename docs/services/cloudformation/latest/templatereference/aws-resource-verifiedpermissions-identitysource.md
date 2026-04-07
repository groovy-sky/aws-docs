This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VerifiedPermissions::IdentitySource

Creates or updates a reference to Amazon Cognito as an external identity
provider.

If you are creating a new identity source, then you must specify a
`Configuration`. If you are updating an existing identity source, then
you must specify an `UpdateConfiguration`.

After you create an identity source, you can use the identities provided by the IdP as
proxies for the principal in authorization queries that use the [IsAuthorizedWithToken](https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_IsAuthorizedWithToken.html) operation. These identities take the form of tokens
that contain claims about the user, such as IDs, attributes and group memberships.
Amazon Cognito provides both identity tokens and access tokens, and Verified Permissions can use either or both. Any combination of identity and access tokens
results in the same Cedar principal. Verified Permissions automatically translates the
information about the identities into the standard Cedar attributes that can be
evaluated by your policies. Because the Amazon Cognito identity and access tokens
can contain different information, the tokens you choose to use determine the attributes
that are available to access in the Cedar principal from your policies.

Amazon Cognito Identity is not available in all of the same AWS Regions as Amazon Verified Permissions. Because of this, the
`AWS::VerifiedPermissions::IdentitySource` type is not available to
create from CloudFormation in Regions where Amazon Cognito Identity is not
currently available. Users can still create
`AWS::VerifiedPermissions::IdentitySource` in those Regions, but only
from the AWS CLI, Amazon Verified Permissions SDK, or from the AWS console.

###### Note

To reference a user from this identity source in your Cedar policies, use the
following syntax.

_IdentityType::"<CognitoUserPoolIdentifier>\|<CognitoClientId>_

Where `IdentityType` is the string that you provide to the
`PrincipalEntityType` parameter for this operation. The
`CognitoUserPoolId` and `CognitoClientId` are defined by
the Amazon Cognito user pool.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::VerifiedPermissions::IdentitySource",
  "Properties" : {
      "Configuration" : IdentitySourceConfiguration,
      "PolicyStoreId" : String,
      "PrincipalEntityType" : String
    }
}

```

### YAML

```yaml

Type: AWS::VerifiedPermissions::IdentitySource
Properties:
  Configuration:
    IdentitySourceConfiguration
  PolicyStoreId: String
  PrincipalEntityType: String

```

## Properties

`Configuration`

Contains configuration information used when creating a new identity source.

_Required_: Yes

_Type_: [IdentitySourceConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-verifiedpermissions-identitysource-identitysourceconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PolicyStoreId`

Specifies the ID of the policy store in which you want to store this identity source. Only policies and
requests made using this policy store can reference identities from the identity provider
configured in the new identity source.

To specify a policy store, use its ID or alias name. When using an alias name, prefix it with `policy-store-alias/`. For example:

- ID: `PSEXAMPLEabcdefg111111`

- Alias name: `policy-store-alias/example-policy-store`

To view aliases, use [ListPolicyStoreAliases](https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_ListPolicyStoreAliases.html).

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9-]*$`

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PrincipalEntityType`

Specifies the namespace and data type of the principals generated for identities
authenticated by the new identity source.

_Required_: No

_Type_: String

_Pattern_: `^.*$`

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the unique id of the new identity source. For
example:

`{ "Ref": "ISEXAMPLEabcdefg111111" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`IdentitySourceId`

The unique ID of the new or updated identity store.

## Examples

### Creating an identity source for a Amazon Cognito user pool

The following example creates an identity source in the specified policy store
that points to a Amazon Cognito user pool using the specified client ID. The
new identity source returns objects of the specified data type.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "AWS CloudFormation sample template for creating an identity source for Verified Permissions.",
    "Parameters": {
        "PolicyStoreId": {
            "Type": "String"
        },
        "UserPoolArn": {
            "Type": "String"
        },
        "ClientIds": {
            "Type": "List<String>"
        },
        "PrincipalEntityType": {
            "Type": "String"
        }
    },
    "Resources": {
        "IdentitySource": {
            "Type": "AWS::VerifiedPermissions::IdentitySource",
            "Properties": {
                "PolicyStoreId": {
                    "Ref": "PolicyStoreId"
                },
                "Configuration": {
                    "CognitoUserPoolConfiguration": {
                        "UserPoolArn": {
                            "Ref": "UserPoolArn"
                        },
                        "ClientIds": {
                            "Ref": "ClientIds"
                        }
                    }
                },
                "PrincipalEntityType": {
                    "Ref": "PrincipalEntityType"
                }
            }
        }
    },
    "Outputs": {
        "IdentitySourceId": {
            "Value": {
                "Fn::GetAtt": [
                    "IdentitySource",
                    "IdentitySourceId"
                ]
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: >-
  Description": "AWS CloudFormation sample template for creating an identity source for Verified Permissions
Parameters:
  PolicyStoreId:
    Type: String
  UserPoolArn:
    Type: String
  ClientIds:
    Type: List<String>
  PrincipalEntityType:
    Type: String
Resources:
  IdentitySource:
    Type: AWS::VerifiedPermissions::IdentitySource
    Properties:
      PolicyStoreId: !Ref PolicyStoreId
      Configuration:
        CognitoUserPoolConfiguration:
          UserPoolArn: !Ref UserPoolArn
          ClientIds: !Ref ClientIds
      PrincipalEntityType: !Ref PrincipalEntityType
Outputs:
  IdentitySourceId:
    Value: !GetAtt IdentitySource.IdentitySourceId
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon Verified Permissions

CognitoGroupConfiguration
