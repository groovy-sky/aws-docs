This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::ConnectorProfile SalesforceConnectorProfileProperties

The connector-specific profile properties required when using Salesforce.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InstanceUrl" : String,
  "isSandboxEnvironment" : Boolean,
  "usePrivateLinkForMetadataAndAuthorization" : Boolean
}

```

### YAML

```yaml

  InstanceUrl: String
  isSandboxEnvironment: Boolean
  usePrivateLinkForMetadataAndAuthorization: Boolean

```

## Properties

`InstanceUrl`

The location of the Salesforce resource.

_Required_: No

_Type_: String

_Pattern_: `\S+`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`isSandboxEnvironment`

Indicates whether the connector profile applies to a sandbox or production environment.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`usePrivateLinkForMetadataAndAuthorization`

If the connection mode for the connector profile is private, this parameter sets whether
Amazon AppFlow uses the private network to send metadata and authorization calls to
Salesforce. Amazon AppFlow sends private calls through AWS PrivateLink. These
calls travel through AWS infrastructure without being exposed to the public
internet.

Set either of the following values:

true

Amazon AppFlow sends all calls to Salesforce over the private network.

These private calls are:

- Calls to get metadata about your Salesforce records. This metadata describes
your Salesforce objects and their fields.

- Calls to get or refresh access tokens that allow Amazon AppFlow to access
your Salesforce records.

- Calls to transfer your Salesforce records as part of a flow run.

false

The default value. Amazon AppFlow sends some calls to Salesforce privately and
other calls over the public internet.

The public calls are:

- Calls to get metadata about your Salesforce records.

- Calls to get or refresh access tokens.

The private calls are:

- Calls to transfer your Salesforce records as part of a flow run.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [SalesforceConnectorProfileProperties](../../../../reference/appflow/1-0/apireference/api-salesforceconnectorprofileproperties.md) in the _Amazon AppFlow API_
_Reference_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SalesforceConnectorProfileCredentials

SAPODataConnectorProfileCredentials
