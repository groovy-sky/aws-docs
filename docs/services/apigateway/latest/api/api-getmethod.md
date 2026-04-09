# GetMethod

Describe an existing Method resource.

## Request Syntax

```nohighlight

GET /restapis/restapi_id/resources/resource_id/methods/http_method HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[http\_method](#API_GetMethod_RequestSyntax)**

Specifies the method request's HTTP method type.

Required: Yes

**[resource\_id](#API_GetMethod_RequestSyntax)**

The Resource identifier for the Method resource.

Required: Yes

**[restapi\_id](#API_GetMethod_RequestSyntax)**

The string identifier of the associated RestApi.

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
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
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[apiKeyRequired](#API_GetMethod_ResponseSyntax)**

A boolean flag specifying whether a valid ApiKey is required to invoke this method.

Type: Boolean

**[authorizationScopes](#API_GetMethod_ResponseSyntax)**

A list of authorization scopes configured on the method. The scopes are used with a `COGNITO_USER_POOLS` authorizer to authorize the method invocation. The authorization works by matching the method scopes against the scopes parsed from the access token in the incoming request. The method invocation is authorized if any method scopes matches a claimed scope in the access token. Otherwise, the invocation is not authorized. When the method scope is configured, the client must provide an access token instead of an identity token for authorization purposes.

Type: Array of strings

**[authorizationType](#API_GetMethod_ResponseSyntax)**

The method's authorization type. Valid values are `NONE` for open access, `AWS_IAM` for using AWS IAM permissions, `CUSTOM` for using a custom authorizer, or `COGNITO_USER_POOLS` for using a Cognito user pool.

Type: String

**[authorizerId](#API_GetMethod_ResponseSyntax)**

The identifier of an authorizer to use on this method. The method's authorization type must be `CUSTOM` or `COGNITO_USER_POOLS`.

Type: String

**[httpMethod](#API_GetMethod_ResponseSyntax)**

The method's HTTP verb.

Type: String

**[methodIntegration](#API_GetMethod_ResponseSyntax)**

Gets the method's integration responsible for passing the client-submitted request to the back end and performing necessary transformations to make the request compliant with the back end.

Type: [Integration](api-integration.md) object

**[methodResponses](#API_GetMethod_ResponseSyntax)**

Gets a method response associated with a given HTTP status code.

Type: String to [MethodResponse](api-methodresponse.md) object map

**[operationName](#API_GetMethod_ResponseSyntax)**

A human-friendly operation identifier for the method. For example, you can assign the `operationName` of `ListPets` for the `GET /pets` method in the `PetStore` example.

Type: String

**[requestModels](#API_GetMethod_ResponseSyntax)**

A key-value map specifying data schemas, represented by Model resources, (as the mapped value) of the request payloads of given content types (as the mapping key).

Type: String to string map

**[requestParameters](#API_GetMethod_ResponseSyntax)**

A key-value map defining required or optional method request parameters that can be accepted by API Gateway. A key is a method request parameter name matching the pattern of `method.request.{location}.{name}`, where `location` is `querystring`, `path`, or `header` and `name` is a valid and unique parameter name. The value associated with the key is a Boolean flag indicating whether the parameter is required ( `true`) or optional ( `false`). The method request parameter names defined here are available in Integration to be mapped to integration request parameters or templates.

Type: String to boolean map

**[requestValidatorId](#API_GetMethod_ResponseSyntax)**

The identifier of a RequestValidator for request validation.

Type: String

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

### View the detailed information about the GET method on an API resource

This example illustrates one usage of GetMethod.

#### Sample Request

```

GET /restapis/uojnr9hd57/resources/0cjtch/methods/GET HTTP/1.1
Content-Type: application/json
Host: apigateway.us-east-1.amazonaws.com
Content-Length: 117
X-Amz-Date: 20160613T205752Z
Authorization: AWS4-HMAC-SHA256 Credential={access_key_ID}/20160613/us-east-1/apigateway/aws4_request, SignedHeaders=content-type;host;x-amz-date, Signature={sig4_hash}

```

#### Sample Response

```

