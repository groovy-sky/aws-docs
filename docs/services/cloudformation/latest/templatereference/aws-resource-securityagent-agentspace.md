This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityAgent::AgentSpace

The `AWS::SecurityAgent::AgentSpace` resource specifies an agent space for security testing. An agent space defines the scope of resources, integrations, and settings available to security testing operations.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SecurityAgent::AgentSpace",
  "Properties" : {
      "AwsResources" : AWSResources,
      "CodeReviewSettings" : CodeReviewSettings,
      "Description" : String,
      "IntegratedResources" : [ IntegratedResource, ... ],
      "KmsKeyId" : String,
      "Name" : String,
      "Tags" : [ Tag, ... ],
      "TargetDomainIds" : [ String, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::SecurityAgent::AgentSpace
Properties:
  AwsResources:
    AWSResources
  CodeReviewSettings:
    CodeReviewSettings
  Description: String
  IntegratedResources:
    - IntegratedResource
  KmsKeyId: String
  Name: String
  Tags:
    - Tag
  TargetDomainIds:
    - String

```

## Properties

`AwsResources`

The Amazon Web Services resources to associate with the agent space.

_Required_: No

_Type_: [AWSResources](aws-properties-securityagent-agentspace-awsresources.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CodeReviewSettings`

The code review settings for the agent space.

_Required_: No

_Type_: [CodeReviewSettings](aws-properties-securityagent-agentspace-codereviewsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the agent space.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IntegratedResources`

The list of integrated resource items to update.

_Required_: No

_Type_: Array of [IntegratedResource](aws-properties-securityagent-agentspace-integratedresource.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyId`

The identifier of the Amazon Web Services KMS key to use for encrypting data in the agent space.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the agent space.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to associate with the agent space.

_Required_: No

_Type_: Array of [Tag](aws-properties-securityagent-agentspace-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetDomainIds`

The list of target domain identifiers to associate with the agent space.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the agent space ID. For example:

`{ "Ref": "MyAgentSpace" }`

For the agent space `MyAgentSpace`, `Ref` returns the unique identifier of the agent space.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AgentSpaceId`

The unique identifier of the agent space. For example: `as-0123456789abcdef0`.

`CreatedAt`

The date and time when the agent space was created, in ISO 8601 format. For example: `2024-01-01T00:00:00Z`.

`UpdatedAt`

The date and time when the agent space was last updated, in ISO 8601 format. For example: `2024-01-01T00:00:00Z`.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Security Agent

AWSResources

All content copied from https://docs.aws.amazon.com/.
