This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EKS::Cluster EncryptionConfig

The encryption configuration for the cluster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Provider" : Provider,
  "Resources" : [ String, ... ]
}

```

### YAML

```yaml

  Provider:
    Provider
  Resources:
    - String

```

## Properties

`Provider`

The encryption provider for the cluster.

_Required_: No

_Type_: [Provider](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-eks-cluster-provider.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Resources`

Specifies the resources to be encrypted. The only supported value is
`secrets`.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ElasticLoadBalancing

KubernetesNetworkConfig
