This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RDS::DBProxy AuthFormat

Specifies the details of authentication used by a proxy to log in as a specific database user.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthScheme" : String,
  "ClientPasswordAuthType" : String,
  "Description" : String,
  "IAMAuth" : String,
  "SecretArn" : String
}

```

### YAML

```yaml

  AuthScheme: String
  ClientPasswordAuthType: String
  Description: String
  IAMAuth: String
  SecretArn: String

```

## Properties

`AuthScheme`

The type of authentication that the proxy uses for connections from the proxy to the underlying database.

_Required_: No

_Type_: String

_Allowed values_: `SECRETS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientPasswordAuthType`

Specifies the details of authentication used by a proxy to log in as a specific database user.

_Required_: No

_Type_: String

_Allowed values_: `MYSQL_NATIVE_PASSWORD | MYSQL_CACHING_SHA2_PASSWORD | POSTGRES_SCRAM_SHA_256 | POSTGRES_MD5 | SQL_SERVER_AUTHENTICATION`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A user-specified description about the authentication used by a proxy to log in as a specific database user.

_Required_: No

_Type_: String

_Pattern_: `.*`

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IAMAuth`

A value that indicates whether to require or disallow AWS Identity and Access Management (IAM) authentication for connections to the proxy.
The `ENABLED` value is valid only for proxies with RDS for Microsoft SQL Server.

_Required_: No

_Type_: String

_Allowed values_: `DISABLED | REQUIRED | ENABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretArn`

The Amazon Resource Name (ARN) representing the secret that the proxy uses to authenticate
to the RDS DB instance or Aurora DB cluster. These secrets are stored within Amazon Secrets Manager.

_Required_: No

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

The following example specifies authentication details for a proxy.

### Authentication details

#### JSON

```json

"ProcessorFeatures":[
   {
      "AuthScheme": "SECRETS",
      "ClientPasswordAuthType": "MYSQL_NATIVE_PASSWORD",
      "Description": "Proxy authentication for MySQL",
      "IAMAuth": "DISABLED",
      "SecretArn": "arn:aws:secretsmanager:us-west-2:111122223333:secret:aes128-1a2b3c"
   }
]
```

#### YAML

```yaml

Auth:
  AuthScheme: SECRETS
  ClientPasswordAuthType: MYSQL_NATIVE_PASSWORD
  Description: Proxy authentication for MySQL
  IAMAuth: DISABLED
  SecretArn: arn:aws:secretsmanager:us-west-2:111122223333:secret:aes128-1a2b3c
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::RDS::DBProxy

TagFormat

All content copied from https://docs.aws.amazon.com/.
