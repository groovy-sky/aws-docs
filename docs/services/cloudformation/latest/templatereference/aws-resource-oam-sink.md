This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Oam::Sink

Creates or updates a _sink_ in the current account, so that it can be used as a
monitoring account in CloudWatch cross-account observability. A sink is a resource that represents an
attachment point in a monitoring account, which source accounts can link to to be able to send observability data.

After you create a sink, you must create a sink policy that allows source accounts to attach to it.
For more information, see [PutSinkPolicy](https://docs.aws.amazon.com/OAM/latest/APIReference/API_PutSinkPolicy.html).

An account can have one sink.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Oam::Sink",
  "Properties" : {
      "Name" : String,
      "Policy" : Json,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::Oam::Sink
Properties:
  Name: String
  Policy: Json
  Tags:
    Key: Value

```

## Properties

`Name`

A name for the sink.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_.-]+$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Policy`

The IAM policy that grants permissions to source accounts to link to this sink. The policy can grant permission
in the following ways:

- Include organization
IDs or organization paths to permit all accounts in an organization

- Include account IDs to permit the specified accounts

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to the sink.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).

_Required_: No

_Type_: Object of String

_Pattern_: `^(?!aws:.*).{1,128}$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the link. For example,
`arn:aws:oam:us-west-1:111111111111:link:abcd1234-a123-456a-a12b-a123b456c789`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN of the sink. For example, `arn:aws:oam:us-west-1:111111111111:sink:abcd1234-a123-456a-a12b-a123b456c789`

## Examples

- [Sample sink to connect that permits links to all accounts in an organization](#aws-resource-oam-sink--examples--Sample_sink_to_connect_that_permits_links_to_all_accounts_in_an_organization)

- [Sample sink that permits a link to an individual account](#aws-resource-oam-sink--examples--Sample_sink_that_permits_a_link_to_an_individual_account)

- [Sample sink for CloudWatch Application Insights applications support](#aws-resource-oam-sink--examples--Sample_sink_for_applications_support)

### Sample sink to connect that permits links to all accounts in an organization

This example creates a sink that allows all accounts in a specified organization
to create links to share metric and log data.

#### JSON

```json

"Name": "SampleSink",
"Policy": {
   "Version": "2012-10-17",
   "Statement": [{
       "Effect": "Allow",
       "Principal": "*",
       "Resource": "*",
       "Action": [ "oam:CreateLink", "oam:UpdateLink" ],
       "Condition": {
        "StringEquals": {"aws:PrincipalOrgID":"o-xxxxxxxxxxx"},
        "ForAllValues:StringEquals": {
          "oam:ResourceTypes": [
            "AWS::CloudWatch::Metric",
            "AWS::Logs::LogGroup"
          ]
        }
       }
    }]
}
```

#### YAML

```yaml

Name: "SampleSink"
Policy:
  Version: '2012-10-17'
  Statement:
  - Effect: Allow
    Principal: "*"
    Resource: "*"
    Action:
    - "oam:CreateLink"
    - "oam:UpdateLink"
    Condition:
      StringEquals:
        aws:PrincipalOrgID: o-xxxxxxxxxxx
      ForAllValues:StringEquals:
        oam:ResourceTypes:
          - "AWS::CloudWatch::Metric"
          - "AWS::Logs::LogGroup"

```

### Sample sink that permits a link to an individual account

This example creates a sink that allows the account with the ID `111111111111` to create
a link to share metrics, logs, and traces.

#### JSON

```json

"Name": "SampleSink",
"Policy": {
   "Version": "2012-10-17",
   "Statement": [{
       "Effect": "Allow",
       "Resource": "*",
       "Action": "oam:*",
       "Principal": {
        "AWS": [ "1111111111111" ]
       },
       "Condition": {
        "ForAllValues:StringEquals": {
          "oam:ResourceTypes": [
            "AWS::CloudWatch::Metric",
            "AWS::Logs::LogGroup",
            "AWS::XRay::Trace"
          ]
        }
       }
    }]
}
```

#### YAML

```yaml

Name: "SampleSink"
Policy:
  Version: '2012-10-17'
  Statement:
  - Effect: Allow
    Resource: "*"
    Action: "oam:*"
    Principal:
      AWS:
      - '1111111111111'
    Condition:
      ForAllValues:StringEquals:
        oam:ResourceTypes:
          - "AWS::CloudWatch::Metric"
          - "AWS::Logs::LogGroup"
          - "AWS::XRay::Trace"

```

### Sample sink for CloudWatch Application Insights applications support

This example creates a sink that allows the account with the ID `111111111111` to create
a link to share metrics, logs, traces, and Application Insights applications.

#### JSON

```json

"Name": "SampleSink",
 "Policy": {
    "Version": "2012-10-17",
    "Statement": [{
        "Effect": "Allow",
        "Resource": "*",
        "Action": "oam:*",
        "Principal": {
         "AWS": [ "1111111111111" ]
        },
        "Condition": {
         "ForAllValues:StringEquals": {
           "oam:ResourceTypes": [
             "AWS::CloudWatch::Metric",
             "AWS::Logs::LogGroup",
             "AWS::XRay::Trace",
             "AWS::ApplicationInsights::Application"
           ]
         }
        }
     }]
 }
```

#### YAML

```yaml

Name: "SampleSink"
 Policy:
   Version: '2012-10-17'
   Statement:
   - Effect: Allow
     Resource: "*"
     Action: "oam:*"
     Principal:
       AWS:
       - '1111111111111'
     Condition:
       ForAllValues:StringEquals:
         oam:ResourceTypes:
           - "AWS::CloudWatch::Metric"
           - "AWS::Logs::LogGroup"
           - "AWS::XRay::Trace"
           - "AWS::ApplicationInsights::Application"

```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

LinkFilter

Next
