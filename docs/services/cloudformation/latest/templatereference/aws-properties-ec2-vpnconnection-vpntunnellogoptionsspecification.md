This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::VPNConnection VpnTunnelLogOptionsSpecification

Options for logging VPN tunnel activity.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CloudwatchLogOptions" : CloudwatchLogOptionsSpecification
}

```

### YAML

```yaml

  CloudwatchLogOptions:
    CloudwatchLogOptionsSpecification

```

## Properties

`CloudwatchLogOptions`

Options for sending VPN tunnel logs to CloudWatch.

_Required_: No

_Type_: [CloudwatchLogOptionsSpecification](aws-properties-ec2-vpnconnection-cloudwatchlogoptionsspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

VpnTunnelOptionsSpecification

All content copied from https://docs.aws.amazon.com/.
