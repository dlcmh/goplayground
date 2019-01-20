Gorilla is less of a framework and more of a set of very useful tools that are generally bundled in frameworks.

Specifically, Gorilla contains:

- __gorilla/context__: This is a package for creating a globally-accessible variable from the request. It's useful for sharing a value from the URL without repeating the code to access it across your application.

- __gorilla/rpc__: This implements RPC-JSON, which is a system for remote code services and communication without implementing specific protocols. This relies on the JSON format to define the intentions of any request.

- __gorilla/schema__: This is a package that allows simple packing of form variables into a struct, which is an otherwise cumbersome process.

- __gorilla/securecookie__: This, unsurprisingly, implements authenticated and encrypted cookies for your application.

- __gorilla/sessions__: Similar to cookies, this provides unique, long-term, and repeatable data stores by utilizing a file-based and/or cookie-based session system.

- __gorilla/mux__: This is intended to create flexible routes that allow regular expressions to dictate available variables for routers.

- The last package is the one we're most interested in here, and it comes with a related package called __gorilla/reverse__, which essentially allows you to reverse the process of creating regular expression-based muxes. We will cover that topic in detail in the later section.
