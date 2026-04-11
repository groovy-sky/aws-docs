# CloudTrail Lake integrations event schema

The following table describes the required and optional schema elements that match
those in CloudTrail event records. The contents of `eventData` are provided by
your events; other fields are provided by CloudTrail after ingestion.

CloudTrail event record contents are described in more detail in [CloudTrail record contents for management, data, and network activity events](cloudtrail-event-reference-record-contents.md).

- [Fields that are provided by CloudTrail after\
ingestion](#fields-cloudtrail)

- [Fields that are provided by your\
events](#fields-event)

The following fields are provided by CloudTrail after ingestion:

Field nameInput typeRequirementDescriptioneventVersionstringRequired

The event version.

eventCategorystringRequired

The event category. For non-AWS events, the value is
`ActivityAuditLog`.

eventTypestringRequired

The event type. For non-AWS events, the valid value is
`ActivityLog`.

eventIDstringRequiredA unique ID for an event.eventTime

string

Required

Event timestamp, in `yyyy-MM-DDTHH:mm:ss` format, in
Universal Coordinated Time (UTC).

awsRegionstringRequired

The AWS Region where the `PutAuditEvents` call was
made.

recipientAccountIdstringRequired

Represents the account ID that received this event. CloudTrail populates
this field by calculating it from event payload.

addendum

-

Optional

Shows information about why event processing was delayed. If information
was missing from an existing event, the addendum block includes the
missing information and a reason for why it was missing.

- reason

stringOptional

The reason that the event or some of its contents were
missing.

- updatedFields

stringOptional

The event record fields that are updated by the addendum. This is
only provided if the reason is `UPDATED_DATA`.

- originalUID

stringOptional

The original event UID from the source. This is only provided if
the reason is `UPDATED_DATA`.

- originalEventID

stringOptional

The original event ID. This is only provided if the reason is
`UPDATED_DATA`.

metadata

-

Required

Information about the channel that the event used.

- ingestionTime

stringRequired

The timestamp when the event was processed, in `yyyy-MM-DDTHH:mm:ss`
format, in Universal Coordinated Time (UTC).

- channelARN

stringRequired

The ARN of the channel that the event used.

The following fields are provided by customer events:

Field nameInput typeRequirementDescriptioneventData

-

RequiredThe audit data sent to CloudTrail in a `PutAuditEvents` call.

- version

stringRequired

The version of the event from its source.

Length constraints: Maximum length of 256.

- userIdentity

-

Required

Information about the user who made a request.

- - type

string

Required

The type of user identity.

Length constraints: Maximum length of 128.

- - principalId

string

Required

A unique identifier for the actor of the event.

Length constraints: Maximum length of 1024.

- - details

JSON object

Optional

Additional information about the identity.

- userAgent

string

Optional

The agent through which the request was made.

Length constraints: Maximum length of 1024.

- eventSource

string

Required

This is the partner event source, or the custom
application about which events are logged.

Length constraints: Maximum length of 1024.

- eventName

string

Required

The requested action, one of the actions in the API for the source
service or application.

Length constraints: Maximum length of 1024.

- eventTime

string

Required

Event timestamp, in `yyyy-MM-DDTHH:mm:ss` format, in
Universal Coordinated Time (UTC).

- UID

stringRequired

The UID value that identifies the request. The service or
application that is called generates this value.

Length constraints: Maximum length of 1024.

- requestParameters

JSON object

Optional

The parameters, if any, that were sent with the request. This
field has a maximum size of 100 kB, and content exceeding the limit
is rejected.

- responseElements

JSON object

Optional

The response element for actions that make changes (create,
update, or delete actions). This field has a maximum size of 100 kB, and content
exceeding the limit is rejected.

- errorCode

stringOptional

A string representing an error for the event.

Length constraints: Maximum length of 256.

- errorMessage

stringOptional

The description of the error.

Length constraints: Maximum length of 256.

- sourceIPAddress

string

Optional

The IP address from which the request was made. Both IPv4 and IPv6 addresses are accepted.

- recipientAccountId

stringRequired

Represents the account ID that received this event. The account ID must be the same as the AWS account ID that owns the channel.

- additionalEventData

JSON object

Optional

Additional data about the event that was not part of the request
or response. This field has a maximum size of 28 kB, and content
exceeding that limit is rejected.

The following example shows the hierarchy of schema elements that match those in
CloudTrail event records.

```JSON

{
    "eventVersion": String,
    "eventCategory": String,
    "eventType": String,
    "eventID": String,
    "eventTime": String,
    "awsRegion": String,
    "recipientAccountId": String,
    "addendum": {
       "reason": String,
       "updatedFields": String,
       "originalUID": String,
       "originalEventID": String
    },
    "metadata" : {
       "ingestionTime": String,
       "channelARN": String
    },
    "eventData": {
        "version": String,
        "userIdentity": {
          "type": String,
          "principalId": String,
          "details": {
             JSON
          }
        },
        "userAgent": String,
        "eventSource": String,
        "eventName": String,
        "eventTime": String,
        "UID": String,
        "requestParameters": {
           JSON
        },
        "responseElements": {
           JSON
        },
        "errorCode": String,
        "errorMessage": String,
        "sourceIPAddress": String,
        "recipientAccountId": String,
        "additionalEventData": {
           JSON
        }
    }
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Delete a channel with the AWS CLI

Dashboards

All content copied from https://docs.aws.amazon.com/.
