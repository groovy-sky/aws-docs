This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Grafana::Workspace

Specifies a _workspace_. In a workspace, you can create Grafana
dashboards and visualizations to analyze your metrics, logs, and traces. You don't have to
build, package, or deploy any hardware to run the Grafana server.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Grafana::Workspace",
  "Properties" : {
      "AccountAccessType" : String,
      "AuthenticationProviders" : [ String, ... ],
      "ClientToken" : String,
      "DataSources" : [ String, ... ],
      "Description" : String,
      "GrafanaVersion" : String,
      "Name" : String,
      "NetworkAccessControl" : NetworkAccessControl,
      "NotificationDestinations" : [ String, ... ],
      "OrganizationalUnits" : [ String, ... ],
      "OrganizationRoleName" : String,
      "PermissionType" : String,
      "PluginAdminEnabled" : Boolean,
      "RoleArn" : String,
      "SamlConfiguration" : SamlConfiguration,
      "StackSetName" : String,
      "VpcConfiguration" : VpcConfiguration
    }
}

```

### YAML

```yaml

Type: AWS::Grafana::Workspace
Properties:
  AccountAccessType: String
  AuthenticationProviders:
    - String
  ClientToken: String
  DataSources:
    - String
  Description: String
  GrafanaVersion: String
  Name: String
  NetworkAccessControl:
    NetworkAccessControl
  NotificationDestinations:
    - String
  OrganizationalUnits:
    - String
  OrganizationRoleName: String
  PermissionType: String
  PluginAdminEnabled: Boolean
  RoleArn: String
  SamlConfiguration:
    SamlConfiguration
  StackSetName: String
  VpcConfiguration:
    VpcConfiguration

```

## Properties

`AccountAccessType`

Specifies whether the workspace can access AWS resources in this AWS account only, or whether it can also access AWS resources in
other accounts in the same organization. If this is `ORGANIZATION`, the
`OrganizationalUnits` parameter specifies which organizational units
the workspace can access.

_Required_: Yes

_Type_: String

_Allowed values_: `CURRENT_ACCOUNT | ORGANIZATION`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuthenticationProviders`

Specifies whether this workspace uses SAML 2.0, AWS IAM Identity Center, or both to
authenticate users for using the Grafana console within a workspace. For more
information, see [User authentication in\
Amazon Managed Grafana](../../../grafana/latest/userguide/authentication-in-amg.md).

_Allowed Values_: `AWS_SSO | SAML`

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientToken`

A unique, case-sensitive, user-provided identifier to ensure the idempotency of the
request.

_Required_: No

_Type_: String

_Pattern_: `^[!-~]{1,64}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DataSources`

Specifies the AWS data sources that have been configured to have
IAM roles and permissions created to allow Amazon Managed Grafana to read
data from these sources.

This list is only used when the workspace was created through the AWS
console, and the `permissionType` is `SERVICE_MANAGED`.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The user-defined description of the workspace.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GrafanaVersion`

Specifies the version of Grafana to support in the workspace. Defaults to
the latest version on create (for example, 9.4), or the current version of the
workspace on update.

Can only be used to upgrade (for example, from 8.4 to 9.4), not
downgrade (for example, from 9.4 to 8.4).

To know what versions are available to upgrade to for a specific workspace, see
the [ListVersions](../../../../reference/grafana/latest/apireference/api-listversions.md) operation.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the workspace.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9-._~]{1,255}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkAccessControl`

The configuration settings for network access to your workspace.

_Required_: No

_Type_: [NetworkAccessControl](aws-properties-grafana-workspace-networkaccesscontrol.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NotificationDestinations`

The AWS notification channels that Amazon Managed Grafana can automatically
create IAM roles and permissions for, to allow Amazon Managed Grafana to use
these channels.

_AllowedValues_: `SNS`

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OrganizationalUnits`

Specifies the organizational units that this workspace is allowed to use data sources
from, if this workspace is in an account that is part of an organization.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OrganizationRoleName`

