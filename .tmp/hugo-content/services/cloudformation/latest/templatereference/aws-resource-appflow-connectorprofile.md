This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::ConnectorProfile

The `AWS::AppFlow::ConnectorProfile` resource is an Amazon AppFlow resource
type that specifies the configuration profile for an instance of a connector. This includes
the provided name, credentials ARN, connection-mode, and so on. The fields that are common to
all types of connector profiles are explicitly specified under the `Properties`
field. The rest of the connector-specific properties are specified under
`Properties/ConnectorProfileConfig`.

###### Note

If you want to use CloudFormation to create a connector profile for connectors
that implement OAuth (such as Salesforce, Slack, Zendesk, and Google Analytics), you must
fetch the access and refresh tokens. You can do this by implementing your own UI for OAuth,
or by retrieving the tokens from elsewhere. Alternatively, you can use the Amazon AppFlow
console to create the connector profile, and then use that connector profile in the flow
creation CloudFormation template.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppFlow::ConnectorProfile",
  "Properties" : {
      "ConnectionMode" : String,
      "ConnectorLabel" : String,
      "ConnectorProfileConfig" : ConnectorProfileConfig,
      "ConnectorProfileName" : String,
      "ConnectorType" : String,
      "KMSArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::AppFlow::ConnectorProfile
Properties:
  ConnectionMode: String
  ConnectorLabel: String
  ConnectorProfileConfig:
    ConnectorProfileConfig
  ConnectorProfileName: String
  ConnectorType: String
  KMSArn: String

```

## Properties

`ConnectionMode`

Indicates the connection mode and if it is public or private.

_Required_: Yes

_Type_: String

_Allowed values_: `Public | Private`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConnectorLabel`

The label for the connector profile being created.

_Required_: No

_Type_: String

_Pattern_: `[\w!@#.-]+`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ConnectorProfileConfig`

Defines the connector-specific configuration and credentials.

_Required_: No

_Type_: [ConnectorProfileConfig](aws-properties-appflow-connectorprofile-connectorprofileconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConnectorProfileName`

The name of the connector profile. The name is unique for each
`ConnectorProfile` in the AWS account.

_Required_: Yes

_Type_: String

_Pattern_: `[\w/!@#+=.-]+`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ConnectorType`

The type of connector, such as Salesforce, Amplitude, and so on.

_Required_: Yes

_Type_: String

_Allowed values_: `Salesforce | Pardot | Singular | Slack | Redshift | Marketo | Googleanalytics | Zendesk | Servicenow | SAPOData | Datadog | Trendmicro | Snowflake | Dynatrace | Infornexus | Amplitude | Veeva | CustomConnector`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KMSArn`

The ARN (Amazon Resource Name) of the Key Management Service (KMS) key you provide for
encryption. This is required if you do not want to use the Amazon AppFlow-managed KMS
key. If you don't provide anything here, Amazon AppFlow uses the Amazon AppFlow-managed KMS key.

_Required_: No

_Type_: String

_Pattern_: `arn:aws:kms:.*:[0-9]+:.*`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the connector profile name. For example:

`{ "Ref": "myConnectorProfile" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ConnectorProfileArn`

The Amazon Resource Name (ARN) of the connector profile.

`CredentialsArn`

The Amazon Resource Name (ARN) of the connector profile credentials.

## See also

- [CreateConnectorProfile](../../../../reference/appflow/1-0/apireference/api-createconnectorprofile.md) in the _Amazon AppFlow API_
_Reference_.

- [DescribeConnectorProfiles](../../../../reference/appflow/1-0/apireference/api-describeconnectorprofiles.md) in the _Amazon AppFlow API_
_Reference_.

- [DeleteConnectorProfile](../../../../reference/appflow/1-0/apireference/api-deleteconnectorprofile.md) in the _Amazon AppFlow API_
_Reference_.

- [UpdateConnectorProfile](../../../../reference/appflow/1-0/apireference/api-updateconnectorprofile.md) in the _Amazon AppFlow API_
_Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LambdaConnectorProvisioningConfig

AmplitudeConnectorProfileCredentials

All content copied from https://docs.aws.amazon.com/.
