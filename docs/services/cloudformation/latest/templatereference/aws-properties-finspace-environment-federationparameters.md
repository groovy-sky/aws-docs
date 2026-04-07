This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FinSpace::Environment FederationParameters

Configuration information when authentication mode is FEDERATED.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApplicationCallBackURL" : String,
  "AttributeMap" : [ AttributeMapItems, ... ],
  "FederationProviderName" : String,
  "FederationURN" : String,
  "SamlMetadataDocument" : String,
  "SamlMetadataURL" : String
}

```

### YAML

```yaml

  ApplicationCallBackURL: String
  AttributeMap:
    - AttributeMapItems
  FederationProviderName: String
  FederationURN: String
  SamlMetadataDocument: String
  SamlMetadataURL: String

```

## Properties

`ApplicationCallBackURL`

The redirect or sign-in URL that should be entered into the SAML 2.0 compliant identity provider configuration
(IdP).

_Required_: No

_Type_: String

_Pattern_: `^https?://[-a-zA-Z0-9+&amp;@#/%?=~_|!:,.;]*[-a-zA-Z0-9+&amp;@#/%=~_|]{1,1000}`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AttributeMap`

SAML attribute name and value. The name must always be `Email` and the value should be set to
the attribute definition in which user email is set. For example, name would be `Email` and
value `http://schemas.xmlsoap.org/ws/2005/05/identity/claims/emailaddress`.
Please check your SAML 2.0 compliant identity provider (IdP) documentation for details.

_Required_: No

_Type_: Array of [AttributeMapItems](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-finspace-environment-attributemapitems.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FederationProviderName`

Name of the identity provider (IdP).

_Required_: No

_Type_: String

_Pattern_: `[^_\p{Z}][\p{L}\p{M}\p{S}\p{N}\p{P}][^_\p{Z}]+`

_Minimum_: `1`

_Maximum_: `32`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FederationURN`

The Uniform Resource Name (URN). Also referred as Service Provider URN or Audience URI or Service Provider Entity ID.

_Required_: No

_Type_: String

_Pattern_: `^[A-Za-z0-9._\-:\/#\+]+$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SamlMetadataDocument`

SAML 2.0 Metadata document from identity provider (IdP).

_Required_: No

_Type_: String

_Pattern_: `.*`

_Minimum_: `1000`

_Maximum_: `10000000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SamlMetadataURL`

Provide the metadata URL from your SAML 2.0 compliant identity provider (IdP).

_Required_: No

_Type_: String

_Pattern_: `^https?://[-a-zA-Z0-9+&amp;@#/%?=~_|!:,.;]*[-a-zA-Z0-9+&amp;@#/%=~_|]{1,1000}`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AttributeMapItems

SuperuserParameters
