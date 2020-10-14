#!/usr/bin/env python
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

from setuptools import setup, find_packages

with open('requirements.txt') as f:
    requirements = [line.strip() for line in f.readlines()]

with open('README.md', 'r') as fh:
    long_description = fh.read()

setup(
    name='nimbella',
    version='1.0.0',
    author='Nimbella Corporation',
    author_email='info@nimbella.com',
    description='A python package to interact with nimbella.com services.',
    long_description=long_description,
    long_description_content_type="text/markdown",
    url="https://github.com/nimbella/nimbella-sdk-python",
    license='Apache-2.0',
    zip_safe=True,
    packages=find_packages(),
    install_requires=requirements,
    classifiers=[
        'Operating System :: OS Independent',
        'Programming Language :: Python',
        'License :: OSI Approved :: Apache Software License',
    ],
)
