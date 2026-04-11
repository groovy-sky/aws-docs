This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::HealthLake::FHIRDatastore

Creates a Data Store that can ingest and export FHIR formatted data.

###### Important

Please note that when a user tries to do an Update operation via CloudFormation, changes to the Data Store
name, Type Version, PreloadDataConfig, or SSEConfiguration will delete their existing Data Store for the stack
and create a new one. This will lead to potential loss of data.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::HealthLake::FHIRDatastore",
  "Properties" : {
      "DatastoreName" : String,
      "DatastoreTypeVersion" : String,
      "IdentityProviderConfiguration" : IdentityProviderConfiguration,
      "PreloadDataConfig" : PreloadDataConfig,
      "SseConfiguration" : SseConfiguration,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::HealthLake::FHIRDatastore
Properties:
  DatastoreName: String
  DatastoreTypeVersion: String
  IdentityProviderConfiguration:
    IdentityProviderConfiguration
  PreloadDataConfig:
    PreloadDataConfig
  SseConfiguration:
    SseConfiguration
  Tags:
    - Tag

```

## Properties

`DatastoreName`

The data store name (user-generated).

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DatastoreTypeVersion`

The FHIR release version supported by the data store. Current support is for version
`R4`.

_Required_: Yes

_Type_: String

_Allowed values_: `R4`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IdentityProviderConfiguration`

The identity provider configuration selected when the data store was created.

_Required_: No

_Type_: [IdentityProviderConfiguration](aws-properties-healthlake-fhirdatastore-identityproviderconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PreloadDataConfig`

The preloaded Synthea data configuration for the data store.

_Required_: No

_Type_: [PreloadDataConfig](aws-properties-healthlake-fhirdatastore-preloaddataconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SseConfiguration`

The server-side encryption key configuration for a customer-provided encryption key
specified for creating a data store.

_Required_: No

_Type_: [SseConfiguration](aws-properties-healthlake-fhirdatastore-sseconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-healthlake-fhirdatastore-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`DatastoreArn`

The Data Store ARN is generated during the creation of the Data Store and can be found in the output from the
initial Data Store creation request.

`DatastoreEndpoint`

The endpoint for the created Data Store.

`DatastoreId`

The Amazon generated Data Store id. This id is in the output from the initial Data Store creation call.

`DatastoreStatus`

The status of the FHIR Data Store. Possible statuses are ‘CREATING’, ‘ACTIVE’, ‘DELETING’, ‘DELETED’.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS HealthLake

CreatedAt

All content copied from https://docs.aws.amazon.com/.
