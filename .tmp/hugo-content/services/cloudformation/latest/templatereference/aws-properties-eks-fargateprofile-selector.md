This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EKS::FargateProfile Selector

An object representing an AWS Fargate profile selector.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Labels" : [ Label, ... ],
  "Namespace" : String
}

```

### YAML

```yaml

  Labels:
    - Label
  Namespace: String

```

## Properties

`Labels`

The Kubernetes labels that the selector should match. A pod must contain all of the labels
that are specified in the selector for it to be considered a match.

_Required_: No

_Type_: Array of [Label](aws-properties-eks-fargateprofile-label.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Namespace`

The Kubernetes `namespace` that the selector should match.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Label

Tag

All content copied from https://docs.aws.amazon.com/.
