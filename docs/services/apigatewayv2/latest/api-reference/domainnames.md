# DomainNames

Represents a collection of domain names. See [Set up Custom Domain Name for an API in API Gateway](../../../apigateway/latest/developerguide/how-to-custom-domains.md).

## URI

`/v2/domainnames`

## HTTP methods

### GET

**Operation ID:** `GetDomainNames`

Gets the domain names for an AWS account.

Query parametersNameTypeRequiredDescription`nextToken`StringFalse

The next page of elements from this collection. Not valid for the last element of the collection.

`maxResults`StringFalse

The maximum number of elements to be returned for this resource.

ResponsesStatus codeResponse modelDescription`200``DomainNames`

Success

`400``BadRequestException`

One of the parameters in the request is invalid.

`404``NotFoundException`

The resource specified in the request was not found.

`429``LimitExceededException`

The client is sending more than the allowed number of requests per unit of time.

### POST

**Operation ID:** `CreateDomainName`

Creates a domain name.

ResponsesStatus codeResponse modelDescription`201``DomainName`

The request has succeeded and has resulted in the creation of a resource.

`400``BadRequestException`

One of the parameters in the request is invalid.

`403``AccessDeniedException`

403 response

`404``NotFoundException`

The resource specified in the request was not found.

`409``ConflictException`

The resource already exists.

`429``LimitExceededException`

The client is sending more than the allowed number of requests per unit of time.

## Schemas

### Request bodies

```json

{
  "domainName": "string",
  "tags": {
  },
  "domainNameConfigurations": [
    {
      "endpointType": enum,
      "ipAddressType": enum,
      "certificateName": "string",
      "certificateArn": "string",
      "ownershipVerificationCertificateArn": "string",
      "apiGatewayDomainName": "string",
      "hostedZoneId": "string",
      "certificateUploadDate": "string",
      "securityPolicy": enum,
      "domainNameStatus": enum,
      "domainNameStatusMessage": "string"
    }
  ],
  "mutualTlsAuthentication": {
    "truststoreUri": "string",
    "truststoreVersion": "string"
  },
  "routingMode": enum
}
```

### Response bodies

```json

{
  "items": [
    {
      "domainName": "string",
      "domainNameArn": "string",
      "domainNameConfigurations": [
        {
          "endpointType": enum,
          "ipAddressType": enum,
          "certificateName": "string",
          "certificateArn": "string",
          "ownershipVerificationCertificateArn": "string",
          "apiGatewayDomainName": "string",
          "hostedZoneId": "string",
          "certificateUploadDate": "string",
          "securityPolicy": enum,
          "domainNameStatus": enum,
          "domainNameStatusMessage": "string"
        }
      ],
      "apiMappingSelectionExpression": "string",
      "tags": {
      },
      "mutualTlsAuthentication": {
        "truststoreUri": "string",
        "truststoreVersion": "string",
        "truststoreWarnings": [
          "string"
        ]
      },
      "routingMode": enum
    }
  ],
  "nextToken": "string"
}
```

```json

{
  "domainName": "string",
  "domainNameArn": "string",
  "domainNameConfigurations": [
    {
      "endpointType": enum,
      "ipAddressType": enum,
      "certificateName": "string",
      "certificateArn": "string",
      "ownershipVerificationCertificateArn": "string",
      "apiGatewayDomainName": "string",
      "hostedZoneId": "string",
      "certificateUploadDate": "string",
      "securityPolicy": enum,
      "domainNameStatus": enum,
      "domainNameStatusMessage": "string"
    }
  ],
  "apiMappingSelectionExpression": "string",
  "tags": {
  },
  "mutualTlsAuthentication": {
    "truststoreUri": "string",
    "truststoreVersion": "string",
    "truststoreWarnings": [
      "string"
    ]
  },
  "routingMode": enum
}
```

```json

{
  "message": "string"
}
```

```json

{
  "message": "string"
}
```

```json

{
  "message": "string",
  "resourceType": "string"
}
```

```json

{
  "message": "string"
}
```

```json

{
  "message": "string",
  "limitType": "string"
}
```

## Properties

### AccessDeniedException

PropertyTypeRequiredDescription`message`

string

False

### BadRequestException

The request is not valid, for example, the input is incomplete or incorrect. See the accompanying error message for details.

PropertyTypeRequiredDescription`message`

string

False

Describes the error encountered.

### ConflictException

