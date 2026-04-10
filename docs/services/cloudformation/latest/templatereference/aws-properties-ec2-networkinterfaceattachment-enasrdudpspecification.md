This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::NetworkInterfaceAttachment EnaSrdUdpSpecification

ENA Express is compatible with both TCP and UDP transport protocols. When it's enabled, TCP traffic
automatically uses it. However, some UDP-based applications are designed to handle network packets that are
out of order, without a need for retransmission, such as live video broadcasting or other near-real-time
applications. For UDP traffic, you can specify whether to use ENA Express, based on your application
environment needs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EnaSrdUdpEnabled" : Boolean
}

```

### YAML

```yaml

  EnaSrdUdpEnabled: Boolean

```

## Properties

`EnaSrdUdpEnabled`

Indicates whether UDP traffic to and from the instance uses ENA Express. To specify this setting,
you must first enable ENA Express.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EnaSrdSpecification

AWS::EC2::NetworkInterfacePermission

All content copied from https://docs.aws.amazon.com/.
