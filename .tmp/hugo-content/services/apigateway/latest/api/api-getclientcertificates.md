# GetClientCertificates

Gets a collection of ClientCertificate resources.

## Request Syntax

```nohighlight

GET /clientcertificates?limit=limit&position=position HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[limit](#API_GetClientCertificates_RequestSyntax)**

The maximum number of returned results per page. The default value is 25 and the maximum value is 500.

**[position](#API_GetClientCertificates_RequestSyntax)**

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
         "clientCertificateId": "string",
         "createdDate": number,
         "description": "string",
         "expirationDate": number,
         "pemEncodedCertificate": "string",
         "tags": {
            "string" : "string"
         }
      }
   ],
   "position": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[item](#API_GetClientCertificates_ResponseSyntax)**

The current page of elements from this collection.

Type: Array of [ClientCertificate](api-clientcertificate.md) objects

**[position](#API_GetClientCertificates_ResponseSyntax)**

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

### Retrieve client certificates

The following example request retrieves the available client certificates in the caller's
AWS account.

A successful response returns the requested `ClientCertificate` resources that can be navigated to by following the linked item or examining the embedded item resource.

#### Sample Request

```

GET /clientcertificates HTTP/1.1
Content-Type: application/json
Host: apigateway.us-east-1.amazonaws.com
X-Amz-Date: 20160601T175605Z
Authorization: AWS4-HMAC-SHA256 Credential={access_key_ID}/us-east-1/apigateway/aws4_request, SignedHeaders=content-type;host;x-amz-date, Signature={sig4_hash}
```

#### Sample Response

```

{
  "_links": {
    "curies": {
      "href": "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-clientcertificate-{rel}.html",
      "name": "clientcertificate",
      "templated": true
    },
    "self": {
      "href": "/clientcertificates"
    },
    "clientcertificate:by-id": {
      "href": "/clientcertificates/{clientcertificate_id}",
      "templated": true
    },
    "clientcertificate:generate": {
      "href": "/clientcertificates"
    },
    "item": {
      "href": "/clientcertificates/xmbiqp"
    }
  },
  "_embedded": {
    "item": {
      "_links": {
        "self": {
          "href": "/clientcertificates/xmbiqp"
        },
        "clientcertificate:delete": {
          "href": "/clientcertificates/xmbiqp"
        },
        "clientcertificate:update": {
          "href": "/clientcertificates/xmbiqp"
        }
      },
      "clientCertificateId": "xmbiqp",get
      "createdDate": "2015-12-08T18:02:16Z",
      "description": "test-client-cert-2",
      "expirationDate": "2016-12-07T18:02:16Z",
      "pemEncodedCertificate": "-----BEGIN CERTIFICATE-----\r\nMIIC6DC...XuHVdZ5r27XRRXEjg==\r\n-----END CERTIFICATE-----"
    }
  }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apigateway-2015-07-09/getclientcertificates.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigateway-2015-07-09/getclientcertificates.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/getclientcertificates.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigateway-2015-07-09/getclientcertificates.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/getclientcertificates.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigateway-2015-07-09/getclientcertificates.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigateway-2015-07-09/getclientcertificates.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigateway-2015-07-09/getclientcertificates.md)

- [AWS SDK for Python](../../../goto/boto3/apigateway-2015-07-09/getclientcertificates.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/getclientcertificates.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetClientCertificate

GetDeployment

All content copied from https://docs.aws.amazon.com/.
