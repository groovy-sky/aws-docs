# AuthorizationData

An authorization token data object that corresponds to a public registry.

## Contents

**authorizationToken**

A base64-encoded string that contains authorization data for a public Amazon ECR registry.
When the string is decoded, it's presented in the format `user:password` for
public registry authentication using `docker login`.

Type: String

Pattern: `^\S+$`

Required: No

**expiresAt**

The Unix time in seconds and milliseconds when the authorization token expires.
Authorization tokens are valid for 12 hours.

Type: Timestamp

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-public-2020-10-30/authorizationdata.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-public-2020-10-30/authorizationdata.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-public-2020-10-30/authorizationdata.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Data Types

Image

All content copied from https://docs.aws.amazon.com/.
