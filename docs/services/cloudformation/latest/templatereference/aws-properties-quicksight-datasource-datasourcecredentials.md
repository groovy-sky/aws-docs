This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSource DataSourceCredentials

Data source credentials. This is a variant type structure. For this structure to be
valid, only one of the attributes can be non-null.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CopySourceArn" : String,
  "CredentialPair" : CredentialPair,
  "KeyPairCredentials" : KeyPairCredentials,
  "SecretArn" : String
}

```

### YAML

```yaml

  CopySourceArn: String
  CredentialPair:
    CredentialPair
  KeyPairCredentials:
    KeyPairCredentials
  SecretArn: String

```

## Properties

`CopySourceArn`

The Amazon Resource Name (ARN) of a data source that has the credential pair that you
want to use. When `CopySourceArn` is not null, the credential pair from the
data source in the ARN is used as the credentials for the
`DataSourceCredentials` structure.

_Required_: No

_Type_: String

_Pattern_: `^arn:[-a-z0-9]*:quicksight:[-a-z0-9]*:[0-9]{12}:datasource/.+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CredentialPair`

Credential pair. For more information, see
`
                            CredentialPair
                        `.

_Required_: No

_Type_: [CredentialPair](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-datasource-credentialpair.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KeyPairCredentials`

The credentials for connecting using key-pair.

_Required_: No

_Type_: [KeyPairCredentials](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-datasource-keypaircredentials.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretArn`

The Amazon Resource Name (ARN) of the secret associated with the data source in AWS Secrets Manager.

_Required_: No

_Type_: String

_Pattern_: `^arn:[-a-z0-9]*:secretsmanager:[-a-z0-9]*:[0-9]{12}:secret:.+$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DatabricksParameters

DataSourceErrorInfo
