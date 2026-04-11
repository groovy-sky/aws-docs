# VpcLink

An API Gateway VPC link for a RestApi to access resources in an Amazon Virtual Private Cloud (VPC).

## Contents

**description**

The description of the VPC link.

Type: String

Required: No

**id**

The identifier of the VpcLink. It is used in an Integration to reference this VpcLink.

Type: String

Required: No

**name**

The name used to label and identify the VPC link.

Type: String

Required: No

**status**

The status of the VPC link. The valid values are `AVAILABLE`, `PENDING`, `DELETING`, or `FAILED`. Deploying an API will wait if the status is `PENDING` and will fail if the status is `DELETING`.

Type: String

Valid Values: `AVAILABLE | PENDING | DELETING | FAILED`

Required: No

**statusMessage**

A description about the VPC link status.

Type: String

Required: No

**tags**

The collection of tags. Each tag element is associated with a given resource.

Type: String to string map

Required: No

**targetArns**

The ARN of the network load balancer of the VPC targeted by the VPC link. The network load balancer must be owned by the same AWS account of the API owner.

Type: Array of strings

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/vpclink.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/vpclink.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/vpclink.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UsagePlanKey

Common Parameters

All content copied from https://docs.aws.amazon.com/.
