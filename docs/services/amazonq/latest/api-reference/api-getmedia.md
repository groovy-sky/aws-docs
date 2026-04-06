# GetMedia

Returns the image bytes corresponding to a media object. If you have implemented your own application with the Chat and ChatSync APIs, and
have enabled content extraction from visual data in Amazon Q Business, you use the GetMedia API operation to download
the images so you can show them in your UI with responses.

For more information, see [Extracting semantic meaning from images and visuals](../qbusiness-ug/extracting-meaning-from-images.md).

## Request Syntax

```nohighlight

GET /applications/applicationId/conversations/conversationId/messages/messageId/media/mediaId HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[applicationId](#API_GetMedia_RequestSyntax)**

The identifier of the Amazon Q Business which contains the media object.

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: Yes

**[conversationId](#API_GetMedia_RequestSyntax)**

The identifier of the Amazon Q Business conversation.

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: Yes

**[mediaId](#API_GetMedia_RequestSyntax)**

The identifier of the media object. You can find this in the `sourceAttributions` returned by the `Chat`, `ChatSync`, and `ListMessages` API responses.

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: Yes

**[messageId](#API_GetMedia_RequestSyntax)**

The identifier of the Amazon Q Business message.

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "mediaBytes": blob,
   "mediaMimeType": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[mediaBytes](#API_GetMedia_ResponseSyntax)**

The base64-encoded bytes of the media object.

Type: Base64-encoded binary data object

**[mediaMimeType](#API_GetMedia_ResponseSyntax)**

The MIME type of the media object (image/png).

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

You don't have access to perform this action. Make sure you have the required
permission policies and user accounts and try again.

HTTP Status Code: 403

**InternalServerException**

An issue occurred with the internal server used for your Amazon Q Business service. Wait
some minutes and try again, or contact [Support](http://aws.amazon.com/contact-us) for help.

HTTP Status Code: 500

**LicenseNotFoundException**

You don't have permissions to perform the action because your license is inactive. Ask
your admin to activate your license and try again after your licence is active.

HTTP Status Code: 400

**MediaTooLargeException**

The requested media object is too large to be returned.

HTTP Status Code: 400

**ResourceNotFoundException**

The application or plugin resource you want to use doesn’t exist. Make sure you have
provided the correct resource and try again.

**message**

The message describing a `ResourceNotFoundException`.

**resourceId**

The identifier of the resource affected.

**resourceType**

The type of the resource affected.

HTTP Status Code: 404

**ThrottlingException**

The request was denied due to throttling. Reduce the number of requests and try
again.

HTTP Status Code: 429

**ValidationException**

The input doesn't meet the constraints set by the Amazon Q Business service. Provide the
correct input and try again.

**fields**

The input field(s) that failed validation.

**message**

The message describing the `ValidationException`.

**reason**

The reason for the `ValidationException`.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/qbusiness-2023-11-27/GetMedia)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/qbusiness-2023-11-27/GetMedia)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qbusiness-2023-11-27/GetMedia)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/qbusiness-2023-11-27/GetMedia)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qbusiness-2023-11-27/GetMedia)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/qbusiness-2023-11-27/GetMedia)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/qbusiness-2023-11-27/GetMedia)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/qbusiness-2023-11-27/GetMedia)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/qbusiness-2023-11-27/GetMedia)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qbusiness-2023-11-27/GetMedia)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GetIndex

GetPlugin
