This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::GatewayTarget

After creating a gateway, you can add targets, which define the tools that your
gateway will host.

For more information about adding gateway targets, see [Add\
targets to an existing gateway](https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/gateway-building-adding-targets.html).

See the **Properties** section below for descriptions of
both the required and optional properties.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::BedrockAgentCore::GatewayTarget",
  "Properties" : {
      "CredentialProviderConfigurations" : [ CredentialProviderConfiguration, ... ],
      "Description" : String,
      "GatewayIdentifier" : String,
      "MetadataConfiguration" : MetadataConfiguration,
      "Name" : String,
      "TargetConfiguration" : TargetConfiguration
    }
}

```

### YAML

```yaml

Type: AWS::BedrockAgentCore::GatewayTarget
Properties:
  CredentialProviderConfigurations:
    - CredentialProviderConfiguration
  Description: String
  GatewayIdentifier: String
  MetadataConfiguration:
    MetadataConfiguration
  Name: String
  TargetConfiguration:
    TargetConfiguration

```

## Properties

`CredentialProviderConfigurations`

The credential provider configurations.

_Required_: No

_Type_: Array of [CredentialProviderConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrockagentcore-gatewaytarget-credentialproviderconfiguration.html)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description for the gateway target.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GatewayIdentifier`

The gateway ID for the gateway target.

_Required_: No

_Type_: String

_Pattern_: `^([0-9a-z][-]?){1,100}-[0-9a-z]{10}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MetadataConfiguration`

The metadata configuration for HTTP header and query parameter propagation to and from this gateway target.

_Required_: No

_Type_: [MetadataConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrockagentcore-gatewaytarget-metadataconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the gateway target.

_Required_: Yes

_Type_: String

_Pattern_: `^([0-9a-zA-Z][-]?){1,100}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetConfiguration`

The configuration for a gateway target. This structure defines how the gateway connects to and interacts with the target endpoint.

_Required_: Yes

_Type_: [TargetConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrockagentcore-gatewaytarget-targetconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the gateway identifier and target ID. For example:

`my-gateway-a1b2c3d4e5|a1B2c3D4e5`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedAt`

The date and time at which the target was created.

`GatewayArn`

The Amazon Resource Name (ARN) of the gateway target.

`LastSynchronizedAt`

The timestamp when the target was last synchronized.

`Status`

The status of the gateway target.

`StatusReasons`

The status reasons for the target status.

`TargetId`

The target ID.

`UpdatedAt`

The date and time at which the target was updated.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

WorkloadIdentityDetails

ApiGatewayTargetConfiguration
