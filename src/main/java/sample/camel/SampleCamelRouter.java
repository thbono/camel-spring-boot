package sample.camel;

import org.apache.camel.builder.RouteBuilder;
import org.springframework.stereotype.Component;

@Component
public class SampleCamelRouter extends RouteBuilder {

    @Override
    public void configure() throws Exception {
        from("timer:hello?period=2000")
                .transform(constant("teste"))
                .to("log:hello");
    }

}
