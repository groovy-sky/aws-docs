This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ManagedBlockchain::Member MemberConfiguration

Configuration properties of the member.

Applies only to Hyperledger Fabric.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Description" : String,
  "MemberFrameworkConfiguration" : MemberFrameworkConfiguration,
  "Name" : String
}

```

### YAML

```yaml

  Description: String
  MemberFrameworkConfiguration:
    MemberFrameworkConfiguration
  Name: String

```

## Properties

`Description`

An optional description of the member.

_Required_: No

_Type_: String

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MemberFrameworkConfiguration`

Configuration properties of the blockchain framework relevant to the member.

_Required_: No

_Type_: [MemberFrameworkConfiguration](aws-properties-managedblockchain-member-memberframeworkconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the member.

_Required_: Yes

_Type_: String

_Pattern_: `^(?!-|[0-9])(?!.*-$)(?!.*?--)[a-zA-Z0-9-]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ApprovalThresholdPolicy

MemberFabricConfiguration

All content copied from https://docs.aws.amazon.com/.
