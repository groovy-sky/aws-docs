This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::AutomatedReasoningPolicyVersion

Creates a new version of an existing Automated Reasoning policy. This allows you to iterate on your policy rules while maintaining previous versions for rollback or comparison purposes.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Bedrock::AutomatedReasoningPolicyVersion",
  "Properties" : {
      "LastUpdatedDefinitionHash" : String,
      "PolicyArn" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Bedrock::AutomatedReasoningPolicyVersion
Properties:
  LastUpdatedDefinitionHash: String
  PolicyArn: String
  Tags:
    - Tag

```

## Properties

`LastUpdatedDefinitionHash`

The hash of the policy definition that was last updated.

_Required_: No

_Type_: String

_Pattern_: `^[0-9a-z]{128}$`

_Minimum_: `128`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PolicyArn`

The Amazon Resource Name (ARN) of the policy.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws(-[^:]+)?:bedrock:[a-z0-9-]{1,20}:[0-9]{12}:automated-reasoning-policy\/[a-z0-9]{12}$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags associated with the Automated Reasoning policy version.

_Required_: No

_Type_: Array of [Tag](aws-properties-bedrock-automatedreasoningpolicyversion-tag.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

Returns the policy version ID.

### Fn::GetAtt

Returns the value of an attribute from a resource in the template.

`CreatedAt`

The timestamp when the policy version was created.

`DefinitionHash`

A hash of the policy definition used to identify the version.

`Description`

The description of the policy version.

`Name`

The name of the policy version.

`PolicyId`

The unique identifier of the policy.

`UpdatedAt`

The timestamp when the policy version was last updated.

`Version`

The version number of the policy version.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
