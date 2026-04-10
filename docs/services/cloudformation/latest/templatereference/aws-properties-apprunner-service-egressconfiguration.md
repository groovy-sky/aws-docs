This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppRunner::Service EgressConfiguration

Describes configuration settings related to outbound network traffic of an AWS App Runner service.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EgressType" : String,
  "VpcConnectorArn" : String
}

```

### YAML

```yaml

  EgressType: String
  VpcConnectorArn: String

```

## Properties

`EgressType`

The type of egress configuration.

Set to `DEFAULT` for access to resources hosted on public networks.

Set to `VPC` to associate your service to a custom VPC specified by `VpcConnectorArn`.

_Required_: Yes

_Type_: String

_Allowed values_: `DEFAULT | VPC`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcConnectorArn`

The Amazon Resource Name (ARN) of the App Runner VPC connector that you want to associate with your App Runner service. Only valid when `EgressType =
        VPC`.

_Required_: No

_Type_: String

_Pattern_: `arn:aws(-[\w]+)*:[a-z0-9-\\.]{0,63}:[a-z0-9-\\.]{0,63}:[0-9]{12}:(\w|\/|-){1,1011}`

_Minimum_: `44`

_Maximum_: `1011`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CodeRepository

EncryptionConfiguration

All content copied from https://docs.aws.amazon.com/.
