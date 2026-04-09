# GetGatewayResponses

Gets the GatewayResponses collection on the given RestApi. If an API developer has not added any definitions for gateway responses, the result will be the API Gateway-generated default GatewayResponses collection for the supported response types.

## Request Syntax

```nohighlight

GET /restapis/restapi_id/gatewayresponses?limit=limit&position=position HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[limit](#API_GetGatewayResponses_RequestSyntax)**

The maximum number of returned results per page. The default value is 25 and the maximum value is 500. The GatewayResponses collection does not support pagination and the limit does not apply here.

**[position](#API_GetGatewayResponses_RequestSyntax)**

The current pagination position in the paged result set. The GatewayResponse collection does not support pagination and the position does not apply here.

**[restapi\_id](#API_GetGatewayResponses_RequestSyntax)**

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
         "defaultResponse": boolean,
         "responseParameters": {
            "string" : "string"
         },
         "responseTemplates": {
            "string" : "string"
         },
         "responseType": "string",
         "statusCode": "string"
      }
   ],
   "position": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[item](#API_GetGatewayResponses_ResponseSyntax)**

Returns the entire collection, because of no pagination support.

Type: Array of [GatewayResponse](api-gatewayresponse.md) objects

**[position](#API_GetGatewayResponses_ResponseSyntax)**

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

### Get the collection of gateway responses of an API

This example illustrates one usage of GetGatewayResponses.

#### Sample Request

```

GET /restapis/o81lxisefl/gatewayresponses HTTP/1.1
Host: beta-apigateway.us-east-1.amazonaws.com
Content-Type: application/json
X-Amz-Date: 20170503T220604Z
Authorization: AWS4-HMAC-SHA256 Credential={access-key-id}/20170503/us-east-1/apigateway/aws4_request, SignedHeaders=content-type;host;x-amz-date, Signature=59b42fe54a76a5de8adf2c67baa6d39206f8e9ad49a1d77ccc6a5da3103a398a
Cache-Control: no-cache
Postman-Token: 5637af27-dc29-fc5c-9dfe-0645d52cb515
```

#### Sample Response

```

