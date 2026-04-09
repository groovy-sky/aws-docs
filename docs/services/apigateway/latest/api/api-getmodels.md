# GetModels

Describes existing Models defined for a RestApi resource.

## Request Syntax

```nohighlight

GET /restapis/restapi_id/models?limit=limit&position=position HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[limit](#API_GetModels_RequestSyntax)**

The maximum number of returned results per page. The default value is 25 and the maximum value is 500.

**[position](#API_GetModels_RequestSyntax)**

The current pagination position in the paged result set.

**[restapi\_id](#API_GetModels_RequestSyntax)**

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
         "contentType": "string",
         "description": "string",
         "id": "string",
         "name": "string",
         "schema": "string"
      }
   ],
   "position": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[item](#API_GetModels_ResponseSyntax)**

The current page of elements from this collection.

Type: Array of [Model](api-model.md) objects

**[position](#API_GetModels_ResponseSyntax)**

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

**TooManyRequestsException**

The request has reached its throttling limit. Retry after the specified time period.

HTTP Status Code: 429

**UnauthorizedException**

The request is denied because the caller has insufficient permissions.

HTTP Status Code: 401

## Examples

### Get the collection of models defined for an API

This example illustrates one usage of GetModels.

#### Sample Request

```

GET /restapis/l9kujxkzq2/models HTTP/1.1
Content-Type: application/json
Host: apigateway.us-east-1.amazonaws.com
Content-Length: 60
X-Amz-Date: 20170223T172652Z
Authorization: AWS4-HMAC-SHA256 Credential={access_key_ID}/20170223/us-east-1/apigateway/aws4_request, SignedHeaders=content-type;host;x-amz-date, Signature={sig4_hash}
```

#### Sample Response

```

{
  "_links": {
    "curies": {
      "href": "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-model-{rel}.html",
      "name": "model",
      "templated": true
    },
    "self": {
      "href": "/restapis/l9kujxkzq2/models"
    },
    "item": [
      {
        "href": "/restapis/l9kujxkzq2/models/Empty?flatten=false"
      },
      {
        "href": "/restapis/l9kujxkzq2/models/Error?flatten=false"
      }
    ],
    "model:by-name": {
      "href": "/restapis/l9kujxkzq2/models/{model_name}?flatten=false",
      "templated": true
    }
  },
  "_embedded": {
    "item": [
      {
        "_links": {
          "self": {
            "href": "/restapis/l9kujxkzq2/models/Empty?flatten=false"
          },
          "model:create": {
            "href": "/restapis/l9kujxkzq2/models"
          },
          "model:delete": {
            "href": "/restapis/l9kujxkzq2/models/Empty"
          },
          "model:generate-template": {
            "href": "/restapis/l9kujxkzq2/models/Empty/default_template"
          },
          "model:update": {
            "href": "/restapis/l9kujxkzq2/models/Empty"
          }
        },
        "contentType": "application/json",
        "description": "This is a default empty schema model",
        "id": "71l0yh",
        "name": "Empty",
        "schema": "{\n  \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n  \"title\" : \"Empty Schema\",\n  \"type\" : \"object\"\n}"
      },
      {
        "_links": {
          "self": {
            "href": "/restapis/l9kujxkzq2/models/Error?flatten=false"
          },
          "model:create": {
            "href": "/restapis/l9kujxkzq2/models"
          },
          "model:delete": {
            "href": "/restapis/l9kujxkzq2/models/Error"
          },
          "model:generate-template": {
            "href": "/restapis/l9kujxkzq2/models/Error/default_template"
          },
          "model:update": {
            "href": "/restapis/l9kujxkzq2/models/Error"
          }
        },
        "contentType": "application/json",
        "description": "This is a default error schema model",
        "id": "sddpxt",
        "name": "Error",
        "schema": "{\n  \"$schema\" : \"http://json-schema.org/draft-04/schema#\",\n  \"title\" : \"Error Schema\",\n  \"type\" : \"object\",\n  \"properties\" : {\n    \"message\" : { \"type\" : \"string\" }\n  }\n}"
      }
    ]
  }
}

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apigateway-2015-07-09/getmodels.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigateway-2015-07-09/getmodels.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/getmodels.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigateway-2015-07-09/getmodels.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/getmodels.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigateway-2015-07-09/getmodels.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigateway-2015-07-09/getmodels.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigateway-2015-07-09/getmodels.md)

- [AWS SDK for Python](../../../goto/boto3/apigateway-2015-07-09/getmodels.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/getmodels.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetModel

GetModelTemplate

All content copied from https://docs.aws.amazon.com/.
