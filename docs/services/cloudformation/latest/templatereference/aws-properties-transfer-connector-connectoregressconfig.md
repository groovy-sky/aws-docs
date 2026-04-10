This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Transfer::Connector ConnectorEgressConfig

Configuration structure that defines how traffic is routed from the connector to the
SFTP server. Contains VPC Lattice settings when using VPC\_LATTICE egress type for
private connectivity through customer VPCs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "VpcLattice" : ConnectorVpcLatticeEgressConfig
}

```

### YAML

```yaml

  VpcLattice:
    ConnectorVpcLatticeEgressConfig

```

## Properties

`VpcLattice`

VPC\_LATTICE configuration for routing connector traffic through customer VPCs. Enables
private connectivity to SFTP servers without requiring public internet access or complex
network configurations.

_Required_: Yes

_Type_: [ConnectorVpcLatticeEgressConfig](aws-properties-transfer-connector-connectorvpclatticeegressconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConnectorAsyncMdnConfig

ConnectorVpcLatticeEgressConfig

All content copied from https://docs.aws.amazon.com/.
