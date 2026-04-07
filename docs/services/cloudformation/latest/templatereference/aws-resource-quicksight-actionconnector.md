This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::ActionConnector

Creates an action connector that enables Amazon Quick Sight to connect to external services and perform actions.
Action connectors support various authentication methods and can be configured with specific actions from supported connector types
like Amazon S3, Salesforce, JIRA.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::QuickSight::ActionConnector",
  "Properties" : {
      "ActionConnectorId" : String,
      "AuthenticationConfig" : AuthConfig,
      "AwsAccountId" : String,
      "Description" : String,
      "Name" : String,
      "Permissions" : [ ResourcePermission, ... ],
      "Tags" : [ Tag, ... ],
      "Type" : String,
      "VpcConnectionArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::QuickSight::ActionConnector
Properties:
  ActionConnectorId: String
  AuthenticationConfig:
    AuthConfig
  AwsAccountId: String
  Description: String
  Name: String
  Permissions:
    - ResourcePermission
  Tags:
    - Tag
  Type: String
  VpcConnectionArn: String

```

## Properties

`ActionConnectorId`

The unique identifier of the action connector.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AuthenticationConfig`

The authentication configuration used to connect to the external service.

_Required_: No

_Type_: [AuthConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-actionconnector-authconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AwsAccountId`

The AWS account ID associated with the action connector.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9]{12}$`

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

The description of the action connector.

_Required_: No

_Type_: String

_Pattern_: `^[A-Za-z0-9 _.,!?-]*$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the action connector.

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Za-z0-9](?:[\w- ]*[A-Za-z0-9])?$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Permissions`

The permissions configuration that defines which users, groups, or namespaces can access this action connector and what operations they can perform.

_Required_: No

_Type_: Array of [ResourcePermission](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-actionconnector-resourcepermission.html)

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-actionconnector-tag.html)

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of action connector.

_Required_: Yes

_Type_: String

_Allowed values_: `GENERIC_HTTP | SERVICENOW_NOW_PLATFORM | SALESFORCE_CRM | MICROSOFT_OUTLOOK | PAGERDUTY_ADVANCE | JIRA_CLOUD | ATLASSIAN_CONFLUENCE | AMAZON_S3 | AMAZON_BEDROCK_AGENT_RUNTIME | AMAZON_BEDROCK_RUNTIME | AMAZON_BEDROCK_DATA_AUTOMATION_RUNTIME | AMAZON_TEXTRACT | AMAZON_COMPREHEND | AMAZON_COMPREHEND_MEDICAL | MICROSOFT_ONEDRIVE | MICROSOFT_SHAREPOINT | MICROSOFT_TEAMS | SAP_BUSINESSPARTNER | SAP_PRODUCTMASTERDATA | SAP_PHYSICALINVENTORY | SAP_BILLOFMATERIALS | SAP_MATERIALSTOCK | ZENDESK_SUITE | SMARTSHEET | SLACK | ASANA | BAMBOO_HR`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VpcConnectionArn`

The ARN of the VPC connection used for secure connectivity to the external service.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) of the action connector.

`CreatedTime`

The timestamp when the action connector was created.

`EnabledActions`

The list of actions that are enabled for this connector.

`LastUpdatedTime`

The date and time when the action connector was last updated.

`Status`

The current status of the action connector.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Quick Sight

APIKeyConnectionMetadata
