# DatadogConnectorProfileCredentials

The connector-specific credentials required by Datadog.

## Contents

**apiKey**

A unique alphanumeric identifier used to authenticate a user, developer, or calling
program to your API.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `\S+`

Required: Yes

**applicationKey**

Application keys, in conjunction with your API key, give you full access to Datadog’s
programmatic API. Application keys are associated with the user account that created them. The
application key is used to log all requests made to the API.

Type: String

Length Constraints: Maximum length of 512.

Pattern: `\S+`

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/datadogconnectorprofilecredentials.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/datadogconnectorprofilecredentials.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/datadogconnectorprofilecredentials.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomerProfilesMetadata

DatadogConnectorProfileProperties

All content copied from https://docs.aws.amazon.com/.
