This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::Assistant

Specifies an Amazon Connect Wisdom assistant.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Wisdom::Assistant",
  "Properties" : {
      "Description" : String,
      "Name" : String,
      "ServerSideEncryptionConfiguration" : ServerSideEncryptionConfiguration,
      "Tags" : [ Tag, ... ],
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::Wisdom::Assistant
Properties:
  Description: String
  Name: String
  ServerSideEncryptionConfiguration:
    ServerSideEncryptionConfiguration
  Tags:
    - Tag
  Type: String

```

## Properties

`Description`

The description of the assistant.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the assistant.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ServerSideEncryptionConfiguration`

The configuration information for the customer managed key used for encryption. The
customer managed key must have a policy that allows `kms:CreateGrant`
and `kms:DescribeKey` permissions to the IAM identity using
the key to invoke Wisdom. To use Wisdom with chat, the key policy must also allow
`kms:Decrypt`, `kms:GenerateDataKey*`, and
`kms:DescribeKey` permissions to the `connect.amazonaws.com`
service principal. For more information about setting up a customer managed key for
Wisdom, see [Enable Amazon Connect Wisdom\
for your instance](../../../connect/latest/adminguide/enable-wisdom.md).

_Required_: No

_Type_: [ServerSideEncryptionConfiguration](aws-properties-wisdom-assistant-serversideencryptionconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags used to organize, track, or control access for this resource.

_Required_: No

_Type_: Array of [Tag](aws-properties-wisdom-assistant-tag.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Type`

The type of assistant.

_Required_: Yes

_Type_: String

_Allowed values_: `AGENT`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the assistant ID.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AssistantArn`

The Amazon Resource Name (ARN) of the assistant.

`AssistantId`

The ID of the Wisdom assistant.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Wisdom::AIPromptVersion

ServerSideEncryptionConfiguration

All content copied from https://docs.aws.amazon.com/.
