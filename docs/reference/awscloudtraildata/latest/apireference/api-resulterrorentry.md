# ResultErrorEntry

Includes the error code and error message for events that could not be ingested by CloudTrail.

## Contents

**errorCode**

The error code for events that could not be ingested by CloudTrail. Possible error codes include: `FieldTooLong`, `FieldNotFound`,
`InvalidChecksum`, `InvalidData`, `InvalidRecipient`, `InvalidEventSource`, `AccountNotSubscribed`,
`Throttling`, and `InternalFailure`.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: Yes

**errorMessage**

The message that describes the error for events that could not be ingested by CloudTrail.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: Yes

**id**

The original event ID from the source event that could not be ingested by CloudTrail.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Pattern: `[-_A-Za-z0-9]+`

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudtrail-data-2021-08-11/resulterrorentry.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudtrail-data-2021-08-11/resulterrorentry.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudtrail-data-2021-08-11/resulterrorentry.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AuditEventResultEntry

Common Parameters

All content copied from https://docs.aws.amazon.com/.
