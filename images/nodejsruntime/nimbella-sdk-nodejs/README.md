# Nimbella SDK for Node.js

A Node.js library to interact with [`nimbella.com`](https://nimbella.com) services.

## Installation

```
npm install @nimbella/sdk
```

## Usage

```js
const nim = require('@nimbella/sdk');

async function main(args) {
  // Redis
  const redis = nim.redis();
  await redis.setAsync('key', 'value');
  const value = await redis.getAsync('key');

  // Storage
  const bucket = await nim.storage();
  const file = bucket.file('hello.txt'); // Filename
  await file.save('Hello world!'); // Contents

  // Database (MySQL)
  const db = await nim.mysql(); // Returns a configured mysql2 connection.
  const [rows, fields] = await db.execute('SELECT * FROM `table`');
}
```

## Support

We're always happy to help you with any issues you encounter. You may want to [join our Slack community](https://nimbella-community.slack.com/) to engage with us for a more rapid response.

## License

Apache-2.0. See [LICENSE](LICENSE) to learn more.
