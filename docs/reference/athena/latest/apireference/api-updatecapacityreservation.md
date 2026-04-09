# UpdateCapacityReservation

Updates the number of requested data processing units for the capacity reservation
with the specified name.

## Request Syntax

```nohighlight

{
   "Name": "string",
   "TargetDpus": number
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[Name](#API_UpdateCapacityReservation_RequestSyntax)**

The name of the capacity reservation.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Pattern: `[a-zA-Z0-9._-]+`

Required: Yes

**[TargetDpus](#API_UpdateCapacityReservation_RequestSyntax)**

The new number of requested data processing units.

Type: Integer

Valid Range: Minimum value of 24.

Required: Yes

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalServerException**

Indicates a platform issue, which may be due to a transient condition or
outage.

HTTP Status Code: 500

**InvalidRequestException**

Indicates that something is wrong with the input to the request. For example, a
required parameter may be missing or out of range.

**AthenaErrorCode**

The error code returned when the query execution failed to process, or when the
processing request for the named query failed.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/athena-2017-05-18/updatecapacityreservation.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/athena-2017-05-18/updatecapacityreservation.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/updatecapacityreservation.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/athena-2017-05-18/updatecapacityreservation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/updatecapacityreservation.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/athena-2017-05-18/updatecapacityreservation.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/athena-2017-05-18/updatecapacityreservation.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/athena-2017-05-18/updatecapacityreservation.md)

- [AWS SDK for Python](../../../../services/goto/boto3/athena-2017-05-18/updatecapacityreservation.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/updatecapacityreservation.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UntagResource

UpdateDataCatalog

All content copied from https://docs.aws.amazon.com/.
