# HostedConfigurationVersionSummary

Information about the configuration.

## Contents

**ApplicationId**

The application ID.

Type: String

Pattern: `[a-z0-9]{4,7}`

Required: No

**ConfigurationProfileId**

The configuration profile ID.

Type: String

Pattern: `[a-z0-9]{4,7}`

Required: No

**ContentType**

A standard MIME type describing the format of the configuration content. For more
information, see [Content-Type](https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html).

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: No

**Description**

A description of the configuration.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1024.

Required: No

**KmsKeyArn**

The Amazon Resource Name of the AWS Key Management Service key that was used to encrypt this
specific version of the configuration data in the AWS AppConfig hosted configuration
store.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `arn:(aws[a-zA-Z-]*)?:[a-z]+:((eusc-)?[a-z]{2}((-gov)|(-iso([a-z]?)))?-[a-z]+-\d{1})?:(\d{12})?:[a-zA-Z0-9-_/:.]+`

Required: No

**VersionLabel**

A user-defined label for an AWS AppConfig hosted configuration version.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Pattern: `.*[^0-9].*`

Required: No

**VersionNumber**

The configuration version.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appconfig-2019-10-09/hostedconfigurationversionsummary.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appconfig-2019-10-09/hostedconfigurationversionsummary.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appconfig-2019-10-09/hostedconfigurationversionsummary.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExtensionSummary

InvalidConfigurationDetail

All content copied from https://docs.aws.amazon.com/.
