# GetResource

Lists information about a resource.

## Request Syntax

```nohighlight

GET /restapis/restapi_id/resources/resource_id?embed=embed HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[embed](#API_GetResource_RequestSyntax)**

A query parameter to retrieve the specified resources embedded in the returned Resource representation in the response. This `embed` parameter value is a list of comma-separated strings. Currently, the request supports only retrieval of the embedded Method resources this way. The query parameter value must be a single-valued list and contain the `"methods"` string. For example, `GET /restapis/{restapi_id}/resources/{resource_id}?embed=methods`.

**[resource\_id](#API_GetResource_RequestSyntax)**

The identifier for the Resource resource.

Required: Yes

**[restapi\_id](#API_GetResource_RequestSyntax)**

The string identifier of the associated RestApi.

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "id": "string",
   "parentId": "string",
   "path": "string",
   "pathPart": "string",
   "resourceMethods": {
      "string" : {
         "apiKeyRequired": boolean,
         "authorizationScopes": [ "string" ],
         "authorizationType": "string",
         "authorizerId": "string",
         "httpMethod": "string",
         "methodIntegration": {
            "cacheKeyParameters": [ "string" ],
            "cacheNamespace": "string",
            "connectionId": "string",
            "connectionType": "string",
            "contentHandling": "string",
            "credentials": "string",
            "httpMethod": "string",
            "integrationResponses": {
               "string" : {
                  "contentHandling": "string",
                  "responseParameters": {
                     "string" : "string"
                  },
                  "responseTemplates": {
                     "string" : "string"
                  },
                  "selectionPattern": "string",
                  "statusCode": "string"
               }
            },
            "integrationTarget": "string",
            "passthroughBehavior": "string",
            "requestParameters": {
               "string" : "string"
            },
            "requestTemplates": {
               "string" : "string"
            },
            "responseTransferMode": "string",
            "timeoutInMillis": number,
            "tlsConfig": {
               "insecureSkipVerification": boolean
            },
            "type": "string",
            "uri": "string"
         },
         "methodResponses": {
            "string" : {
               "responseModels": {
                  "string" : "string"
               },
               "responseParameters": {
                  "string" : boolean
               },
               "statusCode": "string"
            }
         },
         "operationName": "string",
         "requestModels": {
            "string" : "string"
         },
         "requestParameters": {
            "string" : boolean
         },
         "requestValidatorId": "string"
      }
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[id](#API_GetResource_ResponseSyntax)**

The resource's identifier.

Type: String

**[parentId](#API_GetResource_ResponseSyntax)**

The parent resource's identifier.

Type: String

**[path](#API_GetResource_ResponseSyntax)**

The full path for this resource.

Type: String

**[pathPart](#API_GetResource_ResponseSyntax)**

The last path segment for this resource.

Type: String

**[resourceMethods](#API_GetResource_ResponseSyntax)**

Gets an API resource's method of a given HTTP verb.

Type: String to [Method](api-method.md) object map

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**NotFoundException**

The requested resource is not found. Make sure that the request URI is correct.

HTTP Status Code: 404

**TooManyRequestsException**

The request has reached its throttling limit. Retry after the specified time period.

HTTP Status Code: 429

**UnauthorizedException**

The request is denied because the caller has insufficient permissions.

HTTP Status Code: 401

## Examples

### Get an API resource of a given resource identifier

This example illustrates one usage of GetResource.

#### Sample Request

```

GET /restapis/fugvjdxtri/resources/3kzxbg5sa2 HTTP/1.1
Content-Type: application/json
Host: apigateway.us-east-1.amazonaws.com
X-Amz-Date: 20170223T025309Z
Authorization: AWS4-HMAC-SHA256 Credential={access_key_ID}/20170223/us-east-1/apigateway/aws4_request, SignedHeaders=content-type;host;x-amz-date, Signature={sig4_hash}
```

#### Sample Response

```

{
  "_links": {
    "curies": [
      {
        "href": "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-method-{rel}.html",
        "name": "method",c
        "templated": true
      },
      {
        "href": "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-resource-{rel}.html",
        "name": "resource",
        "templated": true
      }
    ],
    "self": {
      "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2"
    },
    "method:by-http-method": {
      "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/{http_method}",
      "templated": true
    },
    "method:put": {
      "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/{http_method}",
      "templated": true
    },
    "resource:create-child": {
      "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2"
    },
    "resource:methods": {
      "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET",
      "name": "GET",
      "title": "GET"
    },
    "resource:update": {
      "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2"
    }
  },
  "id": "3kzxbg5sa2",
  "path": "/"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apigateway-2015-07-09/getresource.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigateway-2015-07-09/getresource.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/getresource.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigateway-2015-07-09/getresource.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/getresource.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigateway-2015-07-09/getresource.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigateway-2015-07-09/getresource.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigateway-2015-07-09/getresource.md)

- [AWS SDK for Python](../../../goto/boto3/apigateway-2015-07-09/getresource.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/getresource.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetRequestValidators

GetResources

All content copied from https://docs.aws.amazon.com/.
