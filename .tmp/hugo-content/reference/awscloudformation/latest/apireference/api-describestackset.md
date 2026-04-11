# DescribeStackSet

Returns the description of the specified StackSet.

###### Note

This API provides _strongly consistent_ reads meaning it will always
return the most up-to-date data.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**CallAs**

\[Service-managed permissions\] Specifies whether you are acting as an account administrator
in the organization's management account or as a delegated administrator in a
member account.

By default, `SELF` is specified. Use `SELF` for StackSets with
self-managed permissions.

- If you are signed in to the management account, specify
`SELF`.

- If you are signed in to a delegated administrator account, specify
`DELEGATED_ADMIN`.

Your AWS account must be registered as a delegated administrator in the management account. For more information, see [Register a\
delegated administrator](../../../../services/cloudformation/latest/userguide/stacksets-orgs-delegated-admin.md) in the _AWS CloudFormation User Guide_.

Type: String

Valid Values: `SELF | DELEGATED_ADMIN`

Required: No

**StackSetName**

The name or unique ID of the StackSet whose description you want.

Type: String

Required: Yes

## Response Elements

The following element is returned by the service.

**StackSet**

The specified StackSet.

Type: [StackSet](api-stackset.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**StackSetNotFound**

The specified StackSet doesn't exist.

HTTP Status Code: 404

## Examples

### DescribeStackSet

This example illustrates one usage of DescribeStackSet.

#### Sample Request

```

https://cloudformation.us-east-1.amazonaws.com/
 ?Action=DescribeStackSet
 &Version=2010-05-15
 &StackSetName=stack-set-example
 &X-Amz-Algorithm=AWS4-HMAC-SHA256
 &X-Amz-Credential=[Access key ID and scope]
 &X-Amz-Date=20170810T233349Z
 &X-Amz-SignedHeaders=content-type;host
 &X-Amz-Signature=[Signature]
```

#### Sample Response

```

<DescribeStackSetResponse xmlns="http://internal.amazon.com/coral/com.amazonaws.maestro.service.v20160713/">
  <DescribeStackSetResult>
    <StackSet>
      <Capabilities>
        <member>CAPABILITY_IAM</member>
      </Capabilities>
      <StackSetId>stack-set-example:c14cd6d1-cd17-40bd-82ed-ff97example</StackSetId>
      <TemplateBody>
        [details omitted]
      </TemplateBody>
      <StackSetName>stack-set-example</StackSetName>
      <Description>Enable AWS Config</Description>
      <Parameters>
        <member>
          <ParameterKey>AllSupported</ParameterKey>
          <UsePreviousValue>false</UsePreviousValue>
          <ParameterValue>true</ParameterValue>
        </member>
        <member>
          <ParameterKey>DeliveryChannelName</ParameterKey>
          <UsePreviousValue>false</UsePreviousValue>
          <ParameterValue><Generated></ParameterValue>
        </member>
        <member>
          <ParameterKey>Frequency</ParameterKey>
          <UsePreviousValue>false</UsePreviousValue>
          <ParameterValue>24hours</ParameterValue>
        </member>
        <member>
          <ParameterKey>IncludeGlobalResourceTypes</ParameterKey>
          <UsePreviousValue>false</UsePreviousValue>
          <ParameterValue>true</ParameterValue>
        </member>
        <member>
          <ParameterKey>NotificationEmail</ParameterKey>
          <UsePreviousValue>false</UsePreviousValue>
          <ParameterValue><None></ParameterValue>
        </member>
        <member>
          <ParameterKey>ResourceTypes</ParameterKey>
          <UsePreviousValue>false</UsePreviousValue>
          <ParameterValue><All></ParameterValue>
        </member>
        <member>
          <ParameterKey>TopicArn</ParameterKey>
          <UsePreviousValue>false</UsePreviousValue>
          <ParameterValue><New Topic></ParameterValue>
        </member>
      </Parameters>
      <Tags>
        <member>
          <Value>marketing</Value>
          <Key>business-unit</Key>
        </member>
      </Tags>
      <StackSetDriftDetectionDetails>
        <DriftDetectionStatus>COMPLETED</DriftDetectionStatus>
        <InSyncStackInstancesCount>5</InSyncStackInstancesCount>
        <FailedStackInstancesCount>0</FailedStackInstancesCount>
        <DriftStatus>IN_SYNC</DriftStatus>
        <TotalStackInstancesCount>5</TotalStackInstancesCount>
        <DriftedStackInstancesCount>0</DriftedStackInstancesCount>
        <InProgressStackInstancesCount>0</InProgressStackInstancesCount>
        <LastDriftCheckTimestamp>2019-12-03T20:00:27.877Z</LastDriftCheckTimestamp>
      </StackSetDriftDetectionDetails>
      <Status>ACTIVE</Status>
    </StackSet>
  </DescribeStackSetResult>
  <ResponseMetadata>
    <RequestId>48d13e76-794b-11e7-95e6-f946example</RequestId>
  </ResponseMetadata>
</DescribeStackSetResponse>
```

### How to view which Regions a StackSet has stack instances deployed to.

Use the `DescribeStackSets` API to output a list of Regions. This list
consists of Regions where a given StackSet has stack instances deployed.

In the following example, the StackSet named "MyStackSet" has stack instances deployed
in Regions "us-east-1" and "us-west-2":

```

{
    "StackSet": {
        "StackSetName": "MyStackSet",
        //...
        "Regions": [
            "us-east-1",
            "us-west-2"
        ]
    }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/cloudformation-2010-05-15/describestackset.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/cloudformation-2010-05-15/describestackset.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/describestackset.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/cloudformation-2010-05-15/describestackset.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/describestackset.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/cloudformation-2010-05-15/describestackset.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/cloudformation-2010-05-15/describestackset.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/cloudformation-2010-05-15/describestackset.md)

- [AWS SDK for Python](../../../../services/goto/boto3/cloudformation-2010-05-15/describestackset.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/describestackset.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeStacks

DescribeStackSetOperation

All content copied from https://docs.aws.amazon.com/.
