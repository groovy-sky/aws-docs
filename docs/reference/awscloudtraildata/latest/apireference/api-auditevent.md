# AuditEvent

An event from a source outside of AWS that you want CloudTrail
to log.

## Contents

**eventData**

The content of an audit event that comes from the event, such as `userIdentity`,
`userAgent`, and `eventSource`.

Type: String

Required: Yes

**id**

The original event ID from the source event.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Pattern: `[-_A-Za-z0-9]+`

Required: Yes

**eventDataChecksum**

A checksum is a base64-SHA256 algorithm that helps you verify that CloudTrail receives the event that matches
with the checksum. Calculate the checksum by running a command like the following:

`printf %s $eventdata | openssl dgst -binary -sha256 | base64`

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudtrail-data-2021-08-11/auditevent.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudtrail-data-2021-08-11/auditevent.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudtrail-data-2021-08-11/auditevent.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Data Types

AuditEventResultEntry

All content copied from https://docs.aws.amazon.com/.
