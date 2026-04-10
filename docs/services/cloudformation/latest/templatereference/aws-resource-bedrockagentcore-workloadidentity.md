This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::WorkloadIdentity

Creates a workload identity for Amazon Bedrock AgentCore. A workload identity provides
OAuth2-based authentication for resources associated with agent runtimes.

For more information about using workload identities in Amazon Bedrock AgentCore, see
[Managing workload\
identities](../../../bedrock-agentcore/latest/devguide/workload-identity.md).

See the **Properties** section below for descriptions of
both the required and optional properties.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::BedrockAgentCore::WorkloadIdentity",
  "Properties" : {
      "AllowedResourceOauth2ReturnUrls" : [ String, ... ],
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::BedrockAgentCore::WorkloadIdentity
Properties:
  AllowedResourceOauth2ReturnUrls:
    - String
  Name: String
  Tags:
    - Tag

```

## Properties

`AllowedResourceOauth2ReturnUrls`

The list of allowed OAuth2 return URLs for resources associated with this workload
identity.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the workload identity. The name must be unique within your account.

_Required_: Yes

_Type_: String

_Pattern_: `[A-Za-z0-9_.-]+`

_Minimum_: `3`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags for the workload identity.

_Required_: No

_Type_: Array of [Tag](aws-properties-bedrockagentcore-workloadidentity-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref`
function, `Ref` returns the workload identity name.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedTime`

The timestamp when the workload identity was created.

`LastUpdatedTime`

The timestamp when the workload identity was last updated.

`WorkloadIdentityArn`

The Amazon Resource Name (ARN) of the workload identity.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::BedrockAgentCore::RuntimeEndpoint

Tag

All content copied from https://docs.aws.amazon.com/.
