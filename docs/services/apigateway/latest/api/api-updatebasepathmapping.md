# UpdateBasePathMapping

Changes information about the BasePathMapping resource.

## Request Syntax

```nohighlight

PATCH /domainnames/domain_name/basepathmappings/base_path?domainNameId=domainNameId HTTP/1.1
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

**[base\_path](#API_UpdateBasePathMapping_RequestSyntax)**

The base path of the BasePathMapping resource to change.

To specify an empty base path, set this parameter to `'(none)'`.

Required: Yes

**[domain\_name](#API_UpdateBasePathMapping_RequestSyntax)**

The domain name of the BasePathMapping resource to change.

Required: Yes

**[domainNameId](#API_UpdateBasePathMapping_RequestSyntax)**

The identifier for the domain name resource. Supported only for private custom domain names.

## Request Body

The request accepts the following data in JSON format.

**[patchOperations](#API_UpdateBasePathMapping_RequestSyntax)**

For more information about supported patch operations, see [Patch Operations](patch-operations.md).

Type: Array of [PatchOperation](api-patchoperation.md) objects

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "basePath": "string",
   "restApiId": "string",
   "stage": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[basePath](#API_UpdateBasePathMapping_ResponseSyntax)**

The base path name that callers of the API must provide as part of the URL after the domain name.

Type: String

**[restApiId](#API_UpdateBasePathMapping_ResponseSyntax)**

The string identifier of the associated RestApi.

Type: String

**[stage](#API_UpdateBasePathMapping_ResponseSyntax)**

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

### Update the base path mapping of an API

The following example request updates the base path ( `TestApi`) of a custom
domain name ( `a.b.c.com`) to map to a different deployment stage ( `stage2 ` of an API
( `fugvjdxtri`).

#### Sample Request

```

PATCH /domainnames/a.b.c.com/basepathmappings/TestApi HTTP/1.1
Content-Type: application/json
Host: apigateway.us-east-1.amazonaws.com
X-Amz-Date: 20160615T025216Z
Authorization: AWS4-HMAC-SHA256 Credential={access_key_ID}/20160615/us-east-1/apigateway/aws4_request, SignedHeaders=content-length;content-type;host;x-amz-date, Signature={sig4_hash}

{
  "patchOperations" : [ {
    "op" : "replace",
    "path" : "/stage",
    "value" : "stage2"
  } ]
}
```

#### Sample Response

```

{
  "_links": {
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
  "stage": "stage2"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apigateway-2015-07-09/updatebasepathmapping.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigateway-2015-07-09/updatebasepathmapping.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/updatebasepathmapping.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigateway-2015-07-09/updatebasepathmapping.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/updatebasepathmapping.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigateway-2015-07-09/updatebasepathmapping.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigateway-2015-07-09/updatebasepathmapping.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigateway-2015-07-09/updatebasepathmapping.md)

- [AWS SDK for Python](../../../goto/boto3/apigateway-2015-07-09/updatebasepathmapping.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/updatebasepathmapping.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UpdateAuthorizer

UpdateClientCertificate

All content copied from https://docs.aws.amazon.com/.
