This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::EncryptionConfiguration

Retrieves the encryption configuration for resources and data of your AWS account in
AWS IoT Core. For more information, see [Data encryption at rest](../../../iot/latest/developerguide/encryption-at-rest.md) in
the _AWS IoT Core Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoT::EncryptionConfiguration",
  "Properties" : {
      "EncryptionType" : String,
      "KmsAccessRoleArn" : String,
      "KmsKeyArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::IoT::EncryptionConfiguration
Properties:
  EncryptionType: String
  KmsAccessRoleArn: String
  KmsKeyArn: String

```

## Properties

`EncryptionType`

The type of the KMS key.

_Required_: Yes

_Type_: String

_Allowed values_: `CUSTOMER_MANAGED_KMS_KEY | AWS_OWNED_KMS_KEY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsAccessRoleArn`

The Amazon Resource Name (ARN) of the IAM role assumed by AWS IoT Core to call AWS KMS on
behalf of the customer.

_Required_: No

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyArn`

The ARN of the customer managed KMS key.

_Required_: No

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the AWS account ID. For example:

`{ "Ref": "MyEncryptionConfiguration" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AccountId`

The unique identifier (ID) of an AWS account.

`LastModifiedDate`

The date when encryption configuration is last updated.

## Examples

The following example creates an encryption configuration using a customer managed AWS KMS key.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Resources": {
    "MyEncryptionConfiguration": {
      "Type": "AWS::IoT::EncryptionConfiguration",
      "Properties": {
        "EncryptionType": "CUSTOMER_MANAGED_KMS_KEY",
        "KmsKeyArn": "arn:aws:kms:us-east-1:123456789012:key/abcd1234-ab12-cd34-ef56-abcdef123456",
        "KmsAccessRoleArn": "arn:aws:iam::123456789012:role/IoTKmsAccessRole"
      }
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: '2010-09-09'
Resources:
  MyEncryptionConfiguration:
    Type: AWS::IoT::EncryptionConfiguration
    Properties:
      EncryptionType: CUSTOMER_MANAGED_KMS_KEY
      KmsKeyArn: arn:aws:kms:us-east-1:123456789012:key/abcd1234-ab12-cd34-ef56-abcdef123456
      KmsAccessRoleArn: arn:aws:iam::123456789012:role/IoTKmsAccessRole
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TlsConfig

ConfigurationDetails

All content copied from https://docs.aws.amazon.com/.
