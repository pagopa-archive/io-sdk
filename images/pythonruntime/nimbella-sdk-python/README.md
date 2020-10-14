# Nimbella SDK for Python

A Python package to interact with [`nimbella.com`](https://nimbella.com) services.

## Installation

```
pip install nimbella
```

## Usage

```py
import nimbella

# Redis
redis = nimbella.redis()
redis.set("key", "value")
value = redis.get("key")

# Storage
bucket = nimbella.storage()
filename = "test.txt"
blob = bucket.blob(filename)
blob.upload_from_string('Expected %s contents' % filename)
```

## Support

We're always happy to help you with any issues you encounter. You may want to [join our Slack community](https://nimbella-community.slack.com/) to engage with us for a more rapid response.

## License

Apache-2.0. See [LICENSE](LICENSE) to learn more.
