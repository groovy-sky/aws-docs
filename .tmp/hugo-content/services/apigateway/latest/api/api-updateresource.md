# UpdateResource

Changes information about a Resource resource.

## Request Syntax

```nohighlight

PATCH /restapis/restapi_id/resources/resource_id HTTP/1.1
Content-type: application/json

{
   "patchOperations": [
      {
         "from": "string",
         "op": "string",
         "path": "string",
         "value": "string"
      }
   ]
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[resource\_id](#API_UpdateResource_RequestSyntax)**

The identifier of the Resource resource.

Required: Yes

**[restapi\_id](#API_UpdateResource_RequestSyntax)**

The string identifier of the associated RestApi.

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[patchOperations](#API_UpdateResource_RequestSyntax)**

For more information about supported patch operations, see [Patch Operations](patch-operations.md).

Type: Array of [PatchOperation](api-patchoperation.md) objects

Required: No

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

**[id](#API_UpdateResource_ResponseSyntax)**

The resource's identifier.

Type: String

**[parentId](#API_UpdateResource_ResponseSyntax)**

The parent resource's identifier.

Type: String

**[path](#API_UpdateResource_ResponseSyntax)**

The full path for this resource.

Type: String

**[pathPart](#API_UpdateResource_ResponseSyntax)**

The last path segment for this resource.

Type: String

**[resourceMethods](#API_UpdateResource_ResponseSyntax)**

Gets an API resource's method of a given HTTP verb.

Type: String to [Method](api-method.md) object map

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**BadRequestException**

The submitted request is not valid, for example, the input is incomplete or incorrect. See the accompanying error message for details.

HTTP Status Code: 400

**ConflictException**

The request configuration has conflicts. For details, see the accompanying error message.

HTTP Status Code: 409

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

### Update a resource

This example illustrates one usage of UpdateResource.

#### Sample Request

```

PATCH /restapis/86l3267lf6/resources/h9m85b HTTP/1.1
Content-Type: application/json
Host: apigateway.us-east-1.amazonaws.com
X-Amz-Date: 20170223T185829Z
Authorization: AWS4-HMAC-SHA256 Credential={access_key_ID}/20170223/us-east-1/apigateway/aws4_request, SignedHeaders=content-length;content-type;host;x-amz-date, Signature={sig4_hash}

{
    "patchOperations" : [{
        "op": "replace",
        "path": "/pathPart",
        "value" : "r1-2"
    },
    {
        "op": "replace",
        "path": "/parentId",
        "value": "nprcay"
    }]
}
```

#### Sample Response

```

{
  "_links": {
    "curies": [
      {
        "href": "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-method-{rel}.html",
        "name": "method",
        "templated": true
      },
      {
        "href": "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-resource-{rel}.html",
        "name": "resource",
        "templated": true
      }
    ],
    "self": {
      "href": "/restapis/86l3267lf6/resources/h9m85b"
    },
    "method:by-http-method": {
      "href": "/restapis/86l3267lf6/resources/h9m85b/methods/{http_method}",
      "templated": true
    },
    "method:put": {
      "href": "/restapis/86l3267lf6/resources/h9m85b/methods/{http_method}",
      "templated": true
    },
    "resource:create-child": {
      "href": "/restapis/86l3267lf6/resources/h9m85b"
    },
    "resource:delete": {
      "href": "/restapis/86l3267lf6/resources/h9m85b"
    },
    "resource:update": {
      "href": "/restapis/86l3267lf6/resources/h9m85b"
    }
  },
  "id": "h9m85b",
  "parentId": "nprcay",
  "path": "/r1/r1-2",
  "pathPart": "r1-2"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apigateway-2015-07-09/updateresource.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigateway-2015-07-09/updateresource.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/updateresource.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigateway-2015-07-09/updateresource.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/updateresource.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigateway-2015-07-09/updateresource.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigateway-2015-07-09/updateresource.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigateway-2015-07-09/updateresource.md)

- [AWS SDK for Python](../../../goto/boto3/apigateway-2015-07-09/updateresource.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/updateresource.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UpdateRequestValidator

UpdateRestApi

All content copied from https://docs.aws.amazon.com/.
