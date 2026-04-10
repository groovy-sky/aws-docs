# ListPublicKeys

Returns all public keys whose private keys were used to sign the digest files within the
specified time range. The public key is needed to validate digest files that were signed
with its corresponding private key.

###### Note

CloudTrail uses different private and public key pairs per Region. Each digest
file is signed with a private key unique to its Region. When you validate a digest file
from a specific Region, you must look in the same Region for its corresponding public
key.

## Request Syntax

```nohighlight

{
   "EndTime": number,
   "NextToken": "string",
   "StartTime": number
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[EndTime](#API_ListPublicKeys_RequestSyntax)**

Optionally specifies, in UTC, the end of the time range to look up public keys for
CloudTrail digest files. If not specified, the current time is used.

Type: Timestamp

Required: No

**[NextToken](#API_ListPublicKeys_RequestSyntax)**

Reserved for future use.

Type: String

Required: No

**[StartTime](#API_ListPublicKeys_RequestSyntax)**

Optionally specifies, in UTC, the start of the time range to look up public keys for
CloudTrail digest files. If not specified, the current time is used, and the
current public key is returned.

Type: Timestamp

Required: No

## Response Syntax

```nohighlight

{
   "NextToken": "string",
   "PublicKeyList": [
      {
         "Fingerprint": "string",
         "ValidityEndTime": number,
         "ValidityStartTime": number,
         "Value": blob
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[NextToken](#API_ListPublicKeys_ResponseSyntax)**

Reserved for future use.

Type: String

**[PublicKeyList](#API_ListPublicKeys_ResponseSyntax)**

Contains an array of PublicKey objects.

###### Note

The returned public keys may have validity time ranges that overlap.

Type: Array of [PublicKey](api-publickey.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidTimeRangeException**

Occurs if the timestamp values are not valid. Either the start time occurs after the end
time, or the time range is outside the range of possible values.

HTTP Status Code: 400

**InvalidTokenException**

Reserved for future use.

HTTP Status Code: 400

**OperationNotPermittedException**

This exception is thrown when the requested operation is not permitted.

HTTP Status Code: 400

**UnsupportedOperationException**

This exception is thrown when the requested operation is not supported.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/cloudtrail-2013-11-01/listpublickeys.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/cloudtrail-2013-11-01/listpublickeys.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudtrail-2013-11-01/listpublickeys.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/cloudtrail-2013-11-01/listpublickeys.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudtrail-2013-11-01/listpublickeys.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/cloudtrail-2013-11-01/listpublickeys.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/cloudtrail-2013-11-01/listpublickeys.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/cloudtrail-2013-11-01/listpublickeys.md)

- [AWS SDK for Python](../../../../services/goto/boto3/cloudtrail-2013-11-01/listpublickeys.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudtrail-2013-11-01/listpublickeys.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListInsightsMetricData

ListQueries

All content copied from https://docs.aws.amazon.com/.
