# CreateRequestValidator

Creates a RequestValidator of a given RestApi.

## Request Syntax

```nohighlight

POST /restapis/restapi_id/requestvalidators HTTP/1.1
Content-type: application/json

{
   "name": "string",
   "validateRequestBody": boolean,
   "validateRequestParameters": boolean
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[restapi\_id](#API_CreateRequestValidator_RequestSyntax)**

The string identifier of the associated RestApi.

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[name](#API_CreateRequestValidator_RequestSyntax)**

The name of the to-be-created RequestValidator.

Type: String

Required: No

**[validateRequestBody](#API_CreateRequestValidator_RequestSyntax)**

A Boolean flag to indicate whether to validate request body according to the configured model schema for the method ( `true`) or not ( `false`).

Type: Boolean

Required: No

**[validateRequestParameters](#API_CreateRequestValidator_RequestSyntax)**

A Boolean flag to indicate whether to validate request parameters, `true`, or not `false`.

Type: Boolean

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 201
Content-type: application/json

{
   "id": "string",
   "name": "string",
   "validateRequestBody": boolean,
   "validateRequestParameters": boolean
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 201 response.

The following data is returned in JSON format by the service.

**[id](#API_CreateRequestValidator_ResponseSyntax)**

The identifier of this RequestValidator.

Type: String

**[name](#API_CreateRequestValidator_ResponseSyntax)**

The name of this RequestValidator

Type: String

**[validateRequestBody](#API_CreateRequestValidator_ResponseSyntax)**

A Boolean flag to indicate whether to validate a request body according to the configured Model schema.

Type: Boolean

**[validateRequestParameters](#API_CreateRequestValidator_ResponseSyntax)**

A Boolean flag to indicate whether to validate request parameters ( `true`) or not ( `false`).

Type: Boolean

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**BadRequestException**

The submitted request is not valid, for example, the input is incomplete or incorrect. See the accompanying error message for details.

HTTP Status Code: 400

**ConflictException**

The request configuration has conflicts. For details, see the accompanying error message.

HTTP Status Code: 409

**LimitExceededException**

The request exceeded the rate limit. Retry after the specified time period.

HTTP Status Code: 429

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

### Create a RequestValidator for an API to validate required request payloads

This example illustrates one usage of CreateRequestValidator.

#### Sample Request

```

POST /restapis/mkhqppt4e4/requestvalidators HTTP/1.1
Content-Type: application/json
Host: apigateway.us-east-1.amazonaws.com
X-Amz-Date: 20170223T172652Z
Authorization: AWS4-HMAC-SHA256 Credential={access_key_ID}/20170223/us-east-1/apigateway/aws4_request, SignedHeaders=content-type;host;x-amz-date, Signature={sig4_hash}

{
  "name" : "body-only",
  "validateRequestBody" : "true",
  "validateRequestParameters" : "false"
}
```

#### Sample Response

```

{
  "_links": {
    "self": {
      "href": "/restapis/mkhqppt4e4/requestvalidators/3n5aa0"
    },
    "request-validator-delete": {
      "href": "/restapis/mkhqppt4e4/requestvalidators/3n5aa0"
    },
    "request-validator-update": {
      "href": "/restapis/mkhqppt4e4/requestvalidators/3n5aa0"
    }
  },
  "id": "3n5aa0",
  "name": "body-only",
  "validateRequestBody": true,
  "validateRequestParameters": false
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apigateway-2015-07-09/createrequestvalidator.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigateway-2015-07-09/createrequestvalidator.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/createrequestvalidator.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigateway-2015-07-09/createrequestvalidator.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/createrequestvalidator.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigateway-2015-07-09/createrequestvalidator.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigateway-2015-07-09/createrequestvalidator.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigateway-2015-07-09/createrequestvalidator.md)

- [AWS SDK for Python](../../../goto/boto3/apigateway-2015-07-09/createrequestvalidator.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/createrequestvalidator.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateModel

CreateResource

All content copied from https://docs.aws.amazon.com/.
