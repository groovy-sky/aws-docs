This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecretsManager::RotationSchedule HostedRotationLambda

Creates a new Lambda rotation
function based on one of the [Secrets Manager rotation function templates](../../../secretsmanager/latest/userguide/reference-available-rotation-templates.md).

You must specify `Transform:
        AWS::SecretsManager-2024-09-16` at the beginning of the CloudFormation
template.

For Amazon RDS master user credentials, see [AWS::RDS::DBCluster MasterUserSecret](../userguide/aws-properties-rds-dbcluster-masterusersecret.md).

For Amazon Redshift admin user credentials, see [AWS::Redshift::Cluster](../userguide/aws-resource-redshift-cluster.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ExcludeCharacters" : String,
  "KmsKeyArn" : String,
  "MasterSecretArn" : String,
  "MasterSecretKmsKeyArn" : String,
  "RotationLambdaName" : String,
  "RotationType" : String,
  "Runtime" : String,
  "SuperuserSecretArn" : String,
  "SuperuserSecretKmsKeyArn" : String,
  "VpcSecurityGroupIds" : String,
  "VpcSubnetIds" : String
}

```

### YAML

```yaml

  ExcludeCharacters: String
  KmsKeyArn: String
  MasterSecretArn: String
  MasterSecretKmsKeyArn: String
  RotationLambdaName: String
  RotationType: String
  Runtime: String
  SuperuserSecretArn: String
  SuperuserSecretKmsKeyArn: String
  VpcSecurityGroupIds: String
  VpcSubnetIds: String

```

## Properties

`ExcludeCharacters`

A string of the characters that you don't want in the password.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyArn`

The ARN of the KMS key that Secrets Manager uses to encrypt the secret. If you don't
specify this value, then Secrets Manager uses the key `aws/secretsmanager`. If
`aws/secretsmanager` doesn't yet exist, then Secrets Manager creates it for you
automatically the first time it encrypts the secret value.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MasterSecretArn`

The ARN of the secret that contains superuser credentials, if you use the
[Alternating users rotation strategy](../../../secretsmanager/latest/userguide/rotating-secrets-strategies.md#rotating-secrets-two-users). CloudFormation grants the execution role for the Lambda rotation function `GetSecretValue` permission to the secret in this property. For more information, see [Lambda rotation function execution role permissions for Secrets Manager](../../../secretsmanager/latest/userguide/rotating-secrets-required-permissions-function.md).

You must create the superuser secret before you can set this property.

You must also include the superuser secret ARN as a key in the JSON of the rotating secret so that the Lambda rotation function can find it. CloudFormation does not hardcode secret ARNs in the Lambda rotation function, so you can use the function to rotate multiple secrets. For more information, see [JSON structure of Secrets Manager secrets](../../../secretsmanager/latest/userguide/reference-secret-json-structure.md).

You can specify `MasterSecretArn` or `SuperuserSecretArn` but not both. They represent the same superuser secret.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MasterSecretKmsKeyArn`

