This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::OAuth2CredentialProvider Oauth2ProviderConfigInput

Contains the input configuration for an OAuth2 provider.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AtlassianOauth2ProviderConfig" : AtlassianOauth2ProviderConfigInput,
  "CustomOauth2ProviderConfig" : CustomOauth2ProviderConfigInput,
  "GithubOauth2ProviderConfig" : GithubOauth2ProviderConfigInput,
  "GoogleOauth2ProviderConfig" : GoogleOauth2ProviderConfigInput,
  "IncludedOauth2ProviderConfig" : IncludedOauth2ProviderConfigInput,
  "LinkedinOauth2ProviderConfig" : LinkedinOauth2ProviderConfigInput,
  "MicrosoftOauth2ProviderConfig" : MicrosoftOauth2ProviderConfigInput,
  "SalesforceOauth2ProviderConfig" : SalesforceOauth2ProviderConfigInput,
  "SlackOauth2ProviderConfig" : SlackOauth2ProviderConfigInput
}

```

### YAML

```yaml

  AtlassianOauth2ProviderConfig:
    AtlassianOauth2ProviderConfigInput
  CustomOauth2ProviderConfig:
    CustomOauth2ProviderConfigInput
  GithubOauth2ProviderConfig:
    GithubOauth2ProviderConfigInput
  GoogleOauth2ProviderConfig:
    GoogleOauth2ProviderConfigInput
  IncludedOauth2ProviderConfig:
    IncludedOauth2ProviderConfigInput
  LinkedinOauth2ProviderConfig:
    LinkedinOauth2ProviderConfigInput
  MicrosoftOauth2ProviderConfig:
    MicrosoftOauth2ProviderConfigInput
  SalesforceOauth2ProviderConfig:
    SalesforceOauth2ProviderConfigInput
  SlackOauth2ProviderConfig:
    SlackOauth2ProviderConfigInput

```

## Properties

`AtlassianOauth2ProviderConfig`

Configuration settings for Atlassian OAuth2 provider integration.

_Required_: No

_Type_: [AtlassianOauth2ProviderConfigInput](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrockagentcore-oauth2credentialprovider-atlassianoauth2providerconfiginput.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomOauth2ProviderConfig`

The configuration for a custom OAuth2 provider.

_Required_: No

_Type_: [CustomOauth2ProviderConfigInput](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrockagentcore-oauth2credentialprovider-customoauth2providerconfiginput.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GithubOauth2ProviderConfig`

The configuration for a GitHub OAuth2 provider.

_Required_: No

_Type_: [GithubOauth2ProviderConfigInput](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrockagentcore-oauth2credentialprovider-githuboauth2providerconfiginput.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GoogleOauth2ProviderConfig`

The configuration for a Google OAuth2 provider.

_Required_: No

_Type_: [GoogleOauth2ProviderConfigInput](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrockagentcore-oauth2credentialprovider-googleoauth2providerconfiginput.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludedOauth2ProviderConfig`

The configuration for a non-custom OAuth2 provider. This includes settings for supported
OAuth2 providers that have built-in integration support.

_Required_: No

_Type_: [IncludedOauth2ProviderConfigInput](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrockagentcore-oauth2credentialprovider-includedoauth2providerconfiginput.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LinkedinOauth2ProviderConfig`

Configuration settings for LinkedIn OAuth2 provider integration.

_Required_: No

_Type_: [LinkedinOauth2ProviderConfigInput](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrockagentcore-oauth2credentialprovider-linkedinoauth2providerconfiginput.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MicrosoftOauth2ProviderConfig`

The configuration for a Microsoft OAuth2 provider.

_Required_: No

_Type_: [MicrosoftOauth2ProviderConfigInput](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrockagentcore-oauth2credentialprovider-microsoftoauth2providerconfiginput.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SalesforceOauth2ProviderConfig`

The configuration for a Salesforce OAuth2 provider.

_Required_: No

_Type_: [SalesforceOauth2ProviderConfigInput](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrockagentcore-oauth2credentialprovider-salesforceoauth2providerconfiginput.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SlackOauth2ProviderConfig`

The configuration for a Slack OAuth2 provider.

_Required_: No

_Type_: [SlackOauth2ProviderConfigInput](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrockagentcore-oauth2credentialprovider-slackoauth2providerconfiginput.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Oauth2Discovery

Oauth2ProviderConfigOutput
