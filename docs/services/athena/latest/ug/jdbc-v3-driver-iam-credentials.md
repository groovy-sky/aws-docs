# IAM credentials

You can use your IAM credentials with the JDBC driver to connect to Amazon Athena by
setting the following connection parameters.

## User

Your AWS access key ID. For information about access keys, see [AWS\
security credentials](../../../iam/latest/userguide/security-creds.md) in the
_IAM User Guide_.

Parameter nameAliasParameter typeDefault valueUserAccessKeyIdRequirednone

## Password

Your AWS secret key ID. For information about access keys, see [AWS\
security credentials](../../../iam/latest/userguide/security-creds.md) in the
_IAM User Guide_.

Parameter nameAliasParameter typeDefault valuePasswordSecretAccessKeyOptionalnone

## Session token

If you use temporary AWS credentials, you must specify a session token. For
information about temporary credentials, see [Temporary security\
credentials in IAM](../../../iam/latest/userguide/id-credentials-temp.md) in the
_IAM User Guide_.

Parameter nameAliasParameter typeDefault valueSessionTokennoneOptionalnone

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Authentication

Default

All content copied from https://docs.aws.amazon.com/.
