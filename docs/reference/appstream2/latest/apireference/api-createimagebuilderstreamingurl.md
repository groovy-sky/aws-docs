# CreateImageBuilderStreamingURL

Creates a URL to start an image builder streaming session.

## Request Syntax

```nohighlight

{
   "Name": "string",
   "Validity": number
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[Name](#API_CreateImageBuilderStreamingURL_RequestSyntax)**

The name of the image builder.

Type: String

Length Constraints: Minimum length of 1.

Required: Yes

**[Validity](#API_CreateImageBuilderStreamingURL_RequestSyntax)**

The time that the streaming URL will be valid, in seconds.
Specify a value between 1 and 604800 seconds. The default is 3600 seconds.

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

**[Expires](#API_CreateImageBuilderStreamingURL_ResponseSyntax)**

The elapsed time, in seconds after the Unix epoch, when this URL expires.

Type: Timestamp

**[StreamingURL](#API_CreateImageBuilderStreamingURL_ResponseSyntax)**

The URL to start the WorkSpaces Applications streaming session.

Type: String

Length Constraints: Minimum length of 1.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**OperationNotPermittedException**

The attempted operation is not permitted.

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appstream-2016-12-01/createimagebuilderstreamingurl.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appstream-2016-12-01/createimagebuilderstreamingurl.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appstream-2016-12-01/createimagebuilderstreamingurl.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appstream-2016-12-01/createimagebuilderstreamingurl.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appstream-2016-12-01/createimagebuilderstreamingurl.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appstream-2016-12-01/createimagebuilderstreamingurl.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appstream-2016-12-01/createimagebuilderstreamingurl.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appstream-2016-12-01/createimagebuilderstreamingurl.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appstream-2016-12-01/createimagebuilderstreamingurl.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appstream-2016-12-01/createimagebuilderstreamingurl.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateImageBuilder

CreateImportedImage

All content copied from https://docs.aws.amazon.com/.
