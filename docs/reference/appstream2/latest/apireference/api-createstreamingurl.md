# CreateStreamingURL

Creates a temporary URL to start an WorkSpaces Applications streaming session for the specified user. A streaming URL enables application streaming to be tested without user setup.

## Request Syntax

```nohighlight

{
   "ApplicationId": "string",
   "FleetName": "string",
   "SessionContext": "string",
   "StackName": "string",
   "UserId": "string",
   "Validity": number
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[ApplicationId](#API_CreateStreamingURL_RequestSyntax)**

The name of the application to launch after the session starts. This is the name that you specified
as **Name** in the Image Assistant. If your fleet is enabled for the **Desktop** stream view, you can also choose to launch directly to the operating system desktop. To do so, specify **Desktop**.

Type: String

Length Constraints: Minimum length of 1.

Required: No

**[FleetName](#API_CreateStreamingURL_RequestSyntax)**

The name of the fleet.

Type: String

Length Constraints: Minimum length of 1.

Required: Yes

**[SessionContext](#API_CreateStreamingURL_RequestSyntax)**

The session context. For more information, see [Session Context](../../../../services/appstream2/latest/developerguide/managing-stacks-fleets.md#managing-stacks-fleets-parameters) in the _Amazon WorkSpaces Applications Administration Guide_.

Type: String

Length Constraints: Minimum length of 1.

Required: No

**[StackName](#API_CreateStreamingURL_RequestSyntax)**

The name of the stack.

Type: String

Length Constraints: Minimum length of 1.

Required: Yes

**[UserId](#API_CreateStreamingURL_RequestSyntax)**

The identifier of the user.

Type: String

Length Constraints: Minimum length of 2. Maximum length of 32.

Pattern: `[\w+=,.@-]*`

Required: Yes

**[Validity](#API_CreateStreamingURL_RequestSyntax)**

The time that the streaming URL will be valid, in seconds.
Specify a value between 1 and 604800 seconds. The default is 60 seconds.

Type: Long

Required: No

## Response Syntax

```nohighlight

{
   "Expires": number,
   "StreamingURL": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[Expires](#API_CreateStreamingURL_ResponseSyntax)**

The elapsed time, in seconds after the Unix epoch, when this URL expires.

Type: Timestamp

**[StreamingURL](#API_CreateStreamingURL_ResponseSyntax)**

The URL to start the WorkSpaces Applications streaming session.

Type: String

Length Constraints: Minimum length of 1.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterCombinationException**

Indicates an incorrect combination of parameters, or a missing parameter.

**Message**

The error message in the exception.

HTTP Status Code: 400

**OperationNotPermittedException**

The attempted operation is not permitted.

**Message**

The error message in the exception.

HTTP Status Code: 400

**ResourceNotAvailableException**

The specified resource exists and is not in use, but isn't available.

**Message**

The error message in the exception.

HTTP Status Code: 400

**ResourceNotFoundException**

The specified resource was not found.

**Message**

The error message in the exception.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appstream-2016-12-01/createstreamingurl.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appstream-2016-12-01/createstreamingurl.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appstream-2016-12-01/createstreamingurl.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appstream-2016-12-01/createstreamingurl.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appstream-2016-12-01/createstreamingurl.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appstream-2016-12-01/createstreamingurl.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appstream-2016-12-01/createstreamingurl.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appstream-2016-12-01/createstreamingurl.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appstream-2016-12-01/createstreamingurl.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appstream-2016-12-01/createstreamingurl.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateStack

CreateThemeForStack

All content copied from https://docs.aws.amazon.com/.
