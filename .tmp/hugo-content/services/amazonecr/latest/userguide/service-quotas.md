# Amazon ECR service quotas

The following table provides the default service quotas for Amazon Elastic Container Registry (Amazon ECR).

NameDefaultAdjustableDescriptionBasic image scans per 24 hoursEach supported Region: 100,000NoThe maximum number of images that can be scanned within a 24 hour period in the current account and region using basic scanning. This limit includes both scan on push and manual scans.Filters per rule in a replication configurationEach supported Region: 100NoThe maximum number of filters per rule in a replication configuration.Images per repositoryEach supported Region: 100,000[Yes](https://console.aws.amazon.com/servicequotas/home/services/ecr/quotas/L-03A36CE1)The maximum number of images per repository.Layer partsEach supported Region: 4,200NoThe maximum number of layer parts. This is only applicable if you are using Amazon ECR API actions directly to initiate multipart uploads for image push operations.Lifecycle policy lengthEach supported Region: 30,720NoThe maximum number of characters in a lifecycle policy.Maximum layer part sizeEach supported Region: 10NoThe maximum size (MiB) of a layer part. This is only applicable if you are using Amazon ECR API actions directly to initiate multipart uploads for image push operations.Maximum layer sizeEach supported Region: 52,000NoThe maximum size (MiB) of a layer.Minimum layer part sizeEach supported Region: 5NoThe minimum size (MiB) of a layer part. This is only applicable if you are using Amazon ECR API actions directly to initiate multipart uploads for image push operations.Pull through cache rules per registryEach supported Region: 50NoThe maximum number of pull-through cache rules.Rate of BatchCheckLayerAvailability requestsEach supported Region: 1,000 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/ecr/quotas/L-B9173138)The maximum number of BatchCheckLayerAvailability requests that you can make per second in the current Region. When an image is pushed to a repository, each image layer is checked to verify if it has been uploaded before. If it has been uploaded, then the image layer is skipped.Rate of BatchGetImage requestsEach supported Region: 2,000 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/ecr/quotas/L-16E70933)The maximum number of BatchGetImage requests that you can make per second in the current Region. When an image is pulled, the BatchGetImage API is called once to retrieve the image manifest. If you request a quota increase for this API, review your GetDownloadUrlForLayer usage as well.Rate of CompleteLayerUpload requestsEach supported Region: 100 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/ecr/quotas/L-44194860)The maximum number of CompleteLayerUpload requests that you can make per second in the current Region. When an image is pushed, the CompleteLayerUpload API is called once per each new image layer to verify that the upload has completed.Rate of GetAuthorizationToken requestsEach supported Region: 500 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/ecr/quotas/L-55A41110)The maximum number of GetAuthorizationToken requests that you can make per second in the current Region.Rate of GetDownloadUrlForLayer requestsEach supported Region: 3,000 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/ecr/quotas/L-A60A366D)The maximum number of GetDownloadUrlForLayer requests that you can make per second in the current Region. When an image is pulled, the GetDownloadUrlForLayer API is called once per image layer that is not already cached. If you request a quota increase for this API, review your BatchGetImage usage as well.Rate of InitiateLayerUpload requestsEach supported Region: 100 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/ecr/quotas/L-95B28F8D)The maximum number of InitiateLayerUpload requests that you can make per second in the current Region. When an image is pushed, the InitiateLayerUpload API is called once per image layer that has not already been uploaded. Whether or not an image layer has been uploaded is determined by the BatchCheckLayerAvailability API action.Rate of PutImage requestsEach supported Region: 10 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/ecr/quotas/L-AD52DFB2)The maximum number of PutImage requests that you can make per second in the current Region. When an image is pushed and all new image layers have been uploaded, the PutImage API is called once to create or update the image manifest and the tags associated with the image.Rate of UploadLayerPart requestsEach supported Region: 500 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/ecr/quotas/L-A1670B10)The maximum number of UploadLayerPart requests that you can make per second in the current Region. When an image is pushed, each new image layer is uploaded in parts and the UploadLayerPart API is called once per each new image layer part.Rate of image scansEach supported Region: 1NoThe maximum number of image scans per image, per 24 hours.Registered repositoriesEach supported Region: 100,000[Yes](https://console.aws.amazon.com/servicequotas/home/services/ecr/quotas/L-CFEB8E8D)The maximum number of repositories that you can create in this account in the current Region.Rules per lifecycle policyEach supported Region: 50NoThe maximum number of rules in a lifecycle policyRules per replication configurationEach supported Region: 10NoThe maximum number of rules in a replication configuration.Tags per imageEach supported Region: 1,000NoThe maximum number of tags per image.Unique destinations across all rules in a replication configurationEach supported Region: 25NoThe maximum number of unique destinations across all rules in a replication configuration.

## Managing your Amazon ECR service quotas in the AWS Management Console

Amazon ECR has integrated with Service Quotas, an AWS service that enables you to view and
manage your quotas from a central location. For more information, see [What Is Service\
Quotas?](../../../servicequotas/latest/userguide/intro.md) in the _Service Quotas User Guide_.

Service Quotas makes it easy to look up the value of all Amazon ECR service quotas.

###### To view Amazon ECR service quotas (AWS Management Console)

1. Open the Service Quotas console at [https://console.aws.amazon.com/servicequotas/](https://console.aws.amazon.com/servicequotas).

2. In the navigation pane, choose **AWS services**.

3. From the **AWS services** list, search for and select
    **Amazon Elastic Container Registry (Amazon ECR)**.

In the **Service quotas** list, you can see the service quota
    name, applied value (if it is available), AWS default quota, and whether the
    quota value is adjustable.

4. To view additional information about a service quota, such as the description,
    choose the quota name.

To request a quota increase, see [Requesting a quota\
increase](../../../servicequotas/latest/userguide/request-increase.md) in the _Service Quotas User Guide_.

### Creating a CloudWatch alarm to monitor API usage metrics

Amazon ECR provides CloudWatch usage metrics that correspond to the AWS service quotas for
each of the APIs involved with the registry authentication, image push, and image
pull actions. In the Service Quotas console, you can visualize your usage on a graph and
configure alarms that alert you when your usage approaches a service quota. For more
information, see [Amazon ECR usage metrics](monitoring-usage.md).

Use the following steps to create a CloudWatch alarm based on one of the Amazon ECR API usage
metrics.

###### To create an alarm based on your Amazon ECR usage quotas (AWS Management Console)

1. Open the Service Quotas console at [https://console.aws.amazon.com/servicequotas/](https://console.aws.amazon.com/servicequotas).

2. In the navigation pane, choose **AWS services**.

3. From the **AWS services** list, search for and select
    **Amazon Elastic Container Registry (Amazon ECR)**.

4. In the **Service quotas** list, select the Amazon ECR usage
    quota you want to create an alarm for.

5. In the Amazon CloudWatch Events alarms section, choose
    **Create**.

6. For **Alarm threshold**, choose the percentage of your
    applied quota value that you want to set as the alarm value.

7. For **Alarm name**, enter a name for the alarm and then
    choose **Create**.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StartLifecyclePolicyPreview

Troubleshooting

All content copied from https://docs.aws.amazon.com/.
