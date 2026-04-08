# UserAuthConfigInfo

Returns the details of authentication used by a proxy to log in as a specific database user.

## Contents

###### Note

In the following list, the required parameters are described first.

**AuthScheme**

The type of authentication that the proxy uses for connections from the proxy to the underlying database.

Type: String

Valid Values: `SECRETS`

Required: No

**ClientPasswordAuthType**

The type of authentication the proxy uses for connections from clients.

Type: String

Valid Values: `MYSQL_NATIVE_PASSWORD | MYSQL_CACHING_SHA2_PASSWORD | POSTGRES_SCRAM_SHA_256 | POSTGRES_MD5 | SQL_SERVER_AUTHENTICATION`

Required: No

**Description**

A user-specified description about the authentication used by a proxy to log in as a specific database user.

Type: String

Required: No

**IAMAuth**

Whether to require or disallow AWS Identity and Access Management (IAM) authentication for connections to the proxy.

Type: String

Valid Values: `DISABLED | REQUIRED | ENABLED`

Required: No

**SecretArn**

The Amazon Resource Name (ARN) representing the secret that the proxy uses to authenticate
to the RDS DB instance or Aurora DB cluster. These secrets are stored within Amazon Secrets Manager.

Type: String

Required: No

**UserName**

The name of the database user to which the proxy connects.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/userauthconfiginfo.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/userauthconfiginfo.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/userauthconfiginfo.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UserAuthConfig

ValidAdditionalStorageOptions

All content copied from https://docs.aws.amazon.com/.