{
  "_links": {
    "curies": {
      "href": "http://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-gatewayresponse-{rel}.html",
      "name": "gatewayresponse",
      "templated": true
    },
    "self": {
      "href": "/restapis/o81lxisefl/gatewayresponses"
    },
    "first": {
      "href": "/restapis/o81lxisefl/gatewayresponses"
    },
    "gatewayresponse:by-type": {
      "href": "/restapis/o81lxisefl/gatewayresponses/{response_type}",
      "templated": true
    },
    "item": [
      {
        "href": "/restapis/o81lxisefl/gatewayresponses/INTEGRATION_FAILURE"
      },
      {
        "href": "/restapis/o81lxisefl/gatewayresponses/RESOURCE_NOT_FOUND"
      },
      {
        "href": "/restapis/o81lxisefl/gatewayresponses/REQUEST_TOO_LARGE"
      },
      {
        "href": "/restapis/o81lxisefl/gatewayresponses/THROTTLED"
      },
      {
        "href": "/restapis/o81lxisefl/gatewayresponses/UNSUPPORTED_MEDIA_TYPE"
      },
      {
        "href": "/restapis/o81lxisefl/gatewayresponses/AUTHORIZER_CONFIGURATION_ERROR"
      },
      {
        "href": "/restapis/o81lxisefl/gatewayresponses/DEFAULT_5XX"
      },
      {
        "href": "/restapis/o81lxisefl/gatewayresponses/DEFAULT_4XX"
      },
      {
        "href": "/restapis/o81lxisefl/gatewayresponses/BAD_REQUEST_PARAMETERS"
      },
      {
        "href": "/restapis/o81lxisefl/gatewayresponses/BAD_REQUEST_BODY"
      },
      {
        "href": "/restapis/o81lxisefl/gatewayresponses/EXPIRED_TOKEN"
      },
      {
        "href": "/restapis/o81lxisefl/gatewayresponses/ACCESS_DENIED"
      },
      {
        "href": "/restapis/o81lxisefl/gatewayresponses/INVALID_API_KEY"
      },
      {
        "href": "/restapis/o81lxisefl/gatewayresponses/UNAUTHORIZED"
      },
      {
        "href": "/restapis/o81lxisefl/gatewayresponses/API_CONFIGURATION_ERROR"
      },
      {
        "href": "/restapis/o81lxisefl/gatewayresponses/QUOTA_EXCEEDED"
      },
      {
        "href": "/restapis/o81lxisefl/gatewayresponses/INTEGRATION_TIMEOUT"
      },
      {
        "href": "/restapis/o81lxisefl/gatewayresponses/MISSING_AUTHENTICATION_TOKEN"
      },
      {
        "href": "/restapis/o81lxisefl/gatewayresponses/INVALID_SIGNATURE"
      },
      {
        "href": "/restapis/o81lxisefl/gatewayresponses/AUTHORIZER_FAILURE"
      },
      {
        "href": "/restapis/o81lxisefl/gatewayresponses/WAF_FILTERED"
      }
    ]
  },
  "_embedded": {
    "item": [
      {
        "_links": {
          "self": {
            "href": "/restapis/o81lxisefl/gatewayresponses/INTEGRATION_FAILURE"
          },
          "gatewayresponse:put": {
            "href": "/restapis/o81lxisefl/gatewayresponses/{response_type}",
            "templated": true
          },
          "gatewayresponse:update": {
            "href": "/restapis/o81lxisefl/gatewayresponses/INTEGRATION_FAILURE"
          }
        },
        "defaultResponse": true,
        "responseParameters": {},
        "responseTemplates": {
          "application/json": "{\"message\":$context.error.messageString}"
        },
        "responseType": "INTEGRATION_FAILURE",
        "statusCode": "504"
      },
      {
        "_links": {
          "self": {
            "href": "/restapis/o81lxisefl/gatewayresponses/RESOURCE_NOT_FOUND"
          },
          "gatewayresponse:put": {
            "href": "/restapis/o81lxisefl/gatewayresponses/{response_type}",
            "templated": true
          },
          "gatewayresponse:update": {
            "href": "/restapis/o81lxisefl/gatewayresponses/RESOURCE_NOT_FOUND"
          }
        },
        "defaultResponse": true,
        "responseParameters": {},
        "responseTemplates": {
          "application/json": "{\"message\":$context.error.messageString}"
        },
        "responseType": "RESOURCE_NOT_FOUND",
        "statusCode": "404"
      },
      {
        "_links": {
          "self": {
            "href": "/restapis/o81lxisefl/gatewayresponses/REQUEST_TOO_LARGE"
          },
          "gatewayresponse:put": {
            "href": "/restapis/o81lxisefl/gatewayresponses/{response_type}",
            "templated": true
          },
          "gatewayresponse:update": {
            "href": "/restapis/o81lxisefl/gatewayresponses/REQUEST_TOO_LARGE"
          }
        },
        "defaultResponse": true,
        "responseParameters": {},
        "responseTemplates": {
          "application/json": "{\"message\":$context.error.messageString}"
        },
        "responseType": "REQUEST_TOO_LARGE",
        "statusCode": "413"
      },
      {
        "_links": {
          "self": {
            "href": "/restapis/o81lxisefl/gatewayresponses/THROTTLED"
          },
          "gatewayresponse:put": {
            "href": "/restapis/o81lxisefl/gatewayresponses/{response_type}",
            "templated": true
          },
          "gatewayresponse:update": {
            "href": "/restapis/o81lxisefl/gatewayresponses/THROTTLED"
          }
        },
        "defaultResponse": true,
        "responseParameters": {},
        "responseTemplates": {
          "application/json": "{\"message\":$context.error.messageString}"
        },
        "responseType": "THROTTLED",
        "statusCode": "429"
      },
      {
        "_links": {
          "self": {
            "href": "/restapis/o81lxisefl/gatewayresponses/UNSUPPORTED_MEDIA_TYPE"
          },
          "gatewayresponse:put": {
            "href": "/restapis/o81lxisefl/gatewayresponses/{response_type}",
            "templated": true
          },
          "gatewayresponse:update": {
            "href": "/restapis/o81lxisefl/gatewayresponses/UNSUPPORTED_MEDIA_TYPE"
          }
        },
        "defaultResponse": true,
        "responseParameters": {},
        "responseTemplates": {
          "application/json": "{\"message\":$context.error.messageString}"
        },
        "responseType": "UNSUPPORTED_MEDIA_TYPE",
        "statusCode": "415"
      },
      {
        "_links": {
          "self": {
            "href": "/restapis/o81lxisefl/gatewayresponses/AUTHORIZER_CONFIGURATION_ERROR"
          },
          "gatewayresponse:put": {
            "href": "/restapis/o81lxisefl/gatewayresponses/{response_type}",
            "templated": true
          },
          "gatewayresponse:update": {
            "href": "/restapis/o81lxisefl/gatewayresponses/AUTHORIZER_CONFIGURATION_ERROR"
          }
        },
        "defaultResponse": true,
        "responseParameters": {},
        "responseTemplates": {
          "application/json": "{\"message\":$context.error.messageString}"
        },
        "responseType": "AUTHORIZER_CONFIGURATION_ERROR",
        "statusCode": "500"
      },
      {
        "_links": {
          "self": {
            "href": "/restapis/o81lxisefl/gatewayresponses/DEFAULT_5XX"
          },
          "gatewayresponse:put": {
            "href": "/restapis/o81lxisefl/gatewayresponses/{response_type}",
            "templated": true
          },
          "gatewayresponse:update": {
            "href": "/restapis/o81lxisefl/gatewayresponses/DEFAULT_5XX"
          }
        },
        "defaultResponse": true,
        "responseParameters": {},
        "responseTemplates": {
          "application/json": "{\"message\":$context.error.messageString}"
        },
        "responseType": "DEFAULT_5XX"
      },
      {
        "_links": {
          "self": {
            "href": "/restapis/o81lxisefl/gatewayresponses/DEFAULT_4XX"
          },
          "gatewayresponse:put": {
            "href": "/restapis/o81lxisefl/gatewayresponses/{response_type}",
            "templated": true
          },
          "gatewayresponse:update": {
            "href": "/restapis/o81lxisefl/gatewayresponses/DEFAULT_4XX"
          }
        },
        "defaultResponse": true,
        "responseParameters": {},
        "responseTemplates": {
          "application/json": "{\"message\":$context.error.messageString}"
        },
        "responseType": "DEFAULT_4XX"
      },
      {
        "_links": {
          "self": {
            "href": "/restapis/o81lxisefl/gatewayresponses/BAD_REQUEST_PARAMETERS"
          },
          "gatewayresponse:put": {
            "href": "/restapis/o81lxisefl/gatewayresponses/{response_type}",
            "templated": true
          },
          "gatewayresponse:update": {
            "href": "/restapis/o81lxisefl/gatewayresponses/BAD_REQUEST_PARAMETERS"
          }
        },
        "defaultResponse": true,
        "responseParameters": {},
        "responseTemplates": {
          "application/json": "{\"message\":$context.error.messageString}"
        },
        "responseType": "BAD_REQUEST_PARAMETERS",
        "statusCode": "400"
      },
      {
        "_links": {
          "self": {
            "href": "/restapis/o81lxisefl/gatewayresponses/BAD_REQUEST_BODY"
          },
          "gatewayresponse:put": {
            "href": "/restapis/o81lxisefl/gatewayresponses/{response_type}",
            "templated": true
          },
          "gatewayresponse:update": {
            "href": "/restapis/o81lxisefl/gatewayresponses/BAD_REQUEST_BODY"
          }
        },
        "defaultResponse": true,
        "responseParameters": {},
        "responseTemplates": {
          "application/json": "{\"message\":$context.error.messageString}"
        },
        "responseType": "BAD_REQUEST_BODY",
        "statusCode": "400"
      },
      {
        "_links": {
          "self": {
            "href": "/restapis/o81lxisefl/gatewayresponses/EXPIRED_TOKEN"
          },
          "gatewayresponse:put": {
            "href": "/restapis/o81lxisefl/gatewayresponses/{response_type}",
            "templated": true
          },
          "gatewayresponse:update": {
            "href": "/restapis/o81lxisefl/gatewayresponses/EXPIRED_TOKEN"
          }
        },
        "defaultResponse": true,
        "responseParameters": {},
        "responseTemplates": {
          "application/json": "{\"message\":$context.error.messageString}"
        },
        "responseType": "EXPIRED_TOKEN",
        "statusCode": "403"
      },
      {
        "_links": {
          "self": {
            "href": "/restapis/o81lxisefl/gatewayresponses/ACCESS_DENIED"
          },
          "gatewayresponse:put": {
            "href": "/restapis/o81lxisefl/gatewayresponses/{response_type}",
            "templated": true
          },
          "gatewayresponse:update": {
            "href": "/restapis/o81lxisefl/gatewayresponses/ACCESS_DENIED"
          }
        },
        "defaultResponse": true,
        "responseParameters": {},
        "responseTemplates": {
          "application/json": "{\"message\":$context.error.messageString}"
        },
        "responseType": "ACCESS_DENIED",
        "statusCode": "403"
      },
      {
        "_links": {
          "self": {
            "href": "/restapis/o81lxisefl/gatewayresponses/INVALID_API_KEY"
          },
          "gatewayresponse:put": {
            "href": "/restapis/o81lxisefl/gatewayresponses/{response_type}",
            "templated": true
          },
          "gatewayresponse:update": {
            "href": "/restapis/o81lxisefl/gatewayresponses/INVALID_API_KEY"
          }
        },
        "defaultResponse": true,
        "responseParameters": {},
        "responseTemplates": {
          "application/json": "{\"message\":$context.error.messageString}"
        },
        "responseType": "INVALID_API_KEY",
        "statusCode": "403"
      },
      {
        "_links": {
          "self": {
            "href": "/restapis/o81lxisefl/gatewayresponses/UNAUTHORIZED"
          },
          "gatewayresponse:put": {
            "href": "/restapis/o81lxisefl/gatewayresponses/{response_type}",
            "templated": true
          },
          "gatewayresponse:update": {
            "href": "/restapis/o81lxisefl/gatewayresponses/UNAUTHORIZED"
          }
        },
        "defaultResponse": true,
        "responseParameters": {},
        "responseTemplates": {
          "application/json": "{\"message\":$context.error.messageString}"
        },
        "responseType": "UNAUTHORIZED",
        "statusCode": "401"
      },
      {
        "_links": {
          "self": {
            "href": "/restapis/o81lxisefl/gatewayresponses/API_CONFIGURATION_ERROR"
          },
          "gatewayresponse:put": {
            "href": "/restapis/o81lxisefl/gatewayresponses/{response_type}",
            "templated": true
          },
          "gatewayresponse:update": {
            "href": "/restapis/o81lxisefl/gatewayresponses/API_CONFIGURATION_ERROR"
          }
        },
        "defaultResponse": true,
        "responseParameters": {},
        "responseTemplates": {
          "application/json": "{\"message\":$context.error.messageString}"
        },
        "responseType": "API_CONFIGURATION_ERROR",
        "statusCode": "500"
      },
      {
        "_links": {
          "self": {
            "href": "/restapis/o81lxisefl/gatewayresponses/QUOTA_EXCEEDED"
          },
          "gatewayresponse:put": {
            "href": "/restapis/o81lxisefl/gatewayresponses/{response_type}",
            "templated": true
          },
          "gatewayresponse:update": {
            "href": "/restapis/o81lxisefl/gatewayresponses/QUOTA_EXCEEDED"
          }
        },
        "defaultResponse": true,
        "responseParameters": {},
        "responseTemplates": {
          "application/json": "{\"message\":$context.error.messageString}"
        },
        "responseType": "QUOTA_EXCEEDED",
        "statusCode": "429"
      },
      {
        "_links": {
          "self": {
            "href": "/restapis/o81lxisefl/gatewayresponses/INTEGRATION_TIMEOUT"
          },
          "gatewayresponse:put": {
            "href": "/restapis/o81lxisefl/gatewayresponses/{response_type}",
            "templated": true
          },
          "gatewayresponse:update": {
            "href": "/restapis/o81lxisefl/gatewayresponses/INTEGRATION_TIMEOUT"
          }
        },
        "defaultResponse": true,
        "responseParameters": {},
        "responseTemplates": {
          "application/json": "{\"message\":$context.error.messageString}"
        },
        "responseType": "INTEGRATION_TIMEOUT",
        "statusCode": "504"
      },
      {
        "_links": {
          "self": {
            "href": "/restapis/o81lxisefl/gatewayresponses/MISSING_AUTHENTICATION_TOKEN"
          },
          "gatewayresponse:put": {
            "href": "/restapis/o81lxisefl/gatewayresponses/{response_type}",
            "templated": true
          },
          "gatewayresponse:update": {
            "href": "/restapis/o81lxisefl/gatewayresponses/MISSING_AUTHENTICATION_TOKEN"
          }
        },
        "defaultResponse": true,
        "responseParameters": {},
        "responseTemplates": {
          "application/json": "{\"message\":$context.error.messageString}"
        },
        "responseType": "MISSING_AUTHENTICATION_TOKEN",
        "statusCode": "403"
      },
      {
        "_links": {
          "self": {
            "href": "/restapis/o81lxisefl/gatewayresponses/INVALID_SIGNATURE"
          },
          "gatewayresponse:put": {
            "href": "/restapis/o81lxisefl/gatewayresponses/{response_type}",
            "templated": true
          },
          "gatewayresponse:update": {
            "href": "/restapis/o81lxisefl/gatewayresponses/INVALID_SIGNATURE"
          }
        },
        "defaultResponse": true,
        "responseParameters": {},
        "responseTemplates": {
          "application/json": "{\"message\":$context.error.messageString}"
        },
        "responseType": "INVALID_SIGNATURE",
        "statusCode": "403"
      },
      {
        "_links": {
          "self": {
            "href": "/restapis/o81lxisefl/gatewayresponses/AUTHORIZER_FAILURE"
          },
          "gatewayresponse:put": {
            "href": "/restapis/o81lxisefl/gatewayresponses/{response_type}",
            "templated": true
          },
          "gatewayresponse:update": {
            "href": "/restapis/o81lxisefl/gatewayresponses/AUTHORIZER_FAILURE"
          }
        },
        "defaultResponse": true,
        "responseParameters": {},
        "responseTemplates": {
          "application/json": "{\"message\":$context.error.messageString}"
        },
        "responseType": "AUTHORIZER_FAILURE",
        "statusCode": "500"
      },
      {
        "_links": {
          "self": {
            "href": "/restapis/o81lxisefl/gatewayresponses/WAF_FILTERED"
          },
          "gatewayresponse:put": {
            "href": "/restapis/o81lxisefl/gatewayresponses/{response_type}",
            "templated": true
          },
          "gatewayresponse:update": {
            "href": "/restapis/o81lxisefl/gatewayresponses/WAF_FILTERED"
          }
        },
        "defaultResponse": true,
        "responseParameters": {},
        "responseTemplates": {
          "application/json": "{\"message\":$context.error.messageString}"
        },
        "responseType": "WAF_FILTERED",
        "statusCode": "403"
      }
    ]
  }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apigateway-2015-07-09/getgatewayresponses.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigateway-2015-07-09/getgatewayresponses.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/getgatewayresponses.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigateway-2015-07-09/getgatewayresponses.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/getgatewayresponses.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigateway-2015-07-09/getgatewayresponses.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigateway-2015-07-09/getgatewayresponses.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigateway-2015-07-09/getgatewayresponses.md)

- [AWS SDK for Python](../../../goto/boto3/apigateway-2015-07-09/getgatewayresponses.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/getgatewayresponses.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetGatewayResponse

GetIntegration

All content copied from https://docs.aws.amazon.com/.
