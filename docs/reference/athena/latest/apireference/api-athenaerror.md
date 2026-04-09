# AthenaError

Provides information about an Athena query error. The
`AthenaError` feature provides standardized error information to help you
understand failed queries and take steps after a query failure occurs.
`AthenaError` includes an `ErrorCategory` field that specifies
whether the cause of the failed query is due to system error, user error, or other
error.

## Contents

**ErrorCategory**

An integer value that specifies the category of a query failure error. The following
list shows the category for each integer value.

**1** \- System

**2** \- User

**3** \- Other

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 3.

Required: No

**ErrorMessage**

Contains a short description of the error that occurred.

Type: String

Required: No

**ErrorType**

An integer value that provides specific information about an Athena query
error. For the meaning of specific values, see the [Error Type Reference](../../../../services/athena/latest/ug/error-reference.md#error-reference-error-type-reference) in the _Amazon Athena User_
_Guide_.

Type: Integer

Valid Range: Minimum value of 0. Maximum value of 9999.

Required: No

**Retryable**

True if the query might succeed if resubmitted.

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/athenaerror.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/athenaerror.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/athenaerror.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ApplicationDPUSizes

BatchGetNamedQueryInput

All content copied from https://docs.aws.amazon.com/.
