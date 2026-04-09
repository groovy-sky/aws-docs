# DeletionProtectionSettings

A parameter to configure deletion protection. Deletion protection prevents a user from
deleting a configuration profile or an environment if AWS AppConfig has called either
[GetLatestConfiguration](api-appconfigdata-getlatestconfiguration.md) or [GetConfiguration](api-getconfiguration.md) for the
configuration profile or from the environment during the specified interval.

The default interval specified by `ProtectionPeriodInMinutes` is 60.
`DeletionProtectionCheck` skips configuration profiles and environments that
were created in the past hour.

## Contents

**Enabled**

A parameter that indicates if deletion protection is enabled or not.

Type: Boolean

Required: No

**ProtectionPeriodInMinutes**

The time interval during which AWS AppConfig monitors for calls to [GetLatestConfiguration](api-appconfigdata-getlatestconfiguration.md) or [GetConfiguration](api-getconfiguration.md) for a
configuration profile or from an environment. AWS AppConfig returns an error if a
user calls [DeleteConfigurationProfile](api-deleteconfigurationprofile.md) or [DeleteEnvironment](api-deleteenvironment.md) for the designated configuration profile or
environment. To bypass the error and delete a configuration profile or an environment,
specify `BYPASS` for the `DeletionProtectionCheck` parameter for
either [DeleteConfigurationProfile](api-deleteconfigurationprofile.md) or [DeleteEnvironment](api-deleteenvironment.md).

Type: Integer

Valid Range: Minimum value of 15. Maximum value of 1440.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appconfig-2019-10-09/deletionprotectionsettings.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appconfig-2019-10-09/deletionprotectionsettings.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appconfig-2019-10-09/deletionprotectionsettings.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConfigurationProfileSummary

DeploymentEvent

All content copied from https://docs.aws.amazon.com/.
