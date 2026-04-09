# CreateBasePathMapping

Creates a new BasePathMapping resource.

## Request Syntax

```nohighlight

POST /domainnames/domain_name/basepathmappings?domainNameId=domainNameId HTTP/1.1
Content-type: application/json

{
   "basePath": "string",
   "restApiId": "string",
   "stage": "string"
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[domain\_name](#API_CreateBasePathMapping_RequestSyntax)**

The domain name of the BasePathMapping resource to create.

Required: Yes

**[domainNameId](#API_CreateBasePathMapping_RequestSyntax)**

The identifier for the domain name resource. Required for private custom domain names.

## Request Body

The request accepts the following data in JSON format.

**[basePath](#API_CreateBasePathMapping_RequestSyntax)**

The base path name that callers of the API must provide as part of the URL after the domain name. This value must be unique for all of the mappings across a single API. Specify '(none)' if you do not want callers to specify a base path name after the domain name.

Type: String

Required: No

**[restApiId](#API_CreateBasePathMapping_RequestSyntax)**

The string identifier of the associated RestApi.

Type: String

Required: Yes

**[stage](#API_CreateBasePathMapping_RequestSyntax)**

The name of the API's stage that you want to use for this mapping. Specify '(none)' if you want callers to explicitly specify the stage name after any base path name.

Type: String

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 201
Content-type: application/json

{
   "basePath": "string",
   "restApiId": "string",
   "stage": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 201 response.

The following data is returned in JSON format by the service.

**[basePath](#API_CreateBasePathMapping_ResponseSyntax)**

The base path name that callers of the API must provide as part of the URL after the domain name.

Type: String

**[restApiId](#API_CreateBasePathMapping_ResponseSyntax)**

The string identifier of the associated RestApi.

Type: String

**[stage](#API_CreateBasePathMapping_ResponseSyntax)**

The name of the associated stage.

Type: String

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

### Creates base path mapping for an API

The following example request creates a `TestApi` base path that is mapped the
`fugvjdxtri` API in the `stage1` stage.

#### Sample Request

```

POST /domainnames/a.b.c.com/basepathmappings HTTP/1.1
Content-Type: application/json
Host: apigateway.us-east-1.amazonaws.com
X-Amz-Date: 20160615T012033Z
Authorization: AWS4-HMAC-SHA256 Credential={access_key_ID}/20160615/us-east-1/apigateway/aws4_request, SignedHeaders=content-length;content-type;host;x-amz-date, Signature={sig4_hash}

{
  "basepath" : "TestApi",
  "restApiId" : "fugvjdxtri",
  "stage" : "stage1"
}
```

#### Sample Response

```

{
  "_links": {upd
    "curies": {
      "href": "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-basepathmapping-{rel}.html",
      "name": "basepathmapping",
      "templated": true
    },
    "self": {
      "href": "/domainnames/a.b.c.com/basepathmappings/TestApi"
    },
    "basepathmapping:create": {
      "href": "/domainnames/a.b.c.com/basepathmappings"
    },
    "basepathmapping:delete": {
      "href": "/domainnames/a.b.c.com/basepathmappings/TestApi"
    },
    "basepathmapping:update": {
      "href": "/domainnames/a.b.c.com/basepathmappings/TestApi"
    }
  },
  "basepath": "TestApi",
  "restApiId": "fugvjdxtri",
  "stage": "stage1"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apigateway-2015-07-09/createbasepathmapping.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigateway-2015-07-09/createbasepathmapping.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/createbasepathmapping.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigateway-2015-07-09/createbasepathmapping.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/createbasepathmapping.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigateway-2015-07-09/createbasepathmapping.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigateway-2015-07-09/createbasepathmapping.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigateway-2015-07-09/createbasepathmapping.md)

- [AWS SDK for Python](../../../goto/boto3/apigateway-2015-07-09/createbasepathmapping.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/createbasepathmapping.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateAuthorizer

CreateDeployment

All content copied from https://docs.aws.amazon.com/.