The name of the IAM role that is used to access resources through
Organizations.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PermissionType`

If this is `SERVICE_MANAGED`, and the workplace was created through the
Amazon Managed Grafana console, then Amazon Managed Grafana automatically creates the
IAM roles and provisions the permissions that the workspace needs to
use AWS data sources and notification channels.

If this is `CUSTOMER_MANAGED`, you must manage those roles and permissions
yourself.

If you are working with a workspace in a member account of an organization and
that account is not a delegated administrator account, and you want the workspace to
access data sources in other AWS accounts in the organization, this
parameter must be set to `CUSTOMER_MANAGED`.

For more information about converting between customer and service managed, see
[Managing permissions for data sources and notification channels](../../../grafana/latest/userguide/amg-datasource-and-notification.md). For more
information about the roles and permissions that must be managed for customer managed
workspaces, see [Amazon Managed Grafana\
permissions and policies for AWS data sources and notification\
channels](../../../grafana/latest/userguide/amg-manage-permissions.md)

_Required_: Yes

_Type_: String

_Allowed values_: `CUSTOMER_MANAGED | SERVICE_MANAGED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PluginAdminEnabled`

Whether plugin administration is enabled in the workspace. Setting to `true`
allows workspace admins to install, uninstall, and update plugins from within the Grafana
workspace.

###### Note

This option is only valid for workspaces that support Grafana version 9 or
newer.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The IAM role that grants permissions to the AWS
resources that the workspace will view data from. This role must already exist.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SamlConfiguration`

If the workspace uses SAML, use this structure to map SAML assertion attributes to
workspace user information and define which groups in the assertion attribute are to
have the `Admin` and `Editor` roles in the workspace.

_Required_: No

