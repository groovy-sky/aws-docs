# IAM credentials

You can use your IAM credentials to connect to Amazon Athena with the ODBC driver using the
connection string parameters described in this section.

## Authentication type

**Connection string name****Parameter type****Default value****Connection string example**AuthenticationTypeRequired`IAM Credentials``AuthenticationType=IAM Credentials;`

## User ID

Your AWS Access Key ID. For more information about access keys, see [AWS security\
credentials](../../../iam/latest/userguide/security-creds.md) in the _IAM User Guide_.

**Connection string name****Parameter type****Default value****Connection string example**UIDRequired`none``UID=AWS_ACCESS_KEY_ID_REDACTED;`

## Password

Your AWS secret key id. For more information about access keys, see [AWS security\
credentials](../../../iam/latest/userguide/security-creds.md) in the _IAM User Guide_.

**Connection string name****Parameter type****Default value****Connection string example**PWDRequired`none``PWD=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKE;`

## Session token

If you use temporary AWS credentials, you must specify a session token. For
information about temporary credentials, see [Temporary security credentials\
in IAM](../../../iam/latest/userguide/id-credentials-temp.md) in the _IAM User Guide_.

**Connection string name****Parameter type****Default value****Connection string example**SessionTokenOptional`none``SessionToken=AQoDYXdzEJr...<remainder of session
                                token>;`

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Authentication

IAM profile

All content copied from https://docs.aws.amazon.com/.
