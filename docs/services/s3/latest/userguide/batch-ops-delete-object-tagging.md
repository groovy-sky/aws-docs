# Delete all object tags

You can use Amazon S3 Batch Operations to perform large-scale batch operations on Amazon S3 objects.
The **Delete all object tags** operation removes all Amazon S3 object tag sets
currently associated with the objects that are listed in the manifest. S3 Batch Operations
doesn't support deleting tags from objects while keeping other tags in place.

If the objects in your manifest are in a versioned bucket, you can remove the tag sets
from a specific version of an object. To do so, you must specify a version ID for every
object in the manifest. If you don't include a version ID for an object, S3 Batch Operations
removes the tag set from the latest version of every object. For more information about
Batch Operations manifests, see [Specifying a manifest](batch-ops-create-job.md#specify-batchjob-manifest).

For more details about object tagging, see [Categorizing your objects using tags](object-tagging.md) in this guide, and [PutObjectTagging](../api/api-putobjecttagging.md), [GetObjectTagging](../api/api-getobjecttagging.md), and [DeleteObjectTagging](../api/api-deleteobjecttagging.md) in the
_Amazon Simple Storage Service API Reference_.

###### Warning

Running this job removes all object tag sets on every object listed in the manifest.

To use the console to create a **Delete all object tags** job, see [Creating an S3 Batch Operations job](batch-ops-create-job.md).

## Restrictions and limitations

When you're using Batch Operations to delete object tags, the following restrictions and
limitations apply:

- The AWS Identity and Access Management (IAM) role that you specify to run the job must have
permissions to perform the underlying Amazon S3 `DeleteObjectTagging`
operation. For more information, see [DeleteObjectTagging](../api/api-deleteobjecttagging.md) in the
_Amazon Simple Storage Service API Reference_.

- S3 Batch Operations uses the Amazon S3 [DeleteObjectTagging](../api/api-deleteobjecttagging.md) operation to remove the tag
sets from every object in the manifest. All restrictions and limitations that
apply to the underlying operation also apply to S3 Batch Operations jobs.

- A single delete object tagging job can support a manifest with up to 20 billion objects.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Compute checksums

Invoke AWS Lambda
function
