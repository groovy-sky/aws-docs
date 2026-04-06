# ServiceNowConnectorProfileCredentials

The connector-specific profile credentials required when using ServiceNow.

## Contents

**oAuth2Credentials**

The OAuth 2.0 credentials required to authenticate the user.

Type: [OAuth2Credentials](api-oauth2credentials.md) object

Required: No

**password**

The password that corresponds to the user name.

Type: String

Length Constraints: Maximum length of 512.

Pattern: `.*`

Required: No

**username**

The name of the user.

Type: String

Length Constraints: Maximum length of 512.

Pattern: `\S+`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/appflow-2020-08-23/ServiceNowConnectorProfileCredentials)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/appflow-2020-08-23/ServiceNowConnectorProfileCredentials)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/appflow-2020-08-23/ServiceNowConnectorProfileCredentials)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ScheduledTriggerProperties

ServiceNowConnectorProfileProperties
