# WorkSpaces Applications Information in CloudTrail

CloudTrail is enabled on your AWS account when you create the account. When supported event
activity occurs in WorkSpaces Applications, that activity is recorded in a CloudTrail event along with other
AWS service events in **Event history**. You can view, search, and download
recent events in your AWS account. For more information, see [Viewing Events with CloudTrail Event\
History](../../../awscloudtrail/latest/userguide/view-cloudtrail-events.md).

For an ongoing record of events in your AWS account, including events for WorkSpaces Applications,
create a trail. A _trail_ enables CloudTrail to deliver log files to an Amazon S3 bucket.
By default, when you create a trail in the console, the trail applies to all AWS Regions. The trail logs events
from all Regions in the AWS partition and delivers the log files to the Amazon S3 bucket that you
specify. Additionally, you can configure other AWS services to further analyze and act upon
the event data collected in CloudTrail logs. For more information, see the following:

- [Overview for\
Creating a Trail](../../../awscloudtrail/latest/userguide/cloudtrail-create-and-update-a-trail.md)

- [CloudTrail Supported Services and Integrations](../../../awscloudtrail/latest/userguide/cloudtrail-aws-service-specific-topics.md#cloudtrail-aws-service-specific-topics-integrations)

- [Configuring Amazon SNS\
Notifications for CloudTrail](../../../awscloudtrail/latest/userguide/getting-notifications-top-level.md)

- [Receiving CloudTrail Log Files from Multiple Regions](../../../awscloudtrail/latest/userguide/receive-cloudtrail-log-files-from-multiple-regions.md) and [Receiving CloudTrail\
Log Files from Multiple Accounts](../../../awscloudtrail/latest/userguide/cloudtrail-receive-logs-from-multiple-accounts.md)

WorkSpaces Applications supports logging the following actions as events in CloudTrail log files:

- [AssociateFleet](../../../../reference/appstream2/latest/apireference/api-associatefleet.md)

- [BatchAssociateUserStack](../../../../reference/appstream2/latest/apireference/api-batchassociateuserstack.md)

- [BatchDisassociateUserStack](../../../../reference/appstream2/latest/apireference/api-batchdisassociateuserstack.md)

- [CopyImage](../../../../reference/appstream2/latest/apireference/api-copyimage.md)

- [CreateDirectoryConfig](../../../../reference/appstream2/latest/apireference/api-createdirectoryconfig.md)

- [CreateFleet](../../../../reference/appstream2/latest/apireference/api-createfleet.md)

- [CreateImageBuilder](../../../../reference/appstream2/latest/apireference/api-createimagebuilder.md)

- [CreateImageBuilderStreamingURL](../../../../reference/appstream2/latest/apireference/api-createimagebuilderstreamingurl.md)

- [CreateStack](../../../../reference/appstream2/latest/apireference/api-createstack.md)

- [CreateStreamingURL](../../../../reference/appstream2/latest/apireference/api-createstreamingurl.md)

- [DeleteDirectoryConfig](../../../../reference/appstream2/latest/apireference/api-deletedirectoryconfig.md)

- [DeleteFleet](../../../../reference/appstream2/latest/apireference/api-deletefleet.md)

- [DeleteImage](../../../../reference/appstream2/latest/apireference/api-deleteimage.md)

- [DeleteImageBuilder](../../../../reference/appstream2/latest/apireference/api-deleteimagebuilder.md)

- [DeleteImagePermissions](../../../../reference/appstream2/latest/apireference/api-deleteimagepermissions.md)

- [DeleteStack](../../../../reference/appstream2/latest/apireference/api-deletestack.md)

- [DescribeDirectoryConfigs](../../../../reference/appstream2/latest/apireference/api-describedirectoryconfigs.md)

- [DescribeFleets](../../../../reference/appstream2/latest/apireference/api-describefleets.md)

- [DescribeImageBuilders](../../../../reference/appstream2/latest/apireference/api-describeimagebuilders.md)

- [DescribeImagePermissions](../../../../reference/appstream2/latest/apireference/api-describeimagepermissions.md)

- [DescribeImages](../../../../reference/appstream2/latest/apireference/api-describeimages.md)

- [DescribeSessions](../../../../reference/appstream2/latest/apireference/api-describesessions.md)

- [DescribeStacks](../../../../reference/appstream2/latest/apireference/api-describestacks.md)

- [DescribeUserStackAssociations](../../../../reference/appstream2/latest/apireference/api-describeuserstackassociations.md)

- [ExpireSession](../../../../reference/appstream2/latest/apireference/api-expiresession.md)

- [ListAssociatedFleets](../../../../reference/appstream2/latest/apireference/api-listassociatedfleets.md)

- [ListAssociatedStacks](../../../../reference/appstream2/latest/apireference/api-listassociatedstacks.md)

- [ListTagsForResource](../../../../reference/appstream2/latest/apireference/api-listtagsforresource.md)

- [StartFleet](../../../../reference/appstream2/latest/apireference/api-startfleet.md)

- [StartImageBuilder](../../../../reference/appstream2/latest/apireference/api-startimagebuilder.md)

- [StopFleet](../../../../reference/appstream2/latest/apireference/api-stopfleet.md)

- [StopImageBuilder](../../../../reference/appstream2/latest/apireference/api-stopimagebuilder.md)

- [TagResource](../../../../reference/appstream2/latest/apireference/api-tagresource.md)

- [UntagResource](../../../../reference/appstream2/latest/apireference/api-untagresource.md)

- [UpdateDirectoryConfig](../../../../reference/appstream2/latest/apireference/api-updatedirectoryconfig.md)

- [UpdateFleet](../../../../reference/appstream2/latest/apireference/api-updatefleet.md)

- [UpdateImagePermissions](../../../../reference/appstream2/latest/apireference/api-updateimagepermissions.md)

- [UpdateStack](../../../../reference/appstream2/latest/apireference/api-updatestack.md)

Every event or log entry contains information about who generated the request. The
identity information helps you determine the following:

- Whether the request was made with root or IAM user credentials.

- Whether the request was made with temporary security credentials for a role or
federated user.

- Whether the request was made by another AWS service.

For more information, see the [CloudTrail userIdentity\
Element](../../../awscloudtrail/latest/userguide/cloudtrail-event-reference-user-identity.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Logging WorkSpaces Applications API Calls

Example: WorkSpaces Applications Log File Entries

All content copied from https://docs.aws.amazon.com/.
