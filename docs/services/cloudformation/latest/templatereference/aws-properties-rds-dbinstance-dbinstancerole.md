This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RDS::DBInstance DBInstanceRole

Information about an AWS Identity and Access Management (IAM) role that is associated with a DB instance.

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
role. IAM roles that are associated with a DB instance grant permission for the DB
instance to access other AWS services on your behalf. For the list of supported feature
names, see the `SupportedFeatureNames` description in [DBEngineVersion](../../../../reference/amazonrds/latest/apireference/api-dbengineversion.md)
in the _Amazon RDS API Reference_.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The Amazon Resource Name (ARN) of the IAM role that is associated with the DB
instance.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

The following specifies a role to associate with a DB instance.

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

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CertificateDetails

DBInstanceStatusInfo

All content copied from https://docs.aws.amazon.com/.
