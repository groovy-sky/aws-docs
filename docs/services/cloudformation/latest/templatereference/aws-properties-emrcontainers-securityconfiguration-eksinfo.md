This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMRContainers::SecurityConfiguration EksInfo

The information about the Amazon EKS cluster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Namespace" : String
}

```

### YAML

```yaml

  Namespace: String

```

## Properties

`Namespace`

The namespaces of the Amazon EKS cluster.

_Required_: No

_Type_: String

_Pattern_: `[a-z0-9]([-a-z0-9]*[a-z0-9])?`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ContainerProvider

EncryptionConfiguration

All content copied from https://docs.aws.amazon.com/.
