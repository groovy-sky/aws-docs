This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecretsManager::Secret

Creates a new secret. A _secret_ can be a password, a set of
credentials such as a user name and password, an OAuth token, or other secret information that
you store in an encrypted form in Secrets Manager.

For Amazon RDS master user credentials, see [AWS::RDS::DBCluster MasterUserSecret](../userguide/aws-properties-rds-dbcluster-masterusersecret.md).

For Amazon Redshift admin user credentials, see [AWS::Redshift::Cluster](../userguide/aws-resource-redshift-cluster.md).

To retrieve a secret in a CloudFormation template, use a _dynamic_
_reference_. For more information, see [Retrieve a secret in an CloudFormation resource](../../../secretsmanager/latest/userguide/cfn-example-reference-secret.md).

For information about creating a secret in the console, see [Create a\
secret](../../../secretsmanager/latest/userguide/manage-create-basic-secret.md). For information about creating a secret using the CLI or SDK, see [CreateSecret](../../../../reference/secretsmanager/latest/apireference/api-createsecret.md).

For information about retrieving a secret in code, see [Retrieve\
secrets from Secrets Manager](../../../secretsmanager/latest/userguide/retrieving-secrets.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SecretsManager::Secret",
  "Properties" : {
      "Description" : String,
      "GenerateSecretString" : GenerateSecretString,
      "KmsKeyId" : String,
      "Name" : String,
      "ReplicaRegions" : [ ReplicaRegion, ... ],
      "SecretString" : String,
      "Tags" : [ Tag, ... ],
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::SecretsManager::Secret
Properties:
  Description: String
  GenerateSecretString:
    GenerateSecretString
  KmsKeyId: String
  Name: String
  ReplicaRegions:
    - ReplicaRegion
  SecretString:
    String
  Tags:
    - Tag
  Type: String

```

## Properties

`Description`

The description of the secret.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GenerateSecretString`

A structure that specifies how to generate a password to encrypt and store in the secret. To include a specific string
in the secret, use `SecretString` instead. If you omit both `GenerateSecretString` and `SecretString`, you create an empty secret. When you make a change to this property, a new secret version is created.

We recommend that you specify the maximum length and include every character type that the
system you are generating a password for can support.

_Required_: No

_Type_: [GenerateSecretString](aws-properties-secretsmanager-secret-generatesecretstring.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyId`

The ARN, key ID, or alias of the AWS KMS key that Secrets Manager uses to
encrypt the secret value in the secret. An alias is always prefixed by `alias/`, for example `alias/aws/secretsmanager`.
For more information, see [About aliases](../../../kms/latest/developerguide/alias-about.md).

To use a AWS KMS key in a different account, use the key ARN or the alias ARN.

If you don't specify this value, then Secrets Manager uses the key `aws/secretsmanager`.
If that key doesn't yet exist, then Secrets Manager creates it for you automatically the first time it
encrypts the secret value.

If the secret is in a different AWS account from the credentials calling the API, then
you can't use `aws/secretsmanager` to encrypt the secret, and you must create
and use a customer managed AWS KMS key.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the new secret.

The secret name can contain ASCII letters, numbers, and the following characters:
/\_+=.@-

Do not end your secret name with a hyphen followed by six characters. If you do so, you
risk confusion and unexpected results when searching for a secret by partial ARN. Secrets Manager
automatically adds a hyphen and six random characters after the secret name at the end of the ARN.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ReplicaRegions`

A custom type that specifies a `Region` and the `KmsKeyId` for a replica secret.

_Required_: No

_Type_: Array of [ReplicaRegion](aws-properties-secretsmanager-secret-replicaregion.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretString`

The text to encrypt and store in the secret. We recommend you use a JSON structure of
key/value pairs for your secret value. To generate a random password, use `GenerateSecretString` instead.
If you omit both `GenerateSecretString` and `SecretString`, you create an empty secret. When you make a change to this property, a new secret version is created.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of tags to attach to the secret. Each tag is a key and value pair of strings in a
JSON text string, for example:

`[{"Key":"CostCenter","Value":"12345"},{"Key":"environment","Value":"production"}]`

Secrets Manager tag key names are case sensitive. A tag with the key "ABC" is a different
tag from one with key "abc".

Stack-level tags, tags you apply to the CloudFormation stack, are also attached to the secret.

If you check tags in permissions policies as part of your security strategy, then adding
or removing a tag can change permissions. If the completion of this operation would result in
you losing your permissions for this secret, then Secrets Manager blocks the operation and
returns an `Access Denied` error. For more information, see [Control\
access to secrets using tags](../../../secretsmanager/latest/userguide/auth-and-access-examples.md#tag-secrets-abac) and [Limit access to identities with tags that match secrets' tags](../../../secretsmanager/latest/userguide/auth-and-access-examples.md#auth-and-access_tags2).

For information about how to format a JSON parameter for the various command line tool
environments, see [Using JSON for\
Parameters](../../../cli/latest/userguide/cli-using-param.md#cli-using-param-json). If your command-line tool or SDK requires quotation marks around the
parameter, you should use single quotes to avoid confusion with the double quotes required in
the JSON text.

The following restrictions apply to tags:

- Maximum number of tags per secret: 50

- Maximum key length: 127 Unicode characters in UTF-8

- Maximum value length: 255 Unicode characters in UTF-8

- Tag keys and values are case sensitive.

- Do not use the `aws:` prefix in your tag names or values because AWS reserves it for AWS use. You can't edit or delete tag names or
values with this prefix. Tags with this prefix do not count against your tags per secret
limit.

- If you use your tagging schema across multiple services and resources, other services
might have restrictions on allowed characters. Generally allowed characters: letters,
spaces, and numbers representable in UTF-8, plus the following special characters: + - = .
\_ : / @.

_Required_: No

_Type_: Array of [Tag](aws-properties-secretsmanager-secret-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The exact string that identifies the third-party partner that holds the external
secret. For more information, see [Managed external secret\
partners](../../../secretsmanager/latest/userguide/mes-partners.md).

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of an `AWS::SecretsManager::Secret` resource to
the intrinsic `Ref` function, the function returns the ARN of the secret configured
such as:

`arn:aws:secretsmanager:us-west-2:123456789012:secret:my-path/my-secret-name-1a2b3c`

If you know the ARN of a secret, you can reference a secret you created in one part of the
stack template from within the definition of another resource in the same template. You
typically use the `Ref` function with the [AWS::SecretsManager::SecretTargetAttachment](../userguide/aws-resource-secretsmanager-secrettargetattachment.md) resource type to get references to both
the secret and its associated database.

For more information about using the `Ref` function, see [Ref](../userguide/intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

The ARN of the secret.

## Examples

- [Creating a secret with a dynamically generated password](#aws-resource-secretsmanager-secret--examples--Creating_a_secret_with_a_dynamically_generated_password)

- [Creating a secret with a hardcoded password](#aws-resource-secretsmanager-secret--examples--Creating_a_secret_with_a_hardcoded_password)

- [Replicating a secret](#aws-resource-secretsmanager-secret--examples--Replicating_a_secret)

### Creating a secret with a dynamically generated password

The following example creates a secret, constructing the secret value from a string
template combined with a dynamically generated random password. The result of this example
is a `SecretString` value that looks like the following:

`{"username": "test-user", "password": "rzDtILsQNfmmHwkJBPsTVhkRvWRtSn"
          }`

#### JSON

```json

{
  "MySecretA": {
    "Type": "AWS::SecretsManager::Secret",
    "Properties": {
      "Name": "MySecretForAppA",
      "Description": "This secret has a dynamically generated secret password.",
      "GenerateSecretString": {
        "SecretStringTemplate": "{\"username\":\"test-user\"}",
        "GenerateStringKey": "password",
        "PasswordLength": 30,
        "ExcludeCharacters": "\"@/\\"
      },
      "Tags": [
        {
          "Key": "AppName",
          "Value": "AppA"
        }
      ]
    }
  }
}

```

#### YAML

```yaml

#This is a Secret resource with a randomly generated password in its SecretString JSON.
MySecretA:
  Type: 'AWS::SecretsManager::Secret'
  Properties:
    Name: MySecretForAppA
    Description: "This secret has a dynamically generated secret password."
    GenerateSecretString:
      SecretStringTemplate: '{"username": "test-user"}'
      GenerateStringKey: "password"
      PasswordLength: 30
      ExcludeCharacters: '"@/\'
    Tags:
      -
        Key: AppName
        Value: AppA
```

### Creating a secret with a hardcoded password

The following example creates a secret and provides the secret value as a literal
string stored in the secret. We recommend that you don't hardcode your password this way.
Instead use the [SecretsManager Secret GenerateSecretString](../userguide/aws-properties-secretsmanager-secret-generatesecretstring.md) property. See the previous example
for the recommended option.

#### JSON

```json

{
  "MySecretB": {
    "Type": "AWS::SecretsManager::Secret",
    "Properties": {
      "Name": "MySecretForAppB",
      "Description": "This secret has a hardcoded password in SecretString (use GenerateSecretString instead)",
      "SecretString": "{\"username\":\"MasterUsername\",\"password\":\"secret-password\"}",
      "Tags": [
        {
          "Key": "AppName",
          "Value": "AppB"
        }
      ]
    }
  }
}
```

#### YAML

```yaml

  # This is another secret that has its password hardcoded into the template (NOT RECOMMENDED)
  MySecretB:
    Type: 'AWS::SecretsManager::Secret'
    Properties:
      Name: MySecretForAppB
      Description: This secret has a hardcoded password in SecretString (use GenerateSecretString instead)
      SecretString: '{"username":"MasterUsername","password":"secret-password"}'
      Tags:
        -
          Key: AppName
          Value: AppB
```

### Replicating a secret

The following example replicates a primary secret to `us-east-1` and
`us-east-2`.

#### JSON

```json

{
   "MyReplicatedSecret":{
      "Type":"AWS::SecretsManager::Secret",
      "Properties":{
         "Name":"MyReplicatedSecret",
         "Description":"This secret is replicated to
        two regions. One with a customer managed key, and one with the AWS managed key for Secrets Manager aws/secretsmanager.",
         "ReplicaRegions":[
            {
               "Region":"us-east-1",
               "KmsKeyId":"alias/exampleAlias"
            },
            {
               "Region":"us-east-2"
            }
         ]
      }
   }
}
```

#### YAML

```yaml

#This is a Secret resource which is replicated to two other regions.
MyReplicatedSecret:
  Type: AWS::SecretsManager::Secret
  Properties:
    Name: MyReplicatedSecret
    Description: 'This secret is replicated to
    two regions. One with a customer managed key, and one with the AWS managed key for Secrets Manager aws/secretsmanager.'
    ReplicaRegions:
    - Region: us-east-1
      KmsKeyId: alias/exampleAlias
    - Region: us-east-2

```

## See also

- [CreateSecret](../../../../reference/secretsmanager/latest/apireference/api-createsecret.md) in
the AWS Secrets Manager API Reference

- [Create and manage secrets](../../../secretsmanager/latest/userguide/managing-secrets.md) in the AWS Secrets Manager User Guide

- [AWS::SecretsManager::ResourcePolicy](../userguide/aws-resource-secretsmanager-resourcepolicy.md)

- [AWS::SecretsManager::RotationSchedule](../userguide/aws-resource-secretsmanager-rotationschedule.md)

- [AWS::SecretsManager::SecretReplicaRegion](../userguide/aws-properties-secretsmanager-secret-replicaregion.md)

- [AWS::SecretsManager::SecretTargetAttachment](../userguide/aws-resource-secretsmanager-secrettargetattachment.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RotationRules

GenerateSecretString

All content copied from https://docs.aws.amazon.com/.
