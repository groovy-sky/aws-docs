# GetRequestValidator

Gets a RequestValidator of a given RestApi.

## Request Syntax

```nohighlight

GET /restapis/restapi_id/requestvalidators/requestvalidator_id HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[requestvalidator\_id](#API_GetRequestValidator_RequestSyntax)**

The identifier of the RequestValidator to be retrieved.

Required: Yes

**[restapi\_id](#API_GetRequestValidator_RequestSyntax)**

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
   "name": "string",
   "validateRequestBody": boolean,
   "validateRequestParameters": boolean
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[id](#API_GetRequestValidator_ResponseSyntax)**

The identifier of this RequestValidator.

Type: String

**[name](#API_GetRequestValidator_ResponseSyntax)**

The name of this RequestValidator

Type: String

**[validateRequestBody](#API_GetRequestValidator_ResponseSyntax)**

A Boolean flag to indicate whether to validate a request body according to the configured Model schema.

Type: Boolean

**[validateRequestParameters](#API_GetRequestValidator_ResponseSyntax)**

A Boolean flag to indicate whether to validate request parameters ( `true`) or not ( `false`).

Type: Boolean

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

### Get the specified RequestValidator of an API

This example illustrates one usage of GetRequestValidator.

#### Sample Request

```

GET /restapis/mkhqppt4e4/requestvalidators/1t3hul HTTP/1.1
Content-Type: application/json
Host: apigateway.us-east-1.amazonaws.com
X-Amz-Date: 20170223T172652Z
Authorization: AWS4-HMAC-SHA256 Credential={access_key_ID}/20170223/us-east-1/apigateway/aws4_request, SignedHeaders=content-type;host;x-amz-date, Signature={sig4_hash}
```

#### Sample Response

```

{
  "_links": {
    "self": {
      "href": "/restapis/mkhqppt4e4/requestvalidators/1t3hul"
    },
    "request-validator-delete": {
      "href": "/restapis/mkhqppt4e4/requestvalidators/1t3hul"
    },
    "request-validator-update": {
      "href": "/restapis/mkhqppt4e4/requestvalidators/1t3hul"
    }
  },
  "id": "1t3hul",
  "name": "params-only",
  "validateRequestBody": false,
  "validateRequestParameters": true
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apigateway-2015-07-09/getrequestvalidator.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigateway-2015-07-09/getrequestvalidator.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/getrequestvalidator.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigateway-2015-07-09/getrequestvalidator.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/getrequestvalidator.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigateway-2015-07-09/getrequestvalidator.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigateway-2015-07-09/getrequestvalidator.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigateway-2015-07-09/getrequestvalidator.md)

- [AWS SDK for Python](../../../goto/boto3/apigateway-2015-07-09/getrequestvalidator.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/getrequestvalidator.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetModelTemplate

GetRequestValidators

All content copied from https://docs.aws.amazon.com/.
