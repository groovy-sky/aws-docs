# RejectedEntityInfo

If an entity is rejected when a `PutLogEvents` request was made, this includes
details about the reason for the rejection.

## Contents

**errorType**

The type of error that caused the rejection of the entity when calling
`PutLogEvents`.

Type: String

Valid Values: `InvalidEntity | InvalidTypeValue | InvalidKeyAttributes | InvalidAttributes | EntitySizeTooLarge | UnsupportedLogGroupType | MissingRequiredFields`

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/rejectedentityinfo.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/rejectedentityinfo.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/rejectedentityinfo.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

RecordField

RejectedLogEventsInfo
