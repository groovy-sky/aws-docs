# UserAuthConfig

Specifies the details of authentication used by a proxy to log in as a specific database user.

## Contents

###### Note

In the following list, the required parameters are described first.

**AuthScheme**

The type of authentication that the proxy uses for connections from the proxy to the underlying database.

Type: String

Valid Values: `SECRETS`

Required: No

**ClientPasswordAuthType**

The type of authentication the proxy uses for connections from clients. The following values are defaults for the corresponding engines:

- RDS for MySQL: `MYSQL_CACHING_SHA2_PASSWORD`

- RDS for SQL Server: `SQL_SERVER_AUTHENTICATION`

- RDS for PostgreSQL: `POSTGRES_SCRAM_SHA2_256`

Type: String

Valid Values: `MYSQL_NATIVE_PASSWORD | MYSQL_CACHING_SHA2_PASSWORD | POSTGRES_SCRAM_SHA_256 | POSTGRES_MD5 | SQL_SERVER_AUTHENTICATION`

Required: No

**Description**

A user-specified description about the authentication used by a proxy to log in as a specific database user.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1000.

Pattern: `.*`

Required: No

**IAMAuth**

A value that indicates whether to require or disallow AWS Identity and Access Management (IAM) authentication for connections to the proxy.
The `ENABLED` value is valid only for proxies with RDS for Microsoft SQL Server.

Type: String

Valid Values: `DISABLED | REQUIRED | ENABLED`

Required: No

**SecretArn**

The Amazon Resource Name (ARN) representing the secret that the proxy uses to authenticate
to the RDS DB instance or Aurora DB cluster. These secrets are stored within Amazon Secrets Manager.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Required: No

**UserName**

The name of the database user to which the proxy connects.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/userauthconfig.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/userauthconfig.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/userauthconfig.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UpgradeTarget

UserAuthConfigInfo

All content copied from https://docs.aws.amazon.com/.
