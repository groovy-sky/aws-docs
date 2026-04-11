This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECR::ReplicationConfiguration ReplicationDestination

An array of objects representing the destination for a replication rule.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Region" : String,
  "RegistryId" : String
}

```

### YAML

```yaml

  Region: String
  RegistryId: String

```

## Properties

`Region`

The Region to replicate to.

_Required_: Yes

_Type_: String

_Pattern_: `[0-9a-z-]{2,25}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RegistryId`

The AWS account ID of the Amazon ECR private registry to replicate to. When configuring
cross-Region replication within your own registry, specify your own account ID.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9]{12}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReplicationConfiguration

ReplicationRule

All content copied from https://docs.aws.amazon.com/.
