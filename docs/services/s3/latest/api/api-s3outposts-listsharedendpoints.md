# ListSharedEndpoints

Lists all endpoints associated with an Outpost that has been shared by AWS Resource Access Manager (RAM).

Related actions include:

- [CreateEndpoint](api-s3outposts-createendpoint.md)

- [DeleteEndpoint](api-s3outposts-deleteendpoint.md)

## Request Syntax

```nohighlight

GET /S3Outposts/ListSharedEndpoints?maxResults=MaxResults&nextToken=NextToken&outpostId=OutpostId HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[MaxResults](#API_s3outposts_ListSharedEndpoints_RequestSyntax)**

The maximum number of endpoints that will be returned in the response.

Valid Range: Minimum value of 0. Maximum value of 100.

**[NextToken](#API_s3outposts_ListSharedEndpoints_RequestSyntax)**

If a previous response from this operation included a `NextToken` value, you
can provide that value here to retrieve the next page of results.

Length Constraints: Minimum length of 1. Maximum length of 1024.

Pattern: `^[A-Za-z0-9\+\:\/\=\?\#-_]+$`

**[OutpostId](#API_s3outposts_ListSharedEndpoints_RequestSyntax)**

The ID of the AWS Outpost.

Pattern: `^(op-[a-f0-9]{17}|\d{12}|ec2)$`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "Endpoints": [
      {
         "AccessType": "string",
         "CidrBlock": "string",
         "CreationTime": number,
         "CustomerOwnedIpv4Pool": "string",
         "EndpointArn": "string",
         "FailedReason": {
            "ErrorCode": "string",
            "Message": "string"
         },
         "NetworkInterfaces": [
            {
               "NetworkInterfaceId": "string"
            }
         ],
         "OutpostsId": "string",
         "SecurityGroupId": "string",
         "Status": "string",
         "SubnetId": "string",
         "VpcId": "string"
      }
   ],
   "NextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[Endpoints](#API_s3outposts_ListSharedEndpoints_ResponseSyntax)**

The list of endpoints associated with the specified Outpost that have been shared by AWS Resource Access Manager (RAM).

Type: Array of [Endpoint](api-s3outposts-endpoint.md) objects

**[NextToken](#API_s3outposts_ListSharedEndpoints_ResponseSyntax)**

If the number of endpoints associated with the specified Outpost exceeds `MaxResults`,
you can include this value in subsequent calls to this operation to retrieve more results.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Pattern: `^[A-Za-z0-9\+\:\/\=\?\#-_]+$`

## Errors

**AccessDeniedException**

Access was denied for this action.

HTTP Status Code: 403

**InternalServerException**

There was an exception with the internal server.

HTTP Status Code: 500

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

- [AWS Command Line Interface V2](../../../goto/cli2/s3outposts-2017-07-25/listsharedendpoints.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3outposts-2017-07-25/listsharedendpoints.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3outposts-2017-07-25/listsharedendpoints.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3outposts-2017-07-25/listsharedendpoints.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3outposts-2017-07-25/listsharedendpoints.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3outposts-2017-07-25/listsharedendpoints.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3outposts-2017-07-25/listsharedendpoints.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3outposts-2017-07-25/listsharedendpoints.md)

- [AWS SDK for Python](../../../goto/boto3/s3outposts-2017-07-25/listsharedendpoints.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3outposts-2017-07-25/listsharedendpoints.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListOutpostsWithS3

Amazon S3 Tables

All content copied from https://docs.aws.amazon.com/.
