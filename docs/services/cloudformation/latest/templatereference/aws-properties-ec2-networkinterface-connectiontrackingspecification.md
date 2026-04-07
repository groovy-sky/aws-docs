This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::NetworkInterface ConnectionTrackingSpecification

Configurable options for connection tracking on a network interface. For more information, see [Connection tracking timeouts](../../../ec2/latest/userguide/security-group-connection-tracking.md#connection-tracking-timeouts) in the _Amazon Elastic Compute Cloud User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TcpEstablishedTimeout" : Integer,
  "UdpStreamTimeout" : Integer,
  "UdpTimeout" : Integer
}

```

### YAML

```yaml

  TcpEstablishedTimeout: Integer
  UdpStreamTimeout: Integer
  UdpTimeout: Integer

```

## Properties

`TcpEstablishedTimeout`

Timeout (in seconds) for idle TCP
connections in an established state. Min: 60 seconds. Max: 432000 seconds (5
days). Default: 432000 seconds. Recommended: Less than 432000 seconds.

_Required_: No

_Type_: Integer

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`UdpStreamTimeout`

Timeout (in seconds) for idle UDP
flows classified as streams which have seen more than one request-response
transaction. Min: 60 seconds. Max: 180 seconds (3 minutes). Default: 180
seconds.

_Required_: No

_Type_: Integer

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`UdpTimeout`

Timeout (in seconds) for idle UDP flows that
have seen traffic only in a single direction or a single request-response
transaction. Min: 30 seconds. Max: 60 seconds. Default: 30 seconds.

_Required_: No

_Type_: Integer

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::EC2::NetworkInterface

InstanceIpv6Address
