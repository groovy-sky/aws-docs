# GetDeployments

Gets information about a Deployments collection.

## Request Syntax

```nohighlight

GET /restapis/restapi_id/deployments?limit=limit&position=position HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[limit](#API_GetDeployments_RequestSyntax)**

The maximum number of returned results per page. The default value is 25 and the maximum value is 500.

**[position](#API_GetDeployments_RequestSyntax)**

The current pagination position in the paged result set.

**[restapi\_id](#API_GetDeployments_RequestSyntax)**

The string identifier of the associated RestApi.

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "item": [
      {
         "apiSummary": {
            "string" : {
               "string" : {
                  "apiKeyRequired": boolean,
                  "authorizationType": "string"
               }
            }
         },
         "createdDate": number,
         "description": "string",
         "id": "string"
      }
   ],
   "position": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[item](#API_GetDeployments_ResponseSyntax)**

The current page of elements from this collection.

Type: Array of [Deployment](api-deployment.md) objects

**[position](#API_GetDeployments_ResponseSyntax)**

The current pagination position in the paged result set.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**BadRequestException**

The submitted request is not valid, for example, the input is incomplete or incorrect. See the accompanying error message for details.

HTTP Status Code: 400

**NotFoundException**

The requested resource is not found. Make sure that the request URI is correct.

HTTP Status Code: 404

**ServiceUnavailableException**

The requested service is not available. For details see the accompanying error message. Retry after the specified time period.

HTTP Status Code: 503

**TooManyRequestsException**

The request has reached its throttling limit. Retry after the specified time period.

HTTP Status Code: 429

**UnauthorizedException**

The request is denied because the caller has insufficient permissions.

HTTP Status Code: 401

## Examples

### Get the deployments of an API

This example illustrates one usage of GetDeployments.

#### Sample Request

```

GET /restapis/fugvjdxtri/deployments?limit=2 HTTP/1.1
Content-Type: application/json
Host: apigateway.us-east-1.amazonaws.com
X-Amz-Date: 20160606T222635Z
Authorization: AWS4-HMAC-SHA256 Credential={secrete_key}/20160606/us-east-1/apigateway/aws4_request, SignedHeaders=content-type;host;x-amz-date, Signature={sig4_hash}
```

#### Sample Response

```

{
  "_links": {
    "curies": {
      "href": "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-deployment-{rel}.html",
      "name": "deployment",
      "templated": true
    },
    "self": {
      "href": "/restapis/fugvjdxtri/deployments?limit=2"
    },
    "deployment:by-id": {
      "href": "/restapis/fugvjdxtri/deployments/{deployment_id}{?embed}",
      "templated": true
    },
    "deployment:create": {
      "href": "/restapis/fugvjdxtri/deployments"
    },
    "item": [
      {
        "href": "/restapis/fugvjdxtri/deployments/4vvti6"
      },
      {
        "href": "/restapis/fugvjdxtri/deployments/a9kdln"
      }
    ],
    "next": {
      "href": "/restapis/fugvjdxtri/deployments?position=aWQ9U2E5a2Rsbg%3D%3D&limit=2"
    }
  },
  "_embedded": {
    "item": [
      {
        "_links": {
          "self": {
            "href": "/restapis/fugvjdxtri/deployments/4vvti6"
          },
          "deployment:delete": {
            "href": "/restapis/fugvjdxtri/deployments/4vvti6"
          },
          "deployment:stages": {
            "href": "/restapis/fugvjdxtri/stages?deployment_id=4vvti6"
          },
          "deployment:update": {
            "href": "/restapis/fugvjdxtri/deployments/4vvti6"
          }
        },
        "createdDate": "2016-06-06T17:42:37Z",
        "id": "4vvti6"
      },
      {
        "_links": {
          "self": {
            "href": "/restapis/fugvjdxtri/deployments/a9kdln"
          },
          "deployment:delete": {
            "href": "/restapis/fugvjdxtri/deployments/a9kdln"
          },
          "deployment:stages": {
            "href": "/restapis/fugvjdxtri/stages?deployment_id=a9kdln"
          },
          "deployment:update": {
            "href": "/restapis/fugvjdxtri/deployments/a9kdln"
          }
        },
        "createdDate": "2016-06-06T22:18:22Z",
        "description": "stage2 deploy2",
        "id": "a9kdln"
      }
    ]
  }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apigateway-2015-07-09/getdeployments.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigateway-2015-07-09/getdeployments.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/getdeployments.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigateway-2015-07-09/getdeployments.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/getdeployments.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigateway-2015-07-09/getdeployments.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigateway-2015-07-09/getdeployments.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigateway-2015-07-09/getdeployments.md)

- [AWS SDK for Python](../../../goto/boto3/apigateway-2015-07-09/getdeployments.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/getdeployments.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetDeployment

GetDocumentationPart

All content copied from https://docs.aws.amazon.com/.
