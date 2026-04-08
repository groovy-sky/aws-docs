# ListServiceOperations

Returns a list of the _operations_ of this service that have been discovered by Application Signals.
Only the operations that were invoked during the specified time range are returned.

## Request Syntax

```nohighlight

POST /service-operations?EndTime=EndTime&MaxResults=MaxResults&NextToken=NextToken&StartTime=StartTime HTTP/1.1
Content-type: application/json

{
   "KeyAttributes": {
      "string" : "string"
   }
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[EndTime](#API_ListServiceOperations_RequestSyntax)**

The end of the time period to retrieve information about. When used in a raw HTTP Query API, it is formatted as
be epoch time in seconds. For example: `1698778057`

Your requested end time will be rounded to the nearest hour.

Required: Yes

**[MaxResults](#API_ListServiceOperations_RequestSyntax)**

The maximum number of results to return in one operation. If you omit this
parameter, the default of 50 is used.

Valid Range: Minimum value of 1. Maximum value of 100.

**[NextToken](#API_ListServiceOperations_RequestSyntax)**

Include this value, if it was returned by the previous operation, to get the next set of service operations.

**[StartTime](#API_ListServiceOperations_RequestSyntax)**

The start of the time period to retrieve information about. When used in a raw HTTP Query API, it is formatted as
be epoch time in seconds. For example: `1698778057`

Your requested start time will be rounded to the nearest hour.

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[KeyAttributes](#API_ListServiceOperations_RequestSyntax)**

Use this field to specify which service you want to retrieve information for. You must specify at least the `Type`,
`Name`, and `Environment` attributes.

This is a string-to-string map. It can
include the following fields.

- `Type` designates the type of object this is.

- `ResourceType` specifies the type of the resource. This field is used only
when the value of the `Type` field is `Resource` or `AWS::Resource`.

- `Name` specifies the name of the object. This is used only if the value of the `Type` field
is `Service`, `RemoteService`, or `AWS::Service`.

- `Identifier` identifies the resource objects of this resource.
This is used only if the value of the `Type` field
is `Resource` or `AWS::Resource`.

- `Environment` specifies the location where this object is hosted, or what it belongs to.

Type: String to string map

Map Entries: Maximum number of 4 items.

Key Pattern: `[a-zA-Z]{1,50}`

Value Length Constraints: Minimum length of 1. Maximum length of 1024.

Value Pattern: `[ -~]*[!-~]+[ -~]*`

Required: Yes

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "EndTime": number,
   "NextToken": "string",
   "ServiceOperations": [
      {
         "MetricReferences": [
            {
               "AccountId": "string",
               "Dimensions": [
                  {
                     "Name": "string",
                     "Value": "string"
                  }
               ],
               "MetricName": "string",
               "MetricType": "string",
               "Namespace": "string"
            }
         ],
         "Name": "string"
      }
   ],
   "StartTime": number
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[EndTime](#API_ListServiceOperations_ResponseSyntax)**

The end of the time period that the returned information applies to. When used in a raw HTTP Query API, it is formatted as
be epoch time in seconds. For example: `1698778057`

This displays the time that Application Signals used for the request. It might not match your request exactly, because
it was rounded to the nearest hour.

Type: Timestamp

**[NextToken](#API_ListServiceOperations_ResponseSyntax)**

Include this value in your next use of this API to get next set
of service operations.

Type: String

**[ServiceOperations](#API_ListServiceOperations_ResponseSyntax)**

An array of structures that each contain information about one operation of this service.

Type: Array of [ServiceOperation](api-serviceoperation.md) objects

Array Members: Minimum number of 0 items. Maximum number of 100 items.

**[StartTime](#API_ListServiceOperations_ResponseSyntax)**

The start of the time period that the returned information applies to. When used in a raw HTTP Query API, it is formatted as
be epoch time in seconds. For example: `1698778057`

This displays the time that Application Signals used for the request. It might not match your request exactly, because
it was rounded to the nearest hour.

Type: Timestamp

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ThrottlingException**

The request was throttled because of quota limits.

HTTP Status Code: 429

**ValidationException**

The resource is not valid.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/application-signals-2024-04-15/listserviceoperations.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/application-signals-2024-04-15/listserviceoperations.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/application-signals-2024-04-15/listserviceoperations.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/application-signals-2024-04-15/listserviceoperations.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/application-signals-2024-04-15/listserviceoperations.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/application-signals-2024-04-15/listserviceoperations.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/application-signals-2024-04-15/listserviceoperations.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/application-signals-2024-04-15/listserviceoperations.md)

- [AWS SDK for Python](../../../../services/goto/boto3/application-signals-2024-04-15/listserviceoperations.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/application-signals-2024-04-15/listserviceoperations.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ListServiceLevelObjectives

ListServices
