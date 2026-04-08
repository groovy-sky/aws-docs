# Amazon Connect Agent Workspace Activity API

The Amazon Connect SDK provides a `SessionExpirationWarningClient` which serves as an
interface that your app in the Amazon Connect agent workspace can use to subscribe to events related to session
expiration due to inactivity and to signal the Amazon Connect that the agent is active.

The `SessionExpirationWarningClient` accepts an optional constructor argument, `
        ConnectClientConfig` which itself is defined as:

```

export type ConnectClientConfig = {
    context?: ModuleContext;
    provider?: AmazonConnectProvider;
};
```

If you do not provide a value for this config, then the client will default to using the **AmazonConnectProvider** set in the global provider scope. You can
also manually configure this using **setGlobalProvider**.

You can instantiate the client as follows:

```

import { SessionExpirationWarningClient } from "@amazon-connect/activity";

const sessionExpirationWarningClient = new SessionExpirationWarningClient();
```

###### Note

You must first instantiate the [AmazonConnectApp](../../../../reference/agentworkspace/latest/devguide/getting-started-initialize-sdk.md) which initializes the default AmazonConnectProvider and returns `
            { provider } `. This is the recommended option.

Alternatively, you can provide a constructor argument:

```

import { SessionExpirationWarningClient } from "@amazon-connect/activity";

const sessionExpirationWarningClient = new SessionExpirationWarningClient({
    context: sampleContext,
    provider: sampleProvider
});
```

The following sections describe the API calls for working with the
SessionExpirationWarning API.

###### Contents

- [onExpirationWarning()](3p-apps-activity-offexpirationwarning.md)

- [offExpirationWarningCleared()](3p-apps-activity-offexpirationwarningcleared.md)

- [offSessionExtensionError()](3p-apps-activity-offsessionextensionerror.md)

- [onExpirationWarning()](3p-apps-activity-onexpirationwarning.md)

- [onExpirationWarningCleared()](3p-apps-activity-onexpirationwarningcleared.md)

- [onSessionExtensionError()](3p-apps-activity-onsessionextensionerror.md)

- [sendActivity()](3p-apps-activity-sendactivity.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Connect SDK API
reference

onExpirationWarning()

All content copied from https://docs.aws.amazon.com/.
