---
title: "Work with flow logs"
---

# Work with flow logs

You can work with flow logs using consoles for Amazon EC2 and Amazon VPC.

###### Tasks

- [1\. Control the use of flow logs with IAM](#controlling-use-of-flow-logs)

- [2\. Create a flow log](#create-flow-log)

- [3\. Tag a flow log](#modify-tags-flow-logs)

- [4\. Delete a flow log](#delete-flow-log)

- [Command line overview](#flow-logs-api-cli)

## 1\. Control the use of flow logs with IAM

By default, users do not have permission to work with flow logs. You can
create an IAM role with a policy attached that grants users the permissions to create, describe,
and delete flow logs.

The following is an example policy that grants users full permissions to create,
describe, and delete flow logs.

JSON

```json

{
  "Version":"2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "ec2:DeleteFlowLogs",
        "ec2:CreateFlowLogs",
        "ec2:DescribeFlowLogs"
      ],
      "Resource": "*"
    }
  ]
}

```

For more information, see [How Amazon VPC works with IAM](security-iam-service-with-iam.md).

## 2\. Create a flow log

You can create flow logs for your VPCs, subnets, or network interfaces. When
you create a flow log, you must specify a destination for the flow log. For more
information, see the following:

- [Create a flow log that publishes to CloudWatch Logs](flow-logs-cwl-create-flow-log.md)

- [Create a flow log that publishes to Amazon S3](flow-logs-s3-create-flow-log.md)

- [Create a flow log that publishes to Amazon Data Firehose](flow-logs-firehose-create-flow-log.md)

## 3\. Tag a flow log

You can add or remove tags for a flow log at any time.

###### To manage tags for a flow log

1. Do one of the following:
   - Open the Amazon EC2 console at [https://console.aws.amazon.com/ec2/](https://console.aws.amazon.com/ec2).
      In the navigation pane, choose **Network Interfaces**.
      Select the checkbox for the network interface.

   - Open the Amazon VPC console at [https://console.aws.amazon.com/vpc/](https://console.aws.amazon.com/vpc).
      In the navigation pane, choose **Your VPCs**.
      Select the checkbox for the VPC.

   - Open the Amazon VPC console at [https://console.aws.amazon.com/vpc/](https://console.aws.amazon.com/vpc).
      In the navigation pane, choose **Subnets**.
      Select the checkbox for the subnet.
2. Choose **Flow Logs**.

3. Choose **Actions**, **Manage tags**.

4. To add a new tag, choose **Add new tag** and enter the key
    and value. To remove a tag, choose **Remove**.

5. When you are finished adding or removing tags, choose **Save**.

## 4\. Delete a flow log

You can delete a flow log at any time. After you delete a flow log, it can take
several minutes to stop collecting data.

Deleting a flow log does not delete the log data from the destination or modify
the destination resource. You must delete the existing flow log data directly from
the destination, and clean up the destination resource, using the console for the
destination service.

###### To delete a flow log

1. Do one of the following:
   - Open the Amazon EC2 console at [https://console.aws.amazon.com/ec2/](https://console.aws.amazon.com/ec2).
      In the navigation pane, choose **Network Interfaces**.
      Select the checkbox for the network interface.

   - Open the Amazon VPC console at [https://console.aws.amazon.com/vpc/](https://console.aws.amazon.com/vpc).
      In the navigation pane, choose **Your VPCs**.
      Select the checkbox for the VPC.

   - Open the Amazon VPC console at [https://console.aws.amazon.com/vpc/](https://console.aws.amazon.com/vpc).
      In the navigation pane, choose **Subnets**.
      Select the checkbox for the subnet.
2. Choose **Flow Logs**.

3. Choose **Actions**, **Delete flow logs**.

4. When prompted for confirmation, type `delete` and then
    choose **Delete**.

## Command line overview

You can perform the tasks described on this page using the command line.

###### Create a flow log

- [create-flow-logs](../../../cli/latest/reference/ec2/create-flow-logs.md) (AWS CLI)

- [New-EC2FlowLog](../../../powershell/latest/reference/items/new-ec2flowlog.md)
(AWS Tools for Windows PowerShell)

###### Describe a flow log

- [describe-flow-logs](../../../cli/latest/reference/ec2/describe-flow-logs.md) (AWS CLI)

- [Get-EC2FlowLog](../../../powershell/latest/reference/items/get-ec2flowlog.md)
(AWS Tools for Windows PowerShell)

###### Tag a flow log

- [create-tags](../../../cli/latest/reference/ec2/create-tags.md) and
[delete-tags](../../../cli/latest/reference/ec2/delete-tags.md)
(AWS CLI)

- [New-EC2Tag](../../../powershell/latest/reference/items/new-ec2tag.md) and
[Remove-EC2Tag](../../../powershell/latest/reference/items/remove-ec2tag.md)
(AWS Tools for Windows PowerShell)

###### Delete a flow log

- [delete-flow-logs](../../../cli/latest/reference/ec2/delete-flow-logs.md) (AWS CLI)

- [Remove-EC2FlowLog](../../../powershell/latest/reference/items/remove-ec2flowlog.md) (AWS Tools for Windows PowerShell)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Flow log limitations

Publish to CloudWatch Logs

All content copied from https://docs.aws.amazon.com/.
