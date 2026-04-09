# UpdateClientCertificate

Changes information about an ClientCertificate resource.

## Request Syntax

```nohighlight

PATCH /clientcertificates/clientcertificate_id HTTP/1.1
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

**[clientcertificate\_id](#API_UpdateClientCertificate_RequestSyntax)**

The identifier of the ClientCertificate resource to be updated.

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[patchOperations](#API_UpdateClientCertificate_RequestSyntax)**

For more information about supported patch operations, see [Patch Operations](patch-operations.md).

Type: Array of [PatchOperation](api-patchoperation.md) objects

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 200
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

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[clientCertificateId](#API_UpdateClientCertificate_ResponseSyntax)**

The identifier of the client certificate.

Type: String

**[createdDate](#API_UpdateClientCertificate_ResponseSyntax)**

The timestamp when the client certificate was created.

Type: Timestamp

**[description](#API_UpdateClientCertificate_ResponseSyntax)**

The description of the client certificate.

Type: String

**[expirationDate](#API_UpdateClientCertificate_ResponseSyntax)**

The timestamp when the client certificate will expire.

Type: Timestamp

**[pemEncodedCertificate](#API_UpdateClientCertificate_ResponseSyntax)**

The PEM-encoded public key of the client certificate, which can be used to configure certificate authentication in the integration endpoint .

Type: String

**[tags](#API_UpdateClientCertificate_ResponseSyntax)**

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

### Update the description of a client-side certificate

This example illustrates one usage of UpdateClientCertificate.

#### Sample Request

```

PATCH /clientcertificates/9ao60f HTTP/1.1
Content-Type: application/json
Host: apigateway.us-east-1.amazonaws.com
X-Amz-Date: 20160615T225025Z
Authorization: AWS4-HMAC-SHA256 Credential={access_key_ID}/20160615/us-east-1/apigateway/aws4_request, SignedHeaders=content-length;content-type;host;x-amz-date, Signature={sig4_hash}

{
  "patchOperations" : [ {
    "op" : "replace",
    "path" : "/description",
    "value" : "my second client-side cert"
  } ]
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
  "description": "my second client-side cert",
  "expirationDate": "2017-06-15T22:33:13Z",
  "pemEncodedCertificate": "-----BEGIN CERTIFICATE-----\r\nMIIC6TC...yQAGEHvs=\r\n-----END CERTIFICATE-----"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apigateway-2015-07-09/updateclientcertificate.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigateway-2015-07-09/updateclientcertificate.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/updateclientcertificate.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigateway-2015-07-09/updateclientcertificate.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/updateclientcertificate.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigateway-2015-07-09/updateclientcertificate.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigateway-2015-07-09/updateclientcertificate.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigateway-2015-07-09/updateclientcertificate.md)

- [AWS SDK for Python](../../../goto/boto3/apigateway-2015-07-09/updateclientcertificate.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/updateclientcertificate.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UpdateBasePathMapping

UpdateDeployment

All content copied from https://docs.aws.amazon.com/.
