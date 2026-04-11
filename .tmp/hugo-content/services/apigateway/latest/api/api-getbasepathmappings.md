# GetBasePathMappings

Represents a collection of BasePathMapping resources.

## Request Syntax

```nohighlight

GET /domainnames/domain_name/basepathmappings?domainNameId=domainNameId&limit=limit&position=position HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[domain\_name](#API_GetBasePathMappings_RequestSyntax)**

The domain name of a BasePathMapping resource.

Required: Yes

**[domainNameId](#API_GetBasePathMappings_RequestSyntax)**

The identifier for the domain name resource. Supported only for private custom domain names.

**[limit](#API_GetBasePathMappings_RequestSyntax)**

The maximum number of returned results per page. The default value is 25 and the maximum value is 500.

**[position](#API_GetBasePathMappings_RequestSyntax)**

The current pagination position in the paged result set.

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "item": [
      {
         "basePath": "string",
         "restApiId": "string",
         "stage": "string"
      }
   ],
   "position": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[item](#API_GetBasePathMappings_ResponseSyntax)**

The current page of elements from this collection.

Type: Array of [BasePathMapping](api-basepathmapping.md) objects

**[position](#API_GetBasePathMappings_ResponseSyntax)**

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

### Get the base path mappings of a custom domain name

This example illustrates one usage of GetBasePathMappings.

#### Sample Request

```

GET /domainnames/a.b.c.com/basepathmappings HTTP/1.1
Content-Type: application/json
Host: apigateway.us-east-1.amazonaws.com
X-Amz-Date: 20160615T221921Z
Authorization: AWS4-HMAC-SHA256 Credential={access_key_ID}/20160615/us-east-1/apigateway/aws4_request, SignedHeaders=content-type;host;x-amz-date, Signature={sig4_hash}
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
      "href": "/domainnames/a.b.c.com/basepathmappings{?limit}",
      "templated": true
    },
    "basepathmapping:by-path": {
      "href": "/domainnames/a.b.c.com/basepathmappings/{base_path}",
      "templated": true
    },
    "basepathmapping:create": {
      "href": "/domainnames/a.b.c.com/basepathmappings"
    },
    "item": [
      {
        "href": "/domainnames/a.b.c.com/basepathmappings/my-service-api"
      },
      {
        "href": "/domainnames/a.b.c.com/basepathmappings/my-sample-api"
      }
    ]
  },
  "_embedded": {
    "item": [
      {
        "_links": {
          "self": {
            "href": "/domainnames/a.b.c.com/basepathmappings/my-service-api"
          },
          "basepathmapping:create": {
            "href": "/domainnames/a.b.c.com/basepathmappings"
          },
          "basepathmapping:delete": {
            "href": "/domainnames/a.b.c.com/basepathmappings/my-service-api"
          },
          "basepathmapping:update": {
            "href": "/domainnames/a.b.c.com/basepathmappings/my-service-api"
          }
        },
        "basepath": "my-service-api",
        "restApiId": "fugvjdxtri",
        "stage": "stage2"
      },
      {
        "_links": {
          "self": {
            "href": "/domainnames/a.b.c.com/basepathmappings/my-sample-api"
          },
          "basepathmapping:create": {
            "href": "/domainnames/a.b.c.com/basepathmappings"
          },
          "basepathmapping:delete": {
            "href": "/domainnames/a.b.c.com/basepathmappings/my-sample-api"
          },
          "basepathmapping:update": {
            "href": "/domainnames/a.b.c.com/basepathmappings/my-sample-api"
          }
        },
        "basepath": "my-sample-api",
        "restApiId": "fugvjdxtri",
        "stage": "stage1"
      }
    ]
  }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apigateway-2015-07-09/getbasepathmappings.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigateway-2015-07-09/getbasepathmappings.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/getbasepathmappings.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigateway-2015-07-09/getbasepathmappings.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/getbasepathmappings.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigateway-2015-07-09/getbasepathmappings.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigateway-2015-07-09/getbasepathmappings.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigateway-2015-07-09/getbasepathmappings.md)

- [AWS SDK for Python](../../../goto/boto3/apigateway-2015-07-09/getbasepathmappings.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/getbasepathmappings.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetBasePathMapping

GetClientCertificate

All content copied from https://docs.aws.amazon.com/.
