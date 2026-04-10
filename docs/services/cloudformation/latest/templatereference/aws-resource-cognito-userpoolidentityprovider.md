This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::UserPoolIdentityProvider

The `AWS::Cognito::UserPoolIdentityProvider` resource creates an identity
provider for a user pool.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Cognito::UserPoolIdentityProvider",
  "Properties" : {
      "AttributeMapping" : {Key: Value, ...},
      "IdpIdentifiers" : [ String, ... ],
      "ProviderDetails" : {Key: Value, ...},
      "ProviderName" : String,
      "ProviderType" : String,
      "UserPoolId" : String
    }
}

```

### YAML

```yaml

Type: AWS::Cognito::UserPoolIdentityProvider
Properties:
  AttributeMapping:
    Key: Value
  IdpIdentifiers:
    - String
  ProviderDetails:
    Key: Value
  ProviderName: String
  ProviderType: String
  UserPoolId: String

```

## Properties

`AttributeMapping`

A mapping of IdP attributes to standard and custom user pool attributes. Specify a
user pool attribute as the key of the key-value pair, and the IdP attribute claim name
as the value.

_Required_: No

_Type_: Object of String

_Pattern_: `^([\p{L}\p{Z}\p{N}_.:/=+\-@]*)$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IdpIdentifiers`

An array of IdP identifiers, for example `"IdPIdentifiers": [ "MyIdP", "MyIdP2"
                ]`. Identifiers are friendly names that you can pass in the
`idp_identifier` query parameter of requests to the [Authorize endpoint](../../../cognito/latest/developerguide/authorization-endpoint.md) to silently redirect to sign-in with the associated IdP.
Identifiers in a domain format also enable the use of [email-address matching with SAML providers](../../../cognito/latest/developerguide/cognito-user-pools-managing-saml-idp-naming.md).

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProviderDetails`

The scopes, URLs, and identifiers for your external identity provider. The following
examples describe the provider detail keys for each IdP type. These values and their
schema are subject to change. Social IdP `authorize_scopes` values must match
the values listed here.

OpenID Connect (OIDC)

Amazon Cognito accepts the following elements when it can't discover endpoint
URLs from `oidc_issuer`: `attributes_url`,
`authorize_url`, `jwks_uri`,
`token_url`.

Create or update request: `"ProviderDetails": {
                    "attributes_request_method": "GET", "attributes_url":
                    "https://auth.example.com/userInfo", "authorize_scopes": "openid profile
                    email", "authorize_url": "https://auth.example.com/authorize",
                    "client_id": "1example23456789", "client_secret":
                    "provider-app-client-secret", "jwks_uri":
                    "https://auth.example.com/.well-known/jwks.json", "oidc_issuer":
                    "https://auth.example.com", "token_url": "https://example.com/token"
                    }`

Describe response: `"ProviderDetails": { "attributes_request_method":
                    "GET", "attributes_url": "https://auth.example.com/userInfo",
                    "attributes_url_add_attributes": "false", "authorize_scopes": "openid
                    profile email", "authorize_url": "https://auth.example.com/authorize",
                    "client_id": "1example23456789", "client_secret":
                    "provider-app-client-secret", "jwks_uri":
                    "https://auth.example.com/.well-known/jwks.json", "oidc_issuer":
                    "https://auth.example.com", "token_url": "https://example.com/token"
                    }`

SAML

Create or update request with Metadata URL: `"ProviderDetails": { "IDPInit": "true",
                    "IDPSignout": "true", "EncryptedResponses" : "true", "MetadataURL":
                    "https://auth.example.com/sso/saml/metadata", "RequestSigningAlgorithm":
                    "rsa-sha256" }`

Create or update request with Metadata file: `"ProviderDetails": { "IDPInit": "true",
                    "IDPSignout": "true", "EncryptedResponses" : "true",
                    "MetadataFile": "[metadata XML]", "RequestSigningAlgorithm":
                    "rsa-sha256" }`

