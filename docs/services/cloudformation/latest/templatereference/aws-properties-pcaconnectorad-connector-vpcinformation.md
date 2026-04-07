This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PCAConnectorAD::Connector VpcInformation

Information about your VPC and security groups used with the connector.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IpAddressType" : String,
  "SecurityGroupIds" : [ String, ... ]
}

```

### YAML

```yaml

  IpAddressType: String
  SecurityGroupIds:
    - String

```

## Properties

`IpAddressType`

The VPC IP address type.

_Required_: No

_Type_: String

_Allowed values_: `IPV4 | DUALSTACK`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SecurityGroupIds`

The security groups used with the connector. You can use a maximum of 4 security groups
with a connector.

_Required_: Yes

_Type_: Array of String

_Minimum_: `11 | 1`

_Maximum_: `20 | 5`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::PCAConnectorAD::Connector

AWS::PCAConnectorAD::DirectoryRegistration