The requested operation would cause a conflict with the current state of a service resource associated with the request. Resolve the conflict before retrying this request. See the accompanying error message for details.

PropertyTypeRequiredDescription`message`

string

False

Describes the error encountered.

### CreateDomainNameInput

Represents the input parameters for a `CreateDomainName` request.

PropertyTypeRequiredDescription`domainName`

string

True

The domain name.

`domainNameConfigurations`

Array of type DomainNameConfiguration

False

The domain name configurations.

`mutualTlsAuthentication`

[MutualTlsAuthenticationInput](#domainnames-model-mutualtlsauthenticationinput)

False

The mutual TLS authentication configuration for a custom domain name.

`routingMode`

[RoutingMode](#domainnames-model-routingmode)

False

The routing mode API Gateway uses to route traffic to your APIs.

`tags`

[Tags](#domainnames-model-tags)

False

The collection of tags associated with a domain name.

### DomainName

Represents a domain name.

PropertyTypeRequiredDescription`apiMappingSelectionExpression`

string

False

The API mapping selection expression.

`domainName`

string

True

The name of the DomainName resource.

`domainNameArn`

string

False

The ARN of the DomainName resource.

`domainNameConfigurations`

Array of type DomainNameConfiguration

False

The domain name configurations.

`mutualTlsAuthentication`

[MutualTlsAuthentication](#domainnames-model-mutualtlsauthentication)

False

The mutual TLS authentication configuration for a custom domain name.

`routingMode`

[RoutingMode](#domainnames-model-routingmode)

False

The routing mode API Gateway uses to route traffic to your APIs.

`tags`

[Tags](#domainnames-model-tags)

False

The collection of tags associated with a domain name.

### DomainNameConfiguration

The domain name configuration.

PropertyTypeRequiredDescription`apiGatewayDomainName`

string

False

A domain name for the API.

`certificateArn`

string

False

An AWS-managed certificate that will be used by the edge-optimized endpoint for this domain name. AWS Certificate Manager is the only supported source.

`certificateName`

string

False

The user-friendly name of the certificate that will be used by the edge-optimized endpoint for this domain name.

`certificateUploadDate`

string

Format: date-time

False

The timestamp when the certificate that was used by edge-optimized endpoint for this domain name was uploaded.

`domainNameStatus`

[DomainNameStatus](#domainnames-model-domainnamestatus)

False

Identifies the status of a domain name migration. Statuses can have a value of AVAILABLE, UPDATING,
PENDING\_CERTIFICATE\_REIMPORT, or PENDING\_OWNERSHIP VERIFICATION. A domain can be modified if its status is
AVAILABLE. If the domain's status is UPDATING, you must wait until the current operation is complete.

`domainNameStatusMessage`

string

False

An optional text message containing detailed information about status of the domain name migration.

`endpointType`

[EndpointType](#domainnames-model-endpointtype)

False

The endpoint type.

`hostedZoneId`

string

False

The Amazon Route 53 Hosted Zone ID of the endpoint.

`ipAddressType`

[IpAddressType](#domainnames-model-ipaddresstype)

False

The IP address types that can invoke the domain name. Use `ipv4` to allow only IPv4 addresses to invoke your domain name, or use `dualstack` to allow both IPv4 and IPv6 addresses to invoke your domain name.

`ownershipVerificationCertificateArn`

string

False

The Amazon resource name (ARN) for the public certificate issued by AWS Certificate Manager. This ARN is used to validate custom domain ownership. It's required only if you configure mutual TLS and use either an ACM-imported or a private CA certificate ARN as the regionalCertificateArn.

`securityPolicy`

[SecurityPolicy](#domainnames-model-securitypolicy)

False

The Transport Layer Security (TLS) version of the security policy for this domain name. The valid values are `TLS_1_0` and `TLS_1_2`.

### DomainNameStatus

Identifies the status of a domain name migration. Statuses can have a value of AVAILABLE, UPDATING,
PENDING\_CERTIFICATE\_REIMPORT, or PENDING\_OWNERSHIP VERIFICATION. A domain can be modified if its status is
AVAILABLE. If the domain's status is UPDATING, you must wait until the current operation is
complete.

- `AVAILABLE`

- `UPDATING`

- `PENDING_CERTIFICATE_REIMPORT`

- `PENDING_OWNERSHIP_VERIFICATION`

### DomainNames

Represents a collection of domain names.

PropertyTypeRequiredDescription`items`

Array of type DomainName

False

The elements from this collection.

`nextToken`

string

False

The next page of elements from this collection. Not valid for the last element of the collection.

### EndpointType

Represents an endpoint type.

- `REGIONAL`

- `EDGE`

### IpAddressType

The IP address types that can invoke your API or domain name.

- `ipv4`

- `dualstack`

### LimitExceededException

A limit has been exceeded. See the accompanying error message for details.

PropertyTypeRequiredDescription`limitType`

string

False

The limit type.

`message`

string

False

Describes the error encountered.

### MutualTlsAuthentication

If specified, API Gateway performs two-way authentication between the client and the server. Clients must present a trusted certificate to access your API.

PropertyTypeRequiredDescription`truststoreUri`

string

False

An Amazon S3 URL that specifies the truststore for mutual TLS
authentication, for example, `s3://bucket-name/key-name
               `.
The truststore can contain certificates from public or private certificate
authorities. To update the truststore, upload a new version to S3, and then
update your custom domain name to use the new version. To update the truststore,
you must have permissions to access the S3 object.

`truststoreVersion`

string

False

The version of the S3 object that contains your truststore. To
specify a version, you must have versioning enabled for the S3 bucket.

`truststoreWarnings`

Array of type string

False

A list of warnings that API Gateway returns while processing your truststore.
Invalid certificates produce warnings. Mutual TLS is still enabled, but some
clients might not be able to access your API. To resolve warnings, upload a new
truststore to S3, and then update you domain name to use the new version.

### MutualTlsAuthenticationInput

If specified, API Gateway performs two-way authentication between the client and the server. Clients must present a trusted certificate to access your API.

PropertyTypeRequiredDescription`truststoreUri`

string

False

An Amazon S3 URL that specifies the truststore for mutual TLS
authentication, for example,
`s3://bucket-name/key-name
               `.
The truststore can contain certificates from public or private certificate
authorities. To update the truststore, upload a new version to S3, and then
update your custom domain name to use the new version. To update the truststore,
you must have permissions to access the S3 object.

`truststoreVersion`

string

False

The version of the S3 object that contains your truststore. To specify a version, you must have versioning enabled for the S3 bucket.

### NotFoundException

The resource specified in the request was not found. See the `message` field for more information.

PropertyTypeRequiredDescription`message`

string

False

Describes the error encountered.

`resourceType`

string

False

The resource type.

### RoutingMode

The routing mode API Gateway uses to route traffic to your APIs.

- `API_MAPPING_ONLY`

- `ROUTING_RULE_ONLY`

- `ROUTING_RULE_THEN_API_MAPPING`

### SecurityPolicy

The Transport Layer Security (TLS) version of the security policy for this domain name. The valid values are `TLS_1_0` and `TLS_1_2`.

- `TLS_1_0`

- `TLS_1_2`

### Tags

Represents a collection of tags associated with the resource.

PropertyTypeRequiredDescription

`*`

string

False

## See also

For more information about using this API in one of the language-specific AWS SDKs and references, see the following:

### GetDomainNames

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/apigatewayv2-2018-11-29/GetDomainNames)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/apigatewayv2-2018-11-29/GetDomainNames)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/apigatewayv2-2018-11-29/GetDomainNames)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/apigatewayv2-2018-11-29/GetDomainNames)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/apigatewayv2-2018-11-29/GetDomainNames)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/apigatewayv2-2018-11-29/GetDomainNames)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/apigatewayv2-2018-11-29/GetDomainNames)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/apigatewayv2-2018-11-29/GetDomainNames)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/apigatewayv2-2018-11-29/GetDomainNames)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/apigatewayv2-2018-11-29/GetDomainNames)

### CreateDomainName

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/apigatewayv2-2018-11-29/CreateDomainName)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/apigatewayv2-2018-11-29/CreateDomainName)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/apigatewayv2-2018-11-29/CreateDomainName)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/apigatewayv2-2018-11-29/CreateDomainName)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/apigatewayv2-2018-11-29/CreateDomainName)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/apigatewayv2-2018-11-29/CreateDomainName)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/apigatewayv2-2018-11-29/CreateDomainName)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/apigatewayv2-2018-11-29/CreateDomainName)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/apigatewayv2-2018-11-29/CreateDomainName)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/apigatewayv2-2018-11-29/CreateDomainName)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DomainName

ExportedAPI