The ARN of the KMS key that Secrets Manager used to encrypt the superuser secret, if
you use the [alternating users strategy](../../../secretsmanager/latest/userguide/rotating-secrets-strategies.md#rotating-secrets-two-users) and the superuser secret is encrypted with a customer managed key. You don't need to specify this property if the superuser secret is encrypted using the key `aws/secretsmanager`. CloudFormation grants the execution role for the Lambda rotation function `Decrypt`, `DescribeKey`, and `GenerateDataKey` permission to the key in this property. For more information, see [Lambda rotation function execution role permissions for Secrets Manager](../../../secretsmanager/latest/userguide/rotating-secrets-required-permissions-function.md).

You can specify `MasterSecretKmsKeyArn` or `SuperuserSecretKmsKeyArn` but not both. They represent the same superuser secret KMS key.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RotationLambdaName`

The name of the Lambda rotation function.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RotationType`

The rotation template to base the rotation function on, one of the following:

- `Db2SingleUser` to use the template [SecretsManagerRDSDb2RotationSingleUser](../../../secretsmanager/latest/userguide/reference-available-rotation-templates.md#sar-template-db2-singleuser).

- `Db2MultiUser` to use the template [SecretsManagerRDSDb2RotationMultiUser](../../../secretsmanager/latest/userguide/reference-available-rotation-templates.md#sar-template-db2-multiuser).

- `MySQLSingleUser` to use the template [SecretsManagerRDSMySQLRotationSingleUser](../../../secretsmanager/latest/userguide/reference-available-rotation-templates.md#sar-template-mysql-singleuser).

- `MySQLMultiUser` to use the template [SecretsManagerRDSMySQLRotationMultiUser](../../../secretsmanager/latest/userguide/reference-available-rotation-templates.md#sar-template-mysql-multiuser).

- `PostgreSQLSingleUser` to use the template [SecretsManagerRDSPostgreSQLRotationSingleUser](../../../secretsmanager/latest/userguide/reference-available-rotation-templates.md#sar-template-postgre-singleuser)

- `PostgreSQLMultiUser` to use the template [SecretsManagerRDSPostgreSQLRotationMultiUser](../../../secretsmanager/latest/userguide/reference-available-rotation-templates.md#sar-template-postgre-multiuser).

- `OracleSingleUser` to use the template [SecretsManagerRDSOracleRotationSingleUser](../../../secretsmanager/latest/userguide/reference-available-rotation-templates.md#sar-template-oracle-singleuser).

- `OracleMultiUser` to use the template [SecretsManagerRDSOracleRotationMultiUser](../../../secretsmanager/latest/userguide/reference-available-rotation-templates.md#sar-template-oracle-multiuser).

- `MariaDBSingleUser` to use the template [SecretsManagerRDSMariaDBRotationSingleUser](../../../secretsmanager/latest/userguide/reference-available-rotation-templates.md#sar-template-mariadb-singleuser).

- `MariaDBMultiUser` to use the template [SecretsManagerRDSMariaDBRotationMultiUser](../../../secretsmanager/latest/userguide/reference-available-rotation-templates.md#sar-template-mariadb-multiuser).

- `SQLServerSingleUser` to use the template [SecretsManagerRDSSQLServerRotationSingleUser](../../../secretsmanager/latest/userguide/reference-available-rotation-templates.md#sar-template-sqlserver-singleuser).

- `SQLServerMultiUser` to use the template [SecretsManagerRDSSQLServerRotationMultiUser](../../../secretsmanager/latest/userguide/reference-available-rotation-templates.md#sar-template-sqlserver-multiuser).

- `RedshiftSingleUser` to use the template [SecretsManagerRedshiftRotationSingleUsr](../../../secretsmanager/latest/userguide/reference-available-rotation-templates.md#sar-template-redshift-singleuser).

- `RedshiftMultiUser` to use the template [SecretsManagerRedshiftRotationMultiUser](../../../secretsmanager/latest/userguide/reference-available-rotation-templates.md#sar-template-redshift-multiuser).

- `MongoDBSingleUser` to use the template [SecretsManagerMongoDBRotationSingleUser](../../../secretsmanager/latest/userguide/reference-available-rotation-templates.md#sar-template-mongodb-singleuser).

- `MongoDBMultiUser` to use the template [SecretsManagerMongoDBRotationMultiUser](../../../secretsmanager/latest/userguide/reference-available-rotation-templates.md#sar-template-mongodb-multiuser).

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Runtime`

###### Important

Do not set this value if you are using `Transform: AWS::SecretsManager-2024-09-16`. Over time, the updated rotation lambda artifacts vended by AWS may not be compatible with the code or shared object files defined in the rotation function deployment package.

Only define the `Runtime` key if:

1. You are using `Transform: AWS::SecretsManager-2020-07-23`.

2. The code or shared object files defined in the rotation function deployment package are incompatible with Python 3.10.

The Python Runtime version for with the rotation function. By default, CloudFormation deploys Python 3.10 binaries for the rotation function. To use a different version of Python, you must do the following two steps:

1. Deploy the matching version Python binaries with your rotation function.

2. Set the version number in this field. For example, for Python 3.10, enter **python3.10**.

If you only do one of the steps, your rotation function will be incompatible with the binaries. For more information, see [Why did my Lambda rotation function fail with a "pg module not found" error](https://repost.aws/knowledge-center/secrets-manager-lambda-rotation).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SuperuserSecretArn`

The ARN of the secret that contains superuser credentials, if you use the
[Alternating users rotation strategy](../../../secretsmanager/latest/userguide/rotating-secrets-strategies.md#rotating-secrets-two-users). CloudFormation grants the execution role for the Lambda rotation function `GetSecretValue` permission to the secret in this property. For more information, see [Lambda rotation function execution role permissions for Secrets Manager](../../../secretsmanager/latest/userguide/rotating-secrets-required-permissions-function.md).

You must create the superuser secret before you can set this property.

You must also include the superuser secret ARN as a key in the JSON of the rotating secret so that the Lambda rotation function can find it. CloudFormation does not hardcode secret ARNs in the Lambda rotation function, so you can use the function to rotate multiple secrets. For more information, see [JSON structure of Secrets Manager secrets](../../../secretsmanager/latest/userguide/reference-secret-json-structure.md).

You can specify `MasterSecretArn` or `SuperuserSecretArn` but not both. They represent the same superuser secret.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SuperuserSecretKmsKeyArn`

The ARN of the KMS key that Secrets Manager used to encrypt the superuser secret, if
you use the [alternating users strategy](../../../secretsmanager/latest/userguide/rotating-secrets-strategies.md#rotating-secrets-two-users) and the superuser secret is encrypted with a customer managed key. You don't need to specify this property if the superuser secret is encrypted using the key `aws/secretsmanager`. CloudFormation grants the execution role for the Lambda rotation function `Decrypt`, `DescribeKey`, and `GenerateDataKey` permission to the key in this property. For more information, see [Lambda rotation function execution role permissions for Secrets Manager](../../../secretsmanager/latest/userguide/rotating-secrets-required-permissions-function.md).

You can specify `MasterSecretKmsKeyArn` or `SuperuserSecretKmsKeyArn` but not both. They represent the same superuser secret KMS key.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcSecurityGroupIds`

A comma-separated list of security group IDs applied to the target database.

The template applies the same security groups as on the Lambda rotation function that is
created as part of this stack.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcSubnetIds`

A comma separated list of VPC subnet IDs of the target database network. The Lambda
rotation function is in the same subnet group.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExternalSecretRotationMetadataItem

RotationRules

All content copied from https://docs.aws.amazon.com/.
