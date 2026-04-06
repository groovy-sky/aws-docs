# CreateDomainName

Creates a new domain name.

## Request Syntax

```nohighlight

POST /domainnames HTTP/1.1
Content-type: application/json

{
   "certificateArn": "string",
   "certificateBody": "string",
   "certificateChain": "string",
   "certificateName": "string",
   "certificatePrivateKey": "string",
   "domainName": "string",
   "endpointAccessMode": "string",
   "endpointConfiguration": {
      "ipAddressType": "string",
      "types": [ "string" ],
      "vpcEndpointIds": [ "string" ]
   },
   "mutualTlsAuthentication": {
      "truststoreUri": "string",
      "truststoreVersion": "string"
   },
   "ownershipVerificationCertificateArn": "string",
   "policy": "string",
   "regionalCertificateArn": "string",
   "regionalCertificateName": "string",
   "routingMode": "string",
   "securityPolicy": "string",
   "tags": {
      "string" : "string"
   }
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[certificateArn](#API_CreateDomainName_RequestSyntax)**

The reference to an AWS-managed certificate that will be used by edge-optimized endpoint or private endpoint for this domain name. AWS Certificate Manager is the only supported source.

Type: String

Required: No

**[certificateBody](#API_CreateDomainName_RequestSyntax)**

\[Deprecated\] The body of the server certificate that will be used by edge-optimized endpoint or private endpoint for this domain name provided by your certificate authority.

Type: String

Required: No

**[certificateChain](#API_CreateDomainName_RequestSyntax)**

\[Deprecated\] The intermediate certificates and optionally the root certificate, one after the other without any blank lines, used by an edge-optimized endpoint for this domain name. If you include the root certificate, your certificate chain must start with intermediate certificates and end with the root certificate. Use the intermediate certificates that were provided by your certificate authority. Do not include any intermediaries that are not in the chain of trust path.

Type: String

Required: No

**[certificateName](#API_CreateDomainName_RequestSyntax)**

The user-friendly name of the certificate that will be used by edge-optimized endpoint or private endpoint for this domain name.

Type: String

Required: No

**[certificatePrivateKey](#API_CreateDomainName_RequestSyntax)**

\[Deprecated\] Your edge-optimized endpoint's domain name certificate's private key.

Type: String

Required: No

**[domainName](#API_CreateDomainName_RequestSyntax)**

The name of the DomainName resource.

Type: String

Required: Yes

**[endpointAccessMode](#API_CreateDomainName_RequestSyntax)**

The endpoint access mode of the DomainName. Only available for DomainNames that use security policies that start with `SecurityPolicy_`.

Type: String

Valid Values: `BASIC | STRICT`

Required: No

**[endpointConfiguration](#API_CreateDomainName_RequestSyntax)**

The endpoint configuration of this DomainName showing the endpoint types and IP address types of the domain name.

Type: [EndpointConfiguration](https://docs.aws.amazon.com/apigateway/latest/api/API_EndpointConfiguration.html) object

Required: No

**[mutualTlsAuthentication](#API_CreateDomainName_RequestSyntax)**

The mutual TLS authentication configuration for a custom domain name. If specified, API Gateway
performs two-way authentication between the client and the server. Clients must present a
trusted certificate to access your API.

Type: [MutualTlsAuthenticationInput](https://docs.aws.amazon.com/apigateway/latest/api/API_MutualTlsAuthenticationInput.html) object

Required: No

**[ownershipVerificationCertificateArn](#API_CreateDomainName_RequestSyntax)**

The ARN of the public certificate issued by ACM to validate ownership of your custom
domain. Only required when configuring mutual TLS and using an ACM imported or private CA
certificate ARN as the regionalCertificateArn.

Type: String

Required: No

**[policy](#API_CreateDomainName_RequestSyntax)**

A stringified JSON policy document that applies to the `execute-api` service for this DomainName regardless of the caller and Method
configuration. Supported only for private custom
domain names.

Type: String

Required: No

**[regionalCertificateArn](#API_CreateDomainName_RequestSyntax)**

The reference to an AWS-managed certificate that will be used by regional endpoint for this domain name. AWS Certificate Manager is the only supported source.

Type: String

Required: No

**[regionalCertificateName](#API_CreateDomainName_RequestSyntax)**

The user-friendly name of the certificate that will be used by regional endpoint for this domain name.

Type: String

Required: No

**[routingMode](#API_CreateDomainName_RequestSyntax)**

The routing mode for this domain name. The routing mode determines how API Gateway sends traffic from your custom domain name to your APIs.

Type: String

Valid Values: `BASE_PATH_MAPPING_ONLY | ROUTING_RULE_ONLY | ROUTING_RULE_THEN_BASE_PATH_MAPPING`

Required: No

**[securityPolicy](#API_CreateDomainName_RequestSyntax)**

The Transport Layer Security (TLS) version + cipher suite for this DomainName.

Type: String

Valid Values: `TLS_1_0 | TLS_1_2 | SecurityPolicy_TLS13_1_3_2025_09 | SecurityPolicy_TLS13_1_3_FIPS_2025_09 | SecurityPolicy_TLS13_1_2_PFS_PQ_2025_09 | SecurityPolicy_TLS13_1_2_FIPS_PQ_2025_09 | SecurityPolicy_TLS13_1_2_PQ_2025_09 | SecurityPolicy_TLS13_1_2_2021_06 | SecurityPolicy_TLS13_2025_EDGE | SecurityPolicy_TLS12_PFS_2025_EDGE | SecurityPolicy_TLS12_2018_EDGE`

Required: No

**[tags](#API_CreateDomainName_RequestSyntax)**

The key-value map of strings. The valid character set is \[a-zA-Z+-=.\_:/\]. The tag key can be up to 128 characters and must not start with `aws:`. The tag value can be up to 256 characters.

Type: String to string map

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 201
Content-type: application/json

{
   "certificateArn": "string",
   "certificateName": "string",
   "certificateUploadDate": number,
   "distributionDomainName": "string",
   "distributionHostedZoneId": "string",
   "domainName": "string",
   "domainNameArn": "string",
   "domainNameId": "string",
   "domainNameStatus": "string",
   "domainNameStatusMessage": "string",
   "endpointAccessMode": "string",
   "endpointConfiguration": {
      "ipAddressType": "string",
      "types": [ "string" ],
      "vpcEndpointIds": [ "string" ]
   },
   "managementPolicy": "string",
   "mutualTlsAuthentication": {
      "truststoreUri": "string",
      "truststoreVersion": "string",
      "truststoreWarnings": [ "string" ]
   },
   "ownershipVerificationCertificateArn": "string",
   "policy": "string",
   "regionalCertificateArn": "string",
   "regionalCertificateName": "string",
   "regionalDomainName": "string",
   "regionalHostedZoneId": "string",
   "routingMode": "string",
   "securityPolicy": "string",
   "tags": {
      "string" : "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 201 response.

The following data is returned in JSON format by the service.

**[certificateArn](#API_CreateDomainName_ResponseSyntax)**

The reference to an AWS-managed certificate that will be used by edge-optimized endpoint or private endpoint for this domain name. AWS Certificate Manager is the only supported source.

Type: String

**[certificateName](#API_CreateDomainName_ResponseSyntax)**

The name of the certificate that will be used by edge-optimized endpoint or private endpoint for this domain name.

Type: String

**[certificateUploadDate](#API_CreateDomainName_ResponseSyntax)**

The timestamp when the certificate that was used by edge-optimized endpoint or private endpoint for this domain name was uploaded.

Type: Timestamp

**[distributionDomainName](#API_CreateDomainName_ResponseSyntax)**

The domain name of the Amazon CloudFront distribution associated with this custom domain name for an edge-optimized endpoint. You set up this association when adding a DNS record pointing the custom domain name to this distribution name. For more information about CloudFront distributions, see the Amazon CloudFront documentation.

Type: String

**[distributionHostedZoneId](#API_CreateDomainName_ResponseSyntax)**

The region-agnostic Amazon Route 53 Hosted Zone ID of the edge-optimized endpoint. The valid value is `Z2FDTNDATAQYW2` for all the regions. For more information, see Set up a Regional Custom Domain Name and AWS Regions and Endpoints for API Gateway.

Type: String

**[domainName](#API_CreateDomainName_ResponseSyntax)**

The custom domain name as an API host name, for example, `my-api.example.com`.

Type: String

**[domainNameArn](#API_CreateDomainName_ResponseSyntax)**

The ARN of the domain name.

Type: String

**[domainNameId](#API_CreateDomainName_ResponseSyntax)**

The identifier for the domain name resource. Supported only for private custom domain names.

Type: String

**[domainNameStatus](#API_CreateDomainName_ResponseSyntax)**

The status of the DomainName migration. The valid values are `AVAILABLE` and `UPDATING`. If the status is `UPDATING`, the domain cannot be modified further until the existing operation is complete. If it is `AVAILABLE`, the domain can be updated.

Type: String

Valid Values: `AVAILABLE | UPDATING | PENDING | PENDING_CERTIFICATE_REIMPORT | PENDING_OWNERSHIP_VERIFICATION | FAILED`

**[domainNameStatusMessage](#API_CreateDomainName_ResponseSyntax)**

An optional text message containing detailed information about status of the DomainName migration.

Type: String

**[endpointAccessMode](#API_CreateDomainName_ResponseSyntax)**

The endpoint access mode of the DomainName.

Type: String

Valid Values: `BASIC | STRICT`

**[endpointConfiguration](#API_CreateDomainName_ResponseSyntax)**

The endpoint configuration of this DomainName showing the endpoint types and IP address types of the domain name.

Type: [EndpointConfiguration](https://docs.aws.amazon.com/apigateway/latest/api/API_EndpointConfiguration.html) object

**[managementPolicy](#API_CreateDomainName_ResponseSyntax)**

A stringified JSON policy document that applies to the API Gateway Management service for this DomainName. This policy document controls access for access association sources to create domain name access associations with this DomainName. Supported only for private custom domain names.

Type: String

**[mutualTlsAuthentication](#API_CreateDomainName_ResponseSyntax)**

The mutual TLS authentication configuration for a custom domain name. If specified, API Gateway
performs two-way authentication between the client and the server. Clients must present a
trusted certificate to access your API.

Type: [MutualTlsAuthentication](https://docs.aws.amazon.com/apigateway/latest/api/API_MutualTlsAuthentication.html) object

**[ownershipVerificationCertificateArn](#API_CreateDomainName_ResponseSyntax)**

The ARN of the public certificate issued by ACM to validate ownership of your custom
domain. Only required when configuring mutual TLS and using an ACM imported or private CA
certificate ARN as the regionalCertificateArn.

Type: String

**[policy](#API_CreateDomainName_ResponseSyntax)**

A stringified JSON policy document that applies to the `execute-api` service for this DomainName regardless of the caller and Method
configuration. Supported only for private custom
domain names.

Type: String

**[regionalCertificateArn](#API_CreateDomainName_ResponseSyntax)**

The reference to an AWS-managed certificate that will be used for validating the regional domain name. AWS Certificate Manager is the only supported source.

Type: String

**[regionalCertificateName](#API_CreateDomainName_ResponseSyntax)**

The name of the certificate that will be used for validating the regional domain name.

Type: String

**[regionalDomainName](#API_CreateDomainName_ResponseSyntax)**

The domain name associated with the regional endpoint for this custom domain name. You set up this association by adding a DNS record that points the custom domain name to this regional domain name. The regional domain name is returned by API Gateway when you create a regional endpoint.

Type: String

**[regionalHostedZoneId](#API_CreateDomainName_ResponseSyntax)**

The region-specific Amazon Route 53 Hosted Zone ID of the regional endpoint. For more information, see Set up a Regional Custom Domain Name and AWS Regions and Endpoints for API Gateway.

Type: String

**[routingMode](#API_CreateDomainName_ResponseSyntax)**

The routing mode for this domain name. The routing mode determines how API Gateway sends traffic from your custom domain name to your APIs.

Type: String

Valid Values: `BASE_PATH_MAPPING_ONLY | ROUTING_RULE_ONLY | ROUTING_RULE_THEN_BASE_PATH_MAPPING`

**[securityPolicy](#API_CreateDomainName_ResponseSyntax)**

The Transport Layer Security (TLS) version + cipher suite for this DomainName.

Type: String

Valid Values: `TLS_1_0 | TLS_1_2 | SecurityPolicy_TLS13_1_3_2025_09 | SecurityPolicy_TLS13_1_3_FIPS_2025_09 | SecurityPolicy_TLS13_1_2_PFS_PQ_2025_09 | SecurityPolicy_TLS13_1_2_FIPS_PQ_2025_09 | SecurityPolicy_TLS13_1_2_PQ_2025_09 | SecurityPolicy_TLS13_1_2_2021_06 | SecurityPolicy_TLS13_2025_EDGE | SecurityPolicy_TLS12_PFS_2025_EDGE | SecurityPolicy_TLS12_2018_EDGE`

**[tags](#API_CreateDomainName_ResponseSyntax)**

The collection of tags. Each tag element is associated with a given resource.

Type: String to string map

## Errors

For information about the errors that are common to all actions, see [Common Error Types](https://docs.aws.amazon.com/apigateway/latest/api/CommonErrors.html).

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

### Create a domain name

This example illustrates one usage of CreateDomainName.

#### Sample Request

```

