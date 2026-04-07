This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTSiteWise::Portal

###### Important

The AWS IoT SiteWise Monitor feature will no longer be open to new customers starting November 7, 2025 . If you would like to use the AWS IoT SiteWise Monitor feature,
sign up prior to that date. Existing customers can continue to use the service as normal. For more information, see
[AWS IoT SiteWise Monitor availability change](https://docs.aws.amazon.com/iot-sitewise/latest/appguide/iotsitewise-monitor-availability-change.html).

Creates a portal, which can contain projects and dashboards. AWS IoT SiteWise Monitor uses IAM Identity Center or IAM
to authenticate portal users and manage user permissions.

###### Note

Before you can sign in to a new portal, you must add at least one identity to that
portal. For more information, see [Adding or removing portal\
administrators](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/administer-portals.html#portal-change-admins) in the _AWS IoT SiteWise User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoTSiteWise::Portal",
  "Properties" : {
      "Alarms" : Alarms,
      "NotificationSenderEmail" : String,
      "PortalAuthMode" : String,
      "PortalContactEmail" : String,
      "PortalDescription" : String,
      "PortalName" : String,
      "PortalType" : String,
      "PortalTypeConfiguration" : {Key: Value, ...},
      "RoleArn" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IoTSiteWise::Portal
Properties:
  Alarms:
    Alarms
  NotificationSenderEmail: String
  PortalAuthMode: String
  PortalContactEmail: String
  PortalDescription: String
  PortalName: String
  PortalType: String
  PortalTypeConfiguration:
    Key: Value
  RoleArn: String
  Tags:
    - Tag

```

## Properties

`Alarms`

Contains the configuration information of an alarm created in an AWS IoT SiteWise Monitor portal.
You can use the alarm to monitor an asset property and get notified when the asset property value is outside a specified range.
For more information, see [Monitoring with alarms](https://docs.aws.amazon.com/iot-sitewise/latest/appguide/monitor-alarms.html) in the _AWS IoT SiteWise Application Guide_.

_Required_: No

_Type_: [Alarms](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotsitewise-portal-alarms.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NotificationSenderEmail`

The email address that sends alarm notifications.

###### Important

If you use the [AWS IoT Events managed Lambda\
function](https://docs.aws.amazon.com/iotevents/latest/developerguide/lambda-support.html) to manage your emails, you must [verify the sender email\
address in Amazon SES](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/verify-email-addresses.html).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PortalAuthMode`

The service to use to authenticate users to the portal. Choose from the following
options:

- `SSO` – The portal uses AWS IAM Identity Center to authenticate users and manage
user permissions. Before you can create a portal that uses IAM Identity Center, you must enable IAM Identity Center.
For more information, see [Enabling IAM Identity Center](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/monitor-get-started.html#mon-gs-sso) in the
_AWS IoT SiteWise User Guide_. This option is only available in AWS Regions other than
the China Regions.

- `IAM` – The portal uses AWS Identity and Access Management to authenticate users and manage
user permissions.

You can't change this value after you create a portal.

Default: `SSO`

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PortalContactEmail`

The AWS administrator's contact email address.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PortalDescription`

A description for the portal.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PortalName`

A friendly name for the portal.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PortalType`

Define the type of portal. The value for AWS IoT SiteWise Monitor (Classic) is `SITEWISE_PORTAL_V1`. The value for AWS IoT SiteWise Monitor (AI-aware) is `SITEWISE_PORTAL_V2`.

_Required_: No

_Type_: String

_Allowed values_: `SITEWISE_PORTAL_V1 | SITEWISE_PORTAL_V2`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PortalTypeConfiguration`

Property description not available.

_Required_: No

_Type_: Object of [PortalTypeEntry](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotsitewise-portal-portaltypeentry.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of a service role that allows the portal's users to access your AWS IoT SiteWise
resources on your behalf. For more information, see [Using service roles for AWS IoT SiteWise Monitor](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/monitor-service-role.html) in the
_AWS IoT SiteWise User Guide_.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of key-value pairs that contain metadata for the portal. For more information, see
[Tagging your AWS IoT SiteWise\
resources](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/tag-resources.html) in the _AWS IoT SiteWise User Guide_.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotsitewise-portal-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `PortalId`.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`PortalArn`

The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the portal, which has the following format.

`arn:${Partition}:iotsitewise:${Region}:${Account}:portal/${PortalId}`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

`PortalClientId`

The IAM Identity Center application generated client ID (used with IAM Identity Center
APIs).

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

`PortalId`

The ID of the created portal.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

`PortalStartUrl`

The public URL for the AWS IoT SiteWise Monitor portal.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Alarms
