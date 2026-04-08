# MasterUserSecret

Contains the secret managed by RDS in AWS Secrets Manager for the master user password.

For more information, see [Password management with AWS Secrets Manager](../../../../services/amazonrds/latest/userguide/rds-secrets-manager.md)
in the _Amazon RDS User Guide_ and [Password management with AWS Secrets Manager](../../../../services/amazonrds/latest/aurorauserguide/rds-secrets-manager.md)
in the _Amazon Aurora User Guide._

## Contents

###### Note

In the following list, the required parameters are described first.

**KmsKeyId**

The AWS KMS key identifier that is used to encrypt the secret.

Type: String

Required: No

**SecretArn**

The Amazon Resource Name (ARN) of the secret.

Type: String

Required: No

**SecretStatus**

The status of the secret.

The possible status values include the following:

- `creating` \- The secret is being created.

- `active` \- The secret is available for normal use and rotation.

- `rotating` \- The secret is being rotated.

- `impaired` \- The secret can be used to access database credentials,
but it can't be rotated. A secret might have this status if, for example,
permissions are changed so that RDS can no longer access either the secret or
the KMS key for the secret.

When a secret has this status, you can correct the condition that caused the
status. Alternatively, modify the DB instance to turn off automatic management
of database credentials, and then modify the DB instance again to turn on
automatic management of database credentials.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/masterusersecret.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/masterusersecret.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/masterusersecret.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

LimitlessDatabase

Metric
