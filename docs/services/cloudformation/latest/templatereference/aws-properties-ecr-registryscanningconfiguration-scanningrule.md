This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECR::RegistryScanningConfiguration ScanningRule

The scanning rules associated with the registry.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RepositoryFilters" : [ RepositoryFilter, ... ],
  "ScanFrequency" : String
}

```

### YAML

```yaml

  RepositoryFilters:
    - RepositoryFilter
  ScanFrequency: String

```

## Properties

`RepositoryFilters`

The details of a scanning repository filter. For more information on how to use
filters, see [Using\
filters](../../../amazonecr/latest/userguide/image-scanning.md#image-scanning-filters) in the _Amazon Elastic Container Registry User Guide_.

_Required_: Yes

_Type_: Array of [RepositoryFilter](aws-properties-ecr-registryscanningconfiguration-repositoryfilter.md)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScanFrequency`

The frequency that scans are performed at for a private registry. When the
`ENHANCED` scan type is specified, the supported scan frequencies are
`CONTINUOUS_SCAN` and `SCAN_ON_PUSH`. When the
`BASIC` scan type is specified, the `SCAN_ON_PUSH` scan
frequency is supported. If scan on push is not specified, then the `MANUAL`
scan frequency is set by default.

_Required_: Yes

_Type_: String

_Allowed values_: `SCAN_ON_PUSH | CONTINUOUS_SCAN`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RepositoryFilter

AWS::ECR::ReplicationConfiguration

All content copied from https://docs.aws.amazon.com/.
