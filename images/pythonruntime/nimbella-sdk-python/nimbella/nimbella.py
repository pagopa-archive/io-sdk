
"""
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
"""

import os
import json

import redis as kv
from google.cloud import storage as gstorage
from google.oauth2 import service_account


def redis():
    redisIP = os.getenv('__NIM_REDIS_IP', os.getenv('__OW_REDIS_IP', ""))
    redisPassword = os.getenv('__NIM_REDIS_PASSWORD', os.getenv('__OW_REDIS_PASSWORD', ""))
    if len(redisIP) == 0:
        raise Exception('Key-Value store is not available.')
    elif len(redisPassword) == 0:
        raise Exception('Key-Value store credentials are not available.')
    else:
        return kv.Redis(host=redisIP, port=6379, password=redisPassword)


def storage(web=False):
    namespace = os.getenv('__OW_NAMESPACE', "")
    apiHost = os.getenv('__OW_API_HOST', "")
    creds = os.getenv('__NIM_STORAGE_KEY', "")

    if len(namespace) == 0 or len(apiHost) == 0:
        raise Exception(
            'Not enough information in the environment to determine the object store bucket name.')

    if len(creds) == 0:
        raise Exception('Object store credentials are not available.')
    try:
        creds = json.loads(creds)
        client = gstorage.Client(
            creds['project_id'], credentials=service_account.Credentials.from_service_account_info(creds))
    except:
        raise Exception(
            'Insufficient information in provided credentials or credentials were invalid.')

    hostpart = '-'.join(apiHost.replace('https://', '').split('.'))
    datapart = '' if web else 'data-'
    bucket = datapart + namespace + '-' + hostpart
    return client.get_bucket(bucket)
