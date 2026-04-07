This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DAX::Cluster SSESpecification

Represents the settings used to enable server-side encryption.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SSEEnabled" : Boolean
}

```

### YAML

```yaml

  SSEEnabled: Boolean

```

## Properties

`SSEEnabled`

Indicates whether server-side encryption is enabled (true) or disabled (false) on
the cluster.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::DAX::Cluster

AWS::DAX::ParameterGroup
