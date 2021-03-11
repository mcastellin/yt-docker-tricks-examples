FROM jenkinsci/jnlp-slave

# switching to root user temporarly to install our agent
USER root

# the `jenkinsci/jnlp-slave` container is bundled with openjdk-8
# so we add the JDK 11 version from the `openjdk:11-jdk` official image
COPY --from=openjdk:11-jdk /usr/local/openjdk-11 /usr/local/openjdk-11

# copy maven 3 from its official container image
COPY --from=maven:3-jdk-8 /usr/share/maven /usr/share/maven
ENV MAVEN_HOME=/usr/share/maven
ENV PATH=${MAVEN_HOME}/bin:$PATH

# install additional packages
RUN apt-get update && apt-get install -qqy git \
    && rm -rf /var/lib/apt/lists

# container will run as jenkins user by default
USER jenkins