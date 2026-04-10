This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Transfer::Connector ConnectorVpcLatticeEgressConfig

VPC\_LATTICE egress configuration that specifies the Resource Configuration ARN and
port for connecting to SFTP servers through customer VPCs. Requires a valid Resource
Configuration with appropriate network access.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PortNumber" : Integer,
  "ResourceConfigurationArn" : String
}

```

### YAML

```yaml

  PortNumber: Integer
  ResourceConfigurationArn: String

```

## Properties

`PortNumber`

Port number for connecting to the SFTP server through VPC\_LATTICE. Defaults to 22 if
not specified. Must match the port on which the target SFTP server is listening.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceConfigurationArn`

ARN of the VPC\_LATTICE Resource Configuration that defines the target SFTP server
location. Must point to a valid Resource Configuration in the customer's VPC with
appropriate network connectivity to the SFTP server.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:[a-z0-9\-]+:vpc-lattice:[a-zA-Z0-9\-]+:\d{12}:resourceconfiguration/rcfg-[0-9a-z]{17}$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConnectorEgressConfig

SftpConfig

All content copied from https://docs.aws.amazon.com/.
