# Deploy from X

Deploy simple Text content from X.... where X is

* Twitter

More to come?

## Usage

### queue

Receives and redistributes events

* `PORT` (_default:_ `3185`) - Port to listen on

### twitter-streamer

Listens for Tweets as configured and sends them to the queue

* `QUEUE_ADDR` (_default:_ `queue:3185`) - Address of the queue service
* `CONSUMER_KEY`, `CONSUMER_SECRET`, `ACCESS_TOKEN`, `ACCESS_TOKEN_SECRET` - obtain from https://developer.twitter.com/en/apps
* `FILTER` - Only tweets matchin the filter will be watched (example: `@appuioli #deployme`). See "track" on https://developer.twitter.com/en/docs/tweets/filter-realtime/guides/basic-stream-parameters


## Development

    make dev-twitter-streamer
    make dev-queue


## License

MIT (see `LICENSE`)

---
> [Manuel Hutter](https://hutter.io/) -
> GitHub [@mhutter](https://github.com/mhutter) -
> Twitter [@dratir](https://twitter.com/dratir)
