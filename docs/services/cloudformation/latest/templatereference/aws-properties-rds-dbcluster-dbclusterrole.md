This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RDS::DBCluster DBClusterRole

Describes an AWS Identity and Access Management (IAM) role that is associated with a DB cluster.

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

The name of the feature associated with the AWS Identity and Access Management (IAM)
role. IAM roles that are associated with a DB cluster grant permission for the DB
cluster to access other AWS services on your behalf. For the list of supported feature
names, see the `SupportedFeatureNames` description in [DBEngineVersion](../../../../reference/amazonrds/latest/apireference/api-dbengineversion.md)
in the _Amazon RDS API Reference_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The Amazon Resource Name (ARN) of the IAM role that is associated with the DB cluster.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

The following specifies a role to associate with a DB cluster.

#### JSON

```json

"AssociatedRoles": [
    {
        "FeatureName": "s3Import",
        "RoleArn": "arn:aws:iam::123456789012:role/RDSLoadFromS3"
    }
]
```

#### YAML

```yaml

AssociatedRoles:
    - FeatureName: s3Import
      RoleArn: 'arn:aws:iam::123456789012:role/RDSLoadFromS3'
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::RDS::DBCluster

Endpoint
