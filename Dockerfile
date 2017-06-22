FROM java:8-jre-alpine
ADD target/camel-spring-boot-1.0.jar /opt/app.jar
ENTRYPOINT ["sh", "-c", "java $JAVA_OPTS -jar /opt/app.jar"]