_Type_: [SamlConfiguration](aws-properties-grafana-workspace-samlconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StackSetName`

The name of the AWS CloudFormation stack set that is used to generate IAM roles to be used for this workspace.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcConfiguration`

The configuration settings for an Amazon VPC that contains data sources for
your Grafana workspace to connect to.

###### Note

Connecting to a private VPC is not yet available in the Asia Pacific (Seoul)
Region (ap-northeast-2).

_Required_: No

_Type_: [VpcConfiguration](aws-properties-grafana-workspace-vpcconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name. For example:

`{ "Ref": "Id" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreationTimestamp`

The date that the workspace was created.

Type: Timestamp

`Endpoint`

The URL that users can use to access the Grafana console in the workspace.

Type: String

`GrafanaVersion`

Specifies the version of Grafana supported by this workspace.

Type: String

`Id`

The unique ID of this workspace.

Type: String

`ModificationTimestamp`

The most recent date that the workspace was modified.

Type: Timestamp

`SamlConfigurationStatus`

Specifies whether the workspace's SAML configuration is complete.

Valid values: `CONFIGURED | NOT_CONFIGURED`

Type: String

`SsoClientId`

The ID of the IAM Identity Center-managed application that is created by Amazon Managed Grafana.

Type: String

`Status`

The current status of the workspace.

Valid values: `ACTIVE | CREATING | DELETING | FAILED | UPDATING | UPGRADING |
                DELETION_FAILED | CREATION_FAILED | UPDATE_FAILED | UPGRADE_FAILED |
                LICENSE_REMOVAL_FAILED`

Type: String

## Examples

### Create a workspace

Create an Amazon Managed Grafana workspace using CloudFormation

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "Amazon Grafana workspace stack",
    "Resources": {
        "AmazonGrafanaWorkspaceIAMRole": {
            "Type": "AWS::IAM::Role",
            "Properties": {
                "ManagedPolicyArns": ["arn:aws:iam::aws:policy/service-role/AmazonGrafanaAthenaAccess"],
                "AssumeRolePolicyDocument": {
                    "Version": "2012-10-17",
                    "Statement": [{
                        "Effect": "Allow",
                        "Principal": {
                            "Service": [
                                "grafana.amazonaws.com"
                            ]
                        },
                        "Action": [
                            "sts:AssumeRole"
                        ]
                    }]
                }
            }
        },
        "AmazonGrafanaWorkspace": {
            "Type": "AWS::Grafana::Workspace",
            "Properties": {
                "AccountAccessType": "CURRENT_ACCOUNT",
                "Name": "AmazonGrafanaWorkspace",
                "Description": "Amazon Grafana Workspace",
                "AuthenticationProviders": ["SAML"],
                "PermissionType": "CUSTOMER_MANAGED",
                "GrafanaVersion": "9.4",
                "RoleArn": {
                    "Fn::GetAtt": [
                        "AmazonGrafanaWorkspaceIAMRole",
                        "Arn"
                    ]
                },
                "SamlConfiguration": {
                    "IdpMetadata": {
                        "Xml": "<md:EntityDescriptor xmlns:md='urn:oasis:names:tc:SAML:2.0:metadata' entityID='entityId'>DATA</md:EntityDescriptor>"
                    },
                    "AssertionAttributes": {
                        "Name": "displayName",
                        "Login": "login",
                        "Email": "email",
                        "Groups": "group",
                        "Role": "role",
                        "Org": "org"
                    },
                    "RoleValues": {
                        "Editor": ["editor1"],
                        "Admin": ["admin1"]
                    },
                    "AllowedOrganizations": ["org1"],
                    "LoginValidityDuration": 60
                }
            }
        }
    },
    "Outputs": {
        "WorkspaceEndpoint": {
            "Value": {
                "Fn::GetAtt": [
                    "AmazonGrafanaWorkspace",
                    "Endpoint"
                ]
            }
        },
        "WorkspaceStatus": {
            "Value": {
                "Fn::GetAtt": [
                    "AmazonGrafanaWorkspace",
                    "Status"
                ]
            }
        },
        "WorkspaceId": {
            "Value": {
                "Ref": "AmazonGrafanaWorkspace"
            }
        },
        "GrafanaVersion": {
            "Value": {
                "Fn::GetAtt": [
                    "AmazonGrafanaWorkspace",
                    "GrafanaVersion"
                ]
            }
        }
    }
}
```

#### YAML

```yaml

Description: Amazon Grafana workspace stack
Resources:
  AmazonGrafanaWorkspaceIAMRole:
    Type: 'AWS::IAM::Role'
    Properties:
      ManagedPolicyArns:
        - 'arn:aws:iam::aws:policy/service-role/AmazonGrafanaAthenaAccess'
      AssumeRolePolicyDocument:
        Version: 2012-10-17
        Statement:
          - Effect: Allow
            Principal:
              Service:
                - grafana.amazonaws.com
            Action:
              - 'sts:AssumeRole'

  AmazonGrafanaWorkspace:
    Type: 'AWS::Grafana::Workspace'
    Properties:
      AccountAccessType: CURRENT_ACCOUNT
      Name: AmazonGrafanaWorkspace
      Description: Amazon Grafana Workspace
      AuthenticationProviders:
        - SAML
      PermissionType: CUSTOMER_MANAGED
      GrafanaVersion: '9.4'
      RoleArn: !GetAtt
        - AmazonGrafanaWorkspaceIAMRole
        - Arn
      SamlConfiguration:
        IdpMetadata:
          Xml: >-
            <md:EntityDescriptor xmlns:md='urn:oasis:names:tc:SAML:2.0:metadata'
            entityID='entityId'>DATA</md:EntityDescriptor>
        AssertionAttributes:
          Name: displayName
          Login: login
          Email: email
          Groups: group
          Role: role
          Org: org
        RoleValues:
          Editor:
            - editor1
          Admin:
            - admin1
        AllowedOrganizations:
          - org1
        LoginValidityDuration: 60

Outputs:
  WorkspaceEndpoint:
    Value: !GetAtt
      - AmazonGrafanaWorkspace
      - Endpoint
  WorkspaceStatus:
    Value: !GetAtt
      - AmazonGrafanaWorkspace
      - Status
  WorkspaceId:
    Value: !Ref AmazonGrafanaWorkspace
  GrafanaVersion:
    Value: !GetAtt
      - AmazonGrafanaWorkspace
      - GrafanaVersion
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon Managed Grafana

AssertionAttributes

All content copied from https://docs.aws.amazon.com/.
