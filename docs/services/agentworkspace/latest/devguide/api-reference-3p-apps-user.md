# Amazon Connect Agent Workspace User API

The Amazon Connect SDK provides an `SettingsClient` which serves as an interface that your app
in Amazon Connect Agent Workspace can use to make data requests on user settings.

The `SettingsClient` accepts an optional constructor argument, `
        ConnectClientConfig` which itself is defined as:

```

        export type ConnectClientConfig = {
            context?: ModuleContext;
            provider?: AmazonConnectProvider;
         };

```

If you do not provide a value for this config, then the client will default to using the **AmazonConnectProvider** set in the global provider scope. You can also manually
configure this using **setGlobalProvider**.

You can instantiate the agent client as follows:

```

        import { SettingsClient } from "@amazon-connect/user";
        const settingsClient = new SettingsClient({ provider });

```

###### Note

You must first instantiate the [AmazonConnectApp](../../../../reference/agentworkspace/latest/devguide/getting-started-initialize-sdk.md) which initializes the default AmazonConnectProvider and returns `
            { provider } `. This is the recommended option.

Alternatively, providing a constructor argument:

```

        import { SettingsClient } from "@amazon-connect/user";

        const settingsClient = new SettingsClient({
            context: sampleContext,
            provider: sampleProvider
    });

```

The following sections describe API calls for working with the User API.

###### Contents

- [getLanguage()](3p-apps-user-requests-getlanguage.md)

- [onLanguageChanged()](3p-apps-user-events-languagechanged-sub.md)

- [offLanguageChanged()](3p-apps-user-events-languagechanged-unsub.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

searchQuickResponses()

getLanguage()

All content copied from https://docs.aws.amazon.com/.