The value of `MetadataFile` must be the plaintext metadata document with all
quote (") characters escaped by backslashes.

Describe response: `"ProviderDetails": { "IDPInit": "true",
                    "IDPSignout": "true", "EncryptedResponses" : "true", "ActiveEncryptionCertificate": "[certificate]",
                    "MetadataURL": "https://auth.example.com/sso/saml/metadata", "RequestSigningAlgorithm":
                    "rsa-sha256", "SLORedirectBindingURI":
                    "https://auth.example.com/slo/saml", "SSORedirectBindingURI":
                    "https://auth.example.com/sso/saml" }`

LoginWithAmazon

Create or update request: `"ProviderDetails": { "authorize_scopes":
                    "profile postal_code", "client_id":
                    "amzn1.application-oa2-client.1example23456789", "client_secret":
                    "provider-app-client-secret"`

Describe response: `"ProviderDetails": { "attributes_url":
                    "https://api.amazon.com/user/profile", "attributes_url_add_attributes":
                    "false", "authorize_scopes": "profile postal_code", "authorize_url":
                    "https://www.amazon.com/ap/oa", "client_id":
                    "amzn1.application-oa2-client.1example23456789", "client_secret":
                    "provider-app-client-secret", "token_request_method": "POST",
                    "token_url": "https://api.amazon.com/auth/o2/token" }`

Google

Create or update request: `"ProviderDetails": { "authorize_scopes":
                    "email profile openid", "client_id":
                    "1example23456789.apps.googleusercontent.com", "client_secret":
                    "provider-app-client-secret" }`

Describe response: `"ProviderDetails": { "attributes_url":
                    "https://people.googleapis.com/v1/people/me?personFields=",
                    "attributes_url_add_attributes": "true", "authorize_scopes": "email
                    profile openid", "authorize_url":
                    "https://accounts.google.com/o/oauth2/v2/auth", "client_id":
                    "1example23456789.apps.googleusercontent.com", "client_secret":
                    "provider-app-client-secret", "oidc_issuer":
                    "https://accounts.google.com", "token_request_method": "POST",
                    "token_url": "https://www.googleapis.com/oauth2/v4/token"
                }`

SignInWithApple

Create or update request: `"ProviderDetails": { "authorize_scopes":
                    "email name", "client_id": "com.example.cognito", "private_key": "1EXAMPLE",
                    "key_id": "2EXAMPLE", "team_id": "3EXAMPLE" }`

Describe response: `"ProviderDetails": {
                    "attributes_url_add_attributes": "false", "authorize_scopes": "email
                    name", "authorize_url": "https://appleid.apple.com/auth/authorize",
                    "client_id": "com.example.cognito", "key_id": "1EXAMPLE", "oidc_issuer":
                    "https://appleid.apple.com", "team_id": "2EXAMPLE",
                    "token_request_method": "POST", "token_url":
                    "https://appleid.apple.com/auth/token" }`

Facebook

Create or update request: `"ProviderDetails": { "api_version": "v17.0",
            "authorize_scopes": "public_profile, email", "client_id": "1example23456789",
            "client_secret": "provider-app-client-secret" }`

Describe response: `"ProviderDetails":
            { "api_version": "v17.0", "attributes_url": "https://graph.facebook.com/v17.0/me?fields=",
            "attributes_url_add_attributes": "true", "authorize_scopes": "public_profile, email",
            "authorize_url": "https://www.facebook.com/v17.0/dialog/oauth", "client_id":
            "1example23456789", "client_secret": "provider-app-client-secret", "token_request_method":
            "GET", "token_url": "https://graph.facebook.com/v17.0/oauth/access_token" }`

_Required_: Yes

_Type_: Object of String

_Pattern_: `^([\p{L}\p{Z}\p{N}_.:/=+\-@]*)$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProviderName`

The name that you want to assign to the IdP. You can pass the identity provider name
in the `identity_provider` query parameter of requests to the [Authorize endpoint](../../../cognito/latest/developerguide/authorization-endpoint.md) to silently redirect to sign-in with the associated
IdP.

_Required_: Yes

_Type_: String

_Pattern_: `[^_\p{Z}][\p{L}\p{M}\p{S}\p{N}\p{P}][^_\p{Z}]+`

_Minimum_: `1`

_Maximum_: `32`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProviderType`

The type of IdP that you want to add. Amazon Cognito supports OIDC, SAML 2.0, Login With
Amazon, Sign In With Apple, Google, and Facebook IdPs.

_Required_: Yes

_Type_: String

_Allowed values_: `SAML | Facebook | Google | LoginWithAmazon | SignInWithApple | OIDC`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UserPoolId`

The Id of the user pool where you want to create an IdP.

_Required_: Yes

_Type_: String

_Pattern_: `[\w-]+_[0-9a-zA-Z]+`

_Minimum_: `1`

_Maximum_: `55`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns physicalResourceId, which is “ProviderName". For
example:

`{ "Ref": "testProvider" }`

For the Amazon Cognito identity provider `testProvider`, Ref returns the
name of the identity provider.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

- [Creating a new Login with Amazon identity provider](#aws-resource-cognito-userpoolidentityprovider--examples--Creating_a_new_Login_with_Amazon_identity_provider)

- [Creating a new Google identity provider](#aws-resource-cognito-userpoolidentityprovider--examples--Creating_a_new_Google_identity_provider)

- [Creating a new Facebook identity provider](#aws-resource-cognito-userpoolidentityprovider--examples--Creating_a_new_Facebook_identity_provider)

- [Creating a new Sign in with Apple identity provider](#aws-resource-cognito-userpoolidentityprovider--examples--Creating_a_new_Sign_in_with_Apple_identity_provider)

- [Creating a new OIDC identity provider](#aws-resource-cognito-userpoolidentityprovider--examples--Creating_a_new_OIDC_identity_provider)

- [Creating a new SAML identity provider](#aws-resource-cognito-userpoolidentityprovider--examples--Creating_a_new_SAML_identity_provider)

### Creating a new Login with Amazon identity provider

The following example creates a Login with Amazon identity provider in the
referenced user pool.

#### JSON

```json

{
  "UserPoolIdentityProvider": {
    "Type": "AWS::Cognito::UserPoolIdentityProvider",
    "Properties": {
      "UserPoolId": {
        "Ref": "UserPool"
      },
      "ProviderName": "LoginWithAmazon",
      "ProviderDetails": {
        "client_id": "YourLoginWithAmazonAppId",
        "client_secret": "YourLoginWithAmazonAppSecret",
        "authorize_scopes": "profile postal_code"
      },
      "ProviderType": "LoginWithAmazon",
      "AttributeMapping": {
        "email": "email"
      }
    }
  }
}
```

#### YAML

```yaml

UserPoolIdentityProvider:
  Type: AWS::Cognito::UserPoolIdentityProvider
  Properties:
    UserPoolId: !Ref UserPool
    ProviderName: "LoginWithAmazon"
    ProviderDetails:
      client_id: "YourLoginWithAmazonAppId"
      client_secret: "YourLoginWithAmazonAppSecret"
      authorize_scopes: "profile postal_code"
    ProviderType: "LoginWithAmazon"
    AttributeMapping:
      email: "email"
```

### Creating a new Google identity provider

The following example creates a Google identity provider in the referenced
user pool.

#### JSON

```json

{
  "UserPoolIdentityProvider": {
    "Type": "AWS::Cognito::UserPoolIdentityProvider",
    "Properties": {
      "UserPoolId": {
        "Ref": "UserPool"
      },
      "ProviderName": "Google",
      "ProviderDetails": {
        "client_id": "YourGoogleAppId",
        "client_secret": "YourGoogleAppSecret",
        "authorize_scopes": "profile email openid"
      },
      "ProviderType": "Google",
      "AttributeMapping": {
        "email": "email"
      }
    }
  }
}
```

#### YAML

```yaml

UserPoolIdentityProvider:
  Type: AWS::Cognito::UserPoolIdentityProvider
  Properties:
    UserPoolId: !Ref UserPool
    ProviderName: "Google"
    ProviderDetails:
      client_id: "YourGoogleAppId"
      client_secret: "YourGoogleAppSecret"
      authorize_scopes: "profile email openid"
    ProviderType: "Google"
    AttributeMapping:
      email: "email"
```

### Creating a new Facebook identity provider

The following example creates a Facebook identity provider in the referenced
user pool.

#### JSON

```json

{
  "UserPoolIdentityProvider": {
    "Type": "AWS::Cognito::UserPoolIdentityProvider",
    "Properties": {
      "UserPoolId": {
        "Ref": "UserPool"
      },
      "ProviderName": "Facebook",
      "ProviderDetails": {
        "client_id": "YourFacebookAppId",
        "client_secret": "YourFacebookAppSecret",
        "authorize_scopes": "public_profile,email"
      },
      "ProviderType": "Facebook",
      "AttributeMapping": {
        "email": "email"
      }
    }
  }
}
```

#### YAML

```yaml

UserPoolIdentityProvider:
  Type: AWS::Cognito::UserPoolIdentityProvider
  Properties:
    UserPoolId: !Ref UserPool
    ProviderName: "Facebook"
    ProviderDetails:
      client_id: "YourFacebookAppId"
      client_secret: "YourFacebookAppSecret"
      authorize_scopes: "public_profile,email"
    ProviderType: "Facebook"
    AttributeMapping:
      email: "email"
```

### Creating a new Sign in with Apple identity provider

The following example creates a Sign in with Apple identity provider in the
referenced user pool.

#### JSON

```json

{
  "UserPoolIdentityProvider": {
    "Type": "AWS::Cognito::UserPoolIdentityProvider",
    "Properties": {
      "UserPoolId": {
        "Ref": "UserPool"
      },
      "ProviderName": "SignInWithApple",
      "ProviderDetails": {
        "client_id": "YourAppleServicesId",
        "team_id": "YourAppleTeamId",
        "key_id": "YourApplePrivateKeyID",
        "private_key": "YourApplePrivateKey",
        "authorize_scopes": "public_profile,email"
      },
      "ProviderType": "SignInWithApple",
      "AttributeMapping": {
        "email": "email"
      }
    }
  }
}
```

#### YAML

```yaml

UserPoolIdentityProvider:
  Type: AWS::Cognito::UserPoolIdentityProvider
  Properties:
    UserPoolId: !Ref UserPool
    ProviderName: "SignInWithApple"
    ProviderDetails:
      client_id: "YourSign"
      team_id: "YourAppleTeamId"
      key_id: "YourApplePrivateKeyID"
      private_key: "YourApplePrivateKey"
      authorize_scopes: "public_profile,email"
    ProviderType: "SignInWithApple"
    AttributeMapping:
      email: "email"
```

### Creating a new OIDC identity provider

The following example creates the OIDC identity provider
"YourOIDCProviderName" in the referenced user pool.

#### JSON

```json

{
  "UserPoolIdentityProvider": {
    "Type": "AWS::Cognito::UserPoolIdentityProvider",
    "Properties": {
      "UserPoolId": {
        "Ref": "UserPool"
      },
      "ProviderName": "YourOIDCProviderName",
      "ProviderDetails": {
        "client_id": "YourOIDCClientId",
        "client_secret": "YourOIDCClientSecret",
        "attributes_request_method": "GET",
        "oidc_issuer": "YourOIDCIssuerURL",
        "authorize_scopes": "email profile openid"
      },
      "ProviderType": "OIDC",
      "AttributeMapping": {
        "email": "email"
      },
      "IdpIdentifiers": [
        "IdpIdentifier"
      ]
    }
  }
}
```

#### YAML

```yaml

UserPoolIdentityProvider:
  Type: AWS::Cognito::UserPoolIdentityProvider
  Properties:
    UserPoolId: !Ref UserPool
    ProviderName: "YourOIDCProviderName"
    ProviderDetails:
      client_id: "YourOIDCClientId"
      client_secret: "YourOIDCClientSecret"
      attributes_request_method: "GET"
      oidc_issuer: "YourOIDCIssuerURL"
      authorize_scopes: "email profile openid"
    ProviderType: "OIDC"
    AttributeMapping:
      email: "email"
    IdpIdentifiers:
      - "IdpIdentifier"
```

### Creating a new SAML identity provider

The following example creates a SAML identity provider "YourProviderName" in
the referenced user pool.

#### JSON

```json

{
   "UserPoolIdentityProvider": {
      "Type": "AWS::Cognito::UserPoolIdentityProvider",
      "Properties": {
         "UserPoolId": {"Ref": "UserPool"},
         "ProviderName": "YourProviderName",
         "ProviderDetails": {
            "MetadataURL": "YourMetadataURL"
         },
         "ProviderType": "SAML",
         "AttributeMapping": {
            "email": "Attribute"
         },
         "IdpIdentifiers": [
            "IdpIdentifier"
         ]
      }
   }
}
```

#### YAML

```yaml

UserPoolIdentityProvider:
  Type: AWS::Cognito::UserPoolIdentityProvider
  Properties:
    UserPoolId: !Ref UserPool
    ProviderName: "YourProviderName"
    ProviderDetails:
      MetadataURL: "YourMetadataURL"
    ProviderType: "SAML"
    AttributeMapping:
      email: "Attribute"
    IdpIdentifiers:
      - "IdpIdentifier"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Cognito::UserPoolGroup

AWS::Cognito::UserPoolResourceServer

All content copied from https://docs.aws.amazon.com/.
