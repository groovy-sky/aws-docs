This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PCS::Cluster Networking

The networking configuration for the cluster's control plane.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NetworkType" : String,
  "SecurityGroupIds" : [ String, ... ],
  "SubnetIds" : [ String, ... ]
}

```

### YAML

```yaml

  NetworkType: String
  SecurityGroupIds:
    - String
  SubnetIds:
    - String

```

## Properties

`NetworkType`

The IP address version the cluster uses. The default is `IPV4`.

_Required_: No

_Type_: String

_Allowed values_: `IPV4 | IPV6`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SecurityGroupIds`

The list of security group IDs associated with the Elastic Network Interface (ENI)
created in subnets.

The following rules are required:

- Inbound rule 1

- Protocol: All

- Ports: All

- Source: Self

- Outbound rule 1

- Protocol: All

- Ports: All

- Destination: 0.0.0.0/0 (IPv4) or ::/0 (IPv6)

- Outbound rule 2

- Protocol: All

- Ports: All

- Destination: Self

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SubnetIds`

The ID of the subnet where AWS PCS creates an Elastic Network Interface (ENI)
to enable communication between managed controllers and AWS PCS resources. The
subnet must have an available IP address, cannot reside in AWS Outposts, AWS Wavelength, or
an AWS Local Zone.

Example: `subnet-abcd1234`

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

JwtKey

Scheduler

All content copied from https://docs.aws.amazon.com/.
