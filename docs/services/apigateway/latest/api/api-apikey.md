# ApiKey

A resource that can be distributed to callers for executing Method resources that require an API key. API keys can be mapped to any Stage on any RestApi, which indicates that the callers with the API key can make requests to that stage.

## Contents

**createdDate**

The timestamp when the API Key was created.

Type: Timestamp

Required: No

**customerId**

An AWS Marketplace customer identifier, when integrating with the AWS SaaS Marketplace.

Type: String

Required: No

**description**

The description of the API Key.

Type: String

Required: No

**enabled**

Specifies whether the API Key can be used by callers.

Type: Boolean

Required: No

**id**

The identifier of the API Key.

Type: String

Required: No

**lastUpdatedDate**

The timestamp when the API Key was last updated.

Type: Timestamp

Required: No

**name**

The name of the API Key.

Type: String

Required: No

**stageKeys**

A list of Stage resources that are associated with the ApiKey resource.

Type: Array of strings

Required: No

**tags**

The collection of tags. Each tag element is associated with a given resource.

Type: String to string map

Required: No

**value**

The value of the API Key.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/apikey.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/apikey.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/apikey.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AccessLogSettings

ApiStage

All content copied from https://docs.aws.amazon.com/.
