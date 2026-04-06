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

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/logs-2014-03-28/RejectedEntityInfo)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/logs-2014-03-28/RejectedEntityInfo)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/logs-2014-03-28/RejectedEntityInfo)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

RecordField

RejectedLogEventsInfo
