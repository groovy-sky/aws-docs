# GenerateClientCertificate

Generates a ClientCertificate resource.

## Request Syntax

```nohighlight

POST /clientcertificates HTTP/1.1
Content-type: application/json

{
   "description": "string",
   "tags": {
      "string" : "string"
   }
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[description](#API_GenerateClientCertificate_RequestSyntax)**

The description of the ClientCertificate.

Type: String

Required: No

**[tags](#API_GenerateClientCertificate_RequestSyntax)**

The key-value map of strings. The valid character set is \[a-zA-Z+-=.\_:/\]. The tag key can be up to 128 characters and must not start with `aws:`. The tag value can be up to 256 characters.

Type: String to string map

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 201
Content-type: application/json

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
```

## Response Elements

If the action is successful, the service sends back an HTTP 201 response.

The following data is returned in JSON format by the service.

**[clientCertificateId](#API_GenerateClientCertificate_ResponseSyntax)**

The identifier of the client certificate.

Type: String

**[createdDate](#API_GenerateClientCertificate_ResponseSyntax)**

The timestamp when the client certificate was created.

Type: Timestamp

**[description](#API_GenerateClientCertificate_ResponseSyntax)**

The description of the client certificate.

Type: String

**[expirationDate](#API_GenerateClientCertificate_ResponseSyntax)**

The timestamp when the client certificate will expire.

Type: Timestamp

**[pemEncodedCertificate](#API_GenerateClientCertificate_ResponseSyntax)**

The PEM-encoded public key of the client certificate, which can be used to configure certificate authentication in the integration endpoint .

Type: String

**[tags](#API_GenerateClientCertificate_ResponseSyntax)**

The collection of tags. Each tag element is associated with a given resource.

Type: String to string map

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

**TooManyRequestsException**

The request has reached its throttling limit. Retry after the specified time period.

HTTP Status Code: 429

**UnauthorizedException**

The request is denied because the caller has insufficient permissions.

HTTP Status Code: 401

## Examples

### Generate a client-side certificate

This example illustrates one usage of GenerateClientCertificate.

#### Sample Request

```

POST /clientcertificates HTTP/1.1
Content-Type: application/json
Host: apigateway.us-east-1.amazonaws.com
X-Amz-Date: 20160615T223313Z
Authorization: AWS4-HMAC-SHA256 Credential={access_key_ID}/20160615/us-east-1/apigateway/aws4_request, SignedHeaders=content-length;content-type;host;x-amz-date, Signature={sig4_hash}

{
    "description": "my-second-client-cert"
}
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
      "href": "/clientcertificates/9ao60f"
    },
    "clientcertificate:delete": {
      "href": "/clientcertificates/9ao60f"
    },
    "clientcertificate:update": {
      "href": "/clientcertificates/9ao60f"
    }
  },
  "clientCertificateId": "9ao60f",
  "createdDate": "2016-06-15T22:33:13Z",
  "description": "my-second-client-cert",
  "expirationDate": "2017-06-15T22:33:13Z",
  "pemEncodedCertificate": "-----BEGIN CERTIFICATE-----\r\nMIIC6T...2yQAGEHvs=\r\n-----END CERTIFICATE-----"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/apigateway-2015-07-09/GenerateClientCertificate)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/apigateway-2015-07-09/GenerateClientCertificate)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/apigateway-2015-07-09/GenerateClientCertificate)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/apigateway-2015-07-09/GenerateClientCertificate)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/apigateway-2015-07-09/GenerateClientCertificate)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/apigateway-2015-07-09/GenerateClientCertificate)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/apigateway-2015-07-09/GenerateClientCertificate)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/apigateway-2015-07-09/GenerateClientCertificate)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/apigateway-2015-07-09/GenerateClientCertificate)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/apigateway-2015-07-09/GenerateClientCertificate)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

FlushStageCache

GetAccount
