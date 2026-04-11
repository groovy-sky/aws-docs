This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::LaunchTemplate Monitoring

Specifies whether detailed monitoring is enabled for an instance. For more information
about detailed monitoring, see [Enable or turn off detailed\
monitoring for your instances](../../../ec2/latest/userguide/using-cloudwatch-new.md) in the _Amazon EC2 User_
_Guide_.

`Monitoring` is a property of [AWS::EC2::LaunchTemplate LaunchTemplateData](../userguide/aws-properties-ec2-launchtemplate-launchtemplatedata.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enabled" : Boolean
}

```

### YAML

```yaml

  Enabled: Boolean

```

## Properties

`Enabled`

Specify `true` to enable detailed monitoring. Otherwise, basic monitoring
is enabled.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [LaunchTemplatesMonitoringRequest](../../../../reference/awsec2/latest/apireference/api-launchtemplatesmonitoringrequest.md) in the _Amazon EC2 API_
_Reference_

- [Create a launch template using advanced settings](../../../autoscaling/ec2/userguide/advanced-settings-for-your-launch-template.md) in the _Amazon EC2 Auto Scaling User Guide_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MetadataOptions

NetworkBandwidthGbps

All content copied from https://docs.aws.amazon.com/.
