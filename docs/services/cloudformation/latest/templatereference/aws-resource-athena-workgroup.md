This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Athena::WorkGroup

The AWS::Athena::WorkGroup resource specifies an Amazon Athena workgroup,
which contains a name, description, creation time, state, and other configuration,
listed under [WorkGroupConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-workgroup.html#cfn-athena-workgroup-workgroupconfiguration). Each workgroup enables you to
isolate queries for you or your group from other queries in the same account. For more
information, see [CreateWorkGroup](https://docs.aws.amazon.com/athena/latest/APIReference/API_CreateWorkGroup.html) in
the _Amazon Athena API Reference_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Athena::WorkGroup",
  "Properties" : {
      "Description" : String,
      "Name" : String,
      "RecursiveDeleteOption" : Boolean,
      "State" : String,
      "Tags" : [ Tag, ... ],
      "WorkGroupConfiguration" : WorkGroupConfiguration
    }
}

```

### YAML

```yaml

Type: AWS::Athena::WorkGroup
Properties:
  Description: String
  Name: String
  RecursiveDeleteOption: Boolean
  State: String
  Tags:
    - Tag
  WorkGroupConfiguration:
    WorkGroupConfiguration

```

## Properties

`Description`

The workgroup description.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The workgroup name.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9._-]{1,128}`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RecursiveDeleteOption`

The option to delete a workgroup and its contents even if the workgroup contains any
named queries. The default is false.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`State`

The state of the workgroup: ENABLED or DISABLED.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags (key-value pairs) to associate with this resource.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-athena-workgroup-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WorkGroupConfiguration`

The configuration of the workgroup, which includes the location in Amazon S3 where
query results are stored, the encryption option, if any, used for query results, whether
Amazon CloudWatch Metrics are enabled for the workgroup, and the limit for the amount of
bytes scanned (cutoff) per query, if it is specified. The [EnforceWorkGroupConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfiguration.html#cfn-athena-workgroup-workgroupconfiguration-enforceworkgroupconfiguration) option determines whether workgroup
settings override client-side query settings.

_Required_: No

_Type_: [WorkGroupConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-athena-workgroup-workgroupconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the WorkGroup. For example:

`{ "Ref": "myWorkGroup" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreationTime`

The date and time the workgroup was created, as a UNIX timestamp in seconds. For
example: `1582761016`.

`WorkGroupConfiguration.EngineVersion.EffectiveEngineVersion`

Property description not available.

`WorkGroupConfigurationUpdates.EngineVersion.EffectiveEngineVersion`

The Athena engine version for running queries, or the PySpark engine
version for running sessions.

## Examples

### Create or update an Athena workgroup

The following example template creates or updates the Athena workgroup
`MyCustomWorkGroup`.

#### JSON

```json

{
    "Resources":{
        "MyAthenaWorkGroup":{
            "Type":"AWS::Athena::WorkGroup",
            "Properties":{
                "Name":"MyCustomWorkGroup",
                "Description":"My WorkGroup",
                "State":"ENABLED",
                "Tags":[
                    {
                        "Key":"key1",
                        "Value":"value1"
                    },
                    {
                        "Key":"key2",
                        "Value":"value2"
                    }
                ],
                "WorkGroupConfiguration":{
                    "BytesScannedCutoffPerQuery":200000000,
                    "EnforceWorkGroupConfiguration":false,
                    "PublishCloudWatchMetricsEnabled":false,
                    "RequesterPaysEnabled":true,
                    "ResultConfiguration":{
                        "OutputLocation":"s3://path/to/my/bucket/"
                    }
                }
            }
        }
    }
}
```

#### YAML

```yaml

Resources:
  MyAthenaWorkGroup:
    Type: AWS::Athena::WorkGroup
    Properties:
      Name: MyCustomWorkGroup
      Description: My WorkGroup
      State: ENABLED
      Tags:
        - Key: "key1"
          Value: "value1"
        - Key: "key2"
          Value: "value2"
      WorkGroupConfiguration:
        BytesScannedCutoffPerQuery: 200000000
        EnforceWorkGroupConfiguration: false
        PublishCloudWatchMetricsEnabled: false
        RequesterPaysEnabled: true
        ResultConfiguration:
          OutputLocation: s3://path/to/my/bucket/
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Athena::PreparedStatement

AclConfiguration
