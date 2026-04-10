This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GuardDuty::Detector CFNS3LogsConfiguration

Describes whether S3 data event logs will be enabled as a data source when the detector
is created.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enable" : Boolean
}

```

### YAML

```yaml

  Enable: Boolean

```

## Properties

`Enable`

The status of S3 data event logs as a data source.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CFNMalwareProtectionConfiguration

CFNScanEc2InstanceWithFindingsConfiguration

All content copied from https://docs.aws.amazon.com/.
