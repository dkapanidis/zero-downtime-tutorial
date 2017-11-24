# Status Page

[Status Page](https://github.com/spiddy/statuspage) is a container that help visualize the availability of a service in graphical web interface.

> NOTE: This is not for production. It is stress-testing the application with hundred requests per second.

## Run Status Page

Run a Status Page pointing to [http://google.com](http://google.com) and exposed to port `:9000`:

```shell
docker run --rm --name statuspage -d -p 9000:9000 -e TARGET_URL=http://google.com spiddy/statuspage
```

> Status Page is prebuilt on [Docker Hub](https://hub.docker.com/r/spiddy/statuspage/)

Open [http://localhost:9000](http://localhost:9000) on a browser, the status page will show the percentage of correct reponses from google.com in a graph.

Clean up the running instance

```shell
docker stop statuspage
```