POST /domainnames HTTP/1.1
Content-Type: application/json
Host: apigateway.us-east-1.amazonaws.com
X-Amz-Date: 20160615T211441Z
Authorization: AWS4-HMAC-SHA256 Credential={access_key_ID}/20160615/us-east-1/apigateway/aws4_request, SignedHeaders=content-length;content-type;host;x-amz-date, Signature={sig4_hash}

{
  "domainName" : "my-api.example.com",
  "certificateName": "my-cert-created-today",
  "certificateArn": "arn:aws:acm:us-east-1:012345678910:certificate/fb1b9770-a305-495d-aefb-27e5e101ff3",
  "endpointConfiguration": {
    "types": ["EDGE"]
  }
}
```

#### Sample Response

```

{
  "_links": {
    "curies": [
      {
        "href": "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-basepathmapping-{rel}.html",
        "name": "basepathmapping",
        "templated": true
      },
      {
        "href": "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-domainname-{rel}.html",
        "name": "domainname",
        "templated": true
      }
    ],
    "self": {
      "href": "/domainnames/my-api.example.com"
    },
    "basepathmapping:by-base-path": {
      "href": "/domainnames/my-api.example.com/basepathmappings/{base_path}",
      "templated": true
    },
    "basepathmapping:create": {
      "href": "/domainnames/my-api.example.com/basepathmappings"
    },
    "domainname:basepathmappings": {
      "href": "/domainnames/my-api.example.com/basepathmappings{?limit}",
      "templated": true
    },
    "domainname:delete": {
      "href": "/domainnames/my-api.example.com"
    },
    "domainname:update": {
      "href": "/domainnames/my-api.example.com"
    }
  },
  "certificateArn": "arn:aws:acm:us-east-1:012345678910:certificate/fb1b9770-a305-495d-aefb-27e5e101ff3",
  "certificateName": "my-cert-created-today",
  "certificateUploadDate": "2016-06-15T21:14:43Z",
  "distributionDomainName": "d2ck2x1vuc8qzh.cloudfront.net",
  "domainName": "my-api.example.com"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/apigateway-2015-07-09/CreateDomainName)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/apigateway-2015-07-09/CreateDomainName)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/apigateway-2015-07-09/CreateDomainName)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/apigateway-2015-07-09/CreateDomainName)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/apigateway-2015-07-09/CreateDomainName)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/apigateway-2015-07-09/CreateDomainName)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/apigateway-2015-07-09/CreateDomainName)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/apigateway-2015-07-09/CreateDomainName)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/apigateway-2015-07-09/CreateDomainName)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/apigateway-2015-07-09/CreateDomainName)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CreateDocumentationVersion

CreateDomainNameAccessAssociation
