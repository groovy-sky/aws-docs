# Scenarios for Amazon S3 using AWS SDKs

The following code examples show you how to implement common scenarios in Amazon S3
with AWS SDKs. These scenarios show you how to accomplish specific tasks by calling multiple functions
within Amazon S3 or combined with other AWS services.
Each scenario includes a link to the complete source code, where you can find instructions on how to set up and run the code.

Scenarios target an intermediate level of experience to help you understand service actions in context.

###### Examples

- [Check if a bucket exists](https://docs.aws.amazon.com/AmazonS3/latest/API/s3_example_s3_Scenario_DoesBucketExist_section.html)

- [Convert text to speech and back to text](https://docs.aws.amazon.com/AmazonS3/latest/API/s3_example_cross_Telephone_section.html)

- [Create a presigned URL](s3-example-s3-scenario-presignedurl-section.md)

- [Create a serverless application to manage photos](https://docs.aws.amazon.com/AmazonS3/latest/API/s3_example_cross_PAM_section.html)

- [Create a web page that lists Amazon S3 objects](https://docs.aws.amazon.com/AmazonS3/latest/API/s3_example_s3_Scenario_ListObjectsWeb_section.html)

- [Create an Amazon Textract explorer application](https://docs.aws.amazon.com/AmazonS3/latest/API/s3_example_cross_TextractExplorer_section.html)

- [Delete all objects in a bucket](https://docs.aws.amazon.com/AmazonS3/latest/API/s3_example_s3_Scenario_DeleteAllObjects_section.html)

- [Delete incomplete multipart uploads](s3-example-s3-scenario-abortmultipartupload-section.md)

- [Detect PPE in images](https://docs.aws.amazon.com/AmazonS3/latest/API/s3_example_cross_RekognitionPhotoAnalyzerPPE_section.html)

- [Detect entities in text extracted from an image](https://docs.aws.amazon.com/AmazonS3/latest/API/s3_example_cross_TextractComprehendDetectEntities_section.html)

- [Detect faces in an image](https://docs.aws.amazon.com/AmazonS3/latest/API/s3_example_cross_DetectFaces_section.html)

- [Detect objects in images](https://docs.aws.amazon.com/AmazonS3/latest/API/s3_example_cross_RekognitionPhotoAnalyzer_section.html)

- [Detect people and objects in a video](https://docs.aws.amazon.com/AmazonS3/latest/API/s3_example_cross_RekognitionVideoDetection_section.html)

- [Download S3 'directories'](https://docs.aws.amazon.com/AmazonS3/latest/API/s3_example_s3_Scenario_DownloadS3Directory_section.html)

- [Download objects to a local directory](https://docs.aws.amazon.com/AmazonS3/latest/API/s3_example_s3_DownloadBucketToDirectory_section.html)

- [Download stream of unknown size](https://docs.aws.amazon.com/AmazonS3/latest/API/s3_example_s3_Scenario_DownloadStream_section.html)

- [Get an object from a Multi-Region Access Point](https://docs.aws.amazon.com/AmazonS3/latest/API/s3_example_s3_GetObject_MRAP_section.html)

- [Get an object from a bucket if it has been modified](https://docs.aws.amazon.com/AmazonS3/latest/API/s3_example_s3_GetObject_IfModifiedSince_section.html)

- [Get started with S3](s3-example-s3-gettingstarted-section.md)

- [Get started with encryption](https://docs.aws.amazon.com/AmazonS3/latest/API/s3_example_s3_Encryption_section.html)

- [Get started with tags](https://docs.aws.amazon.com/AmazonS3/latest/API/s3_example_s3_Scenario_Tagging_section.html)

- [Lock Amazon S3 objects](s3-example-s3-scenario-objectlock-section.md)

- [Make conditional requests](https://docs.aws.amazon.com/AmazonS3/latest/API/s3_example_s3_Scenario_ConditionalRequests_section.html)

- [Manage access control lists (ACLs)](https://docs.aws.amazon.com/AmazonS3/latest/API/s3_example_s3_Scenario_ManageACLs_section.html)

- [Manage large messages using S3](s3-example-sqs-scenario-sqsextendedclient-section.md)

- [Manage versioned objects in batches with a Lambda function](https://docs.aws.amazon.com/AmazonS3/latest/API/s3_example_s3_Scenario_BatchObjectVersioning_section.html)

- [Parse URIs](https://docs.aws.amazon.com/AmazonS3/latest/API/s3_example_s3_Scenario_URIParsing_section.html)

- [Perform a multipart copy](https://docs.aws.amazon.com/AmazonS3/latest/API/s3_example_s3_MultipartCopy_section.html)

- [Process S3 event notifications](https://docs.aws.amazon.com/AmazonS3/latest/API/s3_example_s3_Scenario_ProcessS3EventNotification_section.html)

- [Save EXIF and other image information](https://docs.aws.amazon.com/AmazonS3/latest/API/s3_example_cross_DetectLabels_section.html)

- [Send event notifications to EventBridge](https://docs.aws.amazon.com/AmazonS3/latest/API/s3_example_s3_Scenario_PutBucketNotificationConfiguration_section.html)

- [Track uploads and downloads](https://docs.aws.amazon.com/AmazonS3/latest/API/s3_example_s3_Scenario_TrackUploadDownload_section.html)

- [Transform data with S3 Object Lambda](https://docs.aws.amazon.com/AmazonS3/latest/API/s3_example_cross_ServerlessS3DataTransformation_section.html)

- [Unit and integration test with an SDK](https://docs.aws.amazon.com/AmazonS3/latest/API/s3_example_cross_Testing_section.html)

- [Upload directory to a bucket](https://docs.aws.amazon.com/AmazonS3/latest/API/s3_example_s3_UploadDirectoryToBucket_section.html)

- [Upload or download large files](https://docs.aws.amazon.com/AmazonS3/latest/API/s3_example_s3_Scenario_UsingLargeFiles_section.html)

- [Upload stream of unknown size](https://docs.aws.amazon.com/AmazonS3/latest/API/s3_example_s3_Scenario_UploadStream_section.html)

- [Use checksums](https://docs.aws.amazon.com/AmazonS3/latest/API/s3_example_s3_Scenario_UseChecksums_section.html)

- [Work with Amazon S3 object integrity](https://docs.aws.amazon.com/AmazonS3/latest/API/s3_example_s3_Scenario_ObjectIntegrity_section.html)

- [Work with versioned objects](s3-example-s3-scenario-objectversioningusage-section.md)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

UploadPartCopy

Check if a bucket exists
