FROM openwhisk/actionloop-python-v3.7:nightly
ADD nimbella-sdk-python /tmp/nimbella-sdk-python
RUN apt-get update && apt-get install -y build-essential libsasl2-dev python-dev libldap2-dev libssl-dev
RUN pip install /tmp/nimbella-sdk-python
RUN pip install python-ldap
