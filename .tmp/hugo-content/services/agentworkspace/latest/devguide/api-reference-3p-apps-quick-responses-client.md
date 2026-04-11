# Amazon Connect Agent Workspace Quick Responses API

The Amazon Connect SDK provides a `QuickResponsesClient` which serves as an interface that you
can use to make requests to search your Amazon Connect Quick Responses Knowledge Base.

The `QuickResponsesClient` accepts an optional constructor argument, `
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

import { QuickResponsesClient } from "@amazon-connect/quick-responses";

const quickResponsesClient = new QuickResponsesClient({ provider });
```

###### Note

You must first instantiate the [AmazonConnectApp](../../../../reference/agentworkspace/latest/devguide/getting-started-initialize-sdk.md) which initializes the default AmazonConnectProvider and returns `
            { provider } `. This is the recommended option.

Alternatively, providing a constructor argument:

```typescript

import { QuickResponsesClient } from "@amazon-connect/quick-responses";

const quickResponsesClient = new QuickResponsesClient({
    context: sampleContext,
    provider: sampleProvider
});
```

###### Note

Third-party applications must be configured with \* permission in order to utilize the QuickResponsesClient APIs.

The following sections describe API calls for working with the QuickResponses
API.

###### Contents

- [isEnabled()](3p-apps-quick-responses-requests-isenabled.md)

- [searchQuickResponses()](3p-apps-quick-responses-requests-searchquickresponses.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

searchMessageTemplates()

isEnabled()

All content copied from https://docs.aws.amazon.com/.
