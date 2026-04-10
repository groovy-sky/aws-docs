This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::EC2Fleet TagSpecification

Specifies the tags to apply to a resource when the resource is being created for an EC2
Fleet.

`TagSpecification` is a property of the [AWS::EC2::EC2Fleet](../userguide/aws-resource-ec2-ec2fleet.md) resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ResourceType" : String,
  "Tags" : [ Tag, ... ]
}

```

### YAML

```yaml

  ResourceType: String
  Tags:
    - Tag

```

## Properties

`ResourceType`

The type of resource to tag.

_Required_: No

_Type_: String

_Allowed values_: `client-vpn-endpoint | customer-gateway | dedicated-host | dhcp-options | egress-only-internet-gateway | elastic-gpu | elastic-ip | export-image-task | export-instance-task | fleet | fpga-image | host-reservation | image | import-image-task | import-snapshot-task | instance | internet-gateway | key-pair | launch-template | local-gateway-route-table-vpc-association | natgateway | network-acl | network-insights-analysis | network-insights-path | network-interface | placement-group | reserved-instances | route-table | security-group | snapshot | spot-fleet-request | spot-instances-request | subnet | traffic-mirror-filter | traffic-mirror-session | traffic-mirror-target | transit-gateway | transit-gateway-attachment | transit-gateway-connect-peer | transit-gateway-multicast-domain | transit-gateway-route-table | volume | vpc | vpc-flow-log | vpc-peering-connection | vpn-connection | vpn-gateway`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags to apply to the resource.

_Required_: No

_Type_: Array of [Tag](aws-properties-ec2-ec2fleet-tag.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [TagSpecification](../../../../reference/awsec2/latest/apireference/api-tagspecification.md) in the _Amazon EC2 API_
_Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

TargetCapacitySpecificationRequest

All content copied from https://docs.aws.amazon.com/.
