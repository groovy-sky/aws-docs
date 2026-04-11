# DeleteEndpoint

Deletes an endpoint.

###### Note

It can take up to 5 minutes for this action to finish.

Related actions include:

- [CreateEndpoint](api-s3outposts-createendpoint.md)

- [ListEndpoints](api-s3outposts-listendpoints.md)

## Request Syntax

```nohighlight

DELETE /S3Outposts/DeleteEndpoint?endpointId=EndpointId&outpostId=OutpostId HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[EndpointId](#API_s3outposts_DeleteEndpoint_RequestSyntax)**

The ID of the endpoint.

Pattern: `^[a-zA-Z0-9]{19}$`

Required: Yes

**[OutpostId](#API_s3outposts_DeleteEndpoint_RequestSyntax)**

The ID of the AWS Outposts.

Pattern: `^(op-[a-f0-9]{17}|\d{12}|ec2)$`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```

HTTP/1.1 200

```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

**AccessDeniedException**

Access was denied for this action.

HTTP Status Code: 403

**InternalServerException**

There was an exception with the internal server.

HTTP Status Code: 500

**OutpostOfflineException**

The service link connection to your Outposts home Region is down. Check your connection and try again.

HTTP Status Code: 400

**ResourceNotFoundException**

The requested resource was not found.

HTTP Status Code: 404

**ThrottlingException**

The request was denied due to request throttling.

HTTP Status Code: 429

**ValidationException**

There was an exception validating this data.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3outposts-2017-07-25/deleteendpoint.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3outposts-2017-07-25/deleteendpoint.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3outposts-2017-07-25/deleteendpoint.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3outposts-2017-07-25/deleteendpoint.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3outposts-2017-07-25/deleteendpoint.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3outposts-2017-07-25/deleteendpoint.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3outposts-2017-07-25/deleteendpoint.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3outposts-2017-07-25/deleteendpoint.md)

- [AWS SDK for Python](../../../goto/boto3/s3outposts-2017-07-25/deleteendpoint.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3outposts-2017-07-25/deleteendpoint.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateEndpoint

ListEndpoints

All content copied from https://docs.aws.amazon.com/.
