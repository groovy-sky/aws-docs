# NatGatewayAttachedAppliance

Information about an appliance attached to a NAT Gateway, providing managed security solutions for traffic filtering and inspection.

## Contents

**applianceArn**

The Amazon Resource Name (ARN) of the attached appliance, identifying the specific proxy or security appliance resource.

Type: String

Required: No

**attachmentState**

The current attachment state of the appliance.

Type: String

Valid Values: `attaching | attached | detaching | detached | attach-failed | detach-failed`

Required: No

**failureCode**

The failure code if the appliance attachment or modification operation failed.

Type: String

Required: No

**failureMessage**

A descriptive message explaining the failure if the appliance attachment or modification operation failed.

Type: String

Required: No

**modificationState**

The current modification state of the appliance.

Type: String

Valid Values: `modifying | completed | failed`

Required: No

**type**

The type of appliance attached to the NAT Gateway. For network firewall proxy functionality, this will be "network-firewall-proxy".

Type: String

Valid Values: `network-firewall-proxy`

Required: No

**vpcEndpointId**

The VPC endpoint ID used to route traffic from application VPCs to the proxy for inspection and filtering.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/natgatewayattachedappliance.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/natgatewayattachedappliance.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/natgatewayattachedappliance.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

NatGatewayAddress

NativeApplicationOidcOptions
