# DomainName

Represents a custom domain name as a user-friendly host name of an API (RestApi).

## Contents

**certificateArn**

The reference to an AWS-managed certificate that will be used by edge-optimized endpoint or private endpoint for this domain name. AWS Certificate Manager is the only supported source.

Type: String

Required: No

**certificateName**

The name of the certificate that will be used by edge-optimized endpoint or private endpoint for this domain name.

Type: String

Required: No

**certificateUploadDate**

The timestamp when the certificate that was used by edge-optimized endpoint or private endpoint for this domain name was uploaded.

Type: Timestamp

Required: No

**distributionDomainName**

The domain name of the Amazon CloudFront distribution associated with this custom domain name for an edge-optimized endpoint. You set up this association when adding a DNS record pointing the custom domain name to this distribution name. For more information about CloudFront distributions, see the Amazon CloudFront documentation.

Type: String

Required: No

**distributionHostedZoneId**

The region-agnostic Amazon Route 53 Hosted Zone ID of the edge-optimized endpoint. The valid value is `Z2FDTNDATAQYW2` for all the regions. For more information, see Set up a Regional Custom Domain Name and AWS Regions and Endpoints for API Gateway.

Type: String

Required: No

**domainName**

The custom domain name as an API host name, for example, `my-api.example.com`.

Type: String

Required: No

**domainNameArn**

The ARN of the domain name.

Type: String

Required: No

**domainNameId**

The identifier for the domain name resource. Supported only for private custom domain names.

Type: String

Required: No

**domainNameStatus**

The status of the DomainName migration. The valid values are `AVAILABLE` and `UPDATING`. If the status is `UPDATING`, the domain cannot be modified further until the existing operation is complete. If it is `AVAILABLE`, the domain can be updated.

Type: String

Valid Values: `AVAILABLE | UPDATING | PENDING | PENDING_CERTIFICATE_REIMPORT | PENDING_OWNERSHIP_VERIFICATION | FAILED`

Required: No

**domainNameStatusMessage**

An optional text message containing detailed information about status of the DomainName migration.

Type: String

Required: No

**endpointAccessMode**

The endpoint access mode of the DomainName.

Type: String

Valid Values: `BASIC | STRICT`

Required: No

**endpointConfiguration**

The endpoint configuration of this DomainName showing the endpoint types and IP address types of the domain name.

Type: [EndpointConfiguration](api-endpointconfiguration.md) object

Required: No

**managementPolicy**

A stringified JSON policy document that applies to the API Gateway Management service for this DomainName. This policy document controls access for access association sources to create domain name access associations with this DomainName. Supported only for private custom domain names.

Type: String

Required: No

**mutualTlsAuthentication**

The mutual TLS authentication configuration for a custom domain name. If specified, API Gateway
performs two-way authentication between the client and the server. Clients must present a
trusted certificate to access your API.

Type: [MutualTlsAuthentication](api-mutualtlsauthentication.md) object

Required: No

**ownershipVerificationCertificateArn**

The ARN of the public certificate issued by ACM to validate ownership of your custom
domain. Only required when configuring mutual TLS and using an ACM imported or private CA
certificate ARN as the regionalCertificateArn.

Type: String

Required: No

**policy**

A stringified JSON policy document that applies to the `execute-api` service for this DomainName regardless of the caller and Method
configuration. Supported only for private custom
domain names.

Type: String

Required: No

**regionalCertificateArn**

The reference to an AWS-managed certificate that will be used for validating the regional domain name. AWS Certificate Manager is the only supported source.

Type: String

Required: No

**regionalCertificateName**

The name of the certificate that will be used for validating the regional domain name.

Type: String

Required: No

**regionalDomainName**

The domain name associated with the regional endpoint for this custom domain name. You set up this association by adding a DNS record that points the custom domain name to this regional domain name. The regional domain name is returned by API Gateway when you create a regional endpoint.

Type: String

Required: No

**regionalHostedZoneId**

The region-specific Amazon Route 53 Hosted Zone ID of the regional endpoint. For more information, see Set up a Regional Custom Domain Name and AWS Regions and Endpoints for API Gateway.

Type: String

Required: No

**routingMode**

The routing mode for this domain name. The routing mode determines how API Gateway sends traffic from your custom domain name to your APIs.

Type: String

Valid Values: `BASE_PATH_MAPPING_ONLY | ROUTING_RULE_ONLY | ROUTING_RULE_THEN_BASE_PATH_MAPPING`

Required: No

**securityPolicy**

The Transport Layer Security (TLS) version + cipher suite for this DomainName.

Type: String

Valid Values: `TLS_1_0 | TLS_1_2 | SecurityPolicy_TLS13_1_3_2025_09 | SecurityPolicy_TLS13_1_3_FIPS_2025_09 | SecurityPolicy_TLS13_1_2_PFS_PQ_2025_09 | SecurityPolicy_TLS13_1_2_FIPS_PQ_2025_09 | SecurityPolicy_TLS13_1_2_PQ_2025_09 | SecurityPolicy_TLS13_1_2_2021_06 | SecurityPolicy_TLS13_2025_EDGE | SecurityPolicy_TLS12_PFS_2025_EDGE | SecurityPolicy_TLS12_2018_EDGE`

Required: No

**tags**

The collection of tags. Each tag element is associated with a given resource.

Type: String to string map

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/domainname.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/domainname.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/domainname.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DocumentationVersion

DomainNameAccessAssociation

All content copied from https://docs.aws.amazon.com/.
