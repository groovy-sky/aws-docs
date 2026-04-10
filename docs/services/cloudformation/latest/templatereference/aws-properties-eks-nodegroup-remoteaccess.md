This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EKS::Nodegroup RemoteAccess

An object representing the remote access configuration for the managed node
group.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Ec2SshKey" : String,
  "SourceSecurityGroups" : [ String, ... ]
}

```

### YAML

```yaml

  Ec2SshKey: String
  SourceSecurityGroups:
    - String

```

## Properties

`Ec2SshKey`

The Amazon EC2 SSH key name that provides access for SSH communication with the nodes in
the managed node group. For more information, see [Amazon EC2 key\
pairs and Linux instances](../../../ec2/latest/userguide/ec2-key-pairs.md) in the _Amazon Elastic Compute Cloud User Guide for Linux Instances_. For Windows, an Amazon EC2 SSH
key is used to obtain the RDP password. For more information, see [Amazon EC2 key\
pairs and Windows instances](../../../ec2/latest/windowsguide/ec2-key-pairs.md) in the _Amazon Elastic Compute Cloud User Guide for Windows Instances_.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceSecurityGroups`

The security group IDs that are allowed SSH access (port 22) to the nodes. For
Windows, the port is 3389. If you specify an Amazon EC2 SSH key but don't specify a source
security group when you create a managed node group, then the port on the nodes is
opened to the internet ( `0.0.0.0/0`). For more information, see [Security\
Groups for Your VPC](../../../vpc/latest/userguide/vpc-securitygroups.md) in the _Amazon Virtual Private Cloud User Guide_.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NodeRepairConfigOverrides

ScalingConfig

All content copied from https://docs.aws.amazon.com/.
