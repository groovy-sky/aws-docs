This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::MailManagerIngressPoint

Resource to provision an ingress endpoint for receiving email. An ingress endpoint
serves as the entry point for incoming emails, allowing you to define how emails are
received and processed within your AWS environment.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SES::MailManagerIngressPoint",
  "Properties" : {
      "IngressPointConfiguration" : IngressPointConfiguration,
      "IngressPointName" : String,
      "NetworkConfiguration" : NetworkConfiguration,
      "RuleSetId" : String,
      "StatusToUpdate" : String,
      "Tags" : [ Tag, ... ],
      "TrafficPolicyId" : String,
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::SES::MailManagerIngressPoint
Properties:
  IngressPointConfiguration:
    IngressPointConfiguration
  IngressPointName: String
  NetworkConfiguration:
    NetworkConfiguration
  RuleSetId: String
  StatusToUpdate: String
  Tags:
    - Tag
  TrafficPolicyId: String
  Type: String

```

## Properties

`IngressPointConfiguration`

The configuration of the ingress endpoint resource.

_Required_: No

_Type_: [IngressPointConfiguration](aws-properties-ses-mailmanageringresspoint-ingresspointconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IngressPointName`

A user friendly name for an ingress endpoint resource.

_Required_: No

_Type_: String

_Pattern_: `^[A-Za-z0-9_\-]+$`

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkConfiguration`

The network type (IPv4-only, Dual-Stack, PrivateLink) of the ingress endpoint resource.

_Required_: No

_Type_: [NetworkConfiguration](aws-properties-ses-mailmanageringresspoint-networkconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RuleSetId`

The identifier of an existing rule set that you attach to an ingress endpoint
resource.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StatusToUpdate`

The update status of an ingress endpoint.

_Required_: No

_Type_: String

_Allowed values_: `ACTIVE | CLOSED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags used to organize, track, or control access for the resource. For example, { "tags": {"key1":"value1", "key2":"value2"} }.

_Required_: No

_Type_: Array of [Tag](aws-properties-ses-mailmanageringresspoint-tag.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrafficPolicyId`

The identifier of an existing traffic policy that you attach to an ingress endpoint
resource.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of the ingress endpoint to create.

_Required_: Yes

_Type_: String

_Allowed values_: `OPEN | AUTH`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

`ARecord`

The DNS A Record that identifies your ingress endpoint. Configure your DNS Mail Exchange (MX) record with this value to route emails to Mail Manager.

`IngressPointArn`

The Amazon Resource Name (ARN) of the ingress endpoint resource.

`IngressPointId`

The identifier of the ingress endpoint resource.

`Status`

The status of the ingress endpoint resource.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

IngressPointConfiguration

All content copied from https://docs.aws.amazon.com/.
