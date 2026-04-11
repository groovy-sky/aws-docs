This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataBrew::Job EntityDetectorConfiguration

Configuration of entity detection for a profile job. When undefined, entity
detection is disabled.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllowedStatistics" : AllowedStatistics,
  "EntityTypes" : [ String, ... ]
}

```

### YAML

```yaml

  AllowedStatistics:
    AllowedStatistics
  EntityTypes:
    - String

```

## Properties

`AllowedStatistics`

Configuration of statistics that are allowed to be run on columns that
contain detected entities. When undefined, no statistics will be computed
on columns that contain detected entities.

_Required_: No

_Type_: [AllowedStatistics](aws-properties-databrew-job-allowedstatistics.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EntityTypes`

Entity types to detect. Can be any of the following:

- USA\_SSN

- EMAIL

- USA\_ITIN

- USA\_PASSPORT\_NUMBER

- PHONE\_NUMBER

- USA\_DRIVING\_LICENSE

- BANK\_ACCOUNT

- CREDIT\_CARD

- IP\_ADDRESS

- MAC\_ADDRESS

- USA\_DEA\_NUMBER

- USA\_HCPCS\_CODE

- USA\_NATIONAL\_PROVIDER\_IDENTIFIER

- USA\_NATIONAL\_DRUG\_CODE

- USA\_HEALTH\_INSURANCE\_CLAIM\_NUMBER

- USA\_MEDICARE\_BENEFICIARY\_IDENTIFIER

- USA\_CPT\_CODE

- PERSON\_NAME

- DATE

The Entity type group USA\_ALL is also supported, and includes all of the
above entity types except PERSON\_NAME and DATE.

_Required_: Yes

_Type_: Array of String

_Maximum_: `128`

_Minimum_: `1 | 1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataCatalogOutput

JobSample

All content copied from https://docs.aws.amazon.com/.
