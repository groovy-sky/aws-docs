This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaTailor::SourceLocation SecretsManagerAccessTokenConfiguration

AWS Secrets Manager access token configuration parameters. For information about Secrets Manager access token authentication, see [Working with AWS Secrets Manager access token authentication](../../../mediatailor/latest/ug/channel-assembly-access-configuration-access-token.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HeaderName" : String,
  "SecretArn" : String,
  "SecretStringKey" : String
}

```

### YAML

```yaml

  HeaderName: String
  SecretArn: String
  SecretStringKey:
    String

```

## Properties

`HeaderName`

The name of the HTTP header used to supply the access token in requests to the source location.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretArn`

The Amazon Resource Name (ARN) of the AWS Secrets Manager secret that contains the access token.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretStringKey`

The AWS Secrets Manager [SecretString](../../../../reference/secretsmanager/latest/apireference/api-createsecret.md#SecretsManager-CreateSecret-request-SecretString.html) key associated with the access token. MediaTailor uses the key to look up SecretString key and value pair containing the access token.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HttpConfiguration

SegmentDeliveryConfiguration

All content copied from https://docs.aws.amazon.com/.
