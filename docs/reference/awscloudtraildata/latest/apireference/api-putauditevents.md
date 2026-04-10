# PutAuditEvents

Ingests your application events into CloudTrail Lake. A required parameter,
`auditEvents`, accepts the JSON records (also called
_payload_) of events that you want CloudTrail to ingest. You
can add up to 100 of these events (or up to 1 MB) per `PutAuditEvents`
request.

## Request Syntax

```nohighlight

POST /PutAuditEvents?channelArn=channelArn&externalId=externalId HTTP/1.1
Content-type: application/json

{
   "auditEvents": [
      {
         "eventData": "string",
         "eventDataChecksum": "string",
         "id": "string"
      }
   ]
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[channelArn](#API_PutAuditEvents_RequestSyntax)**

The ARN or ID (the ARN suffix) of a channel.

Pattern: `arn:.*`

Required: Yes

**[externalId](#API_PutAuditEvents_RequestSyntax)**

A unique identifier that is conditionally required when the channel's resource policy includes an external
ID. This value can be any string,
such as a passphrase or account number.

Length Constraints: Minimum length of 2. Maximum length of 1224.

Pattern: `[\w+=,.@:\/-]*`

## Request Body

The request accepts the following data in JSON format.

**[auditEvents](#API_PutAuditEvents_RequestSyntax)**

The JSON payload of events that you want to ingest. You can also point to the JSON event
payload in a file.

Type: Array of [AuditEvent](api-auditevent.md) objects

Array Members: Minimum number of 1 item. Maximum number of 100 items.

Required: Yes

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "failed": [
      {
         "errorCode": "string",
         "errorMessage": "string",
         "id": "string"
      }
   ],
   "successful": [
      {
         "eventID": "string",
         "id": "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[failed](#API_PutAuditEvents_ResponseSyntax)**

Lists events in the provided event payload that could not be
ingested into CloudTrail, and includes the error code and error message
returned for events that could not be ingested.

Type: Array of [ResultErrorEntry](api-resulterrorentry.md) objects

Array Members: Minimum number of 0 items. Maximum number of 100 items.

**[successful](#API_PutAuditEvents_ResponseSyntax)**

Lists events in the provided event payload that were successfully ingested
into CloudTrail.

Type: Array of [AuditEventResultEntry](api-auditeventresultentry.md) objects

Array Members: Minimum number of 0 items. Maximum number of 100 items.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ChannelInsufficientPermission**

The caller's account ID must be the same as the channel owner's account ID.

HTTP Status Code: 400

**ChannelNotFound**

The channel could not be found.

HTTP Status Code: 400

**ChannelUnsupportedSchema**

The schema type of the event is not supported.

HTTP Status Code: 400

**DuplicatedAuditEventId**

Two or more entries in the request have the same event ID.

HTTP Status Code: 400

**InvalidChannelARN**

The specified channel ARN is not a valid
channel ARN.

HTTP Status Code: 400

**UnsupportedOperationException**

The operation requested is not supported in this region or account.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/cloudtrail-data-2021-08-11/putauditevents.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/cloudtrail-data-2021-08-11/putauditevents.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudtrail-data-2021-08-11/putauditevents.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/cloudtrail-data-2021-08-11/putauditevents.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudtrail-data-2021-08-11/putauditevents.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/cloudtrail-data-2021-08-11/putauditevents.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/cloudtrail-data-2021-08-11/putauditevents.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/cloudtrail-data-2021-08-11/putauditevents.md)

- [AWS SDK for Python](../../../../services/goto/boto3/cloudtrail-data-2021-08-11/putauditevents.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudtrail-data-2021-08-11/putauditevents.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Actions

Data Types

All content copied from https://docs.aws.amazon.com/.
