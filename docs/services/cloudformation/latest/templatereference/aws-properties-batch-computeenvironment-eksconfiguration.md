This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::ComputeEnvironment EksConfiguration

Configuration for the Amazon EKS cluster that supports the AWS Batch compute environment. The
cluster must exist before the compute environment can be created.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EksClusterArn" : String,
  "KubernetesNamespace" : String
}

```

### YAML

```yaml

  EksClusterArn: String
  KubernetesNamespace: String

```

## Properties

`EksClusterArn`

The Amazon Resource Name (ARN) of the Amazon EKS cluster. An example is
`arn:aws:eks:us-east-1:123456789012:cluster/ClusterForBatch`.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KubernetesNamespace`

The namespace of the Amazon EKS cluster. AWS Batch manages pods in this namespace. The value
can't left empty or null. It must be fewer than 64 characters long, can't be set to
`default`, can't start with " `kube-`," and must match this regular
expression: `^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`. For more information, see [Namespaces](https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces) in the Kubernetes documentation.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Ec2ConfigurationObject

LaunchTemplateSpecification
