# Event

Contains information about an event that was returned by a lookup request. The result
includes a representation of a CloudTrail event.

## Contents

**AccessKeyId**

The AWS access key ID that was used to sign the request. If the request
was made with temporary security credentials, this is the access key ID of the temporary
credentials.

Type: String

Required: No

**CloudTrailEvent**

A JSON string that contains a representation of the event returned.

Type: String

Required: No

**EventId**

The CloudTrail ID of the event returned.

Type: String

Required: No

**EventName**

The name of the event returned.

Type: String

Required: No

**EventSource**

The AWS service to which the request was made.

Type: String

Required: No

**EventTime**

The date and time of the event returned.

Type: Timestamp

Required: No

**ReadOnly**

Information about whether the event is a write event or a read event.

Type: String

Required: No

**Resources**

A list of resources referenced by the event returned.

Type: Array of [Resource](api-resource.md) objects

Required: No

**Username**

A user name or role name of the requester that called the API in the event
returned.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudtrail-2013-11-01/event.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudtrail-2013-11-01/event.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudtrail-2013-11-01/event.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Destination

EventDataStore

All content copied from https://docs.aws.amazon.com/.
