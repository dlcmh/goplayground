So we might we want to redirect requests to another request? Well there are quite a few reasons, as defined by the HTTP specification that could lead us to implement automatic redirects on any given request. Here are a few of them with their corresponding HTTP status codes:

- A non-canonical address may need to be redirected to the canonical one for SEO purposes or for changes in site architecture. This is handled by __301 Moved Permanently__ or __302 Found__.
- Redirecting after a successful or unsuccessful POST. This helps us to prevent re-POSTing of the same form data accidentally. Typically, this is defined by __307 Temporary Redirect__.
- The page is not necessarily missing, but it now lives in another location. This is handled by the status code __301 Moved Permanently__.
