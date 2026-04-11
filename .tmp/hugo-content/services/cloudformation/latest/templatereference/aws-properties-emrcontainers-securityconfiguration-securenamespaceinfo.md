This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMRContainers::SecurityConfiguration SecureNamespaceInfo

Namespace inputs for the system job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClusterId" : String,
  "Namespace" : String
}

```

### YAML

```yaml

  ClusterId: String
  Namespace: String

```

## Properties

`ClusterId`

The ID of the Amazon EKS cluster where Amazon EMR on EKS jobs run.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Namespace`

The namespace of the Amazon EKS cluster where the system jobs run.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3EncryptionConfiguration

SecurityConfigurationData

All content copied from https://docs.aws.amazon.com/.
