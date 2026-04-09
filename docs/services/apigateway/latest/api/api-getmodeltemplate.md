# GetModelTemplate

Generates a sample mapping template that can be used to transform a payload into the structure of a model.

## Request Syntax

```nohighlight

GET /restapis/restapi_id/models/model_name/default_template HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[model\_name](#API_GetModelTemplate_RequestSyntax)**

The name of the model for which to generate a template.

Required: Yes

**[restapi\_id](#API_GetModelTemplate_RequestSyntax)**

The string identifier of the associated RestApi.

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "value": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[value](#API_GetModelTemplate_ResponseSyntax)**

The Apache Velocity Template Language (VTL) template content used for the template resource.

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

### Generate the sample template from a model

This example illustrates one usage of GetModelTemplate.

#### Sample Request

```

GET /restapis/uojnr9hd57/models/output/default_template HTTP/1.1
Content-Type: application/json
Host: apigateway.us-east-1.amazonaws.com
X-Amz-Date: 20160614T202448Z
Authorization: AWS4-HMAC-SHA256 Credential={access_key_ID}/20160614/us-east-1/apigateway/aws4_request, SignedHeaders=content-type;host;x-amz-date, Signature={sig4_hash}
Response

```

#### Sample Response

```

{
  "_links": {
    "self": {
      "href": "/restapis/uojnr9hd57/models/output/default_template"
    }
  },
  "value": "#set($inputRoot = $input.path('$'))\n{\n  \"a\" : 3.1415,\n  \"b\" : 3.1415,\n  \"op\" : \"foo\",\n  \"c\" : 3.1415\n}"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apigateway-2015-07-09/getmodeltemplate.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigateway-2015-07-09/getmodeltemplate.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/getmodeltemplate.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigateway-2015-07-09/getmodeltemplate.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/getmodeltemplate.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigateway-2015-07-09/getmodeltemplate.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigateway-2015-07-09/getmodeltemplate.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigateway-2015-07-09/getmodeltemplate.md)

- [AWS SDK for Python](../../../goto/boto3/apigateway-2015-07-09/getmodeltemplate.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/getmodeltemplate.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetModels

GetRequestValidator

All content copied from https://docs.aws.amazon.com/.
