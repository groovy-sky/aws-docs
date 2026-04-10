This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VpcLattice::TargetGroup TargetGroupConfig

Describes the configuration of a target group.

For more information, see [Target groups](../../../vpc-lattice/latest/ug/target-groups.md) in the
_Amazon VPC Lattice User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HealthCheck" : HealthCheckConfig,
  "IpAddressType" : String,
  "LambdaEventStructureVersion" : String,
  "Port" : Integer,
  "Protocol" : String,
  "ProtocolVersion" : String,
  "VpcIdentifier" : String
}

```

### YAML

```yaml

  HealthCheck:
    HealthCheckConfig
  IpAddressType: String
  LambdaEventStructureVersion: String
  Port: Integer
  Protocol: String
  ProtocolVersion: String
  VpcIdentifier: String

```

## Properties

`HealthCheck`

The health check configuration. Not supported if the target group type is
`LAMBDA` or `ALB`.

_Required_: No

_Type_: [HealthCheckConfig](aws-properties-vpclattice-targetgroup-healthcheckconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IpAddressType`

The type of IP address used for the target group. Supported only if the target group type is
`IP`. The default is `IPV4`.

_Required_: No

_Type_: String

_Allowed values_: `IPV4 | IPV6`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LambdaEventStructureVersion`

The version of the event structure that your Lambda function receives. Supported only if the
target group type is `LAMBDA`. The default is `V1`.

_Required_: No

_Type_: String

_Allowed values_: `V1 | V2`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Port`

The port on which the targets are listening. For HTTP, the default is 80. For HTTPS, the
default is 443. Not supported if the target group type is `LAMBDA`.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `65535`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Protocol`

The protocol to use for routing traffic to the targets. The default is the protocol of the
target group. Not supported if the target group type is `LAMBDA`.

_Required_: No

_Type_: String

_Allowed values_: `HTTP | HTTPS | TCP`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProtocolVersion`

The protocol version. The default is `HTTP1`. Not supported if the target group
type is `LAMBDA`.

_Required_: No

_Type_: String

_Allowed values_: `HTTP1 | HTTP2 | GRPC`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VpcIdentifier`

The ID of the VPC. Not supported if the target group type is `LAMBDA`.

_Required_: No

_Type_: String

_Pattern_: `^vpc-(([0-9a-z]{8})|([0-9a-z]{17}))$`

_Minimum_: `5`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Target

Next

All content copied from https://docs.aws.amazon.com/.
