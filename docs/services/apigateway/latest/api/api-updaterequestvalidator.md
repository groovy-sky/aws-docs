# UpdateRequestValidator

Updates a RequestValidator of a given RestApi.

## Request Syntax

```nohighlight

PATCH /restapis/restapi_id/requestvalidators/requestvalidator_id HTTP/1.1
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

**[requestvalidator\_id](#API_UpdateRequestValidator_RequestSyntax)**

The identifier of RequestValidator to be updated.

Required: Yes

**[restapi\_id](#API_UpdateRequestValidator_RequestSyntax)**

The string identifier of the associated RestApi.

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[patchOperations](#API_UpdateRequestValidator_RequestSyntax)**

For more information about supported patch operations, see [Patch Operations](patch-operations.md).

Type: Array of [PatchOperation](api-patchoperation.md) objects

Required: No

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

**[id](#API_UpdateRequestValidator_ResponseSyntax)**

The identifier of this RequestValidator.

Type: String

**[name](#API_UpdateRequestValidator_ResponseSyntax)**

The name of this RequestValidator

Type: String

**[validateRequestBody](#API_UpdateRequestValidator_ResponseSyntax)**

A Boolean flag to indicate whether to validate a request body according to the configured Model schema.

Type: Boolean

**[validateRequestParameters](#API_UpdateRequestValidator_ResponseSyntax)**

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

### Update a RequestValidator for an API

This example illustrates one usage of UpdateRequestValidator.

#### Sample Request

```

PATCH /restapis/mkhqppt4e4/requestvalidators/3n5aa0 HTTP/1.1
Content-Type: application/json
Host: apigateway.us-east-1.amazonaws.com
X-Amz-Date: 20170223T172652Z
Authorization: AWS4-HMAC-SHA256 Credential={access_key_ID}/20170223/us-east-1/apigateway/aws4_request, SignedHeaders=content-type;host;x-amz-date, Signature={sig4_hash}

{
  "patchOperations" : [ {
    "op" : "replace",
    "path" : "/name",
    "value" : "body-parameters-switched"
  }, {
    "op" : "replace",
    "path" : "/validateRequestBody",
    "value" : "false"
  }, {
    "op" : "replace",
    "path" : "/validateRequestParameters",
    "value" : "true"
  } ]
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
  "name": "body-parameters-switched",
  "validateRequestBody": false,
  "validateRequestParameters": true
}

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apigateway-2015-07-09/updaterequestvalidator.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigateway-2015-07-09/updaterequestvalidator.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/updaterequestvalidator.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigateway-2015-07-09/updaterequestvalidator.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/updaterequestvalidator.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigateway-2015-07-09/updaterequestvalidator.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigateway-2015-07-09/updaterequestvalidator.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigateway-2015-07-09/updaterequestvalidator.md)

- [AWS SDK for Python](../../../goto/boto3/apigateway-2015-07-09/updaterequestvalidator.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/updaterequestvalidator.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UpdateModel

UpdateResource

All content copied from https://docs.aws.amazon.com/.
