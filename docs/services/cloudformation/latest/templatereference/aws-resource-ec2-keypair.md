This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::KeyPair

Specifies a key pair for use with an Amazon Elastic Compute Cloud instance as follows:

- To import an existing key pair, include the `PublicKeyMaterial` property.

- To create a new key pair, omit the `PublicKeyMaterial` property.

When you import an existing key pair, you specify the public key material for the key. We
assume that you have the private key material for the key. AWS CloudFormation does not
create or return the private key material when you import a key pair.

When you create a new key pair, the private key is saved to AWS Systems Manager
Parameter Store, using a parameter with the following name: `/ec2/keypair/{key_pair_id}`.
For more information about retrieving private key, and the required permissions, see [Create a key pair using CloudFormation](../../../ec2/latest/userguide/create-key-pairs.md#create-key-pair-cloudformation) in the _Amazon EC2 User Guide_.

When CloudFormation deletes a key pair that was created or imported by a stack,
it also deletes the parameter that was used to store the private key material in
Parameter Store.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::KeyPair",
  "Properties" : {
      "KeyFormat" : String,
      "KeyName" : String,
      "KeyType" : String,
      "PublicKeyMaterial" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::EC2::KeyPair
Properties:
  KeyFormat: String
  KeyName: String
  KeyType: String
  PublicKeyMaterial: String
  Tags:
    - Tag

```

## Properties

`KeyFormat`

The format of the key pair.

Default: `pem`

_Required_: No

_Type_: String

_Allowed values_: `pem | ppk`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KeyName`

A unique name for the key pair.

Constraints: Up to 255 ASCII characters

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KeyType`

The type of key pair. Note that ED25519 keys are not supported for Windows instances.

If the `PublicKeyMaterial` property is specified, the `KeyType` property is ignored, and the key type
is inferred from the `PublicKeyMaterial` value.

Default: `rsa`

_Required_: No

_Type_: String

_Allowed values_: `rsa | ed25519`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PublicKeyMaterial`

The public key material. The `PublicKeyMaterial` property is used to import a key pair. If this property is not specified,
then a new key pair will be created.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags to apply to the key pair.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-keypair-tag.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the key pair.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`KeyFingerprint`

If you created the key pair using Amazon EC2:

- For RSA key pairs, the key fingerprint is the SHA-1 digest of the DER encoded private key.

- For ED25519 key pairs, the key fingerprint is the base64-encoded SHA-256 digest, which is the default for OpenSSH, starting with [OpenSSH 6.8](http://www.openssh.com/txt/release-6.8).

If you imported the key pair to Amazon EC2:

- For RSA key pairs, the key fingerprint is the MD5 public key fingerprint as specified in section 4 of RFC 4716.

- For ED25519 key pairs, the key fingerprint is the base64-encoded SHA-256
digest, which is the default for OpenSSH, starting with [OpenSSH 6.8](http://www.openssh.com/txt/release-6.8).

`KeyPairId`

The ID of the key pair.

## Examples

- [Create a new key pair and specify it when launching an instance](#aws-resource-ec2-keypair--examples--Create_a_new_key_pair_and_specify_it_when_launching_an_instance)

- [Import an existing key pair and specify it when launching an instance](#aws-resource-ec2-keypair--examples--Import_an_existing_key_pair_and_specify_it_when_launching_an_instance)

### Create a new key pair and specify it when launching an instance

The following example omits the `PublicKeyMaterial` property to
create a new key pair, and specifies the key pair when launching an instance.

#### JSON

```json

{
    "Resources": {
        "NewKeyPair": {
            "Type": "AWS::EC2::KeyPair",
            "Properties": {
                "KeyName": "MyKeyPair"
            }
        },
        "Ec2Instance": {
            "Type": "AWS::EC2::Instance",
            "Properties": {
                "ImageId": "ami-02b92c281a4d3dc79",
                "KeyName": {
                    "Ref": "NewKeyPair"
                }
            }
        }
    }
}
```

#### YAML

```yaml

Resources:
  NewKeyPair:
    Type: 'AWS::EC2::KeyPair'
    Properties:
      KeyName: MyKeyPair
  Ec2Instance:
    Type: 'AWS::EC2::Instance'
    Properties:
      ImageId: ami-02b92c281a4d3dc79
      KeyName: !Ref NewKeyPair
```

### Import an existing key pair and specify it when launching an instance

The following example uses the `PublicKeyMaterial` property to
import an existing key pair, and specifies the key pair when launching an instance.

#### JSON

```json

{
    "Resources": {
        "ImportedKeyPair": {
            "Type": "AWS::EC2::KeyPair",
            "Properties": {
                "KeyName": "NameForMyImportedKeyPair",
                "PublicKeyMaterial": "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAICfp1F7DhdWZdqkYAUGCzcBsLmJeu9izpIyGpmmg7eCz example"
            }
        },
        "Ec2Instance": {
            "Type": "AWS::EC2::Instance",
            "Properties": {
                "ImageId": "ami-02b92c281a4d3dc79",
                "KeyName": {
                    "Ref": "ImportedKeyPair"
                }
            }
        }
    }
}
```

#### YAML

```yaml

Resources:
  ImportedKeyPair:
    Type: AWS::EC2::KeyPair
    Properties:
      KeyName: NameForMyImportedKeyPair
      PublicKeyMaterial: ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAICfp1F7DhdWZdqkYAUGCzcBsLmJeu9izpIyGpmmg7eCz example
  Ec2Instance:
    Type: AWS::EC2::Instance
    Properties:
      ImageId: ami-02b92c281a4d3dc79
      KeyName:
        Ref: ImportedKeyPair
```

## See also

- [CreateKeyPair](../../../../reference/awsec2/latest/apireference/api-createkeypair.md) in the _Amazon EC2 API Reference_

- [ImportKeyPair](../../../../reference/awsec2/latest/apireference/api-importkeypair.md) in the _Amazon EC2 API Reference_

- [Amazon\
EC2 key pairs](../../../ec2/latest/userguide/ec2-key-pairs.md) in the _Amazon EC2 User Guide_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::EC2::IpPoolRouteTableAssociation

Tag
