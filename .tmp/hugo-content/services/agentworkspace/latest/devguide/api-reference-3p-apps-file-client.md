# Amazon Connect Agent Workspace File API

The Amazon Connect SDK provides a `FileClient` which serves as an interface that you can use
to make file requests to upload, retrieve, and delete attached files.

The `FileClient` accepts an optional constructor argument, `
        ConnectClientConfig` which itself is defined as:

```typescript

export type ConnectClientConfig = {
    context?: ModuleContext;
    provider?: AmazonConnectProvider;
};
```

If you do not provide a value for this config, then the client will default to using the **AmazonConnectProvider** set in the global provider scope. You can also manually
configure this using **setGlobalProvider**.

You can instantiate the agent client as follows:

```typescript

import { FileClient } from "@amazon-connect/file";

const fileClient = new FileClient({ provider });
```

###### Note

You must first instantiate the [AmazonConnectApp](../../../../reference/agentworkspace/latest/devguide/getting-started-initialize-sdk.md) which initializes the default AmazonConnectProvider and returns `
            { provider } `. This is the recommended option.

Alternatively, providing a constructor argument:

```typescript

import { FileClient } from "@amazon-connect/file";

const fileClient = new FileClient({
    context: sampleContext,
    provider: sampleProvider
});
```

###### Note

Third-party applications must be configured with \* permission in order to utilize the FileClient APIs.

The following sections describe API calls for working with the File API.

###### Contents

- [batchGetAttachedFileMetadata()](3p-apps-file-requests-batchgetattachedfilemetadata.md)

- [completeAttachedFileUpload()](3p-apps-file-requests-completeattachedfileupload.md)

- [deleteAttachedFile()](3p-apps-file-requests-deleteattachedfile.md)

- [getAttachedFileUrl()](3p-apps-file-requests-getattachedfileurl.md)

- [startAttachedFileUpload()](3p-apps-file-requests-startattachedfileupload.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

sendEmail()

batchGetAttachedFileMetadata()

All content copied from https://docs.aws.amazon.com/.
