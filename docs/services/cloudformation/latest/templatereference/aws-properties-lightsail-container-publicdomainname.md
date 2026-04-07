This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lightsail::Container PublicDomainName

`PublicDomainName` is a property of the [AWS::Lightsail::Container](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-container.html) resource. It describes the public domain names to use
with a container service, such as `example.com` and
`www.example.com`. It also describes the certificates to use with a container
service.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CertificateName" : String,
  "DomainNames" : [ String, ... ]
}

```

### YAML

```yaml

  CertificateName: String
  DomainNames:
    - String

```

## Properties

`CertificateName`

The name of the certificate for the public domains.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainNames`

The public domain names to use with the container service.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PrivateRegistryAccess

PublicEndpoint
