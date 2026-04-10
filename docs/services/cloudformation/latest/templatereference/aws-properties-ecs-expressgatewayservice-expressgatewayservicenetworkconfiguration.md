This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::ExpressGatewayService ExpressGatewayServiceNetworkConfiguration

The network configuration for an Express service. By default, an Express service utilizes subnets and security groups associated with the default VPC.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SecurityGroups" : [ String, ... ],
  "Subnets" : [ String, ... ]
}

```

### YAML

```yaml

  SecurityGroups:
    - String
  Subnets:
    - String

```

## Properties

`SecurityGroups`

The IDs of the security groups associated with the Express service.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Subnets`

The IDs of the subnets associated with the Express service.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExpressGatewayServiceConfiguration

ExpressGatewayServiceStatus

All content copied from https://docs.aws.amazon.com/.
