# Deleting a metrics filter

You can delete an Amazon CloudWatch request metrics filter if you no longer need it. When you
delete a filter, you are no longer charged for request metrics that use that
_specific filter_. However, you will continue to be charged for any
other filter configurations that exist.

When you delete a filter, you can no longer use the filter for request metrics. Deleting a
filter cannot be undone.

For information about creating a request metrics filter, see the following topics:

- [Creating a CloudWatch metrics configuration for all the objects in your bucket](configure-request-metrics-bucket.md)

- [Creating a metrics configuration that filters by prefix, object tag, or access point](metrics-configurations-filter.md)

1. Sign in to the AWS Management Console and open the Amazon S3 console at
    [https://console.aws.amazon.com/s3/](https://console.aws.amazon.com/s3).

2. In the left navigation pane, choose **General purpose buckets**.

3. In the buckets list, choose the name of the bucket you want to delete a request metrics filter for.

4. Choose the **Metrics** tab.

5. Under **Bucket metrics**, choose **View**
**additional charts**.

6. Choose the **Request metrics** tab.

7. Choose **Manage filters**.

8. Choose your filter.

###### Important

Deleting a filter cannot be undone.

9. Choose **Delete**.

Amazon S3 deletes your filter.

You can also add metrics configurations programmatically with the Amazon S3 REST API.
For more information about adding and working with metrics configurations, see the following topics in the
_Amazon Simple Storage Service API Reference_:

- [PUT Bucket\
Metric Configuration](../api/restbucketputmetricconfiguration.md)

- [GET Bucket\
Metric Configuration](../api/restbucketgetmetricconfiguration.md)

- [List\
Bucket Metric Configuration](../api/restlistbucketmetricsconfiguration.md)

- [DELETE\
Bucket Metric Configuration](../api/restdeletebucketmetricsconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Filtering by prefix, object tag, or access point

Amazon S3 Event Notifications

All content copied from https://docs.aws.amazon.com/.
