This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Neptune::DBCluster DBClusterRole

Describes an Amazon Identity and Access Management (IAM) role that is associated with a DB
cluster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FeatureName" : String,
  "RoleArn" : String
}

```

### YAML

```yaml

  FeatureName: String
  RoleArn: String

```

## Properties

`FeatureName`

The name of the feature associated with the Amazon Identity and Access Management (IAM) role.
For the list of supported feature names, see [DescribeDBEngineVersions](../../../neptune/latest/userguide/api-other-apis.md#DescribeDBEngineVersions).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The Amazon Resource Name (ARN) of the IAM role that is associated with the DB
cluster.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Neptune::DBCluster

ServerlessScalingConfiguration

All content copied from https://docs.aws.amazon.com/.
