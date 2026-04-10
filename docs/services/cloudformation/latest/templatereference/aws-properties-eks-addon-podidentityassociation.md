This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EKS::Addon PodIdentityAssociation

Amazon EKS Pod Identity associations provide the ability to manage credentials for your applications, similar to the way that Amazon EC2 instance profiles provide credentials to Amazon EC2 instances.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RoleArn" : String,
  "ServiceAccount" : String
}

```

### YAML

```yaml

  RoleArn: String
  ServiceAccount: String

```

## Properties

`RoleArn`

The Amazon Resource Name (ARN) of the IAM role to associate with the service account. The EKS Pod Identity
agent manages credentials to assume this role for applications in the containers in the
Pods that use this service account.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws(-cn|-us-gov|-iso(-[a-z])?)?:iam::\d{12}:(role)\/*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceAccount`

The name of the Kubernetes service account inside the cluster to associate the IAM
credentials with.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NamespaceConfig

Tag

All content copied from https://docs.aws.amazon.com/.
