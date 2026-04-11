This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::IdentityPool

The `AWS::Cognito::IdentityPool` resource creates an Amazon Cognito
identity pool.

To avoid deleting the resource accidentally from CloudFormation, use [DeletionPolicy Attribute](../userguide/aws-attribute-deletionpolicy.md) and the [UpdateReplacePolicy Attribute](../userguide/aws-attribute-updatereplacepolicy.md) to retain the resource on deletion or
replacement.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Cognito::IdentityPool",
  "Properties" : {
      "AllowClassicFlow" : Boolean,
      "AllowUnauthenticatedIdentities" : Boolean,
      "CognitoEvents" : Json,
      "CognitoIdentityProviders" : [ CognitoIdentityProvider, ... ],
      "CognitoStreams" : CognitoStreams,
      "DeveloperProviderName" : String,
      "IdentityPoolName" : String,
      "IdentityPoolTags" : [ Tag, ... ],
      "OpenIdConnectProviderARNs" : [ String, ... ],
      "PushSync" : PushSync,
      "SamlProviderARNs" : [ String, ... ],
      "SupportedLoginProviders" : Json
    }
}

```

### YAML

```yaml

Type: AWS::Cognito::IdentityPool
Properties:
  AllowClassicFlow: Boolean
  AllowUnauthenticatedIdentities: Boolean
  CognitoEvents: Json
  CognitoIdentityProviders:
    - CognitoIdentityProvider
  CognitoStreams:
    CognitoStreams
  DeveloperProviderName: String
  IdentityPoolName: String
  IdentityPoolTags:
    - Tag
  OpenIdConnectProviderARNs:
    - String
  PushSync:
    PushSync
  SamlProviderARNs:
    - String
  SupportedLoginProviders: Json

```

## Properties

`AllowClassicFlow`

Enables the Basic (Classic) authentication flow.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AllowUnauthenticatedIdentities`

Specifies whether the identity pool supports unauthenticated logins.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CognitoEvents`

The events to configure.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CognitoIdentityProviders`

The Amazon Cognito user pools and their client IDs.

_Required_: No

_Type_: Array of [CognitoIdentityProvider](aws-properties-cognito-identitypool-cognitoidentityprovider.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CognitoStreams`

Configuration options for configuring Amazon Cognito streams.

_Required_: No

_Type_: [CognitoStreams](aws-properties-cognito-identitypool-cognitostreams.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeveloperProviderName`

The "domain" Amazon Cognito uses when referencing your users. This name acts as a
placeholder that allows your backend and the Amazon Cognito service to communicate about
the developer provider. For the `DeveloperProviderName`, you can use letters
and periods (.), underscores (\_), and dashes (-).

_Minimum length_: 1

_Maximum length_: 100

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IdentityPoolName`

The name of your Amazon Cognito identity pool.

_Minimum length_: 1

_Maximum length_: 128

_Pattern_: `[\w\s+=,.@-]+`

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IdentityPoolTags`

Tags to assign to the identity pool. A tag is a label that you can apply to identity
pools to categorize and manage them in different ways, such as by purpose, owner,
environment, or other criteria.

_Required_: No

_Type_: Array of [Tag](aws-properties-cognito-identitypool-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OpenIdConnectProviderARNs`

The Amazon Resource Names (ARNs) of the OpenID connect providers.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PushSync`

The configuration options to be applied to the identity pool.

_Required_: No

_Type_: [PushSync](aws-properties-cognito-identitypool-pushsync.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SamlProviderARNs`

The Amazon Resource Names (ARNs) of the Security Assertion Markup Language (SAML)
providers.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SupportedLoginProviders`

Key-value pairs that map provider names to provider app IDs.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `IdentityPoolId`, such as
`us-east-2:0d01f4d7-1305-4408-b437-12345EXAMPLE`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Name`

The name of the Amazon Cognito identity pool, returned as a string.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon Cognito

CognitoIdentityProvider

All content copied from https://docs.aws.amazon.com/.
