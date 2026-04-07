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

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/athena-2017-05-18/ResultReuseByAgeConfiguration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/athena-2017-05-18/ResultReuseByAgeConfiguration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/athena-2017-05-18/ResultReuseByAgeConfiguration)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ResultConfigurationUpdates

ResultReuseConfiguration
