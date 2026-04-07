This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLift::Fleet CertificateConfiguration

Determines whether a TLS/SSL certificate is generated for a fleet. This feature must be
enabled when creating the fleet. All instances in a fleet share the same
certificate. The certificate can be retrieved by calling the
[GameLift Server\
SDK](https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-serversdk.html) operation `GetInstanceCertificate`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CertificateType" : String
}

```

### YAML

```yaml

  CertificateType: String

```

## Properties

`CertificateType`

Indicates whether a TLS/SSL certificate is generated for a fleet.

Valid values include:

- **GENERATED** \- Generate a TLS/SSL certificate
for this fleet.

- **DISABLED** \- (default) Do not generate a
TLS/SSL certificate for this fleet.

_Required_: Yes

_Type_: String

_Allowed values_: `DISABLED | GENERATED`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [Create GameLift resources using Amazon CloudFront](https://docs.aws.amazon.com/gamelift/latest/developerguide/resources-cloudformation.html) in the _Amazon_
_GameLift Developer Guide_

- [Deploy a GameLift fleet for a custom game build](https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-creating.html) in the _Amazon_
_GameLift Developer Guide_

- [Deploy a Realtime\
Servers fleet](https://docs.aws.amazon.com/gamelift/latest/developerguide/realtime-fleets-creating.html) in the _Amazon GameLift Developer Guide_

- [CertificateConfiguration](https://docs.aws.amazon.com/gamelift/latest/apireference/API_CertificateConfiguration.html) in the _Amazon GameLift API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AnywhereConfiguration

IpPermission
