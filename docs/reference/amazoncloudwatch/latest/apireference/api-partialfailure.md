# PartialFailure

This array is empty if the API operation was successful for all the rules specified in
the request. If the operation could not process one of the rules, the following data is
returned for each of those rules.

## Contents

**ExceptionType**

The type of error.

Type: String

Required: No

**FailureCode**

The code of the error.

Type: String

Required: No

**FailureDescription**

A description of the error.

Type: String

Required: No

**FailureResource**

The specified rule that could not be deleted.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/monitoring-2010-08-01/partialfailure.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/monitoring-2010-08-01/partialfailure.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/monitoring-2010-08-01/partialfailure.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MuteTargets

Range

All content copied from https://docs.aws.amazon.com/.
