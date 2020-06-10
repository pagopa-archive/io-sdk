cd /tmp
virtualenv virtualenv
source virtualenv/bin/activate
pip3 install redis
apt-get update && apt-get -y install zip
zip -r /mnt/virtualenv.zip virtualenv
