# LogStream

Represents a log stream, which is a sequence of log events from a single emitter of
logs.

## Contents

**arn**

The Amazon Resource Name (ARN) of the log stream.

Type: String

Required: No

**creationTime**

The creation time of the stream, expressed as the number of milliseconds after
`Jan 1, 1970 00:00:00 UTC`.

Type: Long

Valid Range: Minimum value of 0.

Required: No

**firstEventTimestamp**

The time of the first event, expressed as the number of milliseconds after `Jan 1,
        1970 00:00:00 UTC`.

Type: Long

Valid Range: Minimum value of 0.

Required: No

**lastEventTimestamp**

The time of the most recent log event in the log stream in CloudWatch Logs. This number
is expressed as the number of milliseconds after `Jan 1, 1970 00:00:00 UTC`. The
`lastEventTime` value updates on an eventual consistency basis. It typically
updates in less than an hour from ingestion, but in rare situations might take
longer.

Type: Long

Valid Range: Minimum value of 0.

Required: No

**lastIngestionTime**

The ingestion time, expressed as the number of milliseconds after `Jan 1, 1970
        00:00:00 UTC` The `lastIngestionTime` value updates on an eventual
consistency basis. It typically updates in less than an hour after ingestion, but in rare
situations might take longer.

Type: Long

Valid Range: Minimum value of 0.

Required: No

**logStreamName**

The name of the log stream.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 512.

Pattern: `[^:*]*`

Required: No

**storedBytes**

_This member has been deprecated._

The number of bytes stored.

**Important:** As of June 17, 2019, this parameter is no
longer supported for log streams, and is always reported as zero. This change applies only to
log streams. The `storedBytes` parameter for log groups is not affected.

Type: Long

Valid Range: Minimum value of 0.

Required: No

**uploadSequenceToken**

The sequence token.

###### Important

The sequence token is now ignored in `PutLogEvents` actions.
`PutLogEvents` actions are always accepted regardless of receiving an invalid
sequence token. You don't need to obtain `uploadSequenceToken` to use a
`PutLogEvents` action.

Type: String

Length Constraints: Minimum length of 1.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/logstream.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/logstream.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/logstream.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LogGroupSummary

LookupTable

All content copied from https://docs.aws.amazon.com/.
