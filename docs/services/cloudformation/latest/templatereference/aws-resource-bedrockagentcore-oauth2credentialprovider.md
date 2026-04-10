This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::OAuth2CredentialProvider

The `AWS::BedrockAgentCore::OAuth2CredentialProvider` resource Property description not available. for BedrockAgentCore.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::BedrockAgentCore::OAuth2CredentialProvider",
  "Properties" : {
      "CredentialProviderVendor" : String,
      "Name" : String,
      "Oauth2ProviderConfigInput" : Oauth2ProviderConfigInput,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::BedrockAgentCore::OAuth2CredentialProvider
Properties:
  CredentialProviderVendor: String
  Name: String
  Oauth2ProviderConfigInput:
    Oauth2ProviderConfigInput
  Tags:
    - Tag

```

## Properties

`CredentialProviderVendor`

Property description not available.

_Required_: Yes

_Type_: String

_Allowed values_: `GoogleOauth2 | GithubOauth2 | SlackOauth2 | SalesforceOauth2 | MicrosoftOauth2 | CustomOauth2 | AtlassianOauth2 | LinkedinOauth2 | XOauth2 | OktaOauth2 | OneLoginOauth2 | PingOneOauth2 | FacebookOauth2 | YandexOauth2 | RedditOauth2 | ZoomOauth2 | TwitchOauth2 | SpotifyOauth2 | DropboxOauth2 | NotionOauth2 | HubspotOauth2 | CyberArkOauth2 | FusionAuthOauth2 | Auth0Oauth2 | CognitoOauth2`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

Property description not available.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9\-_]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Oauth2ProviderConfigInput`

Contains the input configuration for an OAuth2 provider.

_Required_: No

_Type_: [Oauth2ProviderConfigInput](aws-properties-bedrockagentcore-oauth2credentialprovider-oauth2providerconfiginput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](aws-properties-bedrockagentcore-oauth2credentialprovider-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`CallbackUrl`

Property description not available.

`CreatedTime`

Property description not available.

`CredentialProviderArn`

Property description not available.

`LastUpdatedTime`

Property description not available.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UserPreferenceOverrideExtractionConfigurationInput

AtlassianOauth2ProviderConfigInput

All content copied from https://docs.aws.amazon.com/.
