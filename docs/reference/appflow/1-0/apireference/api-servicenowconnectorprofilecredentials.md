---
title: "ServiceNowConnectorProfileCredentials"
---

# ServiceNowConnectorProfileCredentials
<a name="API_ServiceNowConnectorProfileCredentials"></a>

 The connector-specific profile credentials required when using ServiceNow.

## Contents
<a name="API_ServiceNowConnectorProfileCredentials_Contents"></a>

 ** oAuth2Credentials **   <a name="appflow-Type-ServiceNowConnectorProfileCredentials-oAuth2Credentials"></a>
 The OAuth 2.0 credentials required to authenticate the user.
Type: [OAuth2Credentials](API_OAuth2Credentials.md) object
Required: No

 ** password **   <a name="appflow-Type-ServiceNowConnectorProfileCredentials-password"></a>
 The password that corresponds to the user name.
Type: String
Length Constraints: Maximum length of 512.
Pattern: `.*`
Required: No

 ** username **   <a name="appflow-Type-ServiceNowConnectorProfileCredentials-username"></a>
 The name of the user.
Type: String
Length Constraints: Maximum length of 512.
Pattern: `\S+`
Required: No

## See Also
<a name="API_ServiceNowConnectorProfileCredentials_SeeAlso"></a>

For more information about using this API in one of the language-specific AWS SDKs, see the following:
+  [AWS SDK for C\+\+](https://docs.aws.amazon.com/goto/SdkForCpp/appflow-2020-08-23/ServiceNowConnectorProfileCredentials)
+  [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/appflow-2020-08-23/ServiceNowConnectorProfileCredentials)
+  [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/appflow-2020-08-23/ServiceNowConnectorProfileCredentials)

All content copied from https://docs.aws.amazon.com/.
