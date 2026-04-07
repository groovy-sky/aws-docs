# IAM credentials

You can use your IAM credentials with the JDBC driver to connect to Amazon Athena by
setting the following connection parameters.

## User

Your AWS access key ID. For information about access keys, see [AWS\
security credentials](https://docs.aws.amazon.com/IAM/latest/UserGuide/security-creds.html) in the
_IAM User Guide_.

Parameter nameAliasParameter typeDefault valueUserAccessKeyIdRequirednone

## Password

Your AWS secret key ID. For information about access keys, see [AWS\
security credentials](https://docs.aws.amazon.com/IAM/latest/UserGuide/security-creds.html) in the
_IAM User Guide_.

Parameter nameAliasParameter typeDefault valuePasswordSecretAccessKeyOptionalnone

## Session token

If you use temporary AWS credentials, you must specify a session token. For
information about temporary credentials, see [Temporary security\
credentials in IAM](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp.html) in the
_IAM User Guide_.

Parameter nameAliasParameter typeDefault valueSessionTokennoneOptionalnone

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Authentication

Default
