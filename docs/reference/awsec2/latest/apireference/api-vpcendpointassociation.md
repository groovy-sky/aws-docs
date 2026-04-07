# VpcEndpointAssociation

Describes the VPC resources, VPC endpoint services, Lattice services, or service
networks associated with the VPC endpoint.

## Contents

**associatedResourceAccessibility**

The connectivity status of the resources associated to a VPC endpoint. The resource is
accessible if the associated resource configuration is `AVAILABLE`, otherwise
the resource is inaccessible.

Type: String

Required: No

**associatedResourceArn**

The Amazon Resource Name (ARN) of the associated resource.

Type: String

Required: No

**dnsEntry**

The DNS entry of the VPC endpoint association.

Type: [DnsEntry](api-dnsentry.md) object

Required: No

**failureCode**

An error code related to why an VPC endpoint association failed.

Type: String

Required: No

**failureReason**

A message related to why an VPC endpoint association failed.

Type: String

Required: No

**id**

The ID of the VPC endpoint association.

Type: String

Required: No

**privateDnsEntry**

The private DNS entry of the VPC endpoint association.

Type: [DnsEntry](api-dnsentry.md) object

Required: No

**resourceConfigurationGroupArn**

The Amazon Resource Name (ARN) of the resource configuration group.

Type: String

Required: No

**serviceNetworkArn**

The Amazon Resource Name (ARN) of the service network.

Type: String

Required: No

**serviceNetworkName**

The name of the service network.

Type: String

Required: No

**TagSet.N**

The tags to apply to the VPC endpoint association.

Type: Array of [Tag](api-tag.md) objects

Required: No

**vpcEndpointId**

The ID of the VPC endpoint.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/vpcendpointassociation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/vpcendpointassociation.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/vpcendpointassociation.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

VpcEndpoint

VpcEndpointConnection
