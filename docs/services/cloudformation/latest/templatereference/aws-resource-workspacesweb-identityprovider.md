This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WorkSpacesWeb::IdentityProvider

This resource specifies an identity provider that is then associated with a web portal.
This resource is not required if your portal's `AuthenticationType` is IAM
Identity Center.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::WorkSpacesWeb::IdentityProvider",
  "Properties" : {
      "IdentityProviderDetails" : {Key: Value, ...},
      "IdentityProviderName" : String,
      "IdentityProviderType" : String,
      "PortalArn" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::WorkSpacesWeb::IdentityProvider
Properties:
  IdentityProviderDetails:
    Key: Value
  IdentityProviderName: String
  IdentityProviderType: String
  PortalArn: String
  Tags:
    - Tag

```

## Properties

`IdentityProviderDetails`

The identity provider details. The following list describes the provider detail keys for
each identity provider type.

- For Google and Login with Amazon:

- `client_id`

- `client_secret`

- `authorize_scopes`

- For Facebook:

- `client_id`

- `client_secret`

- `authorize_scopes`

- `api_version`

- For Sign in with Apple:

- `client_id`

- `team_id`

- `key_id`

- `private_key`

- `authorize_scopes`

- For OIDC providers:

- `client_id`

- `client_secret`

- `attributes_request_method`

- `oidc_issuer`

- `authorize_scopes`

- `authorize_url` _if not available from discovery URL specified by oidc\_issuer_
_key_

- `token_url` _if not available from discovery URL specified by oidc\_issuer_
_key_

- `attributes_url` _if not available from discovery URL specified by oidc\_issuer_
_key_

- `jwks_uri` _if not available from discovery URL specified by oidc\_issuer_
_key_

- For SAML providers:

- `MetadataFile` OR `MetadataURL`

- `IDPSignout` (boolean) _optional_

- `IDPInit` (boolean) _optional_

- `RequestSigningAlgorithm` (string) _optional_
\- Only accepts `rsa-sha256`

- `EncryptedResponses` (boolean) _optional_

_Required_: Yes

_Type_: Object of String

_Pattern_: `^[\s\S]*$`

_Minimum_: `0`

_Maximum_: `131072`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IdentityProviderName`

The identity provider name.

_Required_: Yes

_Type_: String

_Pattern_: `^[^_][\p{L}\p{M}\p{S}\p{N}\p{P}][^_]+$`

_Minimum_: `1`

_Maximum_: `32`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IdentityProviderType`

The identity provider type.

_Required_: Yes

_Type_: String

_Allowed values_: `SAML | Facebook | Google | LoginWithAmazon | SignInWithApple | OIDC`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PortalArn`

The ARN of the identity provider.

_Required_: No

_Type_: String

_Pattern_: `^arn:[\w+=\/,.@-]+:[a-zA-Z0-9\-]+:[a-zA-Z0-9\-]*:[a-zA-Z0-9]{1,12}:[a-zA-Z]+(\/[a-fA-F0-9\-]{36})+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](aws-properties-workspacesweb-identityprovider-tag.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function,
`Ref` returns the resource's Amazon Resource Name (ARN).

For more information about using the `Ref` function, see [Ref](../userguide/intrinsic-function-reference-ref.md).

### Fn::GetAtt

`IdentityProviderArn`

The ARN of the identity provider.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
