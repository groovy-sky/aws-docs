This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PCAConnectorAD::ServicePrincipalName

Creates a service principal name (SPN) for the service account in Active Directory. Kerberos
authentication uses SPNs to associate a service instance with a service sign-in
account.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::PCAConnectorAD::ServicePrincipalName",
  "Properties" : {
      "ConnectorArn" : String,
      "DirectoryRegistrationArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::PCAConnectorAD::ServicePrincipalName
Properties:
  ConnectorArn: String
  DirectoryRegistrationArn: String

```

## Properties

`ConnectorArn`

The Amazon Resource Name (ARN) that was returned when you called [CreateConnector.html](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateConnector.html).

_Required_: No

_Type_: String

_Pattern_: `^arn:[\w-]+:pca-connector-ad:[\w-]+:[0-9]+:connector(\/[\w-]+)$`

_Minimum_: `5`

_Maximum_: `200`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DirectoryRegistrationArn`

The Amazon Resource Name (ARN) that was returned when you called [CreateDirectoryRegistration](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateDirectoryRegistration.html).

_Required_: No

_Type_: String

_Pattern_: `^arn:[\w-]+:pca-connector-ad:[\w-]+:[0-9]+:directory-registration(\/[\w-]+)$`

_Minimum_: `5`

_Maximum_: `200`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::PCAConnectorAD::DirectoryRegistration

AWS::PCAConnectorAD::Template
