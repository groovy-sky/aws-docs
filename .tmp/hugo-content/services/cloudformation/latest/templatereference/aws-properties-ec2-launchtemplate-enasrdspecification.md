This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::LaunchTemplate EnaSrdSpecification

ENA Express uses AWS Scalable Reliable Datagram (SRD) technology to increase the
maximum bandwidth used per stream and minimize tail latency of network traffic between EC2 instances.
With ENA Express, you can communicate between two EC2 instances in the same subnet within the same
account, or in different accounts. Both sending and receiving instances must have ENA Express enabled.

To improve the reliability of network packet delivery, ENA Express reorders network packets on the
receiving end by default. However, some UDP-based applications are designed to handle network packets
that are out of order to reduce the overhead for packet delivery at the network layer. When ENA Express
is enabled, you can specify whether UDP network traffic uses it.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EnaSrdEnabled" : Boolean,
  "EnaSrdUdpSpecification" : EnaSrdUdpSpecification
}

```

### YAML

```yaml

  EnaSrdEnabled: Boolean
  EnaSrdUdpSpecification:
    EnaSrdUdpSpecification

```

## Properties

`EnaSrdEnabled`

Indicates whether ENA Express is enabled for the network interface.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnaSrdUdpSpecification`

Configures ENA Express for UDP network traffic.

_Required_: No

_Type_: [EnaSrdUdpSpecification](aws-properties-ec2-launchtemplate-enasrdudpspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Ebs

EnaSrdUdpSpecification

All content copied from https://docs.aws.amazon.com/.
