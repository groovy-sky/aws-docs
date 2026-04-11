This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::Service ServiceConnectTlsConfiguration

The key that encrypts and decrypts your resources for Service Connect TLS.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IssuerCertificateAuthority" : ServiceConnectTlsCertificateAuthority,
  "KmsKey" : String,
  "RoleArn" : String
}

```

### YAML

```yaml

  IssuerCertificateAuthority:
    ServiceConnectTlsCertificateAuthority
  KmsKey: String
  RoleArn: String

```

## Properties

`IssuerCertificateAuthority`

The signer certificate authority.

_Required_: Yes

_Type_: [ServiceConnectTlsCertificateAuthority](aws-properties-ecs-service-serviceconnecttlscertificateauthority.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKey`

The AWS Key Management Service key.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The Amazon Resource Name (ARN) of the IAM role that's associated with the Service
Connect TLS.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ServiceConnectTlsCertificateAuthority

ServiceManagedEBSVolumeConfiguration

All content copied from https://docs.aws.amazon.com/.
