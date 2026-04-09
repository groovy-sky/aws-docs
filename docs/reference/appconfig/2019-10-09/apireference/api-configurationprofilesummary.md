# ConfigurationProfileSummary

A summary of a configuration profile.

## Contents

**ApplicationId**

The application ID.

Type: String

Pattern: `[a-z0-9]{4,7}`

Required: No

**Id**

The ID of the configuration profile.

Type: String

Pattern: `[a-z0-9]{4,7}`

Required: No

**LocationUri**

The URI location of the configuration.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: No

**Name**

The name of the configuration profile.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: No

**Type**

The type of configurations contained in the profile. AWS AppConfig supports
`feature flags` and `freeform` configurations. We recommend you
create feature flag configurations to enable or disable new features and freeform
configurations to distribute configurations to an application. When calling this API, enter
one of the following values for `Type`:

`AWS.AppConfig.FeatureFlags`

`AWS.Freeform`

Type: String

Pattern: `^[a-zA-Z\.]+`

Required: No

**ValidatorTypes**

The types of validators in the configuration profile.

Type: Array of strings

Array Members: Minimum number of 0 items. Maximum number of 2 items.

Valid Values: `JSON_SCHEMA | LAMBDA`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appconfig-2019-10-09/configurationprofilesummary.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appconfig-2019-10-09/configurationprofilesummary.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appconfig-2019-10-09/configurationprofilesummary.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BadRequestDetails

DeletionProtectionSettings

All content copied from https://docs.aws.amazon.com/.
