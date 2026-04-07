This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkFirewall::TLSInspectionConfiguration TLSInspectionConfiguration

The object that defines a TLS inspection configuration. This defines the TLS inspection configuration.

AWS Network Firewall uses a TLS inspection configuration to decrypt traffic. Network Firewall re-encrypts the traffic before sending it to its destination.

To use a TLS inspection configuration, you add it to a new Network Firewall firewall policy, then you apply the firewall policy to a firewall. Network Firewall acts as a proxy service to decrypt and inspect the traffic traveling through your firewalls. You can reference a TLS inspection configuration from more than one firewall policy, and you can use a firewall policy in more than one firewall. For more information about using TLS inspection configurations, see [Inspecting SSL/TLS traffic with TLS\
inspection configurations](https://docs.aws.amazon.com/network-firewall/latest/developerguide/tls-inspection.html) in the _AWS Network Firewall Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ServerCertificateConfigurations" : [ ServerCertificateConfiguration, ... ]
}

```

### YAML

```yaml

  ServerCertificateConfigurations:
    - ServerCertificateConfiguration

```

## Properties

`ServerCertificateConfigurations`

Lists the server certificate configurations that are associated with the TLS configuration.

_Required_: No

_Type_: Array of [ServerCertificateConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-networkfirewall-tlsinspectionconfiguration-servercertificateconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::NetworkFirewall::VpcEndpointAssociation
