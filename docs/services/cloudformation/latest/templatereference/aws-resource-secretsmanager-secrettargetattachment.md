This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecretsManager::SecretTargetAttachment

The `AWS::SecretsManager::SecretTargetAttachment` resource completes the final
link between a Secrets Manager secret and the associated database by adding the database
connection information to the secret JSON. If you want to turn on automatic rotation
for a database credential secret, the secret must contain the database connection information.
For more information, see [JSON structure \
of Secrets Manager database credential secrets](../../../secretsmanager/latest/userguide/reference-secret-json-structure.md).

A single secret resource can only have one target attached to it.

When you remove a `SecretTargetAttachment` from a stack, Secrets Manager removes the database connection information from the secret with a `PutSecretValue` call.

For Amazon RDS master user credentials, see [AWS::RDS::DBCluster MasterUserSecret](../userguide/aws-properties-rds-dbcluster-masterusersecret.md).

For Amazon Redshift admin user credentials, see [AWS::Redshift::Cluster](../userguide/aws-resource-redshift-cluster.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SecretsManager::SecretTargetAttachment",
  "Properties" : {
      "SecretId" : String,
      "TargetId" : String,
      "TargetType" : String
    }
}

```

### YAML

```yaml

Type: AWS::SecretsManager::SecretTargetAttachment
Properties:
  SecretId: String
  TargetId: String
  TargetType: String

```

## Properties

`SecretId`

The ARN or name of the secret. To reference a secret also created in this template, use
the see [Ref](../userguide/intrinsic-function-reference-ref.md)
function with the secret's logical ID. This field is unique for each target attachment definition.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TargetId`

The ID of the database or cluster.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetType`

A string that defines the type of service or database associated with the secret. This
value instructs Secrets Manager how to update the secret with the details of the service or
database. This value must be one of the following:

- AWS::RDS::DBInstance

- AWS::RDS::DBCluster

- AWS::Redshift::Cluster

- AWS::RedshiftServerless::Namespace

- AWS::DocDB::DBInstance

- AWS::DocDB::DBCluster

- AWS::DocDBElastic::Cluster

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of an
`AWS::SecretsManager::SecretTargetAttachment` resource to the intrinsic
`Ref` function, the function returns the ARN of the secret, such as:

`arn:aws:secretsmanager:us-west-2:123456789012:secret:my-path/my-secret-name-1a2b3c`

You can use the ARN to reference a secret you created in one part of the stack template
from within the definition of another resource from a different part of the same
template.

For more information about using the `Ref` function, see [Ref](../userguide/intrinsic-function-reference-ref.md).

### Fn::GetAtt

## See also

- [AWS::SecretsManager::Secret](../userguide/aws-resource-secretsmanager-secret.md)

- [AWS::SecretsManager::RotationSchedule](../userguide/aws-resource-secretsmanager-rotationschedule.md)

- [AWS::SecretsManager::ResourcePolicy](../userguide/aws-resource-secretsmanager-resourcepolicy.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Next

All content copied from https://docs.aws.amazon.com/.
