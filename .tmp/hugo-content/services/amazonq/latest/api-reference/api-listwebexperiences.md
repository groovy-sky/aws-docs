# ListWebExperiences

Lists one or more Amazon Q Business Web Experiences.

## Request Syntax

```nohighlight

GET /applications/applicationId/experiences?maxResults=maxResults&nextToken=nextToken HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[applicationId](#API_ListWebExperiences_RequestSyntax)**

The identifier of the Amazon Q Business application linked to the listed web
experiences.

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: Yes

**[maxResults](#API_ListWebExperiences_RequestSyntax)**

The maximum number of Amazon Q Business Web Experiences to return.

Valid Range: Minimum value of 1. Maximum value of 100.

**[nextToken](#API_ListWebExperiences_RequestSyntax)**

If the `maxResults` response was incomplete because there is more data to
retrieve, Amazon Q Business returns a pagination token in the response. You can use this
pagination token to retrieve the next set of Amazon Q Business conversations.

Length Constraints: Minimum length of 1. Maximum length of 800.

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "nextToken": "string",
   "webExperiences": [
      {
         "createdAt": number,
         "defaultEndpoint": "string",
         "status": "string",
         "updatedAt": number,
         "webExperienceId": "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[nextToken](#API_ListWebExperiences_ResponseSyntax)**

If the response is truncated, Amazon Q Business returns this token, which you can use in a
later request to list the next set of messages.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 800.

**[webExperiences](#API_ListWebExperiences_ResponseSyntax)**

An array of summary information for one or more Amazon Q Business experiences.

Type: Array of [WebExperience](api-webexperience.md) objects

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

- [AWS Command Line Interface V2](../../../goto/cli2/qbusiness-2023-11-27/listwebexperiences.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/qbusiness-2023-11-27/listwebexperiences.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/listwebexperiences.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/qbusiness-2023-11-27/listwebexperiences.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/listwebexperiences.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/qbusiness-2023-11-27/listwebexperiences.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/qbusiness-2023-11-27/listwebexperiences.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/qbusiness-2023-11-27/listwebexperiences.md)

- [AWS SDK for Python](../../../goto/boto3/qbusiness-2023-11-27/listwebexperiences.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/listwebexperiences.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListTagsForResource

PutFeedback

All content copied from https://docs.aws.amazon.com/.
