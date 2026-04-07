# CreateVerifiedAccessEndpoint

An AWS Verified Access endpoint is where you define your application along with an optional endpoint-level access policy.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ApplicationDomain**

The DNS name for users to reach your application.

Type: String

Required: No

**AttachmentType**

The type of attachment.

Type: String

Valid Values: `vpc`

Required: Yes

**CidrOptions**

The CIDR options. This parameter is required if the endpoint type is `cidr`.

Type: [CreateVerifiedAccessEndpointCidrOptions](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateVerifiedAccessEndpointCidrOptions.html) object

Required: No

**ClientToken**

A unique, case-sensitive token that you provide to ensure idempotency of your
modification request. For more information, see [Ensuring idempotency](https://docs.aws.amazon.com/ec2/latest/devguide/ec2-api-idempotency.html).

Type: String

Required: No

**Description**

A description for the Verified Access endpoint.

Type: String

Required: No

**DomainCertificateArn**

The ARN of the public TLS/SSL certificate in AWS Certificate Manager to associate with the endpoint.
The CN in the certificate must match the DNS name your end users will use to reach your
application.

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**EndpointDomainPrefix**

A custom identifier that is prepended to the DNS name that is generated for the
endpoint.

Type: String

Required: No

**EndpointType**

The type of Verified Access endpoint to create.

Type: String

Valid Values: `load-balancer | network-interface | rds | cidr`

Required: Yes

**LoadBalancerOptions**

The load balancer details. This parameter is required if the endpoint type is
`load-balancer`.

Type: [CreateVerifiedAccessEndpointLoadBalancerOptions](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateVerifiedAccessEndpointLoadBalancerOptions.html) object

Required: No

**NetworkInterfaceOptions**

The network interface details. This parameter is required if the endpoint type is
`network-interface`.

Type: [CreateVerifiedAccessEndpointEniOptions](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateVerifiedAccessEndpointEniOptions.html) object

Required: No

**PolicyDocument**

The Verified Access policy document.

Type: String

Required: No

**RdsOptions**

The RDS details. This parameter is required if the endpoint type is `rds`.

Type: [CreateVerifiedAccessEndpointRdsOptions](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateVerifiedAccessEndpointRdsOptions.html) object

Required: No

**SecurityGroupId.N**

The IDs of the security groups to associate with the Verified Access endpoint. Required if `AttachmentType` is set to `vpc`.

Type: Array of strings

Required: No

**SseSpecification**

The options for server side encryption.

Type: [VerifiedAccessSseSpecificationRequest](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_VerifiedAccessSseSpecificationRequest.html) object

Required: No

**TagSpecification.N**

The tags to assign to the Verified Access endpoint.

Type: Array of [TagSpecification](api-tagspecification.md) objects

Required: No

**VerifiedAccessGroupId**

The ID of the Verified Access group to associate the endpoint with.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**verifiedAccessEndpoint**

Details about the Verified Access endpoint.

Type: [VerifiedAccessEndpoint](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_VerifiedAccessEndpoint.html) object

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/CreateVerifiedAccessEndpoint)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/CreateVerifiedAccessEndpoint)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/CreateVerifiedAccessEndpoint)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/CreateVerifiedAccessEndpoint)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/CreateVerifiedAccessEndpoint)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/CreateVerifiedAccessEndpoint)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/CreateVerifiedAccessEndpoint)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/CreateVerifiedAccessEndpoint)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/CreateVerifiedAccessEndpoint)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/CreateVerifiedAccessEndpoint)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CreateTransitGatewayVpcAttachment

CreateVerifiedAccessGroup
