This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeDeploy::DeploymentGroup ECSService

Contains the service and cluster names used to identify an Amazon ECS
deployment's target.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClusterName" : String,
  "ServiceName" : String
}

```

### YAML

```yaml

  ClusterName: String
  ServiceName: String

```

## Properties

`ClusterName`

The name of the cluster that the Amazon ECS service is associated with.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceName`

The name of the target Amazon ECS service.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

EC2TagSetListObject

ELBInfo
