# OAuth2CustomParameter

Custom parameter required for OAuth 2.0 authentication.

## Contents

**connectorSuppliedValues**

Contains default values for this authentication parameter that are supplied by the
connector.

Type: Array of strings

Length Constraints: Maximum length of 256.

Pattern: `\S+`

Required: No

**description**

A description about the custom parameter used for OAuth 2.0 authentication.

Type: String

Length Constraints: Maximum length of 1024.

Pattern: `[\s\w/!@#+=.-]*`

Required: No

**isRequired**

Indicates whether the custom parameter for OAuth 2.0 authentication is required.

Type: Boolean

Required: No

**isSensitiveField**

Indicates whether this authentication custom parameter is a sensitive field.

Type: Boolean

Required: No

**key**

The key of the custom parameter required for OAuth 2.0 authentication.

Type: String

Length Constraints: Maximum length of 512.

Pattern: `\S+`

Required: No

**label**

The label of the custom parameter used for OAuth 2.0 authentication.

Type: String

Length Constraints: Maximum length of 128.

Pattern: `.*`

Required: No

**type**

Indicates whether custom parameter is used with TokenUrl or AuthUrl.

Type: String

Valid Values: `TOKEN_URL | AUTH_URL`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/oauth2customparameter.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/oauth2customparameter.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/oauth2customparameter.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OAuth2Credentials

OAuth2Defaults

All content copied from https://docs.aws.amazon.com/.
