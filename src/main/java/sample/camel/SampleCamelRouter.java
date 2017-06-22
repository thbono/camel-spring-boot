package sample.camel;

import org.apache.camel.builder.RouteBuilder;
import org.apache.camel.component.rabbitmq.RabbitMQConstants;
import org.apache.camel.model.dataformat.JsonLibrary;
import org.springframework.stereotype.Component;

@Component
public class SampleCamelRouter extends RouteBuilder {

    private static final String AQMP_PERSISTENT = "2";

    @Override
    public void configure() throws Exception {
        from("sql:{{app.sql.calculo-margem}}?dataSource=dataSource&consumer.delay={{app.inbound-delay}}&onConsume={{app.sql.calculo-margem-consume}}&maxMessagesPerPoll={{app.inbound-max}}")
                .marshal().json(JsonLibrary.Jackson)
                .log(body().toString())
                .setHeader(RabbitMQConstants.DELIVERY_MODE, simple(AQMP_PERSISTENT))
                .to("rabbitmq:{{spring.rabbitmq.host}}:{{spring.rabbitmq.port}}/{{app.amqp.calculo-margem}}?queue={{app.amqp.calculo-margem}}&username={{spring.rabbitmq.username}}&password={{spring.rabbitmq.password}}&autoDelete=false");
    }

}
