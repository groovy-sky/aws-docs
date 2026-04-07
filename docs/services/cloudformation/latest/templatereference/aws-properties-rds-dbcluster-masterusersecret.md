This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RDS::DBCluster MasterUserSecret

The `MasterUserSecret` return value specifies the secret managed by RDS in AWS Secrets Manager for the master user password.

For more information, see [Password management with AWS Secrets Manager](../../../amazonrds/latest/userguide/rds-secrets-manager.md)
in the _Amazon RDS User Guide_ and [Password management with AWS Secrets Manager](../../../amazonrds/latest/aurorauserguide/rds-secrets-manager.md)
in the _Amazon Aurora User Guide._

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KmsKeyId" : String,
  "SecretArn" : String
}

```

### YAML

```yaml

  KmsKeyId: String
  SecretArn: String

```

## Properties

`KmsKeyId`

The AWS KMS key identifier that is used to encrypt the secret.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretArn`

The Amazon Resource Name (ARN) of the secret. This parameter is a return value that you can retrieve using the `Fn::GetAtt`
intrinsic function. For more information, see [Return values](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#aws-resource-rds-dbcluster-return-values).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Endpoint

ReadEndpoint
