# Amazon ECR Public service quotas

The following table provides the default service quotas for Amazon ECR Public. For unauthenticated customers, Amazon ECR Public supports up to 500GB of data per month. This is the max amount of data supported and isn't adjustable.

NameDefaultAdjustableDescriptionImages per repositoryEach supported Region: 100,000[Yes](https://console.aws.amazon.com/servicequotas/home/services/ecr-public/quotas/L-F0CAD50B)The maximum number of images per repository.Layer partsEach supported Region: 1,000NoThe maximum number of layer parts. This is only applicable if you are using Amazon ECR API actions directly to initiate multipart uploads for image push operations.Maximum layer part sizeEach supported Region: 10NoThe maximum size (MiB) of a layer part. This is only applicable if you are using Amazon ECR API actions directly to initiate multipart uploads for image push operations.Maximum layer sizeEach supported Region: 10,000NoThe maximum size per layer.Minimum layer part sizeEach supported Region: 5NoThe minimum size (MiB) of a layer part. This is only applicable if you are using Amazon ECR API actions directly to initiate multipart uploads for image push operations.Rate of BatchCheckLayerAvailability requestsEach supported Region: 200 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/ecr-public/quotas/L-A1FEAB1E)The maximum number of BatchCheckLayerAvailability requests that you can make per second in the current Region. When an image is pushed to a repository, each image layer is checked to verify if it has been uploaded before. If it has been uploaded, then the image layer is skipped.Rate of CompleteLayerUpload requestsEach supported Region: 10 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/ecr-public/quotas/L-8FA37FDE)The maximum number of CompleteLayerUpload requests that you can make per second in the current Region. When an image is pushed, the CompleteLayerUpload API is called once per each new image layer to verify that the upload has completed.Rate of GetAuthorizationToken requestsEach supported Region: 200 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/ecr-public/quotas/L-B5773E0E)The maximum number of GetAuthorizationToken requests that you can make per second in the current Region.Rate of InitiateLayerUpload requestsEach supported Region: 10 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/ecr-public/quotas/L-1EDF7132)The maximum number of InitiateLayerUpload requests that you can make per second in the current Region. When an image is pushed, the InitiateLayerUpload API is called once per image layer that has not already been uploaded. Whether or not an image layer has been uploaded is determined by the BatchCheckLayerAvailability API action.Rate of PutImage requestsEach supported Region: 10 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/ecr-public/quotas/L-EF0362CE)The maximum number of PutImage requests that you can make per second in the current Region. When an image is pushed and all new image layers have been uploaded, the PutImage API is called once to create or update the image manifest and the tags associated with the image.Rate of UploadLayerPart requestsEach supported Region: 260 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/ecr-public/quotas/L-B6387A7C)The maximum number of UploadLayerPart requests that you can make per second in the current Region. When an image is pushed, each new image layer is uploaded in parts and the UploadLayerPart API is called once per each new image layer part.Rate of authenticated image pullsEach supported Region: 10 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/ecr-public/quotas/L-33392ECC)The maximum number of authenticated image pulls per second.Rate of image pulls to AWS resourcesEach supported Region: 10 per secondNoThe maximum number of image pulls per second to resources running on Amazon ECS, Fargate, or Amazon EC2.Rate of unauthenticated image pullsEach supported Region: 1 per secondNoThe maximum number of unauthenticated image pulls per second.Registered repositoriesEach supported Region: 10,000[Yes](https://console.aws.amazon.com/servicequotas/home/services/ecr-public/quotas/L-502CB705)The maximum number of repositories that you can create in this account in the current Region.Tags per imageEach supported Region: 1,000NoThe maximum number of tags per image.

## Managing your Amazon ECR Public service quotas in the AWS Management Console

Amazon ECR Public has integrated with Service Quotas, an AWS service that enables you to view and
manage your quotas from a central location. For more information, see [What Is Service\
Quotas?](../../../servicequotas/latest/userguide/intro.md) in the _Service Quotas User Guide_.

Service Quotas makes it easy to look up the value of all Amazon ECR service quotas.

###### To view Amazon ECR service quotas (AWS Management Console)

1. Open the Service Quotas console at [https://console.aws.amazon.com/servicequotas/](https://console.aws.amazon.com/servicequotas).

2. In the navigation pane, choose **AWS services**.

3. From the **AWS services** list, search for and select
    **[Amazon Elastic Container Registry Public (Amazon ECR\**
**Public)](https://us-east-1.console.aws.amazon.com/servicequotas/home/services/ecr-public/quotas)**.

In the **Service quotas** list, you can see the service quota
    name, applied value (if it is available), AWS default quota, and whether the
    quota value is adjustable.

4. To view additional information about a service quota, such as the description,
    choose the quota name.

To request a quota increase, see [Requesting a quota\
increase](../../../servicequotas/latest/userguide/request-increase.md) in the _Service Quotas User Guide_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Logging actions with AWS CloudTrail

Troubleshooting

All content copied from https://docs.aws.amazon.com/.
