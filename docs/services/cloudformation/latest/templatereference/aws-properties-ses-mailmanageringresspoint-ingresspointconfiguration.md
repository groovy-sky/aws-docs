This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::MailManagerIngressPoint IngressPointConfiguration

The configuration of the ingress endpoint resource.

###### Important

This data type is a UNION, so only one of the following members can be specified
when used or returned.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SecretArn" : String,
  "SmtpPassword" : String
}

```

### YAML

```yaml

  SecretArn: String
  SmtpPassword: String

```

## Properties

`SecretArn`

The SecretsManager::Secret ARN of the ingress endpoint resource.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws|aws-cn|aws-us-gov|aws-eusc):secretsmanager:[a-z0-9-]+:\d{12}:secret:[a-zA-Z0-9/_+=,.@-]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SmtpPassword`

The password of the ingress endpoint resource.

_Required_: No

_Type_: String

_Pattern_: `^[A-Za-z0-9!@#$%^&*()_+\-=\[\]{}|.,?]+$`

_Minimum_: `8`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SES::MailManagerIngressPoint

NetworkConfiguration

All content copied from https://docs.aws.amazon.com/.
