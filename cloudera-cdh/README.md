# Cloudera CDH

[Download CDH 6.3.2](https://docs.cloudera.com/documentation/enterprise/6/release-notes/topics/rg_cdh_63_download.html#cdh_632-download)

[Install Guide](https://docs.cloudera.com/documentation/enterprise/6/latest/topics/installation.html)

[Pre-install](https://docs.cloudera.com/documentation/enterprise/6/latest/topics/installation_reqts.html#pre-install)

- Java: OpenJDK 8: 1.8u181(tested)

[Hardware requrirement](https://docs.cloudera.com/documentation/enterprise/release-notes/topics/hardware_requirements_guide.html)

# Install

```
systemctl start ntpd.service
systemctl enable ntpd.service
systemctl start nscd.service
systemctl enable nscd.service
systemctl stop firewalld.service
systemctl disable firewalld.service
sed -i 's!SELINUX=enforcing!SELINUX=permissive!g' /etc/selinux/config
setenforce 0

yum install -y wget \
  java-1.8.0-openjdk-devel
wget https://archive.cloudera.com/cm6/6.3.1/redhat7/yum/cloudera-manager.repo -P /etc/yum.repos.d/
yum install -y cloudera-manager-daemons \
  cloudera-manager-agent \
  cloudera-manager-server

yum install -y python-pip
pip install psycopg2==2.7.5 --ignore-installed
```

Check startup script log: /var/log/message
Configure database connection for cloudera-manager
```
sudo /opt/cloudera/cm/schema/scm_prepare_database.sh -h [host] postgresql [db] [username] [password]
```

### (Skipped) Auto-TLS

Auto-TLS is an enterprise feature
```
JAVA_HOME=/usr/lib/jvm/java-1.8.0-openjdk /opt/cloudera/cm-agent/bin/certmanager setup --configure-services
```

If you already enabled it. Disable TLS by
```
psql -h 10.192.161.17 -U devops cdh
delete from CONFIGS where ATTR='web_tls';
```

### Start server

took about 5 mins to start server
```
sudo systemctl start cloudera-scm-server 

go 127.0.0.1:7180
admin/admin
```

# Client

 Installation failed. Failed to receive heartbeat from agent.

- Ensure that the host's hostname is configured properly.
- Ensure that port 7182 is accessible on the Cloudera Manager Server (check firewall rules).
- Ensure that ports 9000 and 9001 are not in use on the host being added.
- Check agent logs in /var/log/cloudera-scm-agent/ on the host being added. (Some of the logs can be found in the installation details).
- If Use TLS Encryption for Agents is enabled in Cloudera Manager (Administration -> Settings -> Security), ensure that /etc/cloudera-scm-agent/config.ini has use_tls=1 on the host being added. Restart the corresponding agent and click the Retry link here.
