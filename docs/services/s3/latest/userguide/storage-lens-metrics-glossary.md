# Amazon S3 Storage Lens metrics glossary

The Amazon S3 Storage Lens metrics glossary provides a complete list of free and advanced metrics for
S3 Storage Lens.

S3 Storage Lens offers free metrics for all dashboards and configurations, with the option to
upgrade to advanced metrics.

- **Free metrics** contain metrics that are relevant to your
storage usage, such as the number of buckets and the objects in your account. Free metrics
also include use-case based metrics, such as cost-optimization and data-protection metrics.
All free metrics are collected daily, and data is available for queries for up to 14 days.

- **Advanced metrics** include all the metrics in free
metrics along with additional metrics, such as advanced performance, advanced data
protection, and advanced cost optimization metrics. Advanced metrics also include additional
metric categories, such as activity metrics and detailed status-code metrics. Advanced
metrics data is available for queries for 15 months.

There are additional charges when you use S3 Storage Lens with advanced metrics and
recommendations. For more information, see [Amazon S3\
pricing](https://aws.amazon.com/s3/pricing). For more information about advanced metrics and recommendations features,
see [Metrics selection](storage-lens-basics-metrics-recommendations.md#storage_lens_basics_metrics_selection).

###### Note

For Storage Lens groups, only free tier storage metrics are available. Advanced tier
metrics are not available at the Storage Lens group level.

###### Metric names

The **Metric name** column in the following table provides
the name of each S3 Storage Lens in the S3 console. The **CloudWatch and**
**export** column provides the name of each metric in Amazon CloudWatch and the metrics export
file that you can configure in your S3 Storage Lens dashboard.

###### Derived metric formulas

Derived metrics are not available for the metrics export and the CloudWatch publishing option.
However, you can use the metrics formulas shown in the **Derived metrics**
**formula** column to compute them.

###### Interpreting the Amazon S3 Storage Lens prefix symbols for metrics unit multiples (K, M, G, and so on)

S3 Storage Lens metrics unit multiples are written with prefix symbols. These prefix symbols
match the International System of Units (SI) symbols that are standardized by the
International Bureau of Weights and Measures (BIPM). These symbols are also used in the
Unified Code for Units of Measure (UCUM). For more information, see [List of SI prefix\
symbols](https://www.bipm.org/en/measurement-units/si-prefixes).

###### Note

- The unit of measurement for S3 storage bytes is in binary gigabytes (GB), where 1 GB
is 230 bytes, 1 TB is 240 bytes, and
1 PB is 250 bytes. This unit of measurement is also known as a
gibibyte (GiB), as defined by the International Electrotechnical Commission (IEC).

- When an object reaches the end of its lifetime based on its lifecycle configuration,
Amazon S3 queues the object for removal and removes it asynchronously. Therefore, there might
be a delay between the expiration date and the date when Amazon S3 removes an object. S3 Storage Lens
doesn't include metrics for objects that have expired but haven't been removed. For more
information about expiration actions in S3 Lifecycle, see [Expiring objects](lifecycle-expire-general-considerations.md).

- Amazon S3 stores metadata (object key, timestamps, etc.) for every object, which requires
minimum storage even for 0KB data files. This is why 0KB objects appear in the (0KB-128KB\]
size range in S3 Storage Lens.

- S3 Storage Lens provides best-effort tracking of cross-region data transfers, primarily
focusing on requests from customer-managed resources like EC2 instances. Requests made
through AWS PrivateLink or certain in-Region requests are unclassified.

The following table shows the S3 Storage Lens metrics glossary.

Metric nameCloudWatch and exportDescriptionTier1Category2DerivedDerived metric formulaStorage Lens groupsTotal storageStorageBytesTotal storage, inclusive of incomplete multipart uploads, object metadata, and
delete markersFreeSummaryN-YObject countObjectCountTotal object countFreeSummaryN-YAverage object size-Average object sizeFreeSummaryYsum(StorageBytes)/sum(ObjectCount)YActive buckets-Number of buckets with storage > 0 bytesFreeSummaryY-YBuckets-Number of bucketsFreeSummaryY-YAccounts-Number of accounts whose storage is in scopeFreeSummaryY-YCurrent version bytesCurrentVersionStorageBytesNumber of bytes that are a current version of an objectFreeCost optimizationN-Y% current version bytes-Percentage of bytes in scope that are current versions of objectsFreeCost optimizationYsum(CurrentVersionStorageBytes)/sum(StorageBytes)YCurrent version object countCurrentVersionObjectCountNumber of current version objectsFreeData protectionN-Y% current version objects-Percentage of objects in scope that are a current versionFreeCost optimizationYsum(CurrentVersionObjectCount)/sum(ObjectCount)YNoncurrent version bytesNonCurrentVersionStorageBytesNumber of noncurrent version bytesFreeCost optimizationN-Y% noncurrent version bytes-Percentage of bytes in scope that are noncurrent versionsFreeCost optimizationYsum(NonCurrentVersionStorageBytes)/sum(StorageBytes)YNoncurrent version object countNonCurrentVersionObjectCountNumber of the noncurrent object versionsFreeCost optimizationN-Y% noncurrent version objects-Percentage of objects in scope that are a noncurrent versionFreeCost optimizationYsum(NonCurrentVersionObjectCount)/sum(ObjectCount)YDelete marker bytesDeleteMarkerStorageBytesNumber of bytes in scope that are delete markersFreeCost optimizationN-Y% delete marker bytes-Percentage of bytes in scope that are delete markersFreeCost optimizationYsum(DeleteMarkerStorageBytes)/sum(StorageBytes)YDelete marker object countDeleteMarkerObjectCountNumber of objects with a delete markerFreeCost optimizationN-Y% delete marker objects-Percentage of objects in scope with a delete markerFreeCost optimizationYsum(DeleteMarkerObjectCount)/sum(ObjectCount)YIncomplete multipart upload bytesIncompleteMultipartUploadStorageBytesTotal bytes in scope for incomplete multipart uploadsFreeCost optimizationN-Y% incomplete multipart upload bytes-Percentage of bytes in scope that are the result of incomplete multipart
uploadsFreeCost optimizationYsum(IncompleteMultipartUploadStorageBytes)/sum(StorageBytes)YIncomplete multipart upload object countIncompleteMultipartUploadObjectCountNumber of objects in scope that are incomplete multipart uploadsFreeCost optimizationN-Y% incomplete multipart upload objects-Percentage of objects in scope that are incomplete multipart uploadsFreeCost optimizationYsum(IncompleteMultipartUploadObjectCount)/sum(ObjectCount)YIncomplete multipart upload storage bytes greater than 7 days oldIncompleteMPUStorageBytesOlderThan7DaysTotal bytes in scope for incomplete multipart uploads that are more than 7 days
oldFreeCost optimizationN-Y% incomplete multipart upload storage bytes greater than 7 days old-Percentage of bytes for incomplete multipart uploads that are more than 7 days
oldFreeCost optimizationYsum(IncompleteMPUStorageBytesOlderThan7Days)/sum(StorageBytes)YIncomplete multipart upload object count greater than 7 days oldIncompleteMPUObjectCountOlderThan7DaysNumber of objects that are incomplete multipart uploads more than 7 days
oldFreeCost optimizationN-Y% incomplete multipart upload object count greater than 7 days old-Percentage of objects that are incomplete multipart uploads more than 7 days
oldFreeCost optimizationYsum(IncompleteMPUObjectCountOlderThan7Days)/sum(ObjectCount)YTransition lifecycle rule countTransitionLifecycleRuleCountNumber of lifecycle rules to transition objects to another storage classAdvancedCost optimizationN-NAverage transition lifecycle rules per bucket-Average number of lifecycle rules to transition objects to another storage
classAdvancedCost optimizationYsum(TransitionLifecycleRuleCount)/sum(DistinctNumberOfBuckets)NExpiration lifecycle rule countExpirationLifecycleRuleCountNumber of lifecycle rules to expire objectsAdvancedCost optimizationN-NAverage expiration lifecycle rules per bucket-Average number of lifecycle rules to expire objectsAdvancedCost optimizationYsum(ExpirationLifecycleRuleCount)/sum(DistinctNumberOfBuckets)NNoncurrent version transition lifecycle rule countNoncurrentVersionTransitionLifecycleRuleCountNumber of lifecycle rules to transition noncurrent object versions to another
storage classAdvancedCost optimizationNNAverage noncurrent version transition lifecycle rules per bucket-Average number of lifecycle rules to transition noncurrent object versions to
another storage classAdvancedCost optimizationYsum(NoncurrentVersionTransitionLifecycleRuleCount)/sum(DistinctNumberOfBuckets)
NNoncurrent version expiration lifecycle rule countNoncurrentVersionExpirationLifecycleRuleCountNumber of lifecycle rules to expire noncurrent object versionsAdvancedCost optimizationN-NAverage noncurrent version expiration lifecycle rules per bucket-Average number of lifecycle rules to expire noncurrent object versionsAdvancedCost optimizationYsum(NoncurrentVersionExpirationLifecycleRuleCount)/sum(DistinctNumberOfBuckets)
NAbort incomplete multipart upload lifecycle rule countAbortIncompleteMPULifecycleRuleCountNumber of lifecycle rules to delete incomplete multipart uploadsAdvancedCost optimizationN-NAverage abort incomplete multipart upload lifecycle rules per bucket-Average number of lifecycle rules to delete incomplete multipart uploadsAdvancedCost optimizationYsum(AbortIncompleteMPULifecycleRuleCount)/sum(DistinctNumberOfBuckets)NExpired object delete marker lifecycle rule countExpiredObjectDeleteMarkerLifecycleRuleCountNumber of lifecycle rules to remove expired object delete markersAdvancedCost optimizationN-NAverage expired object delete marker lifecycle rules per bucket-Average number of lifecycle rules to remove expired object delete markersAdvancedCost optimizationYsum(ExpiredObjectDeleteMarkerLifecycleRuleCount)/sum(DistinctNumberOfBuckets)
NTotal lifecycle rule countTotalLifecycleRuleCountNumber of lifecycle rulesAdvancedCost optimizationN-NAverage lifecycle rule count per bucket-Average number of lifecycle rulesAdvancedCost optimizationYsum(TotalLifecycleRuleCount)/sum(DistinctNumberOfBuckets)NEncrypted bytesEncryptedStorageBytesNumber of encrypted bytesFreeData protectionN-Y% encrypted bytes-Percentage of total bytes that are encryptedFreeData protectionYsum(EncryptedObjectCount)/sum(StorageBytes)YEncrypted object countEncryptedObjectCountNumber of objects that are encryptedFreeData protectionN-Y% encrypted objects-Percentage of objects that are encryptedFreeData protectionYsum(EncryptedStorageBytes)/sum(ObjectCount)YUnencrypted bytesUnencryptedStorageBytesNumber of bytes that are unencryptedFreeData protectionYsum(StorageBytes) - sum(EncryptedStorageBytes)Y% unencrypted bytes-Percentage of bytes that are unencryptedFreeData protectionYsum(UnencryptedStorageBytes)/sum(StorageBytes)YUnencrypted object countUnencryptedObjectCountNumber of objects that are unencryptedFreeData protectionYsum(ObjectCount) - sum(EncryptedObjectCount)Y% unencrypted objects-Percentage of unencrypted objectsFreeData protectionYsum(UnencryptedObjectCount)/sum(ObjectCount)YReplicated storage bytes sourceReplicatedStorageBytesSourceNumber of bytes that are replicated from the source bucketFreeData protectionN-Y% replicated bytes source-Percentage of total bytes that are replicated from the source bucketFreeData protectionYsum(ReplicatedStorageBytesSource)/sum(StorageBytes)YReplicated object count sourceReplicatedObjectCountSourceNumber of replicated objects from the source bucketFreeData protectionN-Y% replicated objects source-Percentage of total objects that are replicated from the source bucketFreeData protectionYsum(ReplicatedStorageObjectCount)/sum(ObjectCount)YReplicated storage bytes destinationReplicatedStorageBytesNumber of bytes that are replicated to the destination bucketFreeData protectionN-N% replicated bytes destination-Percentage of total bytes that are replicated to the destination bucketFreeData protectionYsum(ReplicatedStorageBytes)/sum(StorageBytes)YReplicated object count destinationReplicatedObjectCountNumber of objects that are replicated to the destination bucketFreeData protectionN-Y% replicated objects destination-Percentage of total objects that are replicated to the destination bucketFreeData protectionYsum(ReplicatedObjectCount)/sum(ObjectCount)YObject Lock bytesObjectLockEnabledStorageBytesNumber of Object Lock enabled storage bytesFreeData protectionNsum(UnencryptedStorageBytes)/sum(ObjectLockEnabledStorageCount)-sum(ObjectLockEnabledStorageBytes)Y% Object Lock bytes-Percentage of Object Lock enabled storage bytesFreeData protectionYsum(ObjectLockEnabledStorageBytes)/sum(StorageBytes)YObject Lock object countObjectLockEnabledObjectCountNumber of Object Lock objectsFreeData protectionN-Y% Object Lock objects-Percentage of total objects that have Object Lock enabledFreeData protectionY sum(ObjectLockEnabledObjectCount)/sum(ObjectCount)YVersioning-enabled bucket countVersioningEnabledBucketCountNumber of buckets that have S3 Versioning enabledFreeData protectionN-N% versioning-enabled buckets-Percentage of buckets that have S3 Versioning enabledFreeData protectionYsum(VersioningEnabledBucketCount)/sum(DistinctNumberOfBuckets)NMFA delete-enabled bucket countMFADeleteEnabledBucketCountNumber of buckets that have MFA (multi-factor authentication) delete
enabledFreeData protectionN-N% MFA delete-enabled buckets-Percentage of buckets that have MFA (multi-factor authentication) delete
enabledFreeData protectionYsum(MFADeleteEnabledBucketCount)/sum(DistinctNumberOfBuckets)NSSE-KMS enabled bucket countSSEKMSEnabledBucketCountNumber of buckets that use server-side encryption with AWS Key Management Service keys (SSE-KMS) for
default bucket encryptionFreeData protectionN-N% SSE-KMS enabled buckets-Percentage of buckets that SSE-KMS for default bucket encryptionFreeData protectionYsum(SSEKMSEnabledBucketCount)/sum(DistinctNumberOfBuckets)NAll unsupported signature requestsAllUnsupportedSignatureRequestsTotal number of requests that use unsupported AWS signature versionsAdvancedData protectionN-N% all unsupported signature requests-Percentage of requests that use unsupported AWS signature versionsAdvancedData protectionYsum(AllUnsupportedSignatureRequests)/sum(AllRequests)NAll unsupported TLS requestsAllUnsupportedTLSRequestsTotal number of requests that use unsupported Transport Layer Security (TLS)
versionsAdvancedData protectionN-N% all unsupported TLS requests-Percentage of requests that use unsupported TLS versionsAdvancedData protectionYsum(AllUnsupportedTLSRequests)/sum(AllRequests)NAll SSE-KMS requestsAllSSEKMSRequestsTotal number of requests that specify SSE-KMSAdvancedData protectionN-N% all SSE-KMS requests-Percentage of requests that specify SSE-KMSAdvancedData protectionYsum(AllSSEKMSRequests)/sum(AllRequests)NSame-Region Replication rule countSameRegionReplicationRuleCountNumber of replication rules for Same-Region Replication (SRR)AdvancedData protectionN-NAverage Same-Region Replication rules per bucket-Average number of replication rules for SRRAdvancedData protectionYsum(SameRegionReplicationRuleCount)/sum(DistinctNumberOfBuckets)NCross-Region Replication rule countCrossRegionReplicationRuleCountNumber of replication rules for Cross-Region Replication (CRR)AdvancedData protectionN-NAverage Cross-Region Replication rules per bucket-Average number of replication rules for CRRAdvancedData protectionYsum(CrossRegionReplicationRuleCount)/sum(DistinctNumberOfBuckets)NSame-account replication rule countSameAccountReplicationRuleCountNumber of replication rules for replication within the same accountAdvancedData protectionN-NAverage same-account replication rules per bucket-Average number of replication rules for replication within the same accountAdvancedData protectionYsum(SameAccountReplicationRuleCount)/sum(DistinctNumberOfBuckets)NCross-account replication rule countCrossAccountReplicationRuleCountNumber of replication rules for cross-account replicationAdvancedData protectionN-NAverage cross-account replication rules per bucket-Average number of replication rules for cross-account replicationAdvancedData protectionYsum(CrossAccountReplicationRuleCount)/sum(DistinctNumberOfBuckets)NInvalid destination replication rule countInvalidDestinationReplicationRuleCountNumber of replication rules with a replication destination that's not validAdvancedData protectionN-NAverage invalid destination replication rules per bucket-Average number of replication rules with a replication destination that's not
validAdvancedData protectionYsum(InvalidReplicationRuleCount)/sum(DistinctNumberOfBuckets)NTotal replication rule count-Total replication rule countAdvancedData protectionY-NAverage replication rule count per bucket-Average total replication rule countAdvancedData protectionYsum(all replication rule count metrics)/sum(DistinctNumberOfBuckets)NObject Ownership bucket owner enforced bucket countObjectOwnershipBucketOwnerEnforcedBucketCountNumber of buckets that have access control lists (ACLs) disabled by using the
bucket owner enforced setting for Object OwnershipFreeAccess managementN-N% Object Ownership bucket owner enforced buckets-Percentage of buckets that have ACLs disabled by using the bucket owner enforced
setting for Object OwnershipFreeAccess managementYsum(ObjectOwnershipBucketOwnerEnforcedBucketCount)/sum(DistinctNumberOfBuckets)
NObject Ownership bucket owner preferred bucket countObjectOwnershipBucketOwnerPreferredBucketCountNumber of buckets that use the bucket owner preferred setting for
Object OwnershipFreeAccess managementN-N% Object Ownership bucket owner preferred buckets-Percentage of buckets that use the bucket owner preferred setting for
Object OwnershipFreeAccess managementYsum(ObjectOwnershipBucketOwnerPreferredBucketCount)/sum(DistinctNumberOfBuckets)
NObject Ownership object writer bucket countObjectOwnershipObjectWriterBucketCountNumber of buckets that use the object writer setting for
Object OwnershipFreeAccess managementN-N% Object Ownership object writer buckets-Percentage of buckets that use the object writer setting for
Object OwnershipFreeAccess managementYsum(ObjectOwnershipObjectWriterBucketCount)/sum(DistinctNumberOfBuckets)NTransfer Acceleration enabled bucket countTransferAccelerationEnabledBucketCountNumber of buckets that have Transfer Acceleration enabledFreePerformanceN-N% Transfer Acceleration enabled buckets-Percentage of buckets that have Transfer Acceleration enabledFreePerformanceYsum(TransferAccelerationEnabledBucketCount)/sum(DistinctNumberOfBuckets)NEvent Notification enabled bucket countEventNotificationEnabledBucketCountNumber of buckets that have Event Notifications enabledFreeEventsNN% Event Notification enabled buckets-Percentage of buckets that have Event Notifications enabledFreeEventsYsum(EventNotificationEnabledBucketCount)/sum(DistinctNumberOfBuckets)NAll requestsAllRequests

Total number of requests made

AdvancedActivityN-NGet requestsGetRequests

Total number of `GET` requests made

AdvancedActivityN-NPut requestsPutRequests

Total number of `PUT` requests made

AdvancedActivityN-NHead requestsHeadRequestsNumber of `HEAD` requests madeAdvancedActivityN-NDelete requestsDeleteRequestsNumber of `DELETE` requests madeAdvancedActivityN-NList requestsListRequestsNumber of `LIST` requests madeAdvancedActivityN-NPost requestsPostRequestsNumber of `POST` requests madeAdvancedActivityN-NSelect requestsSelectRequestsNumber of S3 Select requestsAdvancedActivityN-NSelect scanned bytesSelectScannedBytesNumber of S3 Select bytes scannedAdvancedActivityN-NSelect returned bytesSelectReturnedBytesNumber of S3 Select bytes returnedAdvancedActivityN-NBytes downloadedBytesDownloadedNumber of bytes downloadedAdvancedActivityN-N% retrieval rate-Percentage of bytes downloadedAdvancedActivityYsum(BytesDownloaded)/sum(StorageBytes)NBytes uploadedBytesUploadedNumber of bytes uploadedAdvancedActivityN-N% ingest ratio-Percentage of bytes uploadedAdvancedActivityYsum(BytesUploaded)/sum(StorageBytes)N4xx errors4xxErrorsNumber of HTTP 4xx status codesAdvancedActivityN-N5xx errors5xxErrorsNumber of HTTP 5xx status codesAdvancedActivityN-NTotal errors-The sum of all 4xx and 5xx errorsAdvancedActivityYsum(4xxErrors) + sum(5xxErrors)N% error rate-

Total number of 4xx and 5xx errors as a percentage of total requests

AdvancedActivityYsum(TotalErrors)/sum(TotalRequests)N200 OK status count200OKStatusCountNumber of 200 OK status codesAdvancedDetailed status codeN-N% 200 OK status-

Total number of 200 OK status codes as a percentage of total requests

AdvancedDetailed status codeYsum(200OKStatusCount)/sum(AllRequests)N206 Partial Content status count206PartialContentStatusCountNumber of 206 Partial Content status codesAdvancedDetailed status codeN-N% 206 Partial Content status-Number of 206 Partial Content status codes as a percentage of total
requestsAdvancedDetailed status codeYsum(206PartialContentStatusCount)/sum(AllRequests)N400 Bad Request error count

400BadRequestErrorCount

Number of 400 Bad Request status codesAdvancedDetailed status codeN-N% 400 Bad Request errors-Number of 400 Bad Request status codes as a percentage of total requestsAdvancedDetailed status codeYsum(400BadRequestErrorCount)/sum(AllRequests)N403 Forbidden error count

403ForbiddenErrorCount

Number of 403 Forbidden status codesAdvancedDetailed status codeN-N% 403 Forbidden errors-Number of 403 Forbidden status codes as a percentage of total requestsAdvancedDetailed status codeYsum(403ForbiddenErrorCount)/sum(AllRequests)N404 Not Found error count404NotFoundErrorCountNumber of 404 Not Found status codesAdvancedDetailed status codeN-N% 404 Not Found errors-Number of 404 Not Found status codes as a percentage of total requestsAdvancedDetailed status codeYsum(404NotFoundErrorCount)/sum(AllRequests)N500 Internal Server Error count500InternalServerErrorCountNumber of 500 Internal Server Error status codesAdvancedDetailed status codeN-N% 500 Internal Server Errors-Number of 500 Internal Server Error status codes as a percentage of total
requestsAdvancedDetailed status codeYsum(500InternalServerErrorCount)/sum(AllRequests)N503 Service Unavailable error count503ServiceUnavailableErrorCountNumber of 503 Service Unavailable status codesAdvancedDetailed status codeN-N% 503 Service Unavailable errors-Number of 503 Service Unavailable status codes as a percentage of total
requestsAdvancedDetailed status codeYsum(503ServiceUnavailableErrorCount)/sum(AllRequests)N

1 All free tier storage metrics are available at the Storage Lens
group level. Advanced tier metrics are not available at the Storage Lens group level.

2 Rule count metrics and bucket settings metrics aren't available
at the prefix level.

The following table shows the performance metrics available in S3 Storage Lens and their
availability in CloudWatch:

**Metric name****CloudWatch and export****Description****Tier****Category****Derived****Derived metric formula****Storage Lens groups**Average First Byte LatencyAverageFirstByteLatencyAverage per-request time between when an Amazon S3 bucket receives a complete request and when it starts returning the response,
measured over the past 24 hoursAdvancedPerformanceN-NAverage Total Request LatencyAverageTotalRequestLatencyAverage elapsed per-request time between the first byte received and the last byte sent to an Amazon S3 bucket,
measured over the past 24 hoursAdvancedPerformanceN-NRead 0KB request countRead0KBRequestCount\*Number of GetObject requests with data sizes of 0KB, including both range-based
requests and whole object requestsAdvancedPerformanceN-NRead 0KB to 128KB request countRead0KBTo128KBRequestCount\*Number of GetObject requests with data sizes greater than 0KB and up to 128KB,
including both range-based requests and whole object requestsAdvancedPerformanceN-NRead 128KB to 256KB request countRead128KBTo256KBRequestCount\*Number of GetObject requests with data sizes greater than 128KB and up to 256KB,
including both range-based requests and whole object requestsAdvancedPerformanceN-NRead 256KB to 512KB request countRead256KBTo512KBRequestCount\*Number of GetObject requests with data sizes greater than 256KB and up to 512KB,
including both range-based requests and whole object requestsAdvancedPerformanceN-NRead 512KB to 1MB request countRead512KBTo1MBRequestCount\*Number of GetObject requests with data sizes greater than 512KB and up to 1MB,
including both range-based requests and whole object requestsAdvancedPerformanceN-NRead 1MB to 2MB request countRead1MBTo2MBRequestCount\*Number of GetObject requests with data sizes greater than 1MB and up to 2MB,
including both range-based requests and whole object requestsAdvancedPerformanceN-NRead 2MB to 4MB request countRead2MBTo4MBRequestCount\*Number of GetObject requests with data sizes greater than 2MB and up to 4MB,
including both range-based requests and whole object requestsAdvancedPerformanceN-NRead 4MB to 8MB request countRead4MBTo8MBRequestCount\*Number of GetObject requests with data sizes greater than 4MB and up to 8MB,
including both range-based requests and whole object requestsAdvancedPerformanceN-NRead 8MB to 16MB request countRead8MBTo16MBRequestCount\*Number of GetObject requests with data sizes greater than 8MB and up to 16MB,
including both range-based requests and whole object requestsAdvancedPerformanceN-NRead 16MB to 32MB request countRead16MBTo32MBRequestCount\*Number of GetObject requests with data sizes greater than 16MB and up to 32MB,
including both range-based requests and whole object requestsAdvancedPerformanceN-NRead 32MB to 64MB request countRead32MBTo64MBRequestCount\*Number of GetObject requests with data sizes greater than 32MB and up to 64MB,
including both range-based requests and whole object requestsAdvancedPerformanceN-NRead 64MB to 128MB request countRead64MBTo128MBRequestCount\*Number of GetObject requests with data sizes greater than 64MB and up to 128MB,
including both range-based requests and whole object requestsAdvancedPerformanceN-NRead 128MB to 256MB request countRead128MBTo256MBRequestCount\*Number of GetObject requests with data sizes greater than 128MB and up to 256MB,
including both range-based requests and whole object requestsAdvancedPerformanceN-NRead 256MB to 512MB request countRead256MBTo512MBRequestCount\*Number of GetObject requests with data sizes greater than 256MB and up to 512MB,
including both range-based requests and whole object requestsAdvancedPerformanceN-NRead 512MB to 1GB request countRead512MBTo1GBRequestCount\*Number of GetObject requests with data sizes greater than 512MB and up to 1GB,
including both range-based requests and whole object requestsAdvancedPerformanceN-NRead 1GB to 2GB request countRead1GBTo2GBRequestCount\*Number of GetObject requests with data sizes greater than 1GB and up to 2GB,
including both range-based requests and whole object requestsAdvancedPerformanceN-NRead 2GB to 4GB request countRead2GBTo4GBRequestCount\*Number of GetObject requests with data sizes greater than 2GB and up to 4GB,
including both range-based requests and whole object requestsAdvancedPerformanceN-NRead 4GB+ request countReadLargerThan4GBRequestCount\*Number of GetObject requests with data sizes greater than 4GB, including both
range-based requests and whole object requestsAdvancedPerformanceN-NWrite 0KB request countWrite0KBRequestCount\*Number of PutObject, UploadPart, and CreateMultipartUpload requests with data sizes
of 0KBAdvancedPerformanceN-NWrite 0KB to 128KB request countWrite0KBTo128KBRequestCount\*Number of PutObject, UploadPart, and CreateMultipartUpload requests with data sizes
greater than 0KB and up to 128KBAdvancedPerformanceN-NWrite 128KB to 256KB request countWrite128KBTo256KBRequestCount\*Number of PutObject, UploadPart, and CreateMultipartUpload requests with data sizes
greater than 128KB and up to 256KBAdvancedPerformanceN-NWrite 256KB to 512KB request countWrite256KBTo512KBRequestCount\*Number of PutObject, UploadPart, and CreateMultipartUpload requests with data sizes
greater than 256KB and up to 512KBAdvancedPerformanceN-NWrite 512KB to 1MB request countWrite512KBTo1MBRequestCount\*Number of PutObject, UploadPart, and CreateMultipartUpload requests with data sizes
greater than 512KB and up to 1MBAdvancedPerformanceN-NWrite 1MB to 2MB request countWrite1MBTo2MBRequestCount\*Number of PutObject, UploadPart, and CreateMultipartUpload requests with data sizes
greater than 1MB and up to 2MBAdvancedPerformanceN-NWrite 2MB to 4MB request countWrite2MBTo4MBRequestCount\*Number of PutObject, UploadPart, and CreateMultipartUpload requests with data sizes
greater than 2MB and up to 4MBAdvancedPerformanceN-NWrite 4MB to 8MB request countWrite4MBTo8MBRequestCount\*Number of PutObject, UploadPart, and CreateMultipartUpload requests with data sizes
greater than 4MB and up to 8MBAdvancedPerformanceN-NWrite 8MB to 16MB request countWrite8MBTo16MBRequestCount\*Number of PutObject, UploadPart, and CreateMultipartUpload requests with data sizes
greater than 8MB and up to 16MBAdvancedPerformanceN-NWrite 16MB to 32MB request countWrite16MBTo32MBRequestCount\*Number of PutObject, UploadPart, and CreateMultipartUpload requests with data sizes
greater than 16MB and up to 32MBAdvancedPerformanceN-NWrite 32MB to 64MB request countWrite32MBTo64MBRequestCount\*Number of PutObject, UploadPart, and CreateMultipartUpload requests with data sizes
greater than 32MB and up to 64MBAdvancedPerformanceN-NWrite 64MB to 128MB request countWrite64MBTo128MBRequestCount\*Number of PutObject, UploadPart, and CreateMultipartUpload requests with data sizes
greater than 64MB and up to 128MBAdvancedPerformanceN-NWrite 128MB to 256MB request countWrite128MBTo256MBRequestCount\*Number of PutObject, UploadPart, and CreateMultipartUpload requests with data sizes
greater than 128MB and up to 256MBAdvancedPerformanceN-NWrite 256MB to 512MB request countWrite256MBTo512MBRequestCount\*Number of PutObject, UploadPart, and CreateMultipartUpload requests with data sizes
greater than 256MB and up to 512MBAdvancedPerformanceN-NWrite 512MB to 1GB request countWrite512MBTo1GBRequestCount\*Number of PutObject, UploadPart, and CreateMultipartUpload requests with data sizes
greater than 512MB and up to 1GBAdvancedPerformanceN-NWrite 1GB to 2GB request countWrite1GBTo2GBRequestCount\*Number of PutObject, UploadPart, and CreateMultipartUpload requests with data sizes
greater than 1GB and up to 2GBAdvancedPerformanceN-NWrite 2GB to 4GB request countWrite2GBTo4GBRequestCount\*Number of PutObject, UploadPart, and CreateMultipartUpload requests with data sizes
greater than 2GB and up to 4GBAdvancedPerformanceN-NWrite 4GB+ request countWriteLargerThan4GBRequestCount\*Number of PutObject, UploadPart, and CreateMultipartUpload requests with data sizes
greater than 4GBAdvancedPerformanceN-NObject 0KB countObject0KBCountNumber of objects with sizes equal to 0KB, including current version, noncurrent
versions, incomplete multipart uploads, and delete markersAdvancedPerformanceN-NObject 0KB to 128KB countObject0KBTo128KBCountNumber of objects with sizes greater than 0KB and less than equal to 128KB,
including current version, noncurrent versions, incomplete multipart uploads, and delete
markersAdvancedPerformanceN-NObject 128KB to 256KB countObject128KBTo256KBCountNumber of objects with sizes greater than 128KB and less than equal to 256KB,
including current version, noncurrent versions, incomplete multipart uploads, and delete
markersAdvancedPerformanceN-NObject 256KB to 512KB countObject256KBTo512KBCountNumber of objects with sizes greater than 256KB and less than equal to 512KB,
including current version, noncurrent versions, incomplete multipart uploads, and delete
markersAdvancedPerformanceN-NObject 512KB to 1MB countObject512KBTo1MBCountNumber of objects with sizes greater than 512KB and less than equal to 1MB,
including current version, noncurrent versions, incomplete multipart uploads, and delete
markersAdvancedPerformanceN-NObject 1MB to 2MB countObject1MBTo2MBCountNumber of objects with sizes greater than 1MB and less than equal to 2MB, including
current version, noncurrent versions, incomplete multipart uploads, and delete
markersAdvancedPerformanceN-NObject 2MB to 4MB countObject2MBTo4MBCountNumber of objects with sizes greater than 2MB and less than equal to 4MB, including
current version, noncurrent versions, incomplete multipart uploads, and delete
markersAdvancedPerformanceN-NObject 4MB to 8MB countObject4MBTo8MBCountNumber of objects with sizes greater than 4MB and less than equal to 8MB, including
current version, noncurrent versions, incomplete multipart uploads, and delete
markersAdvancedPerformanceN-NObject 8MB to 16MB countObject8MBTo16MBCountNumber of objects with sizes greater than 8MB and less than equal to 16MB,
including current version, noncurrent versions, incomplete multipart uploads, and delete
markersAdvancedPerformanceN-NObject 16MB to 32MB countObject16MBTo32MBCountNumber of objects with sizes greater than 16MB and less than equal to 32MB,
including current version, noncurrent versions, incomplete multipart uploads, and delete
markersAdvancedPerformanceN-NObject 32MB to 64MB countObject32MBTo64MBCountNumber of objects with sizes greater than 32MB and less than equal to 64MB,
including current version, noncurrent versions, incomplete multipart uploads, and delete
markersAdvancedPerformanceN-NObject 64MB to 128MB countObject64MBTo128MBCountNumber of objects with sizes greater than 64MB and less than equal to 128MB,
including current version, noncurrent versions, incomplete multipart uploads, and delete
markersAdvancedPerformanceN-NObject 128MB to 256MB countObject128MBTo256MBCountNumber of objects with sizes greater than 128MB and less than equal to 256MB,
including current version, noncurrent versions, incomplete multipart uploads, and delete
markersAdvancedPerformanceN-NObject 256MB to 512MB countObject256MBTo512MBCountNumber of objects with sizes greater than 256MB and less than equal to 512MB,
including current version, noncurrent versions, incomplete multipart uploads, and delete
markersAdvancedPerformanceN-NObject 512MB to 1GB countObject512MBTo1GBCountNumber of objects with sizes greater than 512MB and less than equal to 1GB,
including current version, noncurrent versions, incomplete multipart uploads, and delete
markersAdvancedPerformanceN-NObject 1GB to 2GB countObject1GBTo2GBCountNumber of objects with sizes greater than 1GB and less than equal to 2GB, including
current version, noncurrent versions, incomplete multipart uploads, and delete
markersAdvancedPerformanceN-NObject 2GB to 4GB countObject2GBTo4GBCountNumber of objects with sizes greater than 2GB and less than equal to 4GB, including
current version, noncurrent versions, incomplete multipart uploads, and delete
markersAdvancedPerformanceN-NObject 4GB+ countObjectLargerThan4GBCountNumber of objects with sizes greater than 4GB, including current version,
noncurrent versions, incomplete multipart uploads, and delete markersAdvancedPerformanceN-NConcurrent Put 503 error countConcurrentPut503ErrorCountNumber of 503 errors that are generated due to concurrent writes to the same
objectAdvancedPerformanceN-N% Concurrent Put 503 errors-Percentage of 503 errors that are generated due to concurrent writes to the same
objectAdvancedPerformanceY100 \* ConcurrentPut503Errors / AllRequestsNCross-Region request countCrossRegionRequestCountNumber of requests that originate from a client in different Region than bucket's
home RegionAdvancedPerformanceN-N% Cross-Region requests-Percentage of requests that originate from a client in different Region than
bucket's home RegionAdvancedPerformanceY100 \* CrossRegionRequestCount / AllRequestsNCross-Region transferred bytesCrossRegionTransferredBytesNumber of bytes that are transferred from calls in different Region than bucket's
home RegionAdvancedPerformanceN-N% Cross-Region transferred bytes-Percentage of bytes transferred that originate from calls in different Region that
bucket's home RegionAdvancedPerformanceY100 \* CrossRegionBytes / (BytesDownloaded + BytesUploaded)NCross-Region without replication request countCrossRegionWithoutReplicationRequestCountNumber of requests that originate from a client in different Region than bucket's
home Region, excluding cross-region replication requestsAdvancedPerformanceN-N% Cross-Region without replication requests-Percentage of requests that originate from a client in different Region that
bucket's home Region, excluding cross-region replication requestsAdvancedPerformanceY100 \* CrossRegionRequestWithoutReplicationCount / AllRequestsNCross-Region without replication transferred bytesCrossRegionWithoutReplicationTransferredBytesNumber of bytes that are transferred from calls in different Region than bucket's
home Region, excluding cross-region replication bytesAdvancedPerformanceN-N% Cross-Region without replication transferred bytes-Number of requests that originate from a Region other than the bucket's home
Region, excluding cross-region replication requestsAdvancedPerformanceY100 \* CrossRegionBytesWithoutReplication / (BytesDownloaded +
BytesUploaded)NIn-Region request countInRegionRequestCountNumber of requests that originate from a client in same Region as bucket's home
RegionAdvancedPerformanceN-N% In-Region requests-Percentage of requests that originate from a client in same Region as bucket's home
RegionAdvancedPerformanceY100 \* InRegionRequestCount / AllRequestsNIn-Region transferred bytesInRegionTransferredBytesNumber of bytes that are transferred from calls from same Region as bucket's home
RegionAdvancedPerformanceN-N% In-Region transferred bytes-Percentage of bytes transferred that originate from calls from same Region as
bucket's home RegionAdvancedPerformanceY100 \* InRegionBytes / (BytesDownloaded + BytesUploaded)NUnique objects accessed count dailyUniqueObjectsAccessedDailyCountNumber of objects that were accessed at least once in last 24 hrsAdvancedPerformanceN-N% Unique objects accessed count daily-Percentage of objects that were accessed at least once in last 24 hrsAdvancedPerformanceY100 \* UniqueObjectsAccessedDailyCount / ObjectCountN

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Understanding S3 Storage Lens

Setting permissions

All content copied from https://docs.aws.amazon.com/.
