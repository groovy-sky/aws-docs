This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ARCRegionSwitch::Plan ReportOutputConfiguration

Configuration for report output destinations used in a Region switch plan.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3Configuration" : S3ReportOutputConfiguration
}

```

### YAML

```yaml

  S3Configuration:
    S3ReportOutputConfiguration

```

## Properties

`S3Configuration`

Configuration for delivering reports to an Amazon S3 bucket.

_Required_: Yes

_Type_: [S3ReportOutputConfiguration](aws-properties-arcregionswitch-plan-s3reportoutputconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReportConfiguration

Route53HealthCheckConfiguration

All content copied from https://docs.aws.amazon.com/.
