# AlarmPromQLCriteria

Contains the configuration that determines how a PromQL alarm evaluates its
contributors, including the query to run and the durations that define when contributors
transition between states.

## Contents

**Query**

The PromQL query that the alarm evaluates. The query must return a result of vector
type. Each entry in the vector result represents an alarm contributor.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 10000.

Required: Yes

**PendingPeriod**

The duration, in seconds, that a contributor must be continuously breaching before
it transitions to the `ALARM` state.

Type: Integer

Valid Range: Minimum value of 0. Maximum value of 86400.

Required: No

**RecoveryPeriod**

The duration, in seconds, that a contributor must continuously not be breaching
before it transitions back to the `OK` state.

Type: Integer

Valid Range: Minimum value of 0. Maximum value of 86400.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/monitoring-2010-08-01/AlarmPromQLCriteria)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/monitoring-2010-08-01/AlarmPromQLCriteria)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/monitoring-2010-08-01/AlarmPromQLCriteria)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AlarmMuteRuleSummary

AnomalyDetector
