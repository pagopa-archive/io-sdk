/**
 * Copyright (c) 2020-present, Nimbella, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

var redis = require('redis');
var bluebird = require('bluebird');

function makeRedisClient() {
  bluebird.promisifyAll(redis.RedisClient.prototype);
  bluebird.promisifyAll(redis.Multi.prototype);
  const redisHost = process.env['__NIM_REDIS_IP'] || process.env['__OW_REDIS_IP'];
  if (!redisHost || redisHost.length == 0) {
    throw new Error('Key-Value store is not available');
  }
  const redisPassword = process.env['__NIM_REDIS_PASSWORD'] || process.env['_OW_REDIS_PASSWORD'];
  if (!redisPassword || redisPassword.length == 0) {
    throw new Error('Key-Value store password is not available');
  }
  const redisParam = {port: 6379, host: redisHost};
  const client = redis.createClient(redisParam);
  if (client == null) {
    throw new Error('Error creating redis client');
  }
  client.auth(redisPassword, function (err, reply) {
    if (err) {
      throw new Error('Error authenticaing redis client' + reply);
    }
  });
  return client;
}

async function makeStorageClient(web = false) {
  let StorageClass = undefined;
  try {
    const {Storage} = require('@google-cloud/storage');
    StorageClass = Storage;
  } catch (err) {
    throw new Error(
      'No object store client is available in this action runtime'
    );
  }
  const creds = process.env['__NIM_STORAGE_KEY'];
  if (!creds || creds.length == 0) {
    throw new Error('Objectstore credentials are not available');
  }
  const namespace = process.env['__OW_NAMESPACE'];
  const apiHost = process.env['__OW_API_HOST'];
  if (!namespace || !apiHost) {
    throw new Error(
      'Not enough information in the environment to determine the object store bucket name'
    );
  }
  const hostpart = apiHost.replace('https://', '').split('.').join('-');
  const datapart = web ? '' : 'data-';
  const bucket = `gs://${datapart}${namespace}-${hostpart}`;
  let parsedCreds = undefined;
  try {
    parsedCreds = JSON.parse(creds);
  } catch {}
  if (parsedCreds) {
    const {client_email, project_id, private_key} = parsedCreds;
    if (client_email && project_id && private_key) {
      const storageOptions = {
        project_id,
        credentials: {client_email, private_key}
      };
      const storage = new StorageClass(storageOptions);
      return storage.bucket(bucket); // a promise, since the function is async
    }
  }
  throw new Error(
    'Insufficient information in provided credentials or credentials were invalid'
  );
}

async function makeSqlClient() {
  const mysql = require('mysql2/promise');
  if (!process.env.__NIM_SQL_KEY) {
    throw new Error('Sql credentials are not available');
  }
  const creds = JSON.parse(process.env.__NIM_SQL_KEY);
  return await mysql.createConnection({
    host: creds.host,
    user: creds.user,
    database: creds.database,
    password: creds.password,
    ssl: {
      ca: creds.serverCaCert,
      cert: creds.clientCert,
      key: creds.clientKey
    }
  });
}

module.exports = {
  redis: makeRedisClient,
  storage: makeStorageClient,
  mysql: makeSqlClient
};