{
  "_links": {
    "curies": [
      {
        "href": "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-integration-{rel}.html",
        "name": "integration",
        "templated": true
      },
      {
        "href": "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-integration-response-{rel}.html",
        "name": "integrationresponse",
        "templated": true
      },
      {
        "href": "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-method-{rel}.html",
        "name": "method",
        "templated": true
      },
      {
        "href": "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-method-response-{rel}.html",
        "name": "methodresponse",
        "templated": true
      }
    ],
    "self": {
      "href": "/restapis/uojnr9hd57/resources/0cjtch/methods/GET",
      "name": "GET",
      "title": "GET"
    },
    "integration:put": {
      "href": "/restapis/uojnr9hd57/resources/0cjtch/methods/GET/integration"
    },
    "method:delete": {
      "href": "/restapis/uojnr9hd57/resources/0cjtch/methods/GET"
    },
    "method:integration": {
      "href": "/restapis/uojnr9hd57/resources/0cjtch/methods/GET/integration"
    },
    "method:responses": {
      "href": "/restapis/uojnr9hd57/resources/0cjtch/methods/GET/responses/200",
      "name": "200",
      "title": "200"
    },
    "method:update": {
      "href": "/restapis/uojnr9hd57/resources/0cjtch/methods/GET"
    },
    "methodresponse:put": {
      "href": "/restapis/uojnr9hd57/resources/0cjtch/methods/GET/responses/{status_code}",
      "templated": true
    }
  },
  "apiKeyRequired": false,
  "authorizationType": "NONE",
  "httpMethod": "GET",
  "requestParameters": {
    "method.request.querystring.operand2": false,
    "method.request.querystring.operator": false,
    "method.request.querystring.operand1": false
  },
  "_embedded": {
    "method:integration": {
      "_links": {
        "self": {
          "href": "/restapis/uojnr9hd57/resources/0cjtch/methods/GET/integration"
        },
        "integration:delete": {
          "href": "/restapis/uojnr9hd57/resources/0cjtch/methods/GET/integration"
        },
        "integration:responses": {
          "href": "/restapis/uojnr9hd57/resources/0cjtch/methods/GET/integration/responses/200",
          "name": "200",
          "title": "200"
        },
        "integration:update": {
          "href": "/restapis/uojnr9hd57/resources/0cjtch/methods/GET/integration"
        },
        "integrationresponse:put": {
          "href": "/restapis/uojnr9hd57/resources/0cjtch/methods/GET/integration/responses/{status_code}",
          "templated": true
        }
      },
      "cacheKeyParameters": [],
      "cacheNamespace": "0cjtch",
      "credentials": "arn:aws:iam::123456789012:role/apigAwsProxyRole",
      "httpMethod": "POST",
      "passthroughBehavior": "WHEN_NO_MATCH",
      "requestTemplates": {
        "application/json": "{\n    \"a\":  \"$input.params('operand1')\",\n    \"b\":  \"$input.params('operand2')\", \n    \"op\": \"$input.params('operator')\"   \n}"
      },
      "type": "AWS",
      "uri": "arn:aws:apigateway:us-west-2:lambda:path//2015-03-31/functions/arn:aws:lambda:us-west-2:123456789012:function:Calc/invocations",
      "_embedded": {
        "integration:responses": {
          "_links": {
            "self": {
              "href": "/restapis/uojnr9hd57/resources/0cjtch/methods/GET/integration/responses/200",
              "name": "200",
              "title": "200"
            },
            "integrationresponse:delete": {
              "href": "/restapis/uojnr9hd57/resources/0cjtch/methods/GET/integration/responses/200"
            },
            "integrationresponse:update": {
              "href": "/restapis/uojnr9hd57/resources/0cjtch/methods/GET/integration/responses/200"
            }
          },
          "responseParameters": {
            "method.response.header.operator": "integration.response.body.op",
            "method.response.header.operand_2": "integration.response.body.b",
            "method.response.header.operand_1": "integration.response.body.a"
          },
          "responseTemplates": {
            "application/json": "#set($res = $input.path('$'))\n{\n    \"result\": \"$res.a, $res.b, $res.op => $res.c\",\n  \"a\" : \"$res.a\",\n  \"b\" : \"$res.b\",\n  \"op\" : \"$res.op\",\n  \"c\" : \"$res.c\"\n}"
          },
          "selectionPattern": "",
          "statusCode": "200"
        }
      }
    },
    "method:responses": {
      "_links": {
        "self": {
          "href": "/restapis/uojnr9hd57/resources/0cjtch/methods/GET/responses/200",
          "name": "200",
          "title": "200"
        },
        "methodresponse:delete": {
          "href": "/restapis/uojnr9hd57/resources/0cjtch/methods/GET/responses/200"
        },
        "methodresponse:update": {
          "href": "/restapis/uojnr9hd57/resources/0cjtch/methods/GET/responses/200"
        }
      },
      "responseModels": {
        "application/json": "Empty"
      },
      "responseParameters": {
        "method.response.header.operator": false,
        "method.response.header.operand_2": false,
        "method.response.header.operand_1": false
      },
      "statusCode": "200"
    }
  }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apigateway-2015-07-09/getmethod.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigateway-2015-07-09/getmethod.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/getmethod.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigateway-2015-07-09/getmethod.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/getmethod.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigateway-2015-07-09/getmethod.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigateway-2015-07-09/getmethod.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigateway-2015-07-09/getmethod.md)

- [AWS SDK for Python](../../../goto/boto3/apigateway-2015-07-09/getmethod.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/getmethod.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetIntegrationResponse

GetMethodResponse

All content copied from https://docs.aws.amazon.com/.
