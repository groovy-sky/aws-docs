# ResultReuseByAgeConfiguration

Specifies whether previous query results are reused, and if so, their maximum
age.

## Contents

**Enabled**

True if previous query results can be reused when the query is run; otherwise, false.
The default is false.

Type: Boolean

Required: Yes

**MaxAgeInMinutes**

Specifies, in minutes, the maximum age of a previous query result that Athena should consider for reuse. The default is 60.

Type: Integer

Valid Range: Minimum value of 0. Maximum value of 10080.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/resultreusebyageconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/resultreusebyageconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/resultreusebyageconfiguration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ResultConfigurationUpdates

ResultReuseConfiguration